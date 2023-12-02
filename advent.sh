#!/usr/bin/env bash

mode=$1
day=$2

if [ -z "$mode" ] || [ -z "$day" ]; then
  echo "Usage: ./advent.sh <run|test> <two-digit day (01)>"
  exit 2
fi

if [ ! -d "$day" ]; then
  echo "Directory for day $day not found" >&2
  exit 1
fi

cd "$day" || exit 1

case "$mode" in
run)
  go run .
  ;;
test)
  go test .
  ;;
esac