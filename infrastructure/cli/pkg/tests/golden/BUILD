load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "golden",
    srcs = ["golden.go"],
    importmap = "infrastructure/pkg/tests/golden",
    importpath = "cli/pkg/tests/golden",
    visibility = ["//visibility:public"],
    deps = ["@com_github_kr_pretty//:go_default_library"],
)
