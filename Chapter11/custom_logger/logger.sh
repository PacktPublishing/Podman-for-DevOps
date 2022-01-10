#!/bin/bash
set -euo pipefail
trap "echo Exited; exit;" SIGINT SIGTERM

# Run an endless loop writing a simple log entry with date
count=1
while true; do
	echo "$(date +%y/%m/%d_%H:%M:%S) - Line #$count" | tee -a /var/log/custom.log 
	count=$((count+1))
	sleep 2
done

