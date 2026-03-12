#!/bin/bash
source ./utils/career_site_filter.sh

./clean.sh

DATE=$(date +%Y-%m-%d)
mkdir $DATE

cp $PWD/utils/clean.sh ./$DATE
cp $PWD/sources/targets.txt ./$DATE
cd $PWD/$DATE
cat targets.txt | cariddi -sr -md 10 -d 100
cd $PWD/output-cariddi
filter
#cat targets.txt | cariddi -sr -oh test2
