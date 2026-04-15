#!/bin/bash

curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r --arg id  "$HERO_ID" '.[] | select(.id|tonumber == ($HERO_ID|tonumber)) | .connections.relatives'