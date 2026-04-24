#!/bin/bash
# Count the number of files in the current directory and all subdirectories
find . \( -type f -o -type d \) | wc -l