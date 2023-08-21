#!/usr/bin/env bash

if [[ -z "$TARGET_DIR" ]]; then
  echo "TARGET_DIR must be set"
  exit 1
fi

if [[ -z "$BUILD_WORKSPACE_DIRECTORY" ]]; then
  echo "BUILD_WORKSPACE_DIRECTORY must be set"
  exit 1
fi

FLAGS="default"
if [[ -n "$SOURCE_REF" ]]; then
  FLAGS="$FLAGS $SOURCE_REF"
fi

cd "$BUILD_WORKSPACE_DIRECTORY" || exit 1

bazel run //third_party:copybara_uberjar -- \
  migrate \
  "$BUILD_WORKSPACE_DIRECTORY/third_party/copybara/$TARGET_DIR/copy.bara.sky" \
  --init-history \
  --folder-dir "$BUILD_WORKSPACE_DIRECTORY" \
  $FLAGS