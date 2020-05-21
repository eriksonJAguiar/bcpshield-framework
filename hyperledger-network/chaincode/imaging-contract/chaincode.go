/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

//dp "github.com/icmc-wines/dp-privacy/diffpriv"

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

// Patinet or research add imaging in blockchain network
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

	// hs := sha1.New()

	// auxPatientID := time.Now().String() + patientFirstname + patientLastname

	// hs.Write([]byte(auxPatientID))

	// hexPatientID := hs.Sum(nil)

	// patientID := hex.EncodeToString(hexPatientID)

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

// get imaging
// One parameter AssetID
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

// Patient Sharing imaging with a doctor
// Add struct for sharing asset in blockchain
// Three args Patient ID, DoctorID, hashIPFS repository and ID files core for sharing Several assets ID for sharing
func (cc *HealthcareChaincode) shareAssetWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 4 && len(args) > 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a numeric string")
	} else if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a numeric string")
	} else if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a numeric string")
	}

	var dicomShared []string

	holder := args[0]
	doctorID := args[1]
	hashIPFS := args[2]
	dicomShared = append(dicomShared, args[3])

	//Configure IPFS
	ipfsReference := hashIPFS
	getTime := time.Now()

	hs := sha1.New()

	concatValues := holder + doctorID + ipfsReference + holder + getTime.String()

	hs.Write([]byte(concatValues))

	hexBatchID := hs.Sum(nil)

	batchID := hex.EncodeToString(hexBatchID)

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
// One args batch ID for get exams shared
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
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + time.Now().String()
	hs := sha1.New()
	hs.Write([]byte(id))
	hexLogID := hs.Sum(nil)
	logID := hex.EncodeToString(hexLogID)

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
// Two attribute "amount images" and "researchID" and "PatientID"
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
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + timeGet.String()
	hs := sha1.New()
	hs.Write([]byte(id))
	hexReqID := hs.Sum(nil)
	reqID := hex.EncodeToString(hexReqID)

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

// Researcher or patients sharing imaging request by others research
//Two attributes holderID, requestID and hashIPFS
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

	timeGet := time.Now()
	sd := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(sd)
	id := strconv.Itoa(rd.Int()) + timeGet.String()
	hs := sha1.New()
	hs.Write([]byte(id))
	auxSharedObjID := hs.Sum(nil)

	sharedDicomObj := &SharedDicom{
		BatchID:       hex.EncodeToString(auxSharedObjID),
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
// Sharing request values with the research
// One attribute batchID
func (cc *HealthcareChaincode) getSharedAssetForResearcher(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var jsonResp string
	var err error

	if len(args) != 1 {
		shim.Error("We expected two element as params")
	}

	sharedDicomID := args[0]
	epsilon := 0.1

	SharedDicomBytes, err := stub.GetState(sharedDicomID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + sharedDicomID + "\"}"
		return shim.Error(jsonResp)
	} else if SharedDicomBytes == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + sharedDicomID + "\"}"
		return shim.Error(jsonResp)
	}

	var sharedDicomValues SharedDicom

	err = json.Unmarshal(SharedDicomBytes, &sharedDicomValues)
	if err != nil {
		return shim.Error("[GET ASSET FOR RESEARCH] Transform sharedDicom to Json error: " + err.Error())
	}

	//Recovery accessed imaging logs
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + time.Now().String() + sharedDicomValues.BatchID
	hs := sha1.New()
	hs.Write([]byte(id))
	hexLogID := hs.Sum(nil)
	logID := hex.EncodeToString(hexLogID)

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
// One param Log ID
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

func (cc *HealthcareChaincode) queryGeneral(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	query := args[0]

	queryRes, err := getQueryResultForQueryString(stub, query)

	if err != nil {
		return shim.Error("Query error: " + err.Error())
	}

	var alldcm []Dicom

	err = json.Unmarshal(queryRes, &alldcm)

	if err != nil {
		shim.Error("Error Unmarshal Dicom to Json: " + err.Error())
	}

	var patientID []int

	for _, dcm := range alldcm {
		id, err := strconv.Atoi(dcm.PatientID)
		if err != nil {
			return shim.Error("Error convert value " + err.Error())
		}
		patientID = append(patientID, id)
	}

	databasePatientID := TransforIntData(patientID)
	if len(databasePatientID) == 0 {
		return shim.Error("Error get Patient ID")
	}

	queryPatientID := Query(databasePatientID, 2)
	if len(queryPatientID) == 0 {
		return shim.Error("Error Dicom ID")
	}

	epsilon := 0.1
	var noise []Matrix
	fmt.Println("Patient Apply Priv")
	valPatientID := DiffPriv(queryPatientID, 2, databasePatientID, epsilon)
	err = json.Unmarshal(valPatientID, &noise)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(noise) == 0 {
		return shim.Error("Error Noise value is Null")
	}
	fmt.Println("Dicom noise result")
	fmt.Println(noise)

	value, err := json.Marshal(noise)
	if err != nil {
		return shim.Error("Error Marshal" + err.Error())
	}

	return shim.Success(value)

}

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

func anonimizeDiffPriv(stub shim.ChaincodeStubInterface, assets []Dicom, amount int, epsilon float64) ([]Dicom, error) {

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

	var patientFirstnames, patientLastnames, patientTelephones []string
	var patientBirth, patientOrganization, patientMothername, patientReligion []string
	var patientSex, patientGender, patientAddress []string
	var patientAge, dicomID, patientID []int
	var patientWeigth, patientHeigth []float64

	for _, dcm := range allDicom {
		id, err := strconv.Atoi(dcm.DicomID)
		if err != nil {
			return nil, err
		}
		dicomID = append(dicomID, id)
		pid, err := strconv.Atoi(dcm.PatientID)
		if err != nil {
			return nil, err
		}
		patientID = append(patientID, pid)
		patientFirstnames = append(patientFirstnames, dcm.PatientFirstname)
		patientLastnames = append(patientLastnames, dcm.PatientLastname)
		patientTelephones = append(patientTelephones, dcm.PatientTelephone)
		patientBirth = append(patientBirth, dcm.PatientBirth)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientMothername = append(patientMothername, dcm.PatientMothername)
		patientReligion = append(patientReligion, dcm.PatientReligion)
		patientSex = append(patientSex, dcm.PatientSex)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientAge = append(patientAge, dcm.PatientAge)
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.PatientGender)
	}

	databaseDicomID := TransforIntData(dicomID)
	databasePatientID := TransforIntData(patientID)
	databaseFirstName := TransforSymbolicData(patientFirstnames)
	databaseLastName := TransforSymbolicData(patientLastnames)
	databaseTelephone := TransforSymbolicData(patientTelephones)
	databaseBirth := TransforSymbolicData(patientBirth)
	databaseOrganization := TransforSymbolicData(patientOrganization)
	databaseMothername := TransforSymbolicData(patientMothername)
	databaseReligion := TransforSymbolicData(patientReligion)
	databaseSex := TransforSymbolicData(patientSex)
	databaseAddress := TransforSymbolicData(patientAddress)
	databaseGender := TransforSymbolicData(patientGender)
	databaseAge := TransforIntData(patientAge)
	databaseWeigth := TransforFloatData(patientWeigth)
	databaseHeigth := TransforFloatData(patientHeigth)

	var patientFirstnamesQ, patientLastnamesQ, patientTelephonesQ []string
	var patientBirthQ, patientOrganizationQ, patientMothernameQ, patientReligionQ []string
	var patientSexQ, patientGenderQ, patientAddressQ []string
	var patientAgeQ, dicomIDQ, patientIDQ []int
	var patientWeigthQ, patientHeigthQ []float64

	for _, dcm := range assets {
		id, err := strconv.Atoi(dcm.DicomID)
		if err != nil {
			return nil, err
		}
		dicomIDQ = append(dicomIDQ, id)
		pid, err := strconv.Atoi(dcm.PatientID)
		if err != nil {
			return nil, err
		}
		patientIDQ = append(patientIDQ, pid)
		patientFirstnamesQ = append(patientFirstnamesQ, dcm.PatientFirstname)
		patientLastnamesQ = append(patientLastnamesQ, dcm.PatientLastname)
		patientTelephonesQ = append(patientTelephonesQ, dcm.PatientTelephone)
		patientBirthQ = append(patientBirthQ, dcm.PatientBirth)
		patientOrganizationQ = append(patientOrganizationQ, dcm.PatientOrganization)
		patientMothernameQ = append(patientMothernameQ, dcm.PatientMothername)
		patientReligionQ = append(patientReligionQ, dcm.PatientReligion)
		patientSexQ = append(patientSexQ, dcm.PatientSex)
		patientAddressQ = append(patientAddressQ, dcm.PatientAddress)
		patientAgeQ = append(patientAgeQ, dcm.PatientAge)
		patientWeigthQ = append(patientWeigthQ, dcm.PatientWeigth)
		patientHeigthQ = append(patientHeigthQ, dcm.PatientHeigth)
		patientGenderQ = append(patientGenderQ, dcm.PatientGender)
	}

	queryDicomID := TransforIntData(dicomIDQ)
	queryPatientID := TransforIntData(patientIDQ)
	queryFirstName := TransforSymbolicData(patientFirstnamesQ)
	queryLastName := TransforSymbolicData(patientLastnamesQ)
	queryTelephone := TransforSymbolicData(patientTelephonesQ)
	queryBirth := TransforSymbolicData(patientBirthQ)
	queryOrganization := TransforSymbolicData(patientOrganizationQ)
	queryMothername := TransforSymbolicData(patientMothernameQ)
	queryReligion := TransforSymbolicData(patientReligionQ)
	querySex := TransforSymbolicData(patientSexQ)
	queryAddress := TransforSymbolicData(patientAddressQ)
	queryGender := TransforSymbolicData(patientGenderQ)
	queryAge := TransforIntData(patientAgeQ)
	queryWeigth := TransforFloatData(patientWeigthQ)
	queryHeigth := TransforFloatData(patientHeigthQ)

	fmt.Println("Query PatientID values")
	fmt.Println(queryDicomID)

	fmt.Println("Query FirtName values")
	fmt.Println(queryFirstName)

	fmt.Println("Start Apply privacy")

	var noisedQueries []interface{}
	var noiseDicomID, noisePatientID, noiseFirstName, noiseLastName, noiseTelephone []Matrix
	var noiseBirth, noiseOrganization, noiseMothername, noiseReligion []Matrix
	var noiseSex, noiseAddress, noiseGender, noiseAge, noiseWeigth, noiseHeigth []Matrix
	var errNoise error

	fmt.Println("Dicom ID")
	valDicomID := DiffPriv(queryDicomID, amount, databaseDicomID, epsilon)
	errNoise = json.Unmarshal(valDicomID, &noiseDicomID)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseDicomID)

	fmt.Println("Patient Noise")
	valPatientID := DiffPriv(queryPatientID, amount, databasePatientID, epsilon)
	errNoise = json.Unmarshal(valPatientID, &noisePatientID)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noisePatientID)

	fmt.Println("First name Noise")
	valFirtName := DiffPriv(queryFirstName, amount, databaseFirstName, epsilon)
	errNoise = json.Unmarshal(valFirtName, &noiseFirstName)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseFirstName)

	fmt.Println("Last Name noise")
	valLastName := DiffPriv(queryLastName, amount, databaseLastName, epsilon)
	errNoise = json.Unmarshal(valLastName, &noiseLastName)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseLastName)

	fmt.Println("Telephone noise")
	valTelephone := DiffPriv(queryTelephone, amount, databaseTelephone, epsilon)
	errNoise = json.Unmarshal(valTelephone, &noiseTelephone)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseTelephone)

	fmt.Println("Birth noise")
	valBirth := DiffPriv(queryBirth, amount, databaseBirth, epsilon)
	errNoise = json.Unmarshal(valBirth, &noiseBirth)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseBirth)

	fmt.Println("Organization noise")
	valOrganization := DiffPriv(queryOrganization, amount, databaseOrganization, epsilon)
	errNoise = json.Unmarshal(valOrganization, &noiseOrganization)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseOrganization)

	fmt.Println("Mothername noise")
	valMothername := DiffPriv(queryMothername, amount, databaseMothername, epsilon)
	errNoise = json.Unmarshal(valMothername, &noiseMothername)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseMothername)

	fmt.Println("Religion noise")
	valReligion := DiffPriv(queryReligion, amount, databaseReligion, epsilon)
	errNoise = json.Unmarshal(valReligion, &noiseReligion)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseReligion)

	fmt.Println("Sex noise")
	valSex := DiffPriv(querySex, amount, databaseSex, epsilon)
	errNoise = json.Unmarshal(valSex, &noiseSex)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseSex)

	fmt.Println("Noise Address")
	valAddress := DiffPriv(queryAddress, amount, databaseAddress, epsilon)
	errNoise = json.Unmarshal(valAddress, &noiseAddress)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseAddress)

	fmt.Println("Noise gender")
	valGender := DiffPriv(queryGender, amount, databaseGender, epsilon)
	errNoise = json.Unmarshal(valGender, &noiseGender)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseGender)

	fmt.Println("Age noise")
	valAge := DiffPriv(queryAge, amount, databaseAge, epsilon)
	errNoise = json.Unmarshal(valAge, &noiseAge)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseAge)

	fmt.Println("Weight noise")
	valWeigth := DiffPriv(queryWeigth, amount, databaseWeigth, epsilon)
	errNoise = json.Unmarshal(valWeigth, &noiseWeigth)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseWeigth)

	fmt.Println("Heigth noise")
	valHeigth := DiffPriv(queryHeigth, amount, databaseHeigth, epsilon)
	errNoise = json.Unmarshal(valHeigth, &noiseHeigth)
	if errNoise != nil {
		return nil, errNoise
	}
	noisedQueries = append(noisedQueries, noiseHeigth)
	fmt.Println(noiseHeigth)

	fmt.Println("End Apply privacy")

	var dicomWithNoise []Dicom

	for i := 0; i < amount; i++ {
		var auxDcm Dicom

		auxDcm.DicomID = strconv.FormatFloat(noiseDicomID[0].Data[i], 'f', -1, 64)
		auxDcm.PatientID = strconv.FormatFloat(noisePatientID[0].Data[i], 'f', -1, 64)
		auxDcm.PatientWeigth = noiseWeigth[0].Data[i]
		auxDcm.PatientHeigth = noiseHeigth[0].Data[i]
		auxDcm.PatientAge = int(noiseAge[0].Data[i])

		auxDcm.PatientFirstname = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseFirstName[i].Data)), " "), "[]")
		auxDcm.PatientLastname = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseLastName[i].Data)), " "), "[]")
		auxDcm.PatientTelephone = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseTelephone[i].Data)), " "), "[]")
		auxDcm.PatientBirth = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseBirth[i].Data)), " "), "[]")
		auxDcm.PatientOrganization = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseOrganization[i].Data)), " "), "[]")
		auxDcm.PatientMothername = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseMothername[i].Data)), " "), "[]")
		auxDcm.PatientReligion = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseReligion[i].Data)), " "), "[]")
		auxDcm.PatientSex = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseSex[i].Data)), " "), "[]")
		auxDcm.PatientAddress = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseAddress[i].Data)), " "), "[]")
		auxDcm.PatientGender = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(noiseGender[i].Data)), " "), "[]")

		dicomWithNoise = append(dicomWithNoise, auxDcm)

	}

	return dicomWithNoise, nil
}

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
	} else if fun == "getAsset" {
		return cc.getAsset(stub, args)
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
	} else if fun == "query" {
		return cc.queryGeneral(stub, args)
	}

	fmt.Println("invoke did not find func: " + fun) //error
	return shim.Error("Received unknown function invocation")
}
