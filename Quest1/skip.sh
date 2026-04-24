#!/bin/bash
# This script lists every second file in the current directory, starting with the first one.
ls -l | sed -n '1~2p'

# Alternative solution using awk:
ls -l | awk 'NR % 2 == 1'