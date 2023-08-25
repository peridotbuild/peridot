export default {
  alias: {
    // We're adding the bazel-bin alias, so we can use generated files with
    // that prefix to make Node tooling happy for IDEs/linters/etc. but
    // also making it "just work" for Bazel.
    'bazel-bin': '.',
  }
}
