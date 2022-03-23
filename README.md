# redispubsubtest
redispubsubtest

for i in 00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19;do cat temp | grep CONFIG:$i | grep set > temp_set; cat temp_set | grep set -n | tail -1;done
