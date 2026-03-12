#!/bin/sh
set -e

SERVICE_DIR="/app/services/parser"
MODULE_NAME="cryptonews/parser"

cd "$SERVICE_DIR"

if [ ! -f go.mod ]; then
  echo "go.mod not found, initializing module..."
  go mod init "$MODULE_NAME"
fi

go mod tidy

exec air -c .air.toml