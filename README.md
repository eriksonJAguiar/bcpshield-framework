# Privacy blockchain architecture - BCShield

It's an architecture blockchain-based to enhancing privacy on healthcare systems for share data between researcher, besides we are using differential privacy and K-anonymity to privacy preserving.

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

## How is intergrate with Hyperledger Blockchain?


## How to use ?

Define de basic Commands

## Team 

* [Erikson Júlio de Aguiar](https://eriksonjaguiar.github.io/)
* Jó Ueyama

## Note

This project was funded by São Paulo Research Foundation - FAPESP. However, all opinion, hypothesis, conclusions, or recommendations contains on this repository are responsibility the author and do not reflect FAPESP's vision. 

