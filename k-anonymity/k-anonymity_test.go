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

	amount := 1
	var data []interface{}

	for i := len(allDicom) - 1; i >= (len(allDicom) - amount); i-- {
		dicomNew := map[string]interface{}{
			"dicomID":      anonymizedDicomID[i],
			"patientID":    anonymizedPatientID[i],
			"FirstName":    anonymizedFirstName[i],
			"LastName":     anonymizedLastName[i],
			"Telephone":    anonymizedTelephone[i],
			"birthday":     anonymizedBirth[i],
			"organization": anonymizedOrganization[i],
			"mothername":   anonymizedMothername[i],
			"relegion":     anonymizedReligion[i],
			"sex":          anonymizedSex[i],
			"address":      anonymizedAddress[i],
			"gender":       anonymizedGender[i],
			"age":          anonymizedAge[i],
			"weigth":       anonymizedWeigth[i],
			"heigth":       anonymizedHeigth[i],
		}

		data = append(data, dicomNew)
	}

	dataValue, err := json.Marshal(data)

	if err != nil {
		panic("Error: " + err.Error())
	}

	var finalValue []map[string]interface{}

	err = json.Unmarshal(dataValue, &finalValue)

	if err != nil {
		panic("Error: " + err.Error())
	}

	fmt.Println("Final Value Anonymized")
	fmt.Println(finalValue)
}
