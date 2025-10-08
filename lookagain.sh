#!/bin/bash

# Find all .sh files, get basename without extension, sort descending
find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r
