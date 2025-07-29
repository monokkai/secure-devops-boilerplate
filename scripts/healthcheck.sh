#!/usr/bin/env bash
set -e

URL="http://locahost:{PORT:-9090}/health"
status=$(curl -s -o /dev/null -w "%{http_code}" "$URL" || true)

if [ "$status" = "200" ]; then
  echo "Service is healthy!"
  exit 0
else
  echo "Healthcheck failed with status $status"
fi