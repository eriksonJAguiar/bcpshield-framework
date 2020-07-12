/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

//dp "github.com/icmc-wines/dp-privacy/diffpriv"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	dp "github.com/eriksonJAguiar/godiffpriv"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the DICOM imaging attributres
type Dicom struct {
	DicomID              string    `json:"dicomID"`
	PatientID            string    `json:"patientID"`
	DocType              string    `json:"docType"`
	PatientFirstname     string    `json:"patientFirstname"`
	PatientLastname      string    `json:"patientLastname"`
	PatientTelephone     string    `json:"patientTelephone"`
	PatientAddress       string    `json:"patientAddress"`
	PatientAge           int       `json:"patientAge"`
	PatientBirth         string    `json:"patientBirth"`
	PatientOrganization  string    `json:"patientOrganization"`
	PatientMothername    string    `json:"patientMothername"`
	PatientReligion      string    `json:"patientReligion"`
	PatientSex           string    `json:"patientSex"`
	PatientGender        string    `json:"patientGender"`
	PatientInsuranceplan string    `json:"patientInsuranceplan"`
	PatientWeigth        float64   `json:"patientWeigth"`
	PatientHeigth        float64   `json:"patientHeigth"`
	MachineModel         string    `json:"machineModel"`
	Timestamp            time.Time `json:"timestamp"`
}

/*
	Type for define the dicom avaliable for sharing
	Atribute isDoctor define when the user want to sharing imaging with
	your personal Doctor
*/
type SharedDicom struct {
	BatchID        string    `json:"batchID"`
	DocType        string    `json:"docType"`
	IpfsReference  string    `json:"ipfsReference"`
	DicomShared    []string  `json:"dicomShared"`
	Holder         string    `json:"holder"`
	HolderAccepted bool      `json:"holderAccepted"`
	Timestamp      time.Time `json:"timestamp"`
	DataAmount     int       `json:"dataAmount"`
	WhoAccessed    string    `json:"whoAccessed"`
	AccessLevel    string    `json:"accessLevel"`
}

// Log type defined when an user access some asset
type Log struct {
	LogID        string    `json:"logID"`
	DocType      string    `json:"docType"`
	AssetToken   string    `json:"assetToken"`
	TypeAsset    string    `json:"typeAsset"`
	HolderAsset  string    `json:"holderAsset"`
	HproviderGet string    `json:"hproviderGet"`
	Timestamp    time.Time `json:"timestamp"`
	WhoAccessed  string    `json:"whoAccessed"`
	AccessLevel  int       `json:"accessLevel"`
}

type Request struct {
	RequestID       string    `json:"requestID"`
	DocType         string    `json:"docType"`
	DataAmount      int       `json:"dataAmount"`
	Timestamp       time.Time `json:"timestamp"`
	HolderRequested string    `json:"holderRequested"`
	UserRequest     string    `json:"userRequest"`
}

// Type for define chaincode
type HealthcareChaincode struct {
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *HealthcareChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Init()", fcn, params)
	return shim.Success(nil)
}

// Patient or HProvider add imaging in blockchain network
// Params: Dicom Type struct
// Returns: True or False for transactions
func (cc *HealthcareChaincode) addAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error

	fmt.Println("Parameter number " + strconv.Itoa(len(args)))
	if len(args) < 17 || len(args) > 17 {
		return shim.Error("Incorrect number of arguments. Expecting 17")
	}

	fmt.Println("- start init Dicom")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	} else if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	} else if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	} else if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	} else if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	} else if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	} else if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	} else if len(args[7]) <= 0 {
		return shim.Error("8th argument must be a non-empty string")
	} else if len(args[8]) <= 0 {
		return shim.Error("9th argument must be a non-empty string")
	} else if len(args[9]) <= 0 {
		return shim.Error("10th argument must be a non-empty string")
	} else if len(args[10]) <= 0 {
		return shim.Error("11th argument must be a non-empty string")
	} else if len(args[11]) <= 0 {
		return shim.Error("12th argument must be a non-empty string")
	} else if len(args[12]) <= 0 {
		return shim.Error("13th argument must be a non-empty string")
	} else if len(args[13]) <= 0 {
		return shim.Error("14th argument must be a non-empty string")
	} else if len(args[14]) <= 0 {
		return shim.Error("15th argument must be a non-empty string")
	} else if len(args[15]) <= 0 {
		return shim.Error("16th argument must be a non-empty string")
	} else if len(args[16]) <= 0 {
		return shim.Error("17th argument must be a non-empty string")
	}

	patientAge, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.ParseFloat(args[14], 64)
	if err != nil {
		return shim.Error("14 argument must be a numeric string")
	}
	patientHeigth, err := strconv.ParseFloat(args[15], 64)
	if err != nil {
		return shim.Error("15 argument must be a numeric string")
	}

	dicomID := args[0]
	patientID := args[1]
	patientFirstname := args[2]
	patientLastname := args[3]
	patientTelephone := args[4]
	patientAddress := args[5]
	patientBirth := args[7]
	patientOrganization := args[8]
	patientMothername := args[9]
	patientReligion := args[10]
	patientSex := args[11]
	patientGender := args[12]
	patientInsuranceplan := args[13]
	machineModel := args[16]

	dicomBytes, err := stub.GetState(dicomID)
	if err != nil {
		return shim.Error("Failed to get pet: " + err.Error())
	} else if dicomBytes != nil {
		return shim.Error("This patient already exists: " + dicomID)
	}

	rec := &Dicom{
		DicomID:              dicomID,
		DocType:              "Dicom",
		PatientID:            patientID,
		PatientFirstname:     patientFirstname,
		PatientLastname:      patientLastname,
		PatientTelephone:     patientTelephone,
		PatientAddress:       patientAddress,
		PatientAge:           patientAge,
		PatientBirth:         patientBirth,
		PatientOrganization:  patientOrganization,
		PatientMothername:    patientMothername,
		PatientReligion:      patientReligion,
		PatientSex:           patientSex,
		PatientGender:        patientGender,
		PatientInsuranceplan: patientInsuranceplan,
		PatientWeigth:        patientWeigth,
		PatientHeigth:        patientHeigth,
		MachineModel:         machineModel,
		Timestamp:            timestamp,
	}

	dicomJSON, err := json.Marshal(rec)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(dicomID, dicomJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- End imaging saved")
	return shim.Success(nil)

}

// Patient or HProvider add imaging in blockchain network using private method HLF
// Params: Dicom Type struct
// Returns: True or False for transactions
func (cc *HealthcareChaincode) addAssetPriv(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error

	fmt.Println("Parameter number " + strconv.Itoa(len(args)))
	if len(args) < 17 || len(args) > 17 {
		return shim.Error("Incorrect number of arguments. Expecting 17")
	}

	fmt.Println("- start init Dicom")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	} else if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	} else if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	} else if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	} else if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	} else if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	} else if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	} else if len(args[7]) <= 0 {
		return shim.Error("8th argument must be a non-empty string")
	} else if len(args[8]) <= 0 {
		return shim.Error("9th argument must be a non-empty string")
	} else if len(args[9]) <= 0 {
		return shim.Error("10th argument must be a non-empty string")
	} else if len(args[10]) <= 0 {
		return shim.Error("11th argument must be a non-empty string")
	} else if len(args[11]) <= 0 {
		return shim.Error("12th argument must be a non-empty string")
	} else if len(args[12]) <= 0 {
		return shim.Error("13th argument must be a non-empty string")
	} else if len(args[13]) <= 0 {
		return shim.Error("14th argument must be a non-empty string")
	} else if len(args[14]) <= 0 {
		return shim.Error("15th argument must be a non-empty string")
	} else if len(args[15]) <= 0 {
		return shim.Error("16th argument must be a non-empty string")
	} else if len(args[16]) <= 0 {
		return shim.Error("17th argument must be a non-empty string")
	}

	patientAge, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.ParseFloat(args[14], 64)
	if err != nil {
		return shim.Error("14 argument must be a numeric string")
	}
	patientHeigth, err := strconv.ParseFloat(args[15], 64)
	if err != nil {
		return shim.Error("15 argument must be a numeric string")
	}

	dicomID := args[0]
	patientID := args[1]
	patientFirstname := args[2]
	patientLastname := args[3]
	patientTelephone := args[4]
	patientAddress := args[5]
	patientBirth := args[7]
	patientOrganization := args[8]
	patientMothername := args[9]
	patientReligion := args[10]
	patientSex := args[11]
	patientGender := args[12]
	patientInsuranceplan := args[13]
	machineModel := args[16]

	dicomBytes, err := stub.GetPrivateData("collectionDicomPrivate", dicomID)
	if err != nil {
		return shim.Error("Failed to get pet: " + err.Error())
	} else if dicomBytes != nil {
		return shim.Error("This patient already exists: " + dicomID)
	}

	rec := &Dicom{
		DicomID:              dicomID,
		DocType:              "Dicom",
		PatientID:            patientID,
		PatientFirstname:     patientFirstname,
		PatientLastname:      patientLastname,
		PatientTelephone:     patientTelephone,
		PatientAddress:       patientAddress,
		PatientAge:           patientAge,
		PatientBirth:         patientBirth,
		PatientOrganization:  patientOrganization,
		PatientMothername:    patientMothername,
		PatientReligion:      patientReligion,
		PatientSex:           patientSex,
		PatientGender:        patientGender,
		PatientInsuranceplan: patientInsuranceplan,
		PatientWeigth:        patientWeigth,
		PatientHeigth:        patientHeigth,
		MachineModel:         machineModel,
		Timestamp:            timestamp,
	}

	dicomJSON, err := json.Marshal(rec)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutPrivateData("collectionDicomPrivate", dicomID, dicomJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- End imaging saved")
	return shim.Success(nil)

}

// get imaging without anyone security
// Params: AssetID to get image
// Returns: Dicom type struct
func (cc *HealthcareChaincode) getAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("We expected one param")
	}

	dicomID := args[0]

	var jsonResp string
	dicomValue, err := stub.GetState(dicomID)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + dicomID + "\"}"
		return shim.Error(jsonResp)
	} else if dicomValue == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + dicomID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(dicomValue)
}

// get imaging using private method HLF
// Params: AssetID to get image
// Returns: Dicom type struct
func (cc *HealthcareChaincode) getAssetPriv(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("We expected one param")
	}

	dicomID := args[0]

	var jsonResp string
	dicomValue, err := stub.GetPrivateData("collectionDicomPrivate", dicomID)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + dicomID + "\"}"
		return shim.Error(jsonResp)
	} else if dicomValue == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + dicomID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(dicomValue)
}

// Patient Sharing imaging with a doctor and add asset for sharing asset in blockchain
// Params: (Four args) batchID, patientID, doctorID, hashIPFS repository and dicomShared represents the ID files to sharing with doctor
// Return: Id hash to represents the request value by a hash
func (cc *HealthcareChaincode) shareAssetWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 5 && len(args) > 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a numeric string")
	} else if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a numeric string")
	} else if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a numeric string")
	} else if len(args[3]) <= 0 {
		return shim.Error("4rd argument must be a numeric string")
	}

	var dicomShared []string

	batchID := args[0]
	holder := args[1]
	doctorID := args[2]
	hashIPFS := args[3]
	dicomShared = append(dicomShared, args[4])

	//Configure IPFS
	ipfsReference := hashIPFS
	getTime := time.Now()

	asset := &SharedDicom{
		BatchID:        batchID,
		DocType:        "SharedDicom",
		IpfsReference:  ipfsReference,
		DicomShared:    dicomShared,
		Holder:         holder,
		HolderAccepted: true,
		Timestamp:      getTime,
		DataAmount:     1,
		WhoAccessed:    doctorID,
		AccessLevel:    "Doctor",
	}

	logJSON, err := json.Marshal(asset)

	if err != nil {
		return shim.Error(err.Error())
	} else if logJSON == nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(batchID, logJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	auxBatch := map[string]interface{}{"id": asset.BatchID}
	idJSON, err := json.Marshal(auxBatch)

	fmt.Println("- End File to sharing done")

	return shim.Success(idJSON)
}

// Doctor can get patient's data shared
// Params: (Two args) batchID for get exams shared
// Returns: Dicom structure type anonymized and IPFS hash value to get real imaging
func (cc *HealthcareChaincode) getSharedAssetWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var batchID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting dicom number of the Record to query")
	}

	batchID = args[0]

	batchValues, err := stub.GetState(batchID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + batchID + "\"}"
		return shim.Error(jsonResp)
	} else if batchValues == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + batchID + "\"}"
		return shim.Error(jsonResp)
	}

	var batch SharedDicom

	err = json.Unmarshal(batchValues, &batch)
	if err != nil {
		shim.Error("Transform batchValues error: " + err.Error())
	}

	dicom := batch.DicomShared[0]

	dicomValue, err := stub.GetState(dicom)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + dicom + "\"}"
		return shim.Error(jsonResp)
	} else if batchValues == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + dicom + "\"}"
		return shim.Error(jsonResp)
	}

	var asset Dicom
	var assets []Dicom
	err = json.Unmarshal(dicomValue, &asset)
	if err != nil {
		return shim.Error("Asset to share error: " + err.Error())
	}
	assets = append(assets, asset)
	fmt.Println("[Log] asset requested: ")
	fmt.Println(assets)
	// Nesse ponto deve ser implementando a comunicação com o IPFS para compartilhar DICOM

	//Create id log image
	// rand.Seed(time.Now().UnixNano())
	// id := strconv.Itoa(rand.Int()) + time.Now().String()
	// hs := sha1.New()
	// hs.Write([]byte(id))
	// hexLogID := hs.Sum(nil)
	// logID := hex.EncodeToString(hexLogID)
	auxid, err := uuid.NewRandom()
	if err != nil {
		return shim.Error("Erro genereta UUID: " + err.Error())
	}
	logID := auxid.String()

	// Deve ser aplicado a K-Anonimity antes de enviar os dados para o médico
	queryString := "{\"selector\":{\"docType\":\"Dicom\"}}"
	allDicomByte, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error("Get values error: " + err.Error())
	}

	var allDicom []Dicom
	err = json.Unmarshal(allDicomByte, &allDicom)
	if err != nil {
		return shim.Error("Get all assets error: " + err.Error())
	}
	fmt.Println("[Log] Query all assets: ")
	fmt.Println(allDicom)

	fmt.Println("[Log] Intializing K-Anonymity process: ")
	anonymizedValues, err := anonimizeKAnonimity(allDicom, assets)
	if err != nil {
		return shim.Error("Anonymized with K-Anonymity error: " + err.Error())
	}

	var anonymizedDicom *interface{}
	err = json.Unmarshal(anonymizedValues, &anonymizedDicom)
	if err != nil {
		return shim.Error("Anaymimzed Dicom with K-anonimity Error: " + err.Error())
	}
	fmt.Println("[Log] Query all assets: ")
	fmt.Println(*anonymizedDicom)

	resultAnonymized, err := json.Marshal(map[string]interface{}{"Dicoms": *anonymizedDicom, "IPFSHash": batch.IpfsReference})
	if err != nil {
		return shim.Error("Convert object to bytes K-Anonymity error: " + err.Error())
	}

	fmt.Println("[Log] Result anoninized success ")

	// Create and store imaging log
	log := Log{
		LogID:        logID,
		DocType:      "Log",
		AssetToken:   dicom,
		HolderAsset:  batch.Holder,
		HproviderGet: "Doctor",
		WhoAccessed:  batch.WhoAccessed,
		AccessLevel:  1,
	}

	if cc.addLog(stub, log) != true {
		jsonResp = "{\"Error\":\" Logs dont record \"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("[Log] Log it was add with success ")

	return shim.Success(resultAnonymized)
}

// Researcher request imaging from patient or other researcher
// Params: (three args) "amount" to represents the images amount requested and "researchID" our on network ID and "PatientID" to represents patient he wants request image
// Returns: requestID to describes the request ID generated
func (cc *HealthcareChaincode) requestAssetForResearcher(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("argument must be a numeric string")
	}

	researchID := args[1]
	patientID := args[2]

	timeGet := time.Now()
	// rand.Seed(time.Now().UnixNano())
	// id := strconv.Itoa(rand.Int()) + timeGet.String()
	// hs := sha1.New()
	// hs.Write([]byte(id))
	// hexReqID := hs.Sum(nil)
	// reqID := hex.EncodeToString(hexReqID)
	auxid, err := uuid.NewRandom()
	if err != nil {
		return shim.Error("Erro genereta UUID: " + err.Error())
	}
	reqID := auxid.String()

	request := &Request{
		RequestID:       reqID,
		DocType:         "Request",
		DataAmount:      amount,
		Timestamp:       timeGet,
		HolderRequested: patientID,
		UserRequest:     researchID,
	}

	requestJSON, err := json.Marshal(request)

	if err != nil {
		return shim.Error("Error convert Request to byte and error: " + err.Error())
	}

	err = stub.PutState(reqID, requestJSON)
	if err != nil {
		return shim.Error("Put request on blockchain Error: " + err.Error())
	}

	idRequest, err := json.Marshal(map[string]interface{}{"requestID": reqID})

	fmt.Println("- Request Sent")
	return shim.Success(idRequest)
}

// Researcher or patients sharing imaging request by others research from researchers requests
// Params: (three args) "holderID" patient ID or researcher ID that will share image with Researcher, "requestID" hash value generated from requestAssetForResearche  and "hashIPFS" hash value IFPS stored on holder database
// Returns:
func (cc *HealthcareChaincode) shareAssetForResearcher(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var jsonResp string
	var err error

	if len(args) != 3 {
		shim.Error("We expected two element as params")
	}

	holderID := args[0]
	requestID := args[1]
	hashIPFS := args[2]

	requestByte, err := stub.GetState(requestID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + requestID + "\"}"
		return shim.Error(jsonResp)
	} else if requestByte == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + requestID + "\"}"
		return shim.Error(jsonResp)
	}

	var requestObj Request
	json.Unmarshal(requestByte, &requestObj)

	//"{\"selector\":{\"owner\":\"tom\"}}"

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Dicom\",\"patientID\":\"%s\"}}", holderID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error("Error get data from queries " + err.Error())
	} else if queryResults == nil {
		return shim.Error("Query result is Null")
	}

	var resultsDicom []Dicom

	err = json.Unmarshal(queryResults, &resultsDicom)

	if err != nil {
		return shim.Error("Unmarshal conveter Json to Object error: " + err.Error())
	}

	i := 0

	var dicomShared []string

	for i < requestObj.DataAmount {
		dicomShared = append(dicomShared, resultsDicom[i].DicomID)
		i++
	}

	// timeGet := time.Now()
	// sd := rand.NewSource(time.Now().UnixNano())
	// rd := rand.New(sd)
	// id := strconv.Itoa(rd.Int()) + timeGet.String()
	// hs := sha1.New()
	// hs.Write([]byte(id))
	// auxSharedObjID := hs.Sum(nil)

	auxid, err := uuid.NewRandom()
	if err != nil {
		return shim.Error("Erro genereta UUID: " + err.Error())
	}

	sharedDicomObj := &SharedDicom{
		BatchID:       auxid.String(),
		DocType:       "SharedDicom",
		IpfsReference: hashIPFS,
		DicomShared:   dicomShared,
		Holder:        holderID,
		Timestamp:     time.Now(),
		DataAmount:    requestObj.DataAmount,
		WhoAccessed:   requestObj.UserRequest,
		AccessLevel:   "Research",
	}

	sharedDicomByte, err := json.Marshal(sharedDicomObj)

	if err != nil {
		return shim.Error("Tranform sharedDicom Json to Bytes Error: " + err.Error())
	}

	err = stub.PutState(sharedDicomObj.BatchID, sharedDicomByte)
	if err != nil {
		return shim.Error("Put SharedDicm Element error " + err.Error())
	}

	return shim.Success(sharedDicomByte)

}

// Researcher get imaging shared
// Params: (two args) "sharedDicomID" describes hash value requested ID; and epsilon to apply defferential privacy
// Returns: Dicom structure type anonymized and Hash IPFS to get images from IFPS structure
func (cc *HealthcareChaincode) getSharedAssetForResearcher(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var jsonResp string
	var err error

	if len(args) != 2 {
		shim.Error("We expected one element as params")
	}

	sharedDicomID := args[0]
	epsilon, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("Faild Convert epsilon value: " + err.Error())
	}

	SharedDicomBytes, err := stub.GetState(sharedDicomID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + sharedDicomID + "\"}"
		return shim.Error(jsonResp)
	} else if len(SharedDicomBytes) == 0 {
		jsonResp = "{\"Error\":\"Record does not exist: " + sharedDicomID + "\"}"
		return shim.Error(jsonResp)
	}

	var sharedDicomValues SharedDicom

	err = json.Unmarshal(SharedDicomBytes, &sharedDicomValues)
	if err != nil {
		return shim.Error("[GET ASSET FOR RESEARCH] Transform sharedDicom to Json error: " + err.Error())
	}

	var dicoms []Dicom

	if len(sharedDicomValues.DicomShared) == 1 {
		sharedID := sharedDicomValues.DicomShared[0]
		dicomValue, err := stub.GetState(sharedID)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + sharedID + "\"}"
			return shim.Error(jsonResp)
		}
		var dcm Dicom
		err = json.Unmarshal(dicomValue, &dcm)
		if err != nil {
			return shim.Error("Convert DCM error: " + err.Error())
		}
		dicoms = append(dicoms, dcm)
	} else {

		selectorString := ""
		for i, sharedID := range sharedDicomValues.DicomShared {
			selectorString += fmt.Sprintf("\"%s\"", sharedID)
			if i+1 <= len(sharedDicomValues.DicomShared)-1 {
				selectorString += ", "
			}
		}

		queryString := fmt.Sprintf("{\"selector\":{ \"_id\": {\"$or\": [%s]}}}", selectorString)
		dicomValue, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for \"}"
			return shim.Error(jsonResp)
		}

		err = json.Unmarshal(dicomValue, &dicoms)
		if err != nil {
			return shim.Error("Convert DCM error: " + err.Error())
		}
	}

	anonimizedDicoms, err := anonimizeDiffPriv(stub, dicoms, sharedDicomValues.DataAmount, epsilon)

	if err != nil {
		return shim.Error("Anonymized records error: " + err.Error())
	} else if anonimizedDicoms == nil {
		return shim.Error("Error anonymized records don't get")
	}

	resultAnonymized, err := json.Marshal(map[string]interface{}{"Dicom": anonimizedDicoms, "IPFSHash": sharedDicomValues.IpfsReference})

	//anonimizedBytes, err := json.Marshal(resultAnonymized)

	if err != nil {
		return shim.Error("Marshal convert erro: " + err.Error())
	}

	//Recovery accessed imaging logs
	// rand.Seed(time.Now().UnixNano())
	// newID := strconv.Itoa(rand.Intn(10000000)) + time.Now().String() + sharedDicomValues.BatchID
	// hs := sha1.New()
	// hs.Write([]byte(newID))
	// hexLogID := hs.Sum(nil)
	// logID := hex.EncodeToString(hexLogID)

	auxid, err := uuid.NewRandom()
	if err != nil {
		return shim.Error("Erro genereta UUID:" + err.Error())
	}
	logID := auxid.String()

	log := Log{
		LogID:        logID,
		DocType:      "Log",
		AssetToken:   sharedDicomValues.BatchID,
		HolderAsset:  sharedDicomValues.Holder,
		HproviderGet: "Research",
		WhoAccessed:  sharedDicomValues.WhoAccessed,
		AccessLevel:  2,
	}

	if cc.addLog(stub, log) != true {
		jsonResp = "{\"Error\":\" Logs dont record \"}"
		return shim.Error(jsonResp)
	}

	//sharedDicomValues.IpfsReference

	return shim.Success(resultAnonymized)
}

// Network admin query access logs as from tokenID into imaging
// Params: (one arg) "logID" represents token value registred in Dicom image
// Returns: Log json
func (cc *HealthcareChaincode) auditLogs(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("We expected just one param")
	}

	var jsonResp string
	logID := args[0]

	logValue, err := stub.GetState(logID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + logID + "\"}"
		return shim.Error(jsonResp)
	} else if logValue == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + logID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(logValue)
}

// Store logs in blockain
// Internal function
func (cc *HealthcareChaincode) addLog(stub shim.ChaincodeStubInterface, log Log) bool {

	logJSON, err := json.Marshal(log)

	if err != nil {
		return false
	}

	err = stub.PutState(log.LogID, logJSON)
	if err != nil {
		return false
	}

	fmt.Println("- End Log Saved")
	return true

}

// Internal function
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	// fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))

		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	// fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

//Apply differential privacy on data
// Internal function
func anonimizeDiffPriv(stub shim.ChaincodeStubInterface, assets []Dicom, amount int, epsilon float64) (map[string]interface{}, error) {

	queryString := "{ \"selector\":{ \"docType\":\"Dicom\" }}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return nil, err
	}

	// Database total
	var allDicom []Dicom

	err = json.Unmarshal(queryResults, &allDicom)

	if err != nil {
		return nil, err
	}

	var patientOrganization, patientRace []string
	var patientSex, patientGender, patientAddress []string
	var dicomID []int
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		id, err := strconv.Atoi(dcm.DicomID)
		if err != nil {
			return nil, err
		}
		dicomID = append(dicomID, id)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientRace = append(patientRace, dcm.PatientReligion)
		patientSex = append(patientSex, dcm.PatientSex)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.PatientGender)
	}
	var dicomWithNoise map[string]interface{}
	var orgNoise, raceNoise, sexNoise, ageNoise, weigthNoise, heigthNoise map[string]float64

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	dicomWithNoise["DicomID"] = id.String()

	privOrg := dp.PrivateDataFactory(patientOrganization)
	auxPrivOrg, err := privOrg.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivOrg, &orgNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Org"] = orgNoise

	privRace := dp.PrivateDataFactory(patientRace)
	auxPrivRace, err := privRace.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivRace, &raceNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Race"] = raceNoise

	privSex := dp.PrivateDataFactory(patientSex)
	auxPrivSex, err := privSex.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivSex, &sexNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Sex"] = sexNoise

	privAge := dp.PrivateDataFactory(patientAge)
	auxPrivAge, err := privAge.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivAge, &ageNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Age"] = ageNoise

	privWeigth := dp.PrivateDataFactory(patientWeigth)
	auxPrivWeigth, err := privWeigth.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivWeigth, &weigthNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Weigth"] = weigthNoise

	privHeigth := dp.PrivateDataFactory(patientHeigth)
	auxPrivHeigth, err := privHeigth.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivHeigth, &heigthNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Heigth"] = heigthNoise

	return dicomWithNoise, nil
}

// Apply K-anonymity privacy on data
// Internal function
func anonimizeKAnonimity(allDicom []Dicom, assets []Dicom) ([]byte, error) {

	fmt.Println("[Log] Assets that will be anonimized")
	fmt.Println(assets)

	for _, ast := range assets {
		allDicom = append(allDicom, ast)
	}

	var patientFirstnames, patientLastnames, patientTelephones, dicomID, patientID []string
	var patientBirth, patientOrganization, patientMothername, patientReligion []string
	var patientInsuranceplan, machineModel []string
	var patientSex, patientGender, patientAddress []string
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		dicomID = append(dicomID, dcm.DicomID)
		patientID = append(patientID, dcm.PatientID)
		patientFirstnames = append(patientFirstnames, dcm.PatientFirstname)
		patientLastnames = append(patientLastnames, dcm.PatientLastname)
		patientTelephones = append(patientTelephones, dcm.PatientTelephone)
		patientBirth = append(patientBirth, dcm.PatientBirth)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientMothername = append(patientMothername, dcm.PatientMothername)
		patientReligion = append(patientReligion, dcm.PatientReligion)
		patientSex = append(patientSex, dcm.PatientSex)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientInsuranceplan = append(patientInsuranceplan, dcm.PatientInsuranceplan)
		machineModel = append(machineModel, dcm.MachineModel)
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.PatientGender)
	}

	anonymizedDicomID := KAnonymitygeneralizationSymbolic(dicomID)
	anonymizedPatientID := KAnonymitygeneralizationSymbolic(patientID)
	anonymizedFirstName := KAnonymitygeneralizationSymbolic(patientFirstnames)
	anonymizedLastName := KAnonymitygeneralizationSymbolic(patientLastnames)
	anonymizedTelephone := KAnonymitygeneralizationSymbolic(patientTelephones)
	anonymizedBirth := KAnonymitygeneralizationSymbolic(patientBirth)
	anonymizedOrganization := KAnonymitygeneralizationSymbolic(patientOrganization)
	anonymizedMothername := KAnonymitygeneralizationSymbolic(patientMothername)
	anonymizedReligion := KAnonymitygeneralizationSymbolic(patientReligion)
	anonymizedSex := KAnonymitygeneralizationSymbolic(patientSex)
	anonymizedAddress := KAnonymitygeneralizationSymbolic(patientAddress)
	anonymizedGender := KAnonymitygeneralizationSymbolic(patientGender)
	anonymizedAge := KAnonymityGeneralizationNumeric(patientAge)
	anonymizedWeigth := KAnonymityGeneralizationNumeric(patientWeigth)
	anonymizedHeigth := KAnonymityGeneralizationNumeric(patientHeigth)
	anonymizedInsuranceplan := KAnonymitygeneralizationSymbolic(patientInsuranceplan)
	anonimizedModelMachine := KAnonymitygeneralizationSymbolic(machineModel)

	amount := len(assets)
	type documentValue struct {
		DicomID              string `json:"dicomID"`
		PatientID            string `json:"patientID"`
		PatientFirstname     string `json:"patientFirstname"`
		PatientLastname      string `json:"patientLastname"`
		PatientTelephone     string `json:"patientTelephone"`
		PatientAddress       string `json:"patientAddress"`
		PatientAge           string `json:"patientAge"`
		PatientBirth         string `json:"patientBirth"`
		PatientOrganization  string `json:"patientOrganization"`
		PatientMothername    string `json:"patientMothername"`
		PatientReligion      string `json:"patientReligion"`
		PatientSex           string `json:"patientSex"`
		PatientGender        string `json:"patientGender"`
		PatientInsuranceplan string `json:"patientInsuranceplan"`
		PatientWeigth        string `json:"patientWeigth"`
		PatientHeigth        string `json:"patientHeigth"`
		MachineModel         string `json:"machineModel"`
		Timestamp            string `json:"timestamp"`
	}

	var data []documentValue

	fmt.Println("[Log] amount assets to anonimized")

	for i := len(allDicom) - 1; i >= (len(allDicom) - amount); i-- {
		var dicomNew documentValue

		dicomNew.DicomID = anonymizedDicomID[i]
		dicomNew.PatientID = anonymizedPatientID[i]
		dicomNew.PatientFirstname = anonymizedFirstName[i]
		dicomNew.PatientLastname = anonymizedLastName[i]
		dicomNew.PatientTelephone = anonymizedTelephone[i]
		dicomNew.PatientAddress = anonymizedAddress[i]
		dicomNew.PatientAge = anonymizedAge[i]
		dicomNew.PatientBirth = anonymizedBirth[i]
		dicomNew.PatientOrganization = anonymizedOrganization[i]
		dicomNew.PatientReligion = anonymizedReligion[i]
		dicomNew.PatientMothername = anonymizedMothername[i]
		dicomNew.PatientSex = anonymizedSex[i]
		dicomNew.PatientGender = anonymizedGender[i]
		dicomNew.PatientInsuranceplan = anonymizedInsuranceplan[i]
		dicomNew.PatientWeigth = anonymizedWeigth[i]
		dicomNew.PatientHeigth = anonymizedHeigth[i]
		dicomNew.MachineModel = anonimizedModelMachine[i]
		dicomNew.Timestamp = time.Now().String()

		data = append(data, dicomNew)
	}

	fmt.Println("[Log] Dataset Anonimized whitin K-Anonymnity function")
	fmt.Println(data)

	dataValue, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return dataValue, nil

}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *HealthcareChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is runing " + fun)

	if fun == "addAsset" {
		return cc.addAsset(stub, args)
	} else if fun == "addAssetPriv" {
		return cc.addAssetPriv(stub, args)
	} else if fun == "getAsset" {
		return cc.getAsset(stub, args)
	} else if fun == "getAssetPriv" {
		return cc.getAssetPriv(stub, args)
	} else if fun == "shareAssetWithDoctor" {
		return cc.shareAssetWithDoctor(stub, args)
	} else if fun == "getSharedAssetWithDoctor" {
		return cc.getSharedAssetWithDoctor(stub, args)
	} else if fun == "shareAssetForResearcher" {
		return cc.shareAssetForResearcher(stub, args)
	} else if fun == "requestAssetForResearcher" {
		return cc.requestAssetForResearcher(stub, args)
	} else if fun == "getSharedAssetForResearcher" {
		return cc.getSharedAssetForResearcher(stub, args)
	} else if fun == "auditLogs" {
		return cc.auditLogs(stub, args)
	}

	fmt.Println("invoke did not find func: " + fun) //error
	return shim.Error("Received unknown function invocation")
}
