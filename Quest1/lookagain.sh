#!/bin/bash
find . -type f -name "*.sh" -exec basename {} .sh + | sort -r

Alternative solution using sed:
find . -type f -name "*.sh" \
| sed 's#.*/##' \
| sed 's/\.sh$//' \
| sort -r

Cleaner alternative solution using sed:
find . 