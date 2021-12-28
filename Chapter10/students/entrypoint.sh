#!/bin/bash

set -euo pipefail

if [ $# -eq 0 ]; then
	exec /usr/local/bin/students
else
	exec $@
fi
