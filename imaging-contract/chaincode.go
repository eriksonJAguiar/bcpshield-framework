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
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the DICOM imaging attributres
type Dicom struct {
	DicomID              string    `json:"dicomID"`
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
	DoctorID       string    `json:"doctorID"`
	HolderAccepted bool      `json:"holderAccepted"`
	Timestamp      time.Time `json:"timestamp"`
	DataAmount     int       `json:"dataAmount"`
}

// Log type defined when an user access some asset
type Log struct {
	LogID        string    `json:"dicomID"`
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
	RequestID      string    `json:"requestID"`
	DocType        string    `json:"docType"`
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

// Add imaging in blockchain network
func (cc *HealthcareChaincode) addImaging(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error

	fmt.Println("Parameter number " + strconv.Itoa(len(args)))
	if len(args) < 16 || len(args) > 16 {
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
	}

	patientAge, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("7 argument must be a numeric string")
	}
	timestamp := time.Now()
	patientWeigth, err := strconv.ParseFloat(args[13], 64)
	if err != nil {
		return shim.Error("14 argument must be a numeric string")
	}
	patientHeigth, err := strconv.ParseFloat(args[14], 64)
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

	dicomBytes, err := stub.GetState(dicomID)
	if err != nil {
		return shim.Error("Failed to get pet: " + err.Error())
	} else if dicomBytes != nil {
		return shim.Error("This patient already exists: " + dicomID)
	}

	rec := &Dicom{
		DicomID:              dicomID,
		DocType:              "Dicom",
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

func (cc *HealthcareChaincode) getImaging(stub shim.ChaincodeStubInterface, args []string) sc.Response {

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

// Sharing imaging with a doctor
// Add struct for sharing asset in blockchain
// Three args Patient ID ,Doctor or research Id and for sharing Several assets ID for sharing
func (cc *HealthcareChaincode) sharingImagingWithDoctor(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 3 && len(args) > 3 {
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
	dicomShared = append(dicomShared, args[2])
	ipfsReference := "IPFS Ref -- ainda sera implementado"
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

	auxBatch := map[string]interface{}{"id": asset.BatchID}
	idJson, err := json.Marshal(auxBatch)

	fmt.Println("- End File to sharing done")

	return shim.Success(idJson)
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

	// if bool(batch.HolderAccepted) == true {
	// 	jsonResp = "{\"Error\":\"Holder not accepted sharing: " + batchID + "\"}"
	// 	return shim.Error(jsonResp)
	// }

	dicom := batch.DicomShared[0]

	dicomValue, err := stub.GetState(dicom)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + dicom + "\"}"
		return shim.Error(jsonResp)
	} else if batchValues == nil {
		jsonResp = "{\"Error\":\"Record Sharing does not exist: " + dicom + "\"}"
		return shim.Error(jsonResp)
	}

	// Nesse ponto deve ser implementando a comunicação com o IPFS para compartilhar DICOM

	//Gravar Logs de acesso a imagem
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + time.Now().String()
	hs := sha1.New()
	hs.Write([]byte(id))
	hexLogID := hs.Sum(nil)
	logID := hex.EncodeToString(hexLogID)

	log := Log{
		LogID:        logID,
		DocType:      "Log",
		AssetToken:   "Falta gererar",
		HolderAsset:  batch.Holder,
		HproviderGet: "Doctor",
		WhoAccessed:  batch.DoctorID,
		AccessLevel:  1,
	}

	if cc.addLog(stub, log) != true {
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
	hexReqID := hs.Sum(nil)
	reqID := hex.EncodeToString(hexReqID)

	request := &Request{
		RequestID:      reqID,
		DocType:        "Request",
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
	return shim.Success(requestJSON)
}

// Sharing request values with the research
// Verify the stub.Query how to use ??? Thus we can return all objets DICOM and using differential privacy
// Two atributes Research ID and BatchShared Dicom ID
func (cc *HealthcareChaincode) sharingImagingForResearchers(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var batchID, jsonResp string
	var err error

	if len(args) != 2 {
		shim.Error("We expected two element as params")
	}

	researchID := args[0]
	batchID = args[1]

	batchValues, err := stub.GetState(batchID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + batchID + "\"}"
		return shim.Error(jsonResp)
	} else if batchValues == nil {
		jsonResp = "{\"Error\":\"Record does not exist: " + batchID + "\"}"
		return shim.Error(jsonResp)
	}

	batchResponse, err := stub.GetState(batchID)

	var batchObj SharedDicom

	json.Unmarshal(batchResponse, &batchObj)

	var dicoms []Dicom

	for _, ast := range batchObj.DicomShared {
		resDicom, err := stub.GetState(ast)
		if err != nil {
			fmt.Println("Error get element ", ast)
			continue
		}
		var img Dicom
		json.Unmarshal(resDicom, &img)
		dicoms = append(dicoms, img)
	}

	//Implementar o envio da imagem pelo IFPS

	//Gravar Logs de acesso a imagem
	rand.Seed(time.Now().UnixNano())
	id := strconv.Itoa(rand.Int()) + time.Now().String()
	hs := sha1.New()
	hs.Write([]byte(id))
	hexLogID := hs.Sum(nil)
	logID := hex.EncodeToString(hexLogID)

	log := Log{
		LogID:        logID,
		DocType:      "Log",
		AssetToken:   "Falta gererar",
		HolderAsset:  batchObj.Holder,
		HproviderGet: "Research",
		WhoAccessed:  researchID,
		AccessLevel:  2,
	}

	if cc.addLog(stub, log) != true {
		jsonResp = "{\"Error\":\" Logs dont record \"}"
		return shim.Error(jsonResp)
	}

	anonimizedDicoms := anonimizeDiffPriv(stub, dicoms, batchObj.DataAmount)

	anonimizedBytes, err := json.Marshal(anonimizedDicoms)

	return shim.Success(anonimizedBytes)
}

// Audit Logs for verify who accessed one image
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

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

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
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func anonimizeDiffPriv(stub shim.ChaincodeStubInterface, assets []Dicom, amount int) []Dicom {

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Dicom\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return nil
	}

	// Database total
	var allDicom []Dicom

	json.Unmarshal(queryResults, &allDicom)
	var dicomID, patientFirstnames, patientLastnames, patientTelephones []string
	var patientBirth, patientOrganization, patientMothername, patientReligion []string
	var patientSex, patientGender, patientAddress []string
	var patientAge []int
	var patientWeigth, patientHeigth []float64

	for _, dcm := range allDicom {
		dicomID = append(dicomID, dcm.DicomID)
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

	databaseDicomID := TransforSymbolicData(dicomID)
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

	queryDicomID := Query(databaseDicomID, amount)
	queryFirstName := Query(databaseFirstName, amount)
	queryLastName := Query(databaseLastName, amount)
	queryTelephone := Query(databaseTelephone, amount)
	queryBirth := Query(databaseBirth, amount)
	queryOrganization := Query(databaseOrganization, amount)
	queryMothername := Query(databaseMothername, amount)
	queryReligion := Query(databaseReligion, amount)
	querySex := Query(databaseSex, amount)
	queryAddress := Query(databaseAddress, amount)
	queryGender := Query(databaseGender, amount)
	queryAge := Query(databaseAge, amount)
	queryWeigth := Query(databaseWeigth, amount)
	queryHeigth := Query(databaseHeigth, amount)

	var noisedQueries []Matrix
	var noiseDicomID, noiseFirstName, noiseLastName, noiseTelephoneName Matrix
	var noiseBirth, noiseOrganization, noiseMothername, noiseReligion Matrix
	var noiseSex, noiseAddress, noiseGender, noiseAge, noiseWeigth, noiseHeigth Matrix

	epsilon := 0.1
	json.Unmarshal(DiffPriv(queryDicomID, amount, databaseDicomID, epsilon), &noiseDicomID)
	noisedQueries = append(noisedQueries, noiseDicomID)

	json.Unmarshal(DiffPriv(queryFirstName, amount, databaseFirstName, epsilon), &noiseFirstName)
	noisedQueries = append(noisedQueries, noiseFirstName)

	json.Unmarshal(DiffPriv(queryLastName, amount, databaseLastName, epsilon), &noiseLastName)
	noisedQueries = append(noisedQueries, noiseLastName)

	json.Unmarshal(DiffPriv(queryTelephone, amount, databaseTelephone, epsilon), &noiseTelephoneName)
	noisedQueries = append(noisedQueries, noiseTelephoneName)

	json.Unmarshal(DiffPriv(queryBirth, amount, databaseBirth, epsilon), &noiseBirth)
	noisedQueries = append(noisedQueries, noiseBirth)

	json.Unmarshal(DiffPriv(queryOrganization, amount, databaseOrganization, epsilon), &noiseOrganization)
	noisedQueries = append(noisedQueries, noiseOrganization)

	json.Unmarshal(DiffPriv(queryMothername, amount, databaseMothername, epsilon), &noiseMothername)
	noisedQueries = append(noisedQueries, noiseMothername)

	json.Unmarshal(DiffPriv(queryReligion, amount, databaseReligion, epsilon), &noiseReligion)
	noisedQueries = append(noisedQueries, noiseReligion)

	json.Unmarshal(DiffPriv(queryReligion, amount, databaseReligion, epsilon), &noiseReligion)
	noisedQueries = append(noisedQueries, noiseReligion)

	json.Unmarshal(DiffPriv(queryReligion, amount, databaseReligion, epsilon), &noiseReligion)
	noisedQueries = append(noisedQueries, noiseReligion)

	json.Unmarshal(DiffPriv(querySex, amount, databaseSex, epsilon), &noiseSex)
	noisedQueries = append(noisedQueries, noiseSex)

	json.Unmarshal(DiffPriv(queryAddress, amount, databaseAddress, epsilon), &noiseAddress)
	noisedQueries = append(noisedQueries, noiseAddress)

	json.Unmarshal(DiffPriv(queryGender, amount, databaseGender, epsilon), &noiseGender)
	noisedQueries = append(noisedQueries, noiseGender)

	json.Unmarshal(DiffPriv(queryAge, amount, databaseAge, epsilon), &noiseAge)
	noisedQueries = append(noisedQueries, noiseAge)

	json.Unmarshal(DiffPriv(queryAge, amount, databaseAge, epsilon), &noiseAge)
	noisedQueries = append(noisedQueries, noiseAge)

	json.Unmarshal(DiffPriv(queryWeigth, amount, databaseWeigth, epsilon), &noiseWeigth)
	noisedQueries = append(noisedQueries, noiseWeigth)

	json.Unmarshal(DiffPriv(queryHeigth, amount, databaseHeigth, epsilon), &noiseHeigth)
	noisedQueries = append(noisedQueries, noiseHeigth)

	var dicomWithNoise []Dicom

	for i := 0; i < amount; i++ {
		var auxDcm Dicom
		auxDcm.DicomID = strconv.FormatFloat(noisedQueries[0].Data[i], 'E', -1, 64)
		auxDcm.PatientFirstname = strconv.FormatFloat(noisedQueries[1].Data[i], 'E', -1, 64)
		auxDcm.PatientLastname = strconv.FormatFloat(noisedQueries[2].Data[i], 'E', -1, 64)
		auxDcm.PatientTelephone = strconv.FormatFloat(noisedQueries[3].Data[i], 'E', -1, 64)
		auxDcm.PatientBirth = strconv.FormatFloat(noisedQueries[4].Data[i], 'E', -1, 64)
		auxDcm.PatientOrganization = strconv.FormatFloat(noisedQueries[5].Data[i], 'E', -1, 64)
		auxDcm.PatientMothername = strconv.FormatFloat(noisedQueries[6].Data[i], 'E', -1, 64)
		auxDcm.PatientReligion = strconv.FormatFloat(noisedQueries[7].Data[i], 'E', -1, 64)
		auxDcm.PatientSex = strconv.FormatFloat(noisedQueries[8].Data[i], 'E', -1, 64)
		auxDcm.PatientAddress = strconv.FormatFloat(noisedQueries[9].Data[i], 'E', -1, 64)
		auxDcm.PatientGender = strconv.FormatFloat(noisedQueries[10].Data[i], 'E', -1, 64)
		auxDcm.PatientAge = int(noisedQueries[11].Data[i])

		auxDcm.PatientWeigth = noisedQueries[12].Data[i]

		auxDcm.PatientHeigth = noisedQueries[13].Data[i]

		dicomWithNoise = append(dicomWithNoise, auxDcm)

	}

	return dicomWithNoise
}

func anonimizeKAnonimity(assets []Dicom) []Dicom {
	return nil
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *HealthcareChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fun, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is runing " + fun)

	if fun == "addImaging" {
		return cc.addImaging(stub, args)
	} else if fun == "getImaging" {
		return cc.getImaging(stub, args)
	} else if fun == "sharingImagingWithDoctor" {
		return cc.sharingImagingWithDoctor(stub, args)
	} else if fun == "getSharedImagingWithDoctor" {
		return cc.getSharedImagingWithDoctor(stub, args)
	} else if fun == "requestImagingForResearchers" {
		return cc.requestImagingForResearchers(stub, args)
	} else if fun == "sharingImagingForResearchers" {
		return cc.sharingImagingForResearchers(stub, args)
	} else if fun == "auditLogs" {
		return cc.auditLogs(stub, args)
	}

	fmt.Println("invoke did not find func: " + fun) //error
	return shim.Error("Received unknown function invocation")
}
