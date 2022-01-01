#!/usr/bin/env bash

set -e

go test -race -v -coverpkg=./internal/... -coverprofile=profile.out ./internal/...
go tool cover -func profile.out
