#!/bin/bash
set -euo pipefail

if [[ $# != 2 ]]
then
    echo "Usage: fetch.sh year day" > /dev/stderr
    exit 1
fi

YEAR=$1
DAY=${2#0}
DAY2=$(printf '%02d' $DAY)
echo "Year: $YEAR"
echo "Day:  $DAY"

if [[ ! -d ./$YEAR/$DAY2 ]]
then
    echo "Directory ./$YEAR/$DAY2 does not exist: creating"
    mkdir -p ./$YEAR/$DAY2
fi

COOKIE=$(cat .cookie)
curl -fsSL "https://adventofcode.com/$YEAR/day/$DAY/input" -H "Cookie: session=$COOKIE" > ./$YEAR/$DAY2/input