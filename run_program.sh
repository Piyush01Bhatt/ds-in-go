#!/bin/sh
set -e # Exit early if any commands fail
(
  cd "$(dirname "$0")"
  go build -o /tmp/ds-in-go cmd/*.go
)

exec /tmp/ds-in-go "$@"

