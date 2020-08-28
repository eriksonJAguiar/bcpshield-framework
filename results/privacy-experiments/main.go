package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	kpriv "privacyexpr/kanonymity"

	dp "github.com/eriksonJAguiar/godiffpriv"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dicom struct {
	DicomID              string    `bson:"dicomID"`
	PatientID            string    `bson:"patientID"`
	PatientFirstname     string    `bson:"patientFirstname"`
	PatientLastname      string    `bson:"patientLastname"`
	PatientTelephone     string    `bson:"patientTelephone"`
	PatientAddress       string    `bson:"patientAddress"`
	PatientAge           int       `bson:"patientAge"`
	PatientOrganization  string    `bson:"patientOrganization"`
	PatientRace          string    `bson:"patientRace"`
	patientGender        string    `bson:"patientGender"`
	PatientInsuranceplan string    `bson:"patientInsuranceplan"`
	PatientWeigth        float64   `bson:"patientWeigth"`
	PatientHeigth        float64   `bson:"patientHeigth"`
	MachineModel         string    `bson:"machineModel"`
	Timestamp            time.Time `bson:"timestamp"`
}

type dicomAux struct {
	DicomID              string    `json:"dicomID"`
	PatientID            string    `json:"patientID"`
	DocType              string    `json:"docType"`
	PatientFirstname     string    `json:"patientFirstname"`
	PatientLastname      string    `json:"patientLastname"`
	PatientTelephone     string    `json:"patientTelephone"`
	PatientAddress       string    `json:"patientAddress"`
	PatientAge           string    `json:"patientAge"`
	PatientOrganization  string    `json:"patientOrganization"`
	PatientRace          string    `json:"patientRace"`
	PatientGender        string    `json:"patientGender"`
	PatientInsuranceplan string    `json:"patientInsuranceplan"`
	PatientWeigth        string    `json:"patientWeigth"`
	PatientHeigth        string    `json:"patientHeigth"`
	MachineModel         string    `json:"machineModel"`
	Timestamp            time.Time `json:"timestamp"`
}

func applyDiffPriv(data dicom, epsilon float64) ([]byte, error) {

	var allDicom []dicom
	allDicom, err := getAllMongo()
	if err != nil {
		return nil, err
	}

	var patientOrganization, patientRace []string
	var patientGender, patientAddress []string
	//var dicomID []string
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		//dicomID = append(dicomID, dcm.DicomID)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientRace = append(patientRace, dcm.PatientRace)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.patientGender)
	}

	var orgNoise, raceNoise, genderNoise, ageNoise, weigthNoise, heigthNoise map[string]float64

	type noise struct {
		DicomID             interface{} `json:"dicomID"`
		PatientAge          interface{} `json:"patientAge"`
		PatientOrganization interface{} `json:"patientOrganization"`
		PatientRace         interface{} `json:"patientRace"`
		PatientGender       interface{} `json:"patientGender"`
		PatientWeigth       interface{} `json:"patientWeigth"`
		PatientHeigth       interface{} `json:"patientHeigth"`
	}

	var dicomWithNoise noise

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	dicomWithNoise.DicomID = id.String()

	privOrg := dp.PrivateDataFactory(patientOrganization)
	auxPrivOrg, err := privOrg.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivOrg, &orgNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientOrganization = orgNoise

	privRace := dp.PrivateDataFactory(patientRace)
	auxPrivRace, err := privRace.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivRace, &raceNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientRace = raceNoise

	privGender := dp.PrivateDataFactory(patientGender)
	auxPrivGender, err := privGender.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivGender, &genderNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientRace = genderNoise

	privAge := dp.PrivateDataFactory(patientAge)
	auxPrivAge, err := privAge.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivAge, &ageNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientAge = ageNoise

	privWeigth := dp.PrivateDataFactory(patientWeigth)
	auxPrivWeigth, err := privWeigth.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivWeigth, &weigthNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientWeigth = weigthNoise

	privHeigth := dp.PrivateDataFactory(patientHeigth)
	auxPrivHeigth, err := privHeigth.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivHeigth, &heigthNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise.PatientWeigth = heigthNoise

	byteNoise, err := json.Marshal(dicomWithNoise)
	if err != nil {
		return nil, err
	}

	return byteNoise, nil
}

func applyKAnonymity(data dicom) ([]byte, error) {
	//fmt.Println("[Log] Assets that will be anonimized")

	allDicom, err := getAllMongo()
	if err != nil {
		return nil, err
	}
	//var allDicom []dicom

	// for _, ast := range assets {
	// 	allDicom = append(allDicom, ast)
	// }

	allDicom = append(allDicom, data)

	var patientFirstnames, patientLastnames, patientTelephones []string
	var patientOrganization, patientRace, patientID []string
	var patientInsuranceplan, machineModel []string
	var patientGender, patientAddress []string
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		patientID = append(patientID, dcm.PatientID)
		patientFirstnames = append(patientFirstnames, dcm.PatientFirstname)
		patientLastnames = append(patientLastnames, dcm.PatientLastname)
		patientTelephones = append(patientTelephones, dcm.PatientTelephone)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientRace = append(patientRace, dcm.PatientRace)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientInsuranceplan = append(patientInsuranceplan, dcm.PatientInsuranceplan)
		machineModel = append(machineModel, dcm.MachineModel)
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.patientGender)
	}

	anonymizedPatientID := kpriv.KAnonymitygeneralizationSymbolic(patientID)
	anonymizedFirstName := kpriv.KAnonymitygeneralizationSymbolic(patientFirstnames)
	anonymizedLastName := kpriv.KAnonymitygeneralizationSymbolic(patientLastnames)
	anonymizedTelephone := kpriv.KAnonymitygeneralizationSymbolic(patientTelephones)
	anonymizedOrganization := kpriv.KAnonymitygeneralizationSymbolic(patientOrganization)
	anonymizedRace := kpriv.KAnonymitygeneralizationSymbolic(patientRace)
	anonymizedAddress := kpriv.KAnonymitygeneralizationSymbolic(patientAddress)
	anonymizedGender := kpriv.KAnonymitygeneralizationSymbolic(patientGender)
	anonymizedAge := kpriv.KAnonymityGeneralizationNumeric(patientAge)
	anonymizedWeigth := kpriv.KAnonymityGeneralizationNumeric(patientWeigth)
	anonymizedHeigth := kpriv.KAnonymityGeneralizationNumeric(patientHeigth)
	anonymizedInsuranceplan := kpriv.KAnonymitygeneralizationSymbolic(patientInsuranceplan)
	anonimizedModelMachine := kpriv.KAnonymitygeneralizationSymbolic(machineModel)

	//var data []documentValue

	fmt.Println("[Log] amount assets to anonimized")

	var dicomNew dicomAux

	id, _ := uuid.NewRandom()
	dicomNew.DicomID = id.String()
	dicomNew.PatientID = anonymizedPatientID[len(allDicom)-1]
	dicomNew.PatientFirstname = anonymizedFirstName[len(allDicom)-1]
	dicomNew.PatientLastname = anonymizedLastName[len(allDicom)-1]
	dicomNew.PatientTelephone = anonymizedTelephone[len(allDicom)-1]
	dicomNew.PatientAddress = anonymizedAddress[len(allDicom)-1]
	dicomNew.PatientAge = anonymizedAge[len(allDicom)-1]
	dicomNew.PatientOrganization = anonymizedOrganization[len(allDicom)-1]
	dicomNew.PatientRace = anonymizedRace[len(allDicom)-1]
	dicomNew.PatientInsuranceplan = anonymizedInsuranceplan[len(allDicom)-1]
	dicomNew.PatientWeigth = anonymizedWeigth[len(allDicom)-1]
	dicomNew.PatientHeigth = anonymizedHeigth[len(allDicom)-1]
	dicomNew.MachineModel = anonimizedModelMachine[len(allDicom)-1]
	dicomNew.PatientGender = anonymizedGender[len(allDicom)-1]
	dicomNew.Timestamp = time.Now()

	// for i := len(allDicom) - 1; i >= (); i-- {
	// 	var dicomNew dicom

	// 	data = append(data, dicomNew)
	// }

	fmt.Println("[Log] Dataset Anonimized whitin K-Anonymnity function")
	//fmt.Println(dicomNew)

	dataValue, err := json.Marshal(dicomNew)

	if err != nil {
		return nil, err
	}

	return dataValue, nil
}

func getAllMongo() ([]dicom, error) {
	//Database local with mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("private")
	col := database.Collection("privData")
	var dcm []dicom
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &dcm); err != nil {
		return nil, err
	}

	return dcm, nil
}

func getRandomMongo() (dicom, error) {
	dcms, _ := getAllMongo()
	// err != nil{
	// 	return nil, err
	// }

	rd := rand.NewSource(time.Now().UnixNano())
	newRd := rand.New(rd)

	s := len(dcms)
	index := newRd.Intn(s - 1)

	return dcms[index], nil

}

func main() {

	var dcm []dicom
	for i := 0; i < 30; i++ {
		aux, _ := getRandomMongo()
		dcm = append(dcm, aux)
	}

	type dicomDiff struct {
		DicomID             string    `json:"dicomID"`
		PatientAge          int       `json:"patientAge"`
		PatientOrganization string    `json:"patientOrganization"`
		PatientRace         string    `json:"patientRace"`
		patientGender       string    `json:"patientGender"`
		PatientWeigth       float64   `json:"patientWeigth"`
		PatientHeigth       float64   `json:"patientHeigth"`
		Timestamp           time.Time `json:"timestamp"`
	}

	//var timeResp []string
	var data []interface{}
	fileTxt, _ := os.OpenFile("./resultados/time-priv-diff_.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fileTxt.Close()
	fileTxtEp, _ := os.OpenFile("./resultados/epsilon-priv-diff.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fileTxtEp.Close()

	episilon := []float64{0.0001, 0.001, 0.01, 1.0, 2.0, 2.5, 3.5, 4.5, 5.0, 6.0, 6.5, 7.0, 7.5, 8, 8.5, 9.0, 9.5, 10.0}

	for _, e := range episilon {
		for _, d := range dcm {
			tinit := time.Now()
			result, _ := applyDiffPriv(d, e)
			tfinal := time.Since(tinit)

			fileTxt.WriteString(tfinal.String() + "\n")
			str := strconv.FormatFloat(e, 'f', -1, 64)
			fileTxtEp.WriteString(str + "\n")

			var dataJs interface{}
			json.Unmarshal(result, &dataJs)
			data = append(data, dataJs)
		}
	}
	//writerTxt.WriteAll(timeResp)
	dataByte, _ := json.Marshal(data)
	originalByte, _ := json.Marshal(dcm)
	//timeByte, _ := json.Marshal(timeResp)
	ioutil.WriteFile("./resultados/data-priv-diff.json", dataByte, 0644)
	ioutil.WriteFile("./resultados/data-original-diff.json", originalByte, 0644)
	//writerJSON.WriteAll(data)

}