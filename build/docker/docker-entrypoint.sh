#!/bin/sh

set -eu

if [ "${1-}" = "shell" ]; then
  shift
  exec /bin/sh "$@"
else
  exec /stona/stona "$@"
fi