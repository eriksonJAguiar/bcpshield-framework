#!/bin/sh

sudo rm -r ./networks/fabric/bcshield/crypto-config
sudo rm -r ./networks/fabric/bcshield/channel-artifacts
sudo rm -r /tmp/hfc-*

cp -r ../../hyperledger-network/crypto-config ./networks/fabric/bcshield/
cp -r ../../hyperledger-network/channel-artifacts ./networks/fabric/bcshield/

cp networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/User1@hprovider.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/User1@hprovider.healthcare.com/msp/keystore/key_sk
cp networks/fabric/bcshield/crypto-config/peerOrganizations/research.healthcare.com/users/User1@research.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/research.healthcare.com/users/User1@research.healthcare.com/msp/keystore/key_sk
cp networks/fabric/bcshield/crypto-config/peerOrganizations/patient.healthcare.com/users/User1@patient.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/patient.healthcare.com/users/User1@patient.healthcare.com/msp/keystore/key_sk 
cp networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/Admin@hprovider.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/Admin@hprovider.healthcare.com/msp/keystore/key_sk
cp networks/fabric/bcshield/crypto-config/peerOrganizations/research.healthcare.com/users/Admin@research.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/research.healthcare.com/users/Admin@research.healthcare.com/msp/keystore/key_sk
cp networks/fabric/bcshield/crypto-config/peerOrganizations/patient.healthcare.com/users/Admin@patient.healthcare.com/msp/keystore/*_sk networks/fabric/bcshield/crypto-config/peerOrganizations/patient.healthcare.com/users/Admin@patient.healthcare.com/msp/keystore/key_sk 

# export CLI_HPROVIDER_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/User1@hprovider.healthcare.com/msp/keystore && ls *_sk)
# export CLI_RESEARCH_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/User1@research.healthcare.com/msp/keystore && ls *_sk)
# export CLI_PATIENT_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/User1@patient.healthcare.com/msp/keystore && ls *_sk)
# export HPROVIDER_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/Admin@hprovider.healthcare.com/msp/keystore && ls *_sk)
# export RESEARCH_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/Admin@research.healthcare.com/msp/keystore && ls *_sk)
# export PATIENT_KEY=$(cd networks/fabric/bcshield/crypto-config/peerOrganizations/hprovider.healthcare.com/users/Admin@patient.healthcare.com/msp/keystore && ls *_sk)

# npx caliper launch master --caliper-bind-sut fabric:1.4.4 --caliper-workspace . --caliper-benchconfig benchmarks/scenario/simple/bcshield/config.yaml --caliper-networkconfig networks/fabric/bcshield/network-config_1.4.yaml


echo
echo "|  ____| \ | |  __ \ "
echo "| |__  |  \| | |  | |"
echo "|  __| | .   | |  | |"
echo "| |____| |\  | |__| |"
echo "|______|_| \_|_____/"
echo  
echo                                     
echo " _____ _   _ _____ _______ _____          _      _____ ______      _______ _____ ____  _   _ "
echo "|_   _| \ | |_   _|__   __|_   _|   /\   | |    |_   _|___  /   /\|__   __|_   _/ __ \| \ | |"
echo "  | | |  \| | | |    | |    | |    /  \  | |      | |    / /   /  \  | |    | || |  | |  \| |"
echo "  | | | .   | | |    | |    | |   / /\ \ | |      | |   / /   / /\ \ | |    | || |  | | .   |"
echo " _| |_| |\  |_| |_   | |   _| |_ / ____ \| |____ _| |_ / /__ / ____ \| |   _| || |__| | |\  |"
echo "|_____|_| \_|_____|  |_|  |_____/_/    \_\______|_____/_____/_/    \_\_|  |_____\____/|_| \_|"
