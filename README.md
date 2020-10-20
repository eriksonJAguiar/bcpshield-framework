# Privacy blockchain architecture - BCPShield

The **B**lock**C**hain to **P**rivacy-preserving in **SH**aring sens**I**tive h**E**a**L**th **D**ata (BCPShield) is an architecture blockchain-based to enhancing privacy on healthcare systems to share data between users. Besides, we are using differential privacy and K-anonymity to privacy-preserving, and the Hyperledger Fabric to build a blockchain network.

![BCPShield Overview](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/architecture-high-level.png)

The modules are:

- **Healthcare provider:** to store and exchange data, also comunicate with blockchain and privacy module.
- **Privacy:** Apply privacy on data.
- **Blockchain:** Blockchain network based on Hyperledger Fabric 1.4.4, and Off-chain network through InterPlanetary File System (IPFS)
- **User inferface:** It is a module make by REST API to establish user, blockchain, Healthcare provider interaction. On the other hand, an interface that applications can connect and uses the architecture.

The user are: patient, doctor, and researcher.

In short, the architecture can enhance data sharing to research and contribute to reliability increase due to joining the blockchain and privacy mechanism.

## How it works?

### Objects

#### __Dicom__

> For our proof of concept we use a Dicom object, which will be transferred between researchers. Dicom object is a  representation blockchain imaging within blockchain database.

>  **Dicom atributes**:

>  * DicomID, PatientID, DocType, PatientFirstname, PatientLastname, PatientTelephone, PatientAddress, PatientAge,       PatientBirth, PatientOrganization, PatientMothername, PatientReligion, PatientSex, PatientGender, PatientInsuranceplan, PatientWeigth, PatientHeigth, MachineModel, Timestamp.
  
 
#### __Log__

>  For our proof of concept we use a Log object, which represents access log on any asset.

> **Log attributes:**
>	* LogID, DocType, AssetToken, TypeAsset, HolderAsset, HproviderGet, Timestamp, WhoAccessed, AccessLevel.


#### __Request__

 > For our proof of concept we use a request object that have aim to request an imaging from anyone patient or researcher.

 > **Request attributes:**
 >	* RequestID, DocType, DataAmount, Timestamp, HolderRequested, UserRequest.

#### __Shared Dicom__

  > For our proof of concept we use a Shared Dicom object that have aim to represents a metadata within blockchain for shared   asset.
  
 > **Shared Dicom attributes:**
 >	* BatchID, DocType, IpfsReference, DicomShared, Holder, HolderAccepted, Timestamp, DataAmount, WhoAccessed.

### Main functions:

  #### addAsset
   * **Objective:** Insert dicom representation wihthin blockchain.
   * **Parameter:** Dicom object.
   * **Return:** bool.
 

  #### getAsset
   * **Objective:** get Dicom object on blockchain.
   * **Parameter:** DicomID.
   * **Return:** Dicom object as Json.
 
  #### shareAssetWithDoctor
   * **Objective:** For patient share data with a doctor.
   * **Parameter:** Patient ID, DoctorID, hashIPFS repository and DicomID files it want to share.
   * **Return:** string hash for doctor to access data.
 
  #### getSharedAssetWithDoctor
   * **Objective:** For doctor gets patient's imaging. 
   * **Parameter:** string hash shared by the patient.
   * **Return:** Dicom object anonymized with K-anonymity and IPFS Hash.

#### requestAssetForResearcher
 * **Objective:** For researcher to request an amount imaging for your research.
 * **Parameter:** Amount data, researchID, and "PatientID.
 * **Return:** RequestID.
 
 #### shareAssetForResearcher
   * **Objective:** For patients or other researchers sharing your data (Observer that when request is activated this  function to trigger).
   * **Parameter:** holderID, requestID and IPFS Hash
   * **Return:** string hash for researcher to access data.
 
#### getSharedAssetForResearcher
 * **Objective:** For researcher to get data requested.
 * **Parameter:** string hash shared by the patient.
 * **Return:** Dicom object anonymized with Differential Privacy and IPFS Hash.
 
#### auditLogs
 * **Objective:** For network's administrator to audit logs when assets leakage.
 * **Parameter:** token Hash inserted within metadata imaging.
 * **Return:** Log object as Json.


## How to use ?

The next steps are about blockchian module: <br \>

- **Step 0:**

a) install docker engine and docker-compose <br \>
b) install golang <br \>
c) download hyperledger fabric binaries <br \>
d) node js version = 10.x and npm version >= 6.x <br \>

- **Step 1:**

a) You should use the folder [hyperledger-network](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/hyperledger-network) <br />
b) In folder run the script  **./byfn.sh generate** to build credentials and keys for organization on Hyperledger Fabric <br />

- **Step 2:**

a) To make a smart contract you might use an extension IBM blockchain platform on VSCode <br />
b) With a smart contract made you should move the contract folder to folder [chaincode](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/hyperledger-network/chaincode) <br />
c) In file [scripts/script.sh](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/scripts/script.sh) change variable **CC_SRC_PATH** to path your smart contract (For instance, CC_SRC_PATH="github.com/chaincode/imaging-contract")

- **Step 3:**

a) Before build a blockchain network, if you intent to make a different network you should change this files: <br />

1. [crypto-config.yaml](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/crypto-config.yaml) 
2. [configtx.yaml](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/configtx.yaml) 
3. [docker-compose-ca.yaml](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/docker-compose-ca.yaml) 
4. [docker-compose-couch.yaml](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/docker-compose-couch.yaml) 
5. [docker-compose-cli.yaml](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/docker-compose-cli.yaml) 
6. [ccp-generate.sh](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/ccp-generate.sh)
7. [utils.sh](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/scripts/utils.sh)
8. [script.sh)](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/scripts/script.sh)
9. [byfn.sh](https://github.com/eriksonJAguiar/bcshield-architecture/blob/master/hyperledger-network/byfn.sh)
<br />
b) To build our network you should run the command **byfn.sh up -s couchdb -a true -l golang -i 1.4.4** <br />
c) The blokchain network is done <br />
d) The API to interact with a blokchain can modify in the file [api-dicom/server/server.js](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/api-dicom/server) <br />
e) Run the script [api-dicom/server/server.js](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/api-dicom/server/server.js) to make available the blockchain API for that communicate with other components. You might execute this command **node server.js**

- **Step 4:**

a) To build a Off-chain network using the IPFS framework we should following this tutorial [IPFS Tutorial](https://medium.com/@s_van_laar/deploy-a-private-ipfs-network-on-ubuntu-in-5-steps-5aad95f7261b) <br \>
b) The configuration has been made in a Linux machine from Google Cloud, but we can use any Linux machine. This machine was a master node on IPFS network <br \>
c) Run the script [ipfs-cli-main.go](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/ipfs-client) using the command **go run ipfs-cli-main.go**. P.s. You should change the IPFS IP <br \>

- **Step 5:** 

a) In Healthcare providers machines you may run components for observer, local database, and privacy methods <br \>
b) First, each machine should have installed Docker, golang, and python 3  <br \>
c) Then, we can run the command to build Docker container for observer. In folder [bcshield-hprovider](https://github.com/eriksonJAguiar/bcshield-architecture/tree/master/bcshield-hprovider) run the script **.\init.sh**  <br \>
d) This script will be build a docker containers for obsever, MongoDB local, and privacy methods (K-anonymity and Differencial privacy)

## Team 

* [Erikson Júlio de Aguiar](https://eriksonjaguiar.github.io/)
* Jó Ueyama

## Note

This project was funded by São Paulo Research Foundation - FAPESP. However, all opinions, hypothesis, conclusions, or recommendations contains on this repository are responsibility the author and do not reflect FAPESP's vision. 

