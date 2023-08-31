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

package mothershipadmin_rpc

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/tools/mothership/migrations"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
	"time"
)

var (
	s        *Server
	userInfo base.UserInfo
)

type fakeTemporalClient struct {
	client.Client
}

func TestMain(m *testing.M) {
	// Create temporary file
	dir, err := os.MkdirTemp("", "test-db-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	scripts, err := base.EmbedFSToOSFS(dir, migrations.UpSQLs, ".")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	pgContainer, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(scripts...),
		postgres.WithDatabase("mshiptest"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		panic(err)
	}
	defer pgContainer.Terminate(ctx)

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}

	db, err := base.NewDB(connStr)
	if err != nil {
		panic(err)
	}

	provider := base.NewTestOidcProvider(&userInfo)

	interceptorDetails := &base.OidcInterceptorDetails{
		Provider: provider,
		Group:    "",
	}
	s, err = NewServer(db, &fakeTemporalClient{}, interceptorDetails)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func testContext() context.Context {
	mdMap := map[string]string{}
	if userInfo != nil {
		mdMap["authorization"] = "bearer " + userInfo.Subject()
	}
	md := metadata.New(mdMap)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	return context.WithValue(ctx, "user", userInfo)
}
