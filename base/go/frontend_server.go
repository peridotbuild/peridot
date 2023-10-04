// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"context"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"html/template"
	"math/rand"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	_ "embed"
)

type FrontendInfo struct {
	// NoRun is a flag to disable running the frontend server
	NoRun bool

	// MuxHandler is the HTTP handler (can be nil)
	MuxHandler http.Handler

	// Title to add to the HTML page
	Title string

	// Port is the port to serve the frontend on
	Port int

	// Self is the URL to the frontend server
	Self string

	// NoAuth is a flag to disable authentication
	NoAuth bool

	// AllowUnauthenticated is a flag to allow unauthenticated users
	AllowUnauthenticated bool

	// OIDCIssuer is the issuer to use for authentication
	OIDCIssuer string

	// OIDCClientID is the client ID to use for authentication
	OIDCClientID string

	// OIDCClientSecret is the client secret to use for authentication
	OIDCClientSecret string

	// OIDCGroup is the group to check for authentication
	OIDCGroup string

	// OIDCUserInfoOverride is a flag to override the userinfo endpoint
	// todo(mustafa): we don't need to use it yet since RESF deploys cluster external Keycloak.
	OIDCUserInfoOverride string

	// AdditionalContent is a map of paths to content to serve
	AdditionalContent map[string][]byte

	// Internal

	// unauthenticatedTemplate is the template for the unauthenticated page
	unauthenticatedTemplate string
}

type frontendTemplateData struct {
	User   UserInfo
	Prefix string
}

//go:embed assets/oh_no_unauthenticated.png
var ohNoGopher []byte
var frontendHtmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, viewport-fit=cover"
    />
    <title>{{-Title-}}</title>

    <link rel="icon" type="image/png" href="/_ga/favicon.png" />

    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      rel="stylesheet"
      href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;600;700&display=swap"
    />

    <style>
        a, a:link, a:active, a:hover {
          /* color: inherit; */
          color: #66b2ff;
        }
    </style>
</head>
<body>
    <div id="app"></div>

    {{if .User}}
    <script>
        window.__peridot_user__ = {
            sub: {{.User.Subject}},
            email: {{.User.Email}},
        }
    </script>
    {{end}}
    {{-Beta-}}
    {{if .Prefix}}
    <script>window.__peridot_prefix__ = '{{.Prefix}}'.replace('\\', '');</script>
    {{end}}

    <script src="{{.BundleJS}}"></script>
</body>
</html>
`
var frontendUnauthenticated = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, viewport-fit=cover"
    />
    <title>{{-Title-}} - Ouch</title>

    <link rel="icon" type="image/png" href="/_ga/favicon.png" />

    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      rel="stylesheet"
      href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;600;700&display=swap"
    />
</head>
<body>
    <div style="font-family:Roboto,sans-serif;display:flex;flex-flow:column;justify-content:center;align-items:center;width:100vw;height:100vh">
        <img height="120px" src="/_ga/oh_no_unauthenticated.png" /><br />
        <code>{{.}}</code><br /><br />

        <a href="/">Go back</a>
    </div>
</body>
</html>
`

const frontendAuthCookieKey = "auth_bearer"

func readDir(embedfs *embed.FS, root string) ([]string, error) {
	var paths []string
	entries, err := embedfs.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subPaths, err := readDir(embedfs, filepath.Join(root, entry.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, subPaths...)
		} else {
			paths = append(paths, filepath.Join(root, entry.Name()))
		}
	}
	return paths, nil
}

func (info *FrontendInfo) renderUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.New("unauthorized.html").Parse(info.unauthenticatedTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// frontendAuthHandler verifies that the user is authenticated
// if not redirects to /auth/oidc/login
func (info *FrontendInfo) frontendAuthHandler(provider OidcProvider, h http.Handler) http.Handler {
	excludedSuffixes := []string{
		"/auth/oidc/login",
		"/auth/oidc/callback",
		"/_ga/oh_no_unauthenticated.png",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the path is excluded, then serve the request
		for _, suffix := range excludedSuffixes {
			if strings.HasSuffix(r.URL.Path, suffix) {
				h.ServeHTTP(w, r)
				return
			}
		}

		ctx := r.Context()

		// get auth cookie
		authCookie, err := r.Cookie(frontendAuthCookieKey)
		if err != nil {
			// only redirect if not allowed unauthenticated
			if !info.AllowUnauthenticated {
				// redirect to login
				http.Redirect(w, r, info.Self+"/auth/oidc/login", http.StatusFound)
				return
			}
		} else {
			// verify the token
			userInfo, err := provider.UserInfo(r.Context(), oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: authCookie.Value,
				TokenType:   "Bearer",
			}))
			if err != nil {
				// redirect to login
				http.Redirect(w, r, info.Self+"/auth/oidc/login", http.StatusFound)
				return
			}

			// Check if the user is in the group
			var claims oidcClaims
			err = userInfo.Claims(&claims)
			if err != nil {
				// redirect to login
				http.Redirect(w, r, info.Self+"/auth/oidc/login", http.StatusFound)
				return
			}

			groups := claims.Groups
			if info.OIDCGroup != "" {
				if !Contains(groups, info.OIDCGroup) {
					// show unauthenticated page
					info.renderUnauthorized(w, fmt.Sprintf("User is not in group %s", info.OIDCGroup))
					return
				}
			}

			// Add the user to the context
			ctx = context.WithValue(ctx, "user", userInfo)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FrontendServer(info *FrontendInfo, embedfs *embed.FS) error {
	if info == nil {
		return errors.New("frontend info is nil")
	}
	port := info.Port

	// Using info.Self, let's determine a path prefix
	// If info.Self is empty, then we'll use the root path
	prefix := ""
	if info.Self != "" {
		parsed, err := url.Parse(info.Self)
		if err != nil {
			return fmt.Errorf("failed to parse self url: %w", err)
		}
		prefix = parsed.Path
	}

	newTemplate := frontendHtmlTemplate
	newUnauthenticatedTemplate := frontendUnauthenticated

	// Set the title
	if info.Title == "" {
		info.Title = "Peridot"
	}
	newTemplate = strings.ReplaceAll(newTemplate, "{{-Title-}}", info.Title)
	newUnauthenticatedTemplate = strings.ReplaceAll(newUnauthenticatedTemplate, "{{-Title-}}", info.Title)

	info.unauthenticatedTemplate = newUnauthenticatedTemplate

	pathToContent := map[string][]byte{
		prefix + "/_ga/oh_no_unauthenticated.png": ohNoGopher,
	}

	// Read the files from the embedfs
	paths, err := readDir(embedfs, "bundle")
	if err != nil {
		return fmt.Errorf("failed to read embedfs: %w", err)
	}

	for _, path := range paths {
		content, err := embedfs.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read embedfs: %w", err)
		}

		// Sha256 hash of the content to add to name
		hash := sha256.New()
		hash.Write(content)
		hashSum := hex.EncodeToString(hash.Sum(nil))

		ext := filepath.Ext(path)
		noExtName := path[:len(path)-len(ext)]
		newPath := fmt.Sprintf("%s/_ga/%s.%s%s", prefix, noExtName, hashSum[:8], ext)

		pathToContent[newPath] = content

		// If name is entrypoint.js, replace the template
		if filepath.Base(path) == "entrypoint.js" {
			newTemplate = strings.ReplaceAll(newTemplate, "{{.BundleJS}}", newPath)
		}
	}

	// Add additional content
	if info.AdditionalContent != nil {
		for path, content := range info.AdditionalContent {
			pathToContent[path] = content
		}
	}

	// Log the paths
	LogInfof("frontend server paths:")
	for path := range pathToContent {
		LogInfof("  %s", path)
	}

	// Serve the content
	http.HandleFunc(prefix+"/_ga/", func(w http.ResponseWriter, r *http.Request) {
		mimeType := mime.TypeByExtension(filepath.Ext(r.URL.Path))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		w.Header().Set("Content-Type", mimeType)
		_, _ = w.Write([]byte(pathToContent[r.URL.Path]))
	})

	http.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// Set beta script (if beta basically set window.__beta__ = true)
		srvTemplate := newTemplate
		betaScript := ""
		if r.Header.Get("x-peridot-beta") == "true" {
			betaScript = "<script>window.__beta__ = true;</script>"
		}
		srvTemplate = strings.ReplaceAll(srvTemplate, "{{-Beta-}}", betaScript)

		tmpl, err := template.New("index.html").Parse(srvTemplate)
		if err != nil {
			info.renderUnauthorized(w, fmt.Sprintf("Failed to parse template: %v", err))
			return
		}

		// If user is in context, then add it to the template
		data := frontendTemplateData{
			Prefix: prefix,
		}

		user := r.Context().Value("user")
		if user != nil {
			data.User = user.(UserInfo)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			info.renderUnauthorized(w, fmt.Sprintf("Failed to execute template: %v", err))
			return
		}
	})

	// Handle other _ga meta routes
	http.HandleFunc(prefix+"/_ga/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("ok"))
	})

	// Handle auth routes
	var provider OidcProvider
	if !info.NoAuth {
		ctx := context.TODO()
		provider2, err := oidc.NewProvider(ctx, info.OIDCIssuer)
		if err != nil {
			return fmt.Errorf("failed to create oidc provider: %w", err)
		}

		provider = &OidcProviderImpl{provider2}

		redirectURL := info.Self + "/auth/oidc/callback"

		oauth2Config := oauth2.Config{
			ClientID:     info.OIDCClientID,
			ClientSecret: info.OIDCClientSecret,
			Endpoint:     provider2.Endpoint(),
			RedirectURL:  redirectURL,
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "groups"},
		}

		http.HandleFunc(prefix+"/auth/oidc/login", func(w http.ResponseWriter, r *http.Request) {
			// Generate a random state
			state := ""
			for i := 0; i < 16; i++ {
				state += strconv.Itoa(rand.Intn(10))
			}

			// Generate the auth url
			authURL := oauth2Config.AuthCodeURL(state)

			// Set the state cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "auth_state",
				Value: state,
				Path:  "/",
				// expires in 2 minutes
				MaxAge: 120,
				// secure if self is https
				Secure: strings.HasPrefix(info.Self, "https://"),
			})

			// Redirect to the auth url
			http.Redirect(w, r, authURL, http.StatusFound)
		})

		http.HandleFunc(prefix+"/auth/oidc/callback", func(w http.ResponseWriter, r *http.Request) {
			// Get the state cookie
			stateCookie, err := r.Cookie("auth_state")
			if err != nil {
				info.renderUnauthorized(w, fmt.Sprintf("Failed to get state cookie: %v", err))
				return
			}

			// Get the state query param
			stateQueryParam := r.URL.Query().Get("state")
			if stateQueryParam == "" {
				info.renderUnauthorized(w, "No state parameter in query")
				return
			}

			// Check if the state cookie and state query param match
			if stateCookie.Value != stateQueryParam {
				info.renderUnauthorized(w, "State cookie and state query param do not match")
				return
			}

			// Exchange the code for a token
			token, err := oauth2Config.Exchange(r.Context(), r.URL.Query().Get("code"))
			if err != nil {
				info.renderUnauthorized(w, fmt.Sprintf("Failed to exchange code: %v", err))
				return
			}

			// Verify the token
			accessToken := token.AccessToken
			userInfo, err := provider.UserInfo(r.Context(), oauth2.StaticTokenSource(token))
			if err != nil {
				info.renderUnauthorized(w, fmt.Sprintf("Failed to get userinfo: %v", err))
				return
			}

			// Check if the user is in the group
			if info.OIDCGroup != "" {
				var claims oidcClaims
				err := userInfo.Claims(&claims)
				if err != nil {
					info.renderUnauthorized(w, fmt.Sprintf("Failed to get claims: %v", err))
					return
				}

				groups := claims.Groups

				found := false
				for _, group := range groups {
					if group == info.OIDCGroup {
						found = true
						break
					}
				}

				if !found {
					info.renderUnauthorized(w, fmt.Sprintf("User is not in group %s", info.OIDCGroup))
					return
				}
			}

			// Set the auth cookie
			http.SetCookie(w, &http.Cookie{
				Name:  frontendAuthCookieKey,
				Value: accessToken,
				Path:  "/",
				// expires in 2 hours
				MaxAge: 7200,
				// secure if self is https
				Secure: strings.HasPrefix(info.Self, "https://"),
			})

			// Redirect to self, this is due to the "root" not being / for all apps
			http.Redirect(w, r, info.Self, http.StatusFound)
		})

		// Handle logout
		http.HandleFunc(prefix+"/auth/oidc/logout", func(w http.ResponseWriter, r *http.Request) {
			// Delete the auth cookie
			http.SetCookie(w, &http.Cookie{
				Name:   frontendAuthCookieKey,
				Value:  "",
				Path:   "/",
				MaxAge: -1,
				// secure if self is https
				Secure: strings.HasPrefix(info.Self, "https://"),
			})

			// Redirect to self, this is due to the "root" not being / for all apps
			http.Redirect(w, r, info.Self, http.StatusFound)
		})
	}

	var handler http.Handler = nil
	// if auth is enabled as well, then wrap the handler with the auth handler
	if !info.NoAuth {
		handler = info.frontendAuthHandler(provider, http.DefaultServeMux)
	} else {
		handler = http.DefaultServeMux
	}
	info.MuxHandler = handler

	if !info.NoRun {
		LogInfof("starting frontend server on port %d", port)
		if err := http.ListenAndServe(":"+strconv.Itoa(port), handler); err != nil {
			return fmt.Errorf("failed to start frontend server: %w", err)
		}
	}

	return nil
}
