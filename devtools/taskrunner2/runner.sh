#!/usr/bin/env bash

ARGS_PATH="$(find . -name "args")"
devtools/taskrunner2/taskrunner2_/taskrunner2 $(cat "${ARGS_PATH}")
