load("//tools:build_rules/shims.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "filtercmd",
    srcs = ["filtercmd.go"],
    deps = [
        "//kythe/go/platform/kzip",
        "//kythe/go/platform/tools/kzip/flags",
        "//kythe/go/platform/vfs",
        "//kythe/go/util/cmdutil",
        "//kythe/go/util/flagutil",
        "//kythe/go/util/log",
        "@com_github_google_subcommands//:go_default_library",
        "@org_bitbucket_creachadair_stringset//:go_default_library",
    ],
)
