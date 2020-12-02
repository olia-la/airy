load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["start.go"],
    importmap = "tools/cmd/core/start",
    importpath = "cli/cmd/core/start",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_spf13_cobra//:go_default_library",
    ],
)
