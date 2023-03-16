#!/bin/bash

set -o pipefail
go test ./... $1 | { grep -v 'no test files'; true; }