#!/usr/bin/env bash

set -euo pipefail

vendor/github.com/temporalio/temporalite/cmd/temporalite/temporalite_/temporalite start --log-format pretty --log-level error
