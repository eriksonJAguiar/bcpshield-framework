package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	kpriv "privacy-mod/kanonymity"

	dp "github.com/eriksonJAguiar/godiffpriv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define the DICOM imaging attributres
type dicom struct {
	DicomID              string    `bson:"dicomID"`
	PatientID            string    `bson:"patientID"`
	DocType              string    `bson:"docType"`
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

func addMong(dcm dicom) bool {
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
	_, err = col.InsertOne(ctx, dcm)
	if err != nil {
		return false
	}

	return true
}

func getOneMongo(dcmID string) (dicom, error) {
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
	var dcm dicom
	err = col.FindOne(ctx, bson.M{"dicomID": dcmID}).Decode(&dcm)
	if err != nil {
		log.Fatal(err)
	}
	// byteResp, err := json.Marshal(dcm)
	// if err != nil {
	// 	return nil, err
	// }

	return dcm, nil
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

func addData(w http.ResponseWriter, r *http.Request) {
	//API Config
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resp dicom
	err := json.Unmarshal(reqBody, &resp)
	if err != nil {
		return
	}

	if !addMong(resp) {
		return
	}

}

func getData(w http.ResponseWriter, r *http.Request) {
	//API get values
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resp map[string]string
	err := json.Unmarshal(reqBody, &resp)
	if err != nil {
		return
	}

	dicomValue, err := getOneMongo(resp["dicomID"])
	if err != nil {
		return
	}

	byteResp, err := json.Marshal(dicomValue)
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteResp)
}

func getDataKanonynmity(w http.ResponseWriter, r *http.Request) {
	//API get values
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resp map[string]string
	err := json.Unmarshal(reqBody, &resp)
	if err != nil {
		return
	}

	dcm, err := getOneMongo(resp["dicomID"])
	if err != nil {
		return
	}

	privDcm, err := applyKAnonymity(dcm)
	if err != nil {
		return
	}

	// privByte, err := json.Marshal(privDcm)
	// if err != nil {
	// 	return
	// }

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(privDcm)
}

func applyKAnonymity(data dicom) ([]byte, error) {
	//fmt.Println("[Log] Assets that will be anonimized")

	assets, err := getAllMongo()
	if err != nil {
		return nil, err
	}
	var allDicom []dicom

	for _, ast := range assets {
		allDicom = append(allDicom, ast)
	}

	allDicom = append(allDicom, data)

	var patientFirstnames, patientLastnames, patientTelephones []string
	var patientOrganization, patientRace, dicomID, patientID []string
	var patientInsuranceplan, machineModel []string
	var patientGender, patientAddress []string
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		dicomID = append(dicomID, dcm.DicomID)
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

	anonymizedDicomID := kpriv.KAnonymitygeneralizationSymbolic(dicomID)
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

	type dicomAux struct {
		DicomID              string    `bson:"dicomID"`
		PatientID            string    `bson:"patientID"`
		DocType              string    `bson:"docType"`
		PatientFirstname     string    `bson:"patientFirstname"`
		PatientLastname      string    `bson:"patientLastname"`
		PatientTelephone     string    `bson:"patientTelephone"`
		PatientAddress       string    `bson:"patientAddress"`
		PatientAge           string    `bson:"patientAge"`
		PatientOrganization  string    `bson:"patientOrganization"`
		PatientRace          string    `bson:"patientRace"`
		patientGender        string    `bson:"patientGender"`
		PatientInsuranceplan string    `bson:"patientInsuranceplan"`
		PatientWeigth        string    `bson:"patientWeigth"`
		PatientHeigth        string    `bson:"patientHeigth"`
		MachineModel         string    `bson:"machineModel"`
		Timestamp            time.Time `bson:"timestamp"`
	}

	var dicomNew dicomAux

	dicomNew.DicomID = anonymizedDicomID[len(allDicom)-1]
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
	dicomNew.patientGender = anonymizedGender[len(allDicom)-1]
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

func getDataDiffPriv(w http.ResponseWriter, r *http.Request) {
	//API get values
	reqBody, _ := ioutil.ReadAll(r.Body)
	var resp map[string]string
	err := json.Unmarshal(reqBody, &resp)
	if err != nil {
		return
	}

	dcm, err := getOneMongo(resp["dicomID"])
	if err != nil {
		panic(err)
	}

	privDcm, err := applyDiffPriv(dcm, 1.0)
	if err != nil {
		panic(err)
	}

	// privByte, err := json.Marshal(privDcm)
	// if err != nil {
	// 	return
	// }

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(privDcm)
}

func applyDiffPriv(data dicom, epsilon float64) ([]byte, error) {

	var allDicom []dicom
	allDicom, err := getAllMongo()
	if err != nil {
		return nil, err
	}

	var patientOrganization, patientRace []string
	var patientGender, patientAddress []string
	var dicomID []string
	var patientWeigth, patientHeigth, patientAge []float64

	for _, dcm := range allDicom {
		dicomID = append(dicomID, dcm.DicomID)
		patientOrganization = append(patientOrganization, dcm.PatientOrganization)
		patientRace = append(patientRace, dcm.PatientRace)
		patientAddress = append(patientAddress, dcm.PatientAddress)
		patientAge = append(patientAge, float64(dcm.PatientAge))
		patientWeigth = append(patientWeigth, dcm.PatientWeigth)
		patientHeigth = append(patientHeigth, dcm.PatientHeigth)
		patientGender = append(patientGender, dcm.patientGender)
	}
	dicomWithNoise := make(map[string]interface{})
	var orgNoise, raceNoise, genderNoise, ageNoise, weigthNoise, heigthNoise map[string]float64

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

	privGender := dp.PrivateDataFactory(patientGender)
	auxPrivGender, err := privGender.ApplyPrivacy(epsilon)
	err = json.Unmarshal(auxPrivGender, &genderNoise)
	if err != nil {
		return nil, err
	}
	dicomWithNoise["Gender"] = genderNoise

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

	byteNoise, err := json.Marshal(dicomWithNoise)
	if err != nil {
		return nil, err
	}

	return byteNoise, nil
}

func api() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/add", addData)
	router.HandleFunc("/api/get", getData)
	router.HandleFunc("/api/getPriv", getDataKanonynmity)
	router.HandleFunc("/api/getPrivDiff", getDataDiffPriv)
	fmt.Println("Stating Server ...")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	api()
}
