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

# check if a process is already running
# we don't have a lock so just use ps and grep
if ps aux | grep -v grep | grep "temporalite start" ; then
  # we're not going to fail, but silently exit with a small warning
  # this is due to different services in this repo including devtemporal as a target
  # in taskrunner2, but some situations may warrant a dedicated devtemporal
  echo "WARNING: devtemporal already running, not starting a new one"
  exit 0
fi

vendor/github.com/temporalio/temporalite/cmd/temporalite/temporalite_/temporalite start --log-format pretty --log-level fatal --ephemeral
