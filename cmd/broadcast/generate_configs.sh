#!/bin/bash
mkdir -p "configs"
./ida -gen_config=true -graph_config=$1
