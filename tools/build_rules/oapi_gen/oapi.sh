#!/usr/bin/env bash
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


set -euo pipefail

java_bin="{{java_bin}}"
cli_jar="{{generator_cli_jar}}"
generator="{{generator}}"
openapi_file="{{openapi_file}}"
output_dir="{{output_dir}}"

# Generate the client.
$java_bin -jar $cli_jar generate -i $openapi_file -g $generator -o $output_dir --skip-validate-spec > /dev/null
