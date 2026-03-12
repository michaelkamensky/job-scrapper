#!/bin/bash
source ./utils/career_site_filter.sh

./clean.sh

DATE=$(date +%Y-%m-%d)
mkdir $DATE

cp $PWD/utils/clean.sh ./$DATE
cp $PWD/sources/targets.txt ./$DATE
cd $PWD/$DATE
cat targets.txt | cariddi -sr -d 10000 -intensive -sd 2 -j 3000 -c 1 -ie png,jpg,jpeg,gif,webp,svg
cd $PWD/output-cariddi
filter
#cat targets.txt | cariddi -sr -oh test2
