#!/bin/bash
# This script is used to identify the user running the script and print a message accordingly.
cat hello.sh

USER=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name')
echo "Your superhero name is: $USER"