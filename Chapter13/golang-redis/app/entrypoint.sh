#!/bin/bash

set -euo pipefail

if [ $# -eq 0 ]; then
	exec /usr/local/bin/golang-redis
else
	exec $@
fi
