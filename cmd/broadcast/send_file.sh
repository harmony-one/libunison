#!/bin/bash
# example: node 4 send file to other peers
# ./send_file 4 test.txt [coopcast|manycast]

mkdir -p received
./ida -nbr_config configs/config_$1.txt -all_config configs/config_allpeers.txt -broadcast -msg_file $2 -mode $3
