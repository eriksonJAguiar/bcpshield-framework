package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	dp "github.com/icmc-wines/dp-privacy/diffpriv"
)

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

func testForNumericValues() {
	originalData := []float64{0.5, 0.9, 1.3, 0.1, 0.3, 0.6}
	//First step is convert any data type to Matrix
	convertedData := dp.TransforFloatData(originalData)
	//define amount data recovered
	amountQuery := 4

	//Send a query in dataset
	DataFromQuery := dp.Query(convertedData, amountQuery)
	//Apply diiferential privacy function
	epsilon := 1.0
	noiseData := dp.DiffPriv(DataFromQuery, amountQuery, convertedData, epsilon)

	fmt.Println("Dataset noise ", noiseData)
}

func testForSymbolicValues() {
	originalData := []string{"So close no matter how far", "It couldn't be much more from the heart", "Forever trusting who we are", "And nothing else matters"}
	//First step is convert any data type to Matrix
	convertedData := dp.TransforSymbolicData(originalData)
	//define amount data recovered
	amountQuery := 2
	//Send a query in dataset
	DataFromQuery := dp.Query(convertedData, amountQuery)
	//Apply diiferential privacy function
	epsilon := 1.0
	noiseData := dp.DiffPriv(DataFromQuery, amountQuery, convertedData, epsilon)
	fmt.Println(convertedData)
	fmt.Println("---------------------------")
	fmt.Println("Dataset noise ", noiseData)
}

func testAnonimizeDiffPriv() ([]Dicom, error) {

	// , , , , , , "23", , , , , "Male", "ASASSAS", "Plan X", "75.5", "1.80", "ASDFG"

	// Database total
	var allDicom []Dicom

	allDicom = append(allDicom, Dicom{
		DicomID:              "10005",
		PatientID:            "11110",
		DocType:              "Dicom",
		PatientFirstname:     "Bob",
		PatientLastname:      "Truth",
		PatientTelephone:     "(43) 0000-0000",
		PatientAddress:       "São Carlos - SP",
		PatientAge:           23,
		PatientBirth:         "1996-31-08",
		PatientOrganization:  "USP",
		PatientMothername:    "AAAA",
		PatientReligion:      "None",
		PatientSex:           "Male",
		PatientGender:        "None",
		PatientInsuranceplan: "AAAAAAAAA",
		PatientWeigth:        1.80,
		PatientHeigth:        70.5,
		MachineModel:         "AXAXAXAXA",
		Timestamp:            time.Now(),
	})

	allDicom = append(allDicom, Dicom{
		DicomID:              "10006",
		PatientID:            "11110",
		DocType:              "Dicom",
		PatientFirstname:     "Bob",
		PatientLastname:      "Truth",
		PatientTelephone:     "(43) 0000-0000",
		PatientAddress:       "São Carlos - SP",
		PatientAge:           23,
		PatientBirth:         "1996-31-08",
		PatientOrganization:  "USP",
		PatientMothername:    "AAAAA",
		PatientReligion:      "None",
		PatientSex:           "Male",
		PatientGender:        "None",
		PatientInsuranceplan: "AAAAAAAAA",
		PatientWeigth:        1.80,
		PatientHeigth:        70.5,
		MachineModel:         "AXAXAXAXA",
		Timestamp:            time.Now(),
	})

	allDicom = append(allDicom, Dicom{
		DicomID:              "10007",
		PatientID:            "11110",
		DocType:              "Dicom",
		PatientFirstname:     "Bob",
		PatientLastname:      "Truth",
		PatientTelephone:     "(43) 0000-0000",
		PatientAddress:       "São Carlos - SP",
		PatientAge:           23,
		PatientBirth:         "1996-31-08",
		PatientOrganization:  "USP",
		PatientMothername:    "AAAA",
		PatientReligion:      "None",
		PatientSex:           "Male",
		PatientGender:        "None",
		PatientInsuranceplan: "AAAAAAAAA",
		PatientWeigth:        1.80,
		PatientHeigth:        70.5,
		MachineModel:         "AXAXAXAXA",
		Timestamp:            time.Now(),
	})

	allDicom = append(allDicom, Dicom{
		DicomID:              "10009",
		PatientID:            "11110",
		DocType:              "Dicom",
		PatientFirstname:     "Bob",
		PatientLastname:      "Truth",
		PatientTelephone:     "(43) 0000-0000",
		PatientAddress:       "São Carlos - SP",
		PatientAge:           23,
		PatientBirth:         "1996-31-08",
		PatientOrganization:  "USP",
		PatientMothername:    "AAAA",
		PatientReligion:      "None",
		PatientSex:           "Male",
		PatientGender:        "None",
		PatientInsuranceplan: "AAAAAAAAA",
		PatientWeigth:        1.80,
		PatientHeigth:        70.5,
		MachineModel:         "AXAXAXAXA",
		Timestamp:            time.Now(),
	})

	var assets []Dicom
	assets = append(assets, allDicom[0])
	assets = append(assets, allDicom[1])
	amount := 2
	epsilon := 0.1

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

	databaseDicomID := dp.TransforIntData(dicomID)
	databasePatientID := dp.TransforIntData(patientID)
	databaseFirstName := dp.TransforSymbolicData(patientFirstnames)
	databaseLastName := dp.TransforSymbolicData(patientLastnames)
	databaseTelephone := dp.TransforSymbolicData(patientTelephones)
	databaseBirth := dp.TransforSymbolicData(patientBirth)
	databaseOrganization := dp.TransforSymbolicData(patientOrganization)
	databaseMothername := dp.TransforSymbolicData(patientMothername)
	databaseReligion := dp.TransforSymbolicData(patientReligion)
	databaseSex := dp.TransforSymbolicData(patientSex)
	databaseAddress := dp.TransforSymbolicData(patientAddress)
	databaseGender := dp.TransforSymbolicData(patientGender)
	databaseAge := dp.TransforIntData(patientAge)
	databaseWeigth := dp.TransforFloatData(patientWeigth)
	databaseHeigth := dp.TransforFloatData(patientHeigth)

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

	queryDicomID := dp.TransforIntData(dicomIDQ)
	queryPatientID := dp.TransforIntData(patientIDQ)
	queryFirstName := dp.TransforSymbolicData(patientFirstnamesQ)
	queryLastName := dp.TransforSymbolicData(patientLastnamesQ)
	queryTelephone := dp.TransforSymbolicData(patientTelephonesQ)
	queryBirth := dp.TransforSymbolicData(patientBirthQ)
	queryOrganization := dp.TransforSymbolicData(patientOrganizationQ)
	queryMothername := dp.TransforSymbolicData(patientMothernameQ)
	queryReligion := dp.TransforSymbolicData(patientReligionQ)
	querySex := dp.TransforSymbolicData(patientSexQ)
	queryAddress := dp.TransforSymbolicData(patientAddressQ)
	queryGender := dp.TransforSymbolicData(patientGenderQ)
	queryAge := dp.TransforIntData(patientAgeQ)
	queryWeigth := dp.TransforFloatData(patientWeigthQ)
	queryHeigth := dp.TransforFloatData(patientHeigthQ)

	fmt.Println("Query PatientID values")
	fmt.Println(queryDicomID)

	fmt.Println("Query FirtName values")
	fmt.Println(queryFirstName)

	fmt.Println("Start Apply privacy")

	var noisedQueries []interface{}
	var noiseDicomID, noisePatientID, noiseFirstName, noiseLastName, noiseTelephone []dp.Matrix
	var noiseBirth, noiseOrganization, noiseMothername, noiseReligion []dp.Matrix
	var noiseSex, noiseAddress, noiseGender, noiseAge, noiseWeigth, noiseHeigth []dp.Matrix
	var errNoise error

	fmt.Println("Dicom ID")
	valDicomID := dp.DiffPriv(queryDicomID, amount, databaseDicomID, epsilon)
	errNoise = json.Unmarshal(valDicomID, &noiseDicomID)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseDicomID)

	fmt.Println("Patient Noise")
	valPatientID := dp.DiffPriv(queryPatientID, amount, databasePatientID, epsilon)
	errNoise = json.Unmarshal(valPatientID, &noisePatientID)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noisePatientID)

	fmt.Println("First name Noise")
	valFirtName := dp.DiffPriv(queryFirstName, amount, databaseFirstName, epsilon)
	errNoise = json.Unmarshal(valFirtName, &noiseFirstName)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseFirstName)

	fmt.Println("Last Name noise")
	valLastName := dp.DiffPriv(queryLastName, amount, databaseLastName, epsilon)
	errNoise = json.Unmarshal(valLastName, &noiseLastName)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseLastName)

	fmt.Println("Telephone noise")
	valTelephone := dp.DiffPriv(queryTelephone, amount, databaseTelephone, epsilon)
	errNoise = json.Unmarshal(valTelephone, &noiseTelephone)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseTelephone)

	fmt.Println("Birth noise")
	valBirth := dp.DiffPriv(queryBirth, amount, databaseBirth, epsilon)
	errNoise = json.Unmarshal(valBirth, &noiseBirth)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseBirth)

	fmt.Println("Organization noise")
	valOrganization := dp.DiffPriv(queryOrganization, amount, databaseOrganization, epsilon)
	errNoise = json.Unmarshal(valOrganization, &noiseOrganization)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseOrganization)

	fmt.Println("Mothername noise")
	valMothername := dp.DiffPriv(queryMothername, amount, databaseMothername, epsilon)
	errNoise = json.Unmarshal(valMothername, &noiseMothername)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseMothername)

	fmt.Println("Religion noise")
	valReligion := dp.DiffPriv(queryReligion, amount, databaseReligion, epsilon)
	errNoise = json.Unmarshal(valReligion, &noiseReligion)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseReligion)

	fmt.Println("Sex noise")
	valSex := dp.DiffPriv(querySex, amount, databaseSex, epsilon)
	errNoise = json.Unmarshal(valSex, &noiseSex)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseSex)

	fmt.Println("Noise Address")
	valAddress := dp.DiffPriv(queryAddress, amount, databaseAddress, epsilon)
	errNoise = json.Unmarshal(valAddress, &noiseAddress)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseAddress)

	fmt.Println("Noise gender")
	valGender := dp.DiffPriv(queryGender, amount, databaseGender, epsilon)
	errNoise = json.Unmarshal(valGender, &noiseGender)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseGender)

	fmt.Println("Age noise")
	valAge := dp.DiffPriv(queryAge, amount, databaseAge, epsilon)
	errNoise = json.Unmarshal(valAge, &noiseAge)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseAge)

	fmt.Println("Weight noise")
	valWeigth := dp.DiffPriv(queryWeigth, amount, databaseWeigth, epsilon)
	errNoise = json.Unmarshal(valWeigth, &noiseWeigth)
	if errNoise != nil {
		return nil, errNoise
	}
	fmt.Println(noiseWeigth)

	fmt.Println("Heigth noise")
	valHeigth := dp.DiffPriv(queryHeigth, amount, databaseHeigth, epsilon)
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

func main() {

	dicomNoise, err := testAnonimizeDiffPriv()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(dicomNoise[0].PatientAddress)

}
