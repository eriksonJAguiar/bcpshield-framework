#!/bin/sh

#Parameter config.yaml caliper benchmark folder


echo "                    ______                          "
echo "                   / _____) _                 _     "
echo "                  ( (____ _| |_ _____  ____ _| |_   "
echo "                   \____ (_   _|____ |/ ___|_   _)  "
echo "                   _____) )| |_/ ___ | |     | |_   "
echo "                  (______/  \__)_____|_|      \__)  "
echo "                                                    "    
echo "                                  _                 "               
echo "                                 (_)                   _             "        
echo "    _____ _   _ ____  _____  ____ _ ____  _____ ____ _| |_  ___      "
echo "   | ___ ( \ / )  _ \| ___ |/ ___) |    \| ___ |  _ (_   _)/___)     "
echo "   | ____|) X (| |_| | ____| |   | | | | | ____| | | || |_|___ |     "
echo "   |_____|_/ \_)  __/|_____)_|   |_|_|_|_|_____)_| |_| \__|___/ .... "
echo "                                                                     "

for i in 1 2 3 4 5
do
    echo "Running round $i"
    npx caliper launch master --caliper-bind-sut fabric:1.4.4 --caliper-workspace . --caliper-benchconfig benchmarks/scenario/simple/bcshield/$1 --caliper-networkconfig networks/fabric/bcshield/network-config_1.4.yaml
    sleep 3
    cp -r report.html ./reports/kanonimity/report-$i.html
done

echo "                      _______ _______ ______        "
echo "                     (_______|_______|______)       "
echo "                      _____   _     _ _     _       "
echo "                     |  ___) | |   | | |   | |      "
echo "                     | |_____| |   | | |__/ /       "
echo "                     |_______)_|   |_|_____/          "
echo "                                                    "    
echo "                                  _                 "       
echo "                                  _                 "               
echo "                                 (_)                   _             "        
echo "    _____ _   _ ____  _____  ____ _ ____  _____ ____ _| |_  ___      "
echo "   | ___ ( \ / )  _ \| ___ |/ ___) |    \| ___ |  _ (_   _)/___)     "
echo "   | ____|) X (| |_| | ____| |   | | | | | ____| | | || |_|___ |     "
echo "   |_____|_/ \_)  __/|_____)_|   |_|_|_|_|_____)_| |_| \__|___/ .... "
echo "  