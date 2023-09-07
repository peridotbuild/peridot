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
    name = "srpm_import",
    srcs = [
        "srpm_import.go",
        "srpmproc_compat.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/worker_server/srpm_import",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go/storage",
        "//vendor/github.com/go-git/go-billy/v5:go-billy",
        "//vendor/github.com/go-git/go-git/v5:go-git",
        "//vendor/github.com/go-git/go-git/v5/config",
        "//vendor/github.com/go-git/go-git/v5/plumbing",
        "//vendor/github.com/go-git/go-git/v5/plumbing/object",
        "//vendor/github.com/go-git/go-git/v5/storage",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/rocky-linux/srpmproc/pb",
        "//vendor/github.com/rocky-linux/srpmproc/pkg/data",
        "//vendor/github.com/rocky-linux/srpmproc/pkg/directives",
        "//vendor/github.com/rocky-linux/srpmproc/pkg/misc",
        "//vendor/github.com/sassoftware/go-rpmutils",
        "//vendor/golang.org/x/crypto/openpgp",
        "@org_golang_google_protobuf//encoding/prototext",
    ],
)

go_test(
    name = "srpm_import_test",
    size = "small",
    srcs = [
        "srpm_import_test.go",
        "srpmproc_compat_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":srpm_import"],
    deps = [
        "//base/go/storage/memory",
        "//vendor/github.com/go-git/go-billy/v5/memfs",
        "//vendor/github.com/go-git/go-billy/v5/osfs",
        "//vendor/github.com/go-git/go-git/v5:go-git",
        "//vendor/github.com/go-git/go-git/v5/config",
        "//vendor/github.com/go-git/go-git/v5/plumbing/cache",
        "//vendor/github.com/go-git/go-git/v5/plumbing/object",
        "//vendor/github.com/go-git/go-git/v5/storage/filesystem",
        "//vendor/github.com/go-git/go-git/v5/storage/memory",
        "//vendor/github.com/stretchr/testify/require",
        "//vendor/golang.org/x/crypto/openpgp",
    ],
)
