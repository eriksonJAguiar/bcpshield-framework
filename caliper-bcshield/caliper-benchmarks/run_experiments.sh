#!/bin/sh


for i in 1 2 3 4 5
do
    echo "Running round $i"
    npx caliper launch master --caliper-bind-sut fabric:1.4.4 --caliper-workspace . --caliper-benchconfig benchmarks/scenario/simple/bcshield/config-simple.yaml --caliper-networkconfig networks/fabric/bcshield/network-config_1.4.yaml
    sleep 3
    cp -r report.html ./reports/report-$i.html
done