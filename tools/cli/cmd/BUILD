load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["root.go"],
    importmap = "tools/cmd",
    importpath = "cli/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//tools/cli/cmd/core:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
    ],
)
