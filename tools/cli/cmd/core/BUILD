load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "core.go",
        "version.go",
    ],
    importmap = "tools/cmd/core",
    importpath = "cli/cmd/core",
    visibility = ["//visibility:public"],
    deps = [
        "//tools/cli/cmd/core/auth:go_default_library",
        "//tools/cli/cmd/core/start:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
    ],
)
