load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "golden_files",
    srcs = [
        "cli.app.reset.no-args.golden",
        "cli.kafka.consumergroup.golden",
        "cli.kafka.status.golden",
        "cli.no-args.golden",
    ],
)

go_library(
    name = "tests",
    srcs = ["mockserver.go"],
    importmap = "infrastructure/pkg/tests",
    importpath = "cli/pkg/tests",
    deps = [
        "@io_goji//:go_default_library",
        "@io_goji//pat:go_default_library",
    ],
)
