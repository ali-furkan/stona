#!/bin/sh

set -eu

make install

make build

mv ./stona $1/stona

echo ✅ Moved to $1 folder