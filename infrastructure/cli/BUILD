# gazelle:prefix cli
# gazelle:importmap_prefix infrastructure
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importmap = "infrastructure",
    importpath = "cli",
    visibility = ["//visibility:private"],
    deps = ["//infrastructure/cli/cmd:go_default_library"],
)

go_binary(
    name = "airy",
    out = "airy",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
