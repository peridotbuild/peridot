package base

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
)

const UserContextKey = "user"

type OidcInterceptorDetails struct {
	Issuer               string
	Group                string
	AllowUnauthenticated bool
}

type oidcClaims struct {
	Groups []string
}

func OidcGrpcInterceptor(details *OidcInterceptorDetails) (grpc.UnaryServerInterceptor, error) {
	ctx := context.TODO()
	provider, err := oidc.NewProvider(ctx, details.Issuer)
	if err != nil {
		return nil, err
	}

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "missing metadata")
		}

		token, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			if details.AllowUnauthenticated {
				return handler(ctx, req)
			}

			// let's check if there is a cookie
			cookie := md.Get("cookie")
			if len(cookie) == 0 {
				if details.AllowUnauthenticated {
					return handler(ctx, req)
				}
				return nil, status.Error(codes.Unauthenticated, "missing auth token")
			}

			// parse the cookie
			header := http.Header{}
			header.Add("Cookie", cookie[0])
			req := http.Request{Header: header}

			// verify the token
			cookieToken, err := req.Cookie(frontendAuthCookieKey)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}

			token = cookieToken.Value
		}

		// verify the token
		userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: token,
			TokenType:   "bearer",
		}))
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// check if the user is in the group
		if details.Group != "" {
			var claims oidcClaims
			if err := userInfo.Claims(&claims); err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}

			if !Contains[string](claims.Groups, details.Group) {
				return nil, status.Error(codes.PermissionDenied, "user not in group")
			}
		}

		// add user to context
		ctx = context.WithValue(ctx, UserContextKey, userInfo)

		return handler(ctx, req)
	}

	return interceptor, nil
}
