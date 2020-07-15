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
	PatientOrganization  string    `json:"patientOrganization"`
	PatientRace          string    `json:"patientRace"`
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
	if len(args) < 14 || len(args) > 14 {
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
	}

	patientAge, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.ParseFloat(args[11], 64)
	if err != nil {
		return shim.Error("12 argument must be a numeric string")
	}
	patientHeigth, err := strconv.ParseFloat(args[12], 64)
	if err != nil {
		return shim.Error("13 argument must be a numeric string")
	}

	dicomID := args[0]
	patientID := args[1]
	patientFirstname := args[2]
	patientLastname := args[3]
	patientTelephone := args[4]
	patientAddress := args[5]
	patientOrganization := args[7]
	patientRace := args[8]
	patientGender := args[9]
	patientInsuranceplan := args[10]
	machineModel := args[13]

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
		PatientOrganization:  patientOrganization,
		PatientRace:          patientRace,
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

	fmt.Println("Parameter number " + strconv.Itoa(len(args)))
	if len(args) < 14 || len(args) > 14 {
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
	}

	patientAge, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.ParseFloat(args[11], 64)
	if err != nil {
		return shim.Error("12 argument must be a numeric string")
	}
	patientHeigth, err := strconv.ParseFloat(args[12], 64)
	if err != nil {
		return shim.Error("13 argument must be a numeric string")
	}

	dicomID := args[0]
	patientID := args[1]
	patientFirstname := args[2]
	patientLastname := args[3]
	patientTelephone := args[4]
	patientAddress := args[5]
	patientOrganization := args[7]
	patientRace := args[8]
	patientGender := args[9]
	patientInsuranceplan := args[10]
	machineModel := args[13]

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
		PatientOrganization:  patientOrganization,
		PatientRace:          patientRace,
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

// Add asset when add Differential Privacy
// Params: (two param) string byte asset pivate
// Returns: True or False for transactions
func (cc *HealthcareChaincode) addAssetDiff(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Error we expected one param")
	}

	id := args[0]
	dicom := []byte(args[1])

	asset, _ := stub.GetState(id)
	if asset != nil || len(asset) > 0 {
		return shim.Error("This asset already exists")
	}

	if err := stub.PutState(id, dicom); err != nil {
		return shim.Error("Error Put State on blockchian")
	}

	return shim.Success([]byte("true"))
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

// Researcher request imaging from patient or other researcher
// Params: (three args) requester ID; patient ID; userType
// "PatientID": to represents patient he wants  request image
// Returns: requestID to describes the request ID generated
func (cc *HealthcareChaincode) requestAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// amount, err := strconv.Atoi(args[0])
	// if err != nil {
	// 	return shim.Error("argument must be a numeric string")
	// }

	requesterID := args[0]
	patientID := args[1]
	userType := args[2]

	timeGet := time.Now()
	auxid, err := uuid.NewRandom()
	if err != nil {
		return shim.Error("Erro genereta UUID: " + err.Error())
	}
	reqID := auxid.String()

	// request := &Request{
	// 	RequestID:       reqID,
	// 	DocType:         "Request",
	// 	DataAmount:      1,
	// 	Timestamp:       timeGet,
	// 	HolderRequested: patientID,
	// 	UserRequest:     researchID,
	// }

	request := &SharedDicom{
		BatchID:        reqID,
		DocType:        "SharedDicom",
		IpfsReference:  "",
		Holder:         patientID,
		HolderAccepted: false,
		Timestamp:      timeGet,
		DataAmount:     1,
		WhoAccessed:    requesterID,
		AccessLevel:    userType,
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

// Getting data requested from healthcare provider
// Params: (One arg) request ID
// Returns: Data requested if it is available
func (cc *HealthcareChaincode) getRequested(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("We expected one param")
	}

	requestID := args[0]

	var jsonResp string
	sharedValue, err := stub.GetState(requestID)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + requestID + "\"}"
		return shim.Error(jsonResp)
	} else if sharedValue == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + requestID + "\"}"
		return shim.Error(jsonResp)
	}

	var shared SharedDicom
	if err = json.Unmarshal(sharedValue, &shared); err != nil {
		return shim.Error("Error to convert shared asset: " + err.Error())
	}

	if shared.DicomShared == nil {
		return shim.Error("There aren't images shared with you!")
	}

	var dicomShared []interface{}

	for _, dcmID := range shared.DicomShared {
		byteDcmAux, err := stub.GetState(dcmID)
		if err != nil {
			continue
		}
		var dcmAux interface{}
		if err = json.Unmarshal(byteDcmAux, &dcmAux); err != nil {
			continue
		}
		dicomShared = append(dicomShared, dcmAux)

	}

	byteDicomShared, err := json.Marshal(dicomShared)
	if err != nil {
		return shim.Error("Error convert bytes Shared Dicom: " + err.Error())
	}

	return shim.Success(byteDicomShared)
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

// Researcher or patients sharing imaging request by others research from researchers requests
// Params: (two args) "holderID": patient holder ID; "timestamp": time last observation
// Returns: requests found
func (cc *HealthcareChaincode) observerRequests(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Params error, we expected 2 paramters")
	}

	holderID := args[0]
	timeReq := args[1]

	query := fmt.Sprintf("{ \"selector\":{ \"docType\":\"SharedDicom\", \"holder\": \"%s\",\"timestamp\": { \"$gte\": \" %s\" } }}", holderID, timeReq)

	queryRes, err := getQueryResultForQueryString(stub, query)
	if err != nil {
		return shim.Error("Query error: " + err.Error())
	}
	// var result interface{}
	// err = json.Unmarshal(queryRes, &result)
	// if err != nil {
	// 	return shim.Error("Error convert query result: " + err.Error())
	// }

	return shim.Success(queryRes)

}

// Researcher or patients sharing imaging request by others research from researchers requests
// Params: (N args) "requestID"; 1 or more assets ID to update on blockchain
// Returns: requests found
func (cc *HealthcareChaincode) notifyRequester(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 2 {
		return shim.Error("We expected 2 param - requestID and SharedAsset IDs")
	}

	requestID := args[0]
	byteReq, err := stub.GetState(requestID)
	if err != nil {
		shim.Error("Error to get asset through requestID")
	}

	var shared SharedDicom
	if err = json.Unmarshal(byteReq, &shared); err != nil {
		return shim.Error("Error convert shared Dicom")
	}

	//shared.DicomShared =

	shared.DicomShared = make([]string, (len(args) - 1))
	for i := 1; i < len(args); i++ {
		shared.DicomShared = append(shared.DicomShared, args[i])
	}

	shared.HolderAccepted = true

	byteShare, _ := json.Marshal(shared)

	stub.PutState(shared.BatchID, byteShare)

	return shim.Success(byteShare)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *HealthcareChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is runing " + fun)

	if fun == "addAsset" {
		return cc.addAsset(stub, args)
	} else if fun == "addAssetPriv" {
		return cc.addAssetPriv(stub, args)
	} else if fun == "addAssetDiff" {
		return cc.addAssetDiff(stub, args)
	} else if fun == "getAsset" {
		return cc.getAsset(stub, args)
	} else if fun == "getAssetPriv" {
		return cc.getAssetPriv(stub, args)
	} else if fun == "getRequested" {
		return cc.getRequested(stub, args)
	} else if fun == "requestAsset" {
		return cc.requestAsset(stub, args)
	} else if fun == "observerRequests" {
		return cc.observerRequests(stub, args)
	} else if fun == "auditLogs" {
		return cc.auditLogs(stub, args)
	}

	fmt.Println("invoke did not find func: " + fun) //error
	return shim.Error("Received unknown function invocation")
}
