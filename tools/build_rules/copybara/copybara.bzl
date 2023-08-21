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

def copybara(name):
    """Copybara is a tool for transforming and moving code between repositories."""
    native.sh_binary(
        name = name,
        srcs = ["//tools/build_rules/copybara:sync.bash"],
        env = {
            "TARGET_DIR": name,
        },
        visibility = ["//visibility:public"],
    )
