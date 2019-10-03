#!/bin/bash

# Exit on error
set -Eeu

mkdir -p build/coverage
echo "mode: count" > build/coverage/tmp.out
for package in $@; do
  go test -covermode=count -coverprofile build/coverage/profile.out "${package}"
  if [ -f build/coverage/profile.out ]; then
    tail -q -n +2 build/coverage/profile.out >> build/coverage/tmp.out
    rm build/coverage/profile.out
  fi
done

# Ignore generated files
cat build/coverage/tmp.out | grep -v ".pb.go" --exclude-dir=examples --exclude-dir=e2e > build/coverage/cover.out

# Generate coverage report in html formart
go tool cover -func=build/coverage/cover.out
go tool cover -html=build/coverage/cover.out -o build/coverage/coverage.html

# Remove temporary file
rm build/coverage/tmp.out build/coverage/cover.out
