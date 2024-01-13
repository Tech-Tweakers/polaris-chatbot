#!/bin/bash

#
# Simple script to display new log files created by llama.cpp 
#

folder_path="."

while true; do
    latest_file=$(ls -t "$folder_path" | grep '^llama' | head -n 1)

    if [ -n "$latest_file" ]; then
        echo "Tailing the latest file: $latest_file"
        tail -f "$folder_path/$latest_file"
    else
        echo "No matching files found. Waiting..."
    fi
    sleep 1  # Adjust the sleep interval as needed
done

