#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Installing new contract ...."
echo

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/healthcare.com/orderers/orderer.healthcare.com/msp/tlscacerts/tlsca.healthcare.com-cert.pem
PEER0_HPROVIDER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hprovider.healthcare.com/peers/peer0.hprovider.healthcare.com/tls/ca.crt
PEER0_RESEARCH_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/research.healthcare.com/peers/peer0.research.healthcare.com/tls/ca.crt
PEER0_ORG3_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.healthcare.com/peers/peer0.org3.healthcare.com/tls/ca.crt

CONTRACT=dicom
rg_name=(hprovider research)
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
NO_CHAINCODE="$6"
NEW_VERSION_CHAINCODE="$7"
echo $NEW_VERSION_CHAINCODE
: ${CHANNEL_NAME:="healthchannel"}
: ${DELAY:="3"}
: ${LANGUAGE:="node"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
: ${NO_CHAINCODE:="false"}
: ${NEW_VERSION_CHAINCODE:="1.0"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=10


CC_SRC_PATH="github.com/chaincode/Dicom-contract"
if [ "$LANGUAGE" = "node" ]; then
	CC_SRC_PATH="/opt/gopath/src/github.com/chaincode/Dicom-contract"
fi

if [ "$LANGUAGE" = "java" ]; then
	CC_SRC_PATH="/opt/gopath/src/github.com/chaincode/Dicom-contract"
fi


verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
    echo
    exit 1
  fi
}


echo "Channel name : "$CHANNEL_NAME
setGlobals() {
  PEER=$1
  ORG=$2
  if [ $ORG -eq 1 ]; then
    CORE_PEER_LOCALMSPID="HProviderMSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_HPROVIDER_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hprovider.healthcare.com/users/Admin@hprovider.healthcare.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.hprovider.healthcare.com:7051
    elif [ $PEER -eq 1 ]; then
      CORE_PEER_ADDRESS=peer1.hprovider.healthcare.com:8051
    elif [ $PEER -eq 2 ]; then
      CORE_PEER_ADDRESS=peer2.hprovider.healthcare.com:11051
    elif [ $PEER -eq 3 ]; then
      CORE_PEER_ADDRESS=peer3.hprovider.healthcare.com:12051
    else
      CORE_PEER_ADDRESS=peer4.hprovider.healthcare.com:13051
    fi
  elif [ $ORG -eq 2 ]; then
    CORE_PEER_LOCALMSPID="ResearchMSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_RESEARCH_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/research.healthcare.com/users/Admin@research.healthcare.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.research.healthcare.com:9051
    elif [ $PEER -eq 1 ]; then
      CORE_PEER_ADDRESS=peer1.research.healthcare.com:10051
    elif [ $PEER -eq 2 ]; then
      CORE_PEER_ADDRESS=peer2.research.healthcare.com:14051
    elif [ $PEER -eq 3 ]; then
      CORE_PEER_ADDRESS=peer3.research.healthcare.com:15051
    else
      CORE_PEER_ADDRESS=peer4.research.healthcare.com:16051
    fi

  elif [ $ORG -eq 3 ]; then
    CORE_PEER_LOCALMSPID="Org3MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.healthcare.com/users/Admin@org3.healthcare.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org3.healthcare.com:11051
    else
      CORE_PEER_ADDRESS=peer1.org3.healthcare.com:12051
    fi
  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

installChaincode() {
  PEER=$1
  ORG=$2
  VER=$3
  setGlobals $PEER $ORG
  VERSION=${3:-${VER}}
  set -x
  peer chaincode install -n ${CONTRACT} -v ${VERSION} -l ${LANGUAGE} -p ${CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on peer${PEER}.${org_name[$ORG-1]} has failed"
  echo "===================== Chaincode is installed on peer${PEER}.${org_name[$ORG-1]} ===================== "
  echo
}

instantiateChaincode() {
  PEER=$1
  ORG=$2
  VER=$3
  setGlobals $PEER $ORG
  VERSION=${3:-${VER}}

  # while 'peer chaincode' command can get the orderer endpoint from the peer
  # (if join was successful), let's supply it directly as we know it using
  # the "-o" option
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer chaincode instantiate -o orderer.healthcare.com:7050 -C $CHANNEL_NAME -n ${CONTRACT} -l ${LANGUAGE} -v ${VERSION} -c '{"Args":[]}' -P "AND ('HProviderMSP.peer','ResearchMSP.peer')" >&log.txt
    res=$?
    set +x
  else
    set -x
    peer chaincode instantiate -o orderer.healthcare.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CONTRACT} -l ${LANGUAGE} -v ${VERSION} -c '{"Args":[]}' -P "AND ('HProviderMSP.peer','ResearchMSP.peer')" >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Chaincode instantiation on peer${PEER}.${org_name[$ORG-1]} on channel '$CHANNEL_NAME' failed"
  echo "===================== Chaincode is instantiated on peer${PEER}.${org_name[$ORG-1]}on channel '$CHANNEL_NAME' ===================== "
  echo
}

upgradeChaincode() {
  PEER=$1
  ORG=$2
  VER=$3
  setGlobals $PEER $ORG
  VERSION=${3:-${VER}}

  set -x
  peer chaincode upgrade -o orderer.healthcare.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CONTRACT} -v ${VERSION} -c '{"Args":[]}' -P "AND ('HProviderMSP.peer','ResearchMSP.peer','Org3MSP.peer')"
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode upgrade on peer${PEER}.${org_name[$ORG-1]} has failed"
  echo "===================== Chaincode is upgraded on peer${PEER}.${org_name[$ORG-1]} on channel '$CHANNEL_NAME' ===================== "
  echo
}

if [ "${NO_CHAINCODE}" != "true" -a "${NEW_VERSION_CHAINCODE}" = "1.0" ]; then

  echo "Installing chaincode on peer0.hprivider..."
  installChaincode 0 1 $NEW_VERSION_CHAINCODE

  echo "Installing chaincode on peer1.hprivider..."
  installChaincode 1 1 $NEW_VERSION_CHAINCODE

  echo "Installing chaincode on peer2.hprivider..."
  installChaincode 2 1 $NEW_VERSION_CHAINCODE


  echo "Installing chaincode on peer3.hprivider..."
  installChaincode 3 1 $NEW_VERSION_CHAINCODE


  echo "Installing chaincode on peer4.hprivider..."
  installChaincode 4 1 $NEW_VERSION_CHAINCODE
    
    
  echo "Install chaincode on peer0.research..."
  installChaincode 0 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer1.research..."
  installChaincode 1 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer2.research..."
  installChaincode 2 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer3.research..."
  installChaincode 3 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer3.research..."
  installChaincode 4 2 $NEW_VERSION_CHAINCODE


  echo "Instantiating chaincode on peer0.hprovider..."
  instantiateChaincode 0 1

  echo
  echo "========= All GOOD, New Contract added =========== "
  echo

elif [ "${NO_CHAINCODE}" != "true" -a "${NEW_VERSION_CHAINCODE}" > "1.0" ]; then

  echo "Installing chaincode on peer0.hprivider..."
  installChaincode 0 1 $NEW_VERSION_CHAINCODE

  echo "Installing chaincode on peer1.hprivider..."
  installChaincode 1 1 $NEW_VERSION_CHAINCODE

  echo "Installing chaincode on peer2.hprivider..."
  installChaincode 2 1 $NEW_VERSION_CHAINCODE


  echo "Installing chaincode on peer3.hprivider..."
  installChaincode 3 1 $NEW_VERSION_CHAINCODE


  echo "Installing chaincode on peer4.hprivider..."
  installChaincode 4 1 $NEW_VERSION_CHAINCODE
    
    
  echo "Install chaincode on peer0.research..."
  installChaincode 0 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer1.research..."
  installChaincode 1 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer2.research..."
  installChaincode 2 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer3.research..."
  installChaincode 3 2 $NEW_VERSION_CHAINCODE

  echo "Install chaincode on peer3.research..."
  installChaincode 4 2 $NEW_VERSION_CHAINCODE


  echo "Instantiating chaincode on peer0.hprovider..."
  upgradeChaincode 0 1 $NEW_VERSION_CHAINCODE

  echo
  echo "========= All GOOD, New Contract added =========== "
  echo

else
  echo
  echo "========= ERROR, New Contract Not added =========== "
  echo
fi


echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0