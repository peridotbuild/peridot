# Copyright 2023 Peridot Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "worker_server",
    srcs = [
        "entry.go",
        "process_rpm.go",
        "retract_entry.go",
        "worker.go",
        "workflows.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/worker_server",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//base/go/storage",
        "//tools/mothership/db",
        "//tools/mothership/proto/admin/v1:pb",
        "//tools/mothership/proto/v1:pb",
        "//tools/mothership/worker_server/forge",
        "//tools/mothership/worker_server/srpm_import",
        "//vendor/github.com/go-git/go-billy/v5:go-billy",
        "//vendor/github.com/go-git/go-billy/v5/memfs",
        "//vendor/github.com/go-git/go-git/v5:go-git",
        "//vendor/github.com/go-git/go-git/v5/config",
        "//vendor/github.com/go-git/go-git/v5/plumbing",
        "//vendor/github.com/go-git/go-git/v5/storage/memory",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/sassoftware/go-rpmutils",
        "//vendor/go.temporal.io/sdk/temporal",
        "//vendor/go.temporal.io/sdk/workflow",
        "//vendor/golang.org/x/crypto/openpgp",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)

go_test(
    name = "worker_server_test",
    size = "small",
    srcs = [
        "entry_test.go",
        "forge_test.go",
        "main_test.go",
        "process_rpm_test.go",
        "retract_entry_test.go",
        "worker_test.go",
        "workflows_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":worker_server"],
    embedsrcs = ["testdata/RPM-GPG-KEY-Rocky-8"],
    deps = [
        "//base/go",
        "//base/go/storage/memory",
        "//tools/mothership/db",
        "//tools/mothership/migrations",
        "//tools/mothership/proto/v1:pb",
        "//tools/mothership/worker_server/forge",
        "//vendor/github.com/go-git/go-billy/v5/osfs",
        "//vendor/github.com/go-git/go-git/v5/plumbing/transport/http",
        "//vendor/github.com/stretchr/testify/mock",
        "//vendor/github.com/stretchr/testify/require",
        "//vendor/github.com/stretchr/testify/suite",
        "//vendor/github.com/testcontainers/testcontainers-go",
        "//vendor/github.com/testcontainers/testcontainers-go/modules/postgres",
        "//vendor/github.com/testcontainers/testcontainers-go/wait",
        "//vendor/go.temporal.io/sdk/log",
        "//vendor/go.temporal.io/sdk/testsuite",
        "//vendor/golang.org/x/crypto/openpgp",
    ],
)
