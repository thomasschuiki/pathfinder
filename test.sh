#! /usr/bin/env sh
PACKAGE=${1:-"..."}
go test -v "./${PACKAGE}" | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
