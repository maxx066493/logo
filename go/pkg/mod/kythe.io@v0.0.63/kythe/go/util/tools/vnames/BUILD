load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "vnames",
    srcs = [
        "apply.go",
        "convert.go",
        "vnames.go",
    ],
    deps = [
        "//kythe/go/platform/vfs",
        "//kythe/go/util/cmdutil",
        "//kythe/go/util/log",
        "//kythe/go/util/vnameutil",
        "//kythe/proto:storage_go_proto",
        "@com_github_google_subcommands//:go_default_library",
        "@org_golang_google_protobuf//encoding/protojson:go_default_library",
    ],
)
