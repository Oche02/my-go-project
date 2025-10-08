#!/bin/bash

# Check if HERO_ID is set
if [ -z  ]; then
    echo "Error: HERO_ID is not set"
    exit 1
fi

# Fetch JSON and filter by HERO_ID
result=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id )) | .connections.relatives | gsub"\\n")')

# Check if result is empty or null
if [ -z "$result" ] || [ "$result" = "null" ]; then
    echo "No relatives found forID"
    exit 1
else
    echo -n "$result"
fi
