#!/bin/bash
# example: start servers with node_id 0,1,2,3 in background
# ./startserver 4 [coopcast|manycast]

mkdir -p logs
unset -v i
i=0
while :
do
    case $((${i} < ${1})) in
    0)
        break
        ;;
    esac
   ./ida -nbr_config configs/config_$i.txt  -all_config configs/config_allpeers.txt > logs/server_$i.out -mode $2 2>&1 &
   i=$((${i} + 1))
done
