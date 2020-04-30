/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

//dp "github.com/icmc-wines/dp-privacy/diffpriv"

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the DICOM imaging attributres
type Dicom struct {
	DicomID              string    `json:"dicomID"`
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
	PatientWeigth        int       `json:"patientWeigth"`
	PatientHeigth        int       `json:"patientHeigth"`
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
	IpfsReference  string    `json:"ipfsReference"`
	DicomShared    string    `json:"dicomShared"`
	Holder         string    `json:"holder"`
	DoctorID       string    `json:"doctorID"`
	HolderAccepted bool      `json:"holderAccepted"`
	Timestamp      time.Time `json:"timestamp"`
}

// Log type defined when an user access some asset
type Log struct {
	LogID        string    `json:"dicomID"`
	AssetToken   string    `json:"assetToken"`
	TypeAsset    string    `json:"typeAsset"`
	HolderAsset  string    `json:"holderAsset"`
	HproviderGet string    `json:"hproviderGet"`
	Timestamp    time.Time `json:"timestamp"`
	WhoAccessed  string    `json:"whoAccessed"`
	AccessLevel  int       `json:"accessLevel"`
}

type Request struct {
	RequestID      string    `json:"requestID"`
	DataAmount     int       `json:"dataAmount"`
	Timestamp      time.Time `json:"timestamp"`
	HolderAccepted bool      `json:"holderAccepted"`
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

func (cc *HealthcareChaincode) addImaging(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error

	fmt.Println("Parameter number" + strconv.Itoa(len(args)))
	if len(args) < 17 || len(args) > 17 {
		return shim.Error("Incorrect number of arguments. Expecting 18")
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

	patientAge, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.Atoi(args[13])
	if err != nil {
		return shim.Error("14 argument must be a numeric string")
	}
	patientHeigth, err := strconv.Atoi(args[14])
	if err != nil {
		return shim.Error("15 argument must be a numeric string")
	}

	dicomID := args[0]
	patientFirstname := args[1]
	patientLastname := args[2]
	patientTelephone := args[3]
	patientAddress := args[4]
	patientBirth := args[6]
	patientOrganization := args[7]
	patientMothername := args[8]
	patientReligion := args[9]
	patientSex := args[10]
	patientGender := args[11]
	patientInsuranceplan := args[12]
	machineModel := args[15]

	dicomBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get pet: " + err.Error())
	} else if dicomBytes != nil {
		return shim.Error("This patient already exists: " + args[0])
	}

	rec := &Dicom{
		DicomID:              dicomID,
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

// func (cc *HealthcareChaincode) addLog(stub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	var err error

// 	fmt.Println("Parameter number" + strconv.Itoa(len(args)))
// 	if len(args) < 6 || len(args) > 6 {
// 		return shim.Error("Incorrect number of arguments. Expecting 6")
// 	}

// 	fmt.Println("- start save Logs")
// 	if len(args[0]) <= 0 {
// 		return shim.Error("1st argument must be a non-empty string")
// 	} else if len(args[1]) <= 0 {
// 		return shim.Error("2nd argument must be a non-empty string")
// 	} else if len(args[2]) <= 0 {
// 		return shim.Error("3rd argument must be a non-empty string")
// 	} else if len(args[3]) <= 0 {
// 		return shim.Error("4th argument must be a non-empty string")
// 	} else if len(args[4]) <= 0 {
// 		return shim.Error("5th argument must be a non-empty string")
// 	} else if len(args[5]) <= 0 {
// 		return shim.Error("6th argument must be a non-empty string")
// 	}

// 	accessLevel, err := strconv.Atoi(args[5])
// 	if err != nil {
// 		return shim.Error("8 argument must be a numeric string")
// 	}

// 	assetToken := args[0]
// 	typeAsset := args[1]
// 	holderAsset := args[2]
// 	hproviderGet := args[3]
// 	timestamp := time.Now()
// 	whoAccessed := args[4]

// 	hs := sha1.New()

// 	concatValues := string(accessLevel) + assetToken + typeAsset + holderAsset + hproviderGet + whoAccessed + timestamp.String()

// 	hs.Write([]byte(concatValues))

// 	logID := string(hs.Sum(nil))

// 	rec := &Log{
// 		LogID:        logID,
// 		AssetToken:   assetToken,
// 		TypeAsset:    typeAsset,
// 		HolderAsset:  holderAsset,
// 		HproviderGet: hproviderGet,
// 		Timestamp:    timestamp,
// 		WhoAccessed:  whoAccessed,
// 		AccessLevel:  accessLevel,
// 	}

// 	logJSON, err := json.Marshal(rec)

// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = stub.PutState(logID, logJSON)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	fmt.Println("- End Log Saved")
// 	return shim.Success(nil)

// }

// Sharing imaging with a doctor
// Add struct for sharing asset in blockchain
// Three args Patient ID ,Doctor or research Id and for sharing Several assets ID for sharing
func (cc *HealthcareChaincode) sharingImagingWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) > 0 && len(args) <= 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a numeric string")
	} else if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a numeric string")
	} else if len(args[3]) <= 0 {
		return shim.Error("3rd argument must be a numeric string")
	}

	//var dicomShared []string

	holder := args[0]
	doctorID := args[1]
	dicomShared := args[2]
	ipfsReference := "IPFS Ref -- ainda sera implementado"
	getTime := time.Now()

	hs := sha1.New()

	concatValues := holder + doctorID + ipfsReference + holder + getTime.String()

	hs.Write([]byte(concatValues))

	batchID := string(hs.Sum(nil))

	asset := &SharedDicom{
		BatchID:        batchID,
		IpfsReference:  ipfsReference,
		DicomShared:    dicomShared,
		Holder:         holder,
		DoctorID:       doctorID,
		HolderAccepted: true,
		Timestamp:      getTime,
	}

	logJSON, err := json.Marshal(asset)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(batchID, logJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- End File to sharing done")

	return shim.Success(nil)
}

// Doctor can get patient's data shared
// One args batch ID for get exams shared
func (cc *HealthcareChaincode) getSharedImagingWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {
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

	json.Unmarshal(batchValues, &batch)

	if batch.HolderAccepted == true {
		jsonResp = "{\"Error\":\"Holder not accepted sharing: " + batchID + "\"}"
		return shim.Error(jsonResp)
	}

	dicomValue, err := stub.GetState(batch.DicomShared)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + batch.DicomShared + "\"}"
		return shim.Error(jsonResp)
	} else if batchValues == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + batch.DicomShared + "\"}"
		return shim.Error(jsonResp)
	}

	// Nesse ponto deve ser implementando a comunicação com o IPFS para compartilhar DICOM

	//Gravar Logs de acesso a imagem
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + time.Now().String()
	hs := sha1.New()
	hs.Write([]byte(id))
	logID := string(hs.Sum(nil))
	log := Log{
		LogID:        logID,
		AssetToken:   "Falta gererar",
		HolderAsset:  batch.Holder,
		HproviderGet: "Doctor",
		WhoAccessed:  batch.DoctorID,
		AccessLevel:  1,
	}

	if addLog(stub, log) != true {
		jsonResp = "{\"Error\":\" Logs dont record \"}"
		return shim.Error(jsonResp)
	}

	// Deve ser aplicado a K-Anonimity antes de enviar os dados para o médico

	return shim.Success(dicomValue)
}

// Request imaging
// One attribute amount images
func (cc *HealthcareChaincode) requestImagingForResearchers(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("argument must be a numeric string")
	}

	timeGet := time.Now()
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + timeGet.String()
	hs := sha1.New()
	hs.Write([]byte(id))
	reqID := string(hs.Sum(nil))

	request := &Request{
		RequestID:      reqID,
		DataAmount:     amount,
		Timestamp:      timeGet,
		HolderAccepted: true,
	}

	requestJSON, err := json.Marshal(request)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(reqID, requestJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- Request Sent")
	return shim.Success(nil)
}

// Sharing request values with the research
// Verify the stub.Query how to use ??? Thus we can return all objets DICOM and using differential privacy
func (cc *HealthcareChaincode) sharingImagingForResearchers(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	return shim.Success(nil)
}

// Store logs in blockain
// Internal function
func addLog(stub shim.ChaincodeStubInterface, log Log) bool {

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

// Audit Logs for verify who accessed one image
func (cc *HealthcareChaincode) auditLogs(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *HealthcareChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is runing " + fun)

	if fun == "addImaging" {
		return cc.addImaging(stub, args)
	} else if fun == "sharingImagingWithDoctor" {
		return cc.sharingImagingWithDoctor(stub, args)
	} else if fun == "getSharedImagingWithDoctor" {
		return cc.getSharedImagingWithDoctor(stub, args)
	} else if fun == "sharingImagingForResearchers" {
		return cc.sharingImagingForResearchers(stub, args)
	} else if fun == "auditLogs" {
		return cc.auditLogs(stub, args)
	}

	fmt.Println("invoke did not find func: " + fun) //error
	return shim.Error("Received unknown function invocation")
}
