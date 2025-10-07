#!/bin/bash

# to-git-or-not-to-git.sh

curl -s "https://acad.learn2earn.ng/assets/superhero/all.json" | \
jq -r '.[] | select(.id == 170) | "\(.name)\n\(.power)\n\(.gender)"'