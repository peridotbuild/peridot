load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "thrift",
    srcs = [
        "application_exception.go",
        "binary_protocol.go",
        "buffered_transport.go",
        "compact_protocol.go",
        "debug_protocol.go",
        "deserializer.go",
        "exception.go",
        "field.go",
        "framed_transport.go",
        "http_client.go",
        "http_transport.go",
        "iostream_transport.go",
        "json_protocol.go",
        "memory_buffer.go",
        "messagetype.go",
        "multiplexed_protocol.go",
        "numeric.go",
        "pointerize.go",
        "processor.go",
        "processor_factory.go",
        "protocol.go",
        "protocol_exception.go",
        "protocol_factory.go",
        "rich_transport.go",
        "serializer.go",
        "server.go",
        "server_socket.go",
        "server_transport.go",
        "simple_json_protocol.go",
        "simple_server.go",
        "socket.go",
        "ssl_server_socket.go",
        "ssl_socket.go",
        "transport.go",
        "transport_exception.go",
        "transport_factory.go",
        "type.go",
        "zlib_transport.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift",
    importpath = "github.com/uber-go/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift",
    visibility = ["//visibility:public"],
)
