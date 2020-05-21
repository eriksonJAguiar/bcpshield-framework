package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestKAnonymity(t *testing.T) {
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
		PatientBirth:         "1980-03-20",
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
		PatientBirth:         "1970-01-20",
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
		PatientBirth:         "1996-08-31",
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
		PatientBirth:         "1996-05-10",
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

	var patientFirstnames, patientLastnames, patientTelephones, dicomID, patientID []string
	var patientBirth, patientOrganization, patientMothername, patientReligion []string
	var patientSex, patientGender, patientAddress []string
	var patientInsuranceplan, machineModel []string
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
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.PatientGender)
		patientInsuranceplan = append(patientInsuranceplan, dcm.PatientInsuranceplan)
		machineModel = append(machineModel, dcm.MachineModel)
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

	amount := 1
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

	dataValue, err := json.Marshal(data)

	if err != nil {
		panic("Error: " + err.Error())
	}

	var finalValue *interface{}

	err = json.Unmarshal(dataValue, &finalValue)

	if err != nil {
		panic("Error: " + err.Error())
	}

	fmt.Println("Final Value Anonymized")
	fmt.Println(*finalValue)
}
