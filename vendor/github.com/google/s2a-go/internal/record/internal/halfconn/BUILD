load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "halfconn",
    srcs = [
        "ciphersuite.go",
        "counter.go",
        "expander.go",
        "halfconn.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/google/s2a-go/internal/record/internal/halfconn",
    importpath = "github.com/google/s2a-go/internal/record/internal/halfconn",
    visibility = [
        "//third_party:__subpackages__",
        "//vendor/github.com/google/s2a-go:__subpackages__",
    ],
    deps = [
        "//vendor/github.com/google/s2a-go/internal/proto/common_go_proto",
        "//vendor/github.com/google/s2a-go/internal/record/internal/aeadcrypter",
        "//vendor/golang.org/x/crypto/cryptobyte",
        "//vendor/golang.org/x/crypto/hkdf",
    ],
)
