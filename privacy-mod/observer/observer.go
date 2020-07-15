package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type dicom struct {
	DicomID              string  `json:"dicomID"`
	PatientID            string  `json:"patientID"`
	PatientFirstname     string  `json:"patientFirstname"`
	PatientLastname      string  `json:"patientLastname"`
	PatientTelephone     string  `json:"patientTelephone"`
	PatientAddress       string  `json:"patientAddress"`
	PatientAge           int     `json:"patientAge"`
	PatientOrganization  string  `json:"patientOrganization"`
	PatientRace          string  `json:"patientRace"`
	PatientGender        string  `json:"patientGender"`
	PatientInsuranceplan string  `json:"patientInsuranceplan"`
	PatientWeigth        float64 `json:"patientWeigth"`
	PatientHeigth        float64 `json:"patientHeigth"`
	MachineModel         string  `json:"machineModel"`
	User                 string  `json:"user"`
}

func observeBlockchain(ip string, port string) {
	//cli := httputil.NewClientConn()
	listen := true
	url := "http://35.211.244.95:3000/api/observerRequests"
	for listen {
		lastObseration := time.Now()

		payload := strings.NewReader(fmt.Sprintf("{\n\t\"user\": \"0590d195-6669-4823-8dc4-1f3118acc57d\", \n\t\"timestamp\": \"%s\"\n}", lastObseration))
		req, _ := http.NewRequest("GET", url, payload)
		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		var resBody map[string]string
		json.Unmarshal(body, &resBody)

		userType := resBody["accessLevel"]
		reqID := resBody["batchID"]
		if userType == "Doctor" {
			dcmID := callPrivacyKanonymity()
			notifyRequester(reqID, dcmID)
		} else if userType == "Researcher" {

		}
		time.Sleep(100)
		//lastObseration = time.Now().Format(time.RFC3339)
	}

}

func callPrivacyKanonymity() []string {

	url := "http://localhost:5000/api/getPriv"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	var resBody []dicom
	json.Unmarshal(body, &resBody)

	var dcmID []string

	url = "http://35.211.244.95:3000/api/addAsset"
	for _, dcm := range resBody {
		dcm.User = "0590d195-6669-4823-8dc4-1f3118acc57d"
		aux, _ := json.Marshal(dcm)
		byteDcm := bytes.NewReader(aux)
		http.Post(url, "application/json", byteDcm)
		dcmID = append(dcmID, dcm.DicomID)
	}

	return dcmID
}

func callDiffPrivacy() []string {

	url := "http://localhost:5000/api/getPrivDiff"
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)

	type noise struct {
		DicomID             interface{} `json:"dicomID"`
		PatientAge          interface{} `json:"patientAge"`
		PatientOrganization interface{} `json:"patientOrganization"`
		PatientRace         interface{} `json:"patientRace"`
		PatientGender       interface{} `json:"patientGender"`
		PatientWeigth       interface{} `json:"patientWeigth"`
		PatientHeigth       interface{} `json:"patientHeigth"`
	}

	var resBody []noise
	json.Unmarshal(body, &resBody)

	var dcmID []string

	// url = "http://35.211.244.95:3000/api/addAsset"
	// for _, dcm := range resBody {
	// 	var auxDcm dicom
	// 	auxDcm.User = "0590d195-6669-4823-8dc4-1f3118acc57d"
	// 	aux, _ := json.Marshal(dcm)
	// 	byteDcm := bytes.NewReader(aux)
	// 	http.Post(url, "application/json", byteDcm)
	// 	dcmID = append(dcmID, auxDcm.DicomID)
	// }

	return dcmID
}

func notifyRequester(reqID string, dcmIDs []string) {
	url := "http://35.211.244.95:3000/api/notify"
	for _, id := range dcmIDs {
		var notify map[string]string
		notify["user"] = "0590d195-6669-4823-8dc4-1f3118acc57d"
		notify["requestID"] = reqID
		notify["assets"] = id

		aux, _ := json.Marshal(notify)
		byteReq := bytes.NewReader(aux)
		http.Post(url, "application/json", byteReq)
	}
}

func main() {
	//fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println("Starting sentinel ...")
	//observeBlockchain("127.0.0.1", "5000")
	callDiffPrivacy()
	fmt.Println("End sentinel ...")
}
