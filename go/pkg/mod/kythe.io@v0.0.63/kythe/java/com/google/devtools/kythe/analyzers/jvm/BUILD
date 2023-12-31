load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(
    default_visibility = ["//kythe:default_visibility"],
    licenses = ["notice"],
)

java_library(
    name = "jvm_graph",
    srcs = ["JvmGraph.java"],
    deps = [
        "//kythe/java/com/google/devtools/kythe/analyzers/base:entryset",
        "//kythe/java/com/google/devtools/kythe/analyzers/base:schema",
        "//kythe/java/com/google/devtools/kythe/platform/shared:statistics",
        "//kythe/proto:common_java_proto",
        "//kythe/proto:storage_java_proto",
        "//third_party/guava",
        "@com_google_protobuf//:protobuf_java",
    ],
)

java_library(
    name = "analyzer",
    srcs = [
        "ClassFileIndexer.java",
        "KytheClassVisitor.java",
    ],
    deps = [
        ":jvm_graph",
        "//kythe/java/com/google/devtools/kythe/analyzers/base:config",
        "//kythe/java/com/google/devtools/kythe/analyzers/base:entryset",
        "//kythe/java/com/google/devtools/kythe/analyzers/base:schema",
        "//kythe/java/com/google/devtools/kythe/analyzers/base:stream_fact_emitter",
        "//kythe/java/com/google/devtools/kythe/common:flogger",
        "//kythe/java/com/google/devtools/kythe/extractors/jvm:jvm_extractor",
        "//kythe/java/com/google/devtools/kythe/extractors/shared",
        "//kythe/java/com/google/devtools/kythe/platform/kzip",
        "//kythe/java/com/google/devtools/kythe/platform/shared:analysis_exception",
        "//kythe/java/com/google/devtools/kythe/platform/shared:file_data",
        "//kythe/java/com/google/devtools/kythe/platform/shared:statistics",
        "//kythe/java/com/google/devtools/kythe/util:json",
        "//kythe/proto:analysis_java_proto",
        "//kythe/proto:java_java_proto",
        "//kythe/proto:storage_java_proto",
        "//third_party/guava",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:protobuf_java",
        "@maven//:com_beust_jcommander",
        "@maven//:org_checkerframework_checker_qual",
        "@maven//:org_ow2_asm_asm",
    ],
)

java_binary(
    name = "class_file_indexer",
    main_class = "com.google.devtools.kythe.analyzers.jvm.ClassFileIndexer",
    runtime_deps = [":analyzer"],
)
