#!/bin/bash

if [ ! -s "$1" ]; then
  echo "Commit message cannot be empty."
  exit 1
fi

if ! grep -qE "^(feat|fix|docs|style|refactor|test|chore)\(.*\): .+" "$1"; then
  echo "Commit message must follow the format: type(scope): message"
  echo "e.g., feat(auth): add login feature"
  exit 1
fi

echo "Commit message is valid."
