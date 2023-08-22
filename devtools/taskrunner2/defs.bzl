# Copyright 2023 Peridot Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
