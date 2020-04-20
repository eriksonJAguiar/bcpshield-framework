#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORGMSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ./connections/ccp-template.json 
}

function yaml_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORGMSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ./connections/ccp-template.yaml | sed -e $'s/\\\\n/\\\n        /g'
}

ORG=hprovider
ORGMSP=HProvider
P0PORT=7051
P1PORT=8051
CAPORT=7054
PEERPEM=crypto-config/peerOrganizations/hprovider.healthcare.com/tlsca/tlsca.hprovider.healthcare.com-cert.pem
CAPEM=crypto-config/peerOrganizations/hprovider.healthcare.com/ca/ca.hprovider.healthcare.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > ./connections/connection-hprovider.json
echo "$(yaml_ccp $ORG $ORGMSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > ./connections/connection-hprovider.yaml

ORG=research
ORGMSP=Research
P0PORT=9051
P1PORT=10051
CAPORT=8054
PEERPEM=crypto-config/peerOrganizations/research.healthcare.com/tlsca/tlsca.research.healthcare.com-cert.pem
CAPEM=crypto-config/peerOrganizations/research.healthcare.com/ca/ca.research.healthcare.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > ./connections/connection-research.json
echo "$(yaml_ccp $ORG $ORGMSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > ./connections/connection-research.yaml
