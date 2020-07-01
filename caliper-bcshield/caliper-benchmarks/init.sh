#!/bin/sh

sudo rm -r ./networks/fabric/bcshield/crypto-config
sudo rm -r ./networks/fabric/bcshield/channel-artifacts

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