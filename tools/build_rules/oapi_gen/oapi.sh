#!/usr/bin/env bash

set -euo pipefail

cli_bin="{{generator_cli_bin}}"
generator="{{generator}}"
openapi_file="{{openapi_file}}"
output_dir="{{output_dir}}"

$cli_bin generate -i $openapi_file -g $generator -o $output_dir --skip-validate-spec > /dev/null
