# taskrunner2
Taskrunner2 is used to specify a list of Bazel targets to run.

Useful for development. Also supports running certain targets with ibazel.

Does NOT require you to have ibazel installed.

### Macro
Services that want to use taskrunner2 in this repository should use the `taskrunner2` macro.

Example:

```
load("//devtools/taskrunner2:defs.bzl", "taskrunner2")

taskrunner2(
    name = "mothership",
    dev_frontend_flags = True,
    targets = [
        "//devtools/devtemporal",
        "//devtools/devdex",
    ],
    watch_targets = [
        "//tools/mothership/cmd/mship_api",
        "//tools/mothership/ui",
    ],
)
```

### Flags
```
name - Name of the taskrunner2 target
dev_frontend_flags - Whether to pass --dev_frontend_flags to taskrunner2. This is useful for running the frontend in dev mode and using the devdex instance.
targets - List of targets to run (bazel)
watch_targets - List of targets to watch for changes (ibazel)
```
