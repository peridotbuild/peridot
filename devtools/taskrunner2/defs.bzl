def taskrunner2(name, targets = [], watch_targets = []):
    args = []

    for target in targets:
        args.append("-t")
        args.append(target)

    for target in watch_targets:
        args.append("-w")
        args.append(target)

    # generate args output
    native.genrule(
        name = name + "_args",
        outs = ["args"],
        cmd = "echo " + " ".join(args) + " > $@",
        visibility = ["//visibility:private"],
    )

    native.sh_binary(
        name = name,
        srcs = [
            "//devtools/taskrunner2:runner.sh",
        ],
        data = [
            "//devtools/taskrunner2",
            ":" + name + "_args",
        ],
    )
