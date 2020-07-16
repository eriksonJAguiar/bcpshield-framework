package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var USER string

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

	type request struct {
		User      string `json:"user"`
		Timestamp string `json:"timestamp"`
	}

	type sharedDicom struct {
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

	lastObseration := time.Now().Format("2020-07-15")
	for listen {
		var reqPayload request
		reqPayload.User = USER
		reqPayload.Timestamp = lastObseration
		aux, _ := json.Marshal(reqPayload)
		payload := bytes.NewReader(aux)
		//payload := strings.NewReader(fmt.Sprintf("{\n\t\"user\": \"%s\", \n\t\"timestamp\": \"%s\"\n}", USER, lastObseration))
		req, _ := http.NewRequest("GET", url, payload)
		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if len(body) == 0 {
			continue
		}
		var resBody string
		json.Unmarshal(body, &resBody)

		var result []sharedDicom
		json.Unmarshal([]byte(resBody), &result)

		for _, resp := range result {
			userType := resp.AccessLevel
			reqID := resp.BatchID
			if userType == "Doctor" {
				dcmID := callPrivacyKanonymity()
				notifyRequester(reqID, dcmID)
			} else if userType == "Researcher" {
				dcmID := callDiffPrivacy()
				notifyRequester(reqID, dcmID)
			}
		}
		//lastObseration := time.Now()
		time.Sleep(5 * 60)
		lastObseration = time.Now().Format(time.RFC3339)
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
		dcm.User = USER
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

	type request struct {
		User    string `json:"user"`
		DicomID string `json:"dicomID"`
		Asset   string `json:"asset"`
	}

	var dcmID []string

	url = "http://35.211.244.95:3000/api/addAssetDiff"
	for _, dcm := range resBody {
		//var auxDcm dicom
		var req request
		req.DicomID = dcm.DicomID.(string)
		req.User = USER
		auxDcm, _ := json.Marshal(dcm)
		req.Asset = string(auxDcm[:])
		auxReq, _ := json.Marshal(req)
		byteReq := bytes.NewReader(auxReq)
		http.Post(url, "application/json", byteReq)
		dcmID = append(dcmID, req.DicomID)
	}

	return dcmID
}

func notifyRequester(reqID string, dcmIDs []string) {
	url := "http://35.211.244.95:3000/api/notify"
	type request struct {
		User      string `json:"user"`
		RequestID string `json:"requestID"`
		Assets    string `json:"assets"`
	}
	for _, id := range dcmIDs {
		var notify request
		notify.User = USER
		notify.RequestID = reqID
		notify.Assets = id

		aux, _ := json.Marshal(notify)
		byteReq := bytes.NewReader(aux)
		http.Post(url, "application/json", byteReq)
	}
}

func main() {
	//fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println("Starting sentinel ...")
	USER = "97921473-d222-4dc4-8bf8-6ade5272504a"
	observeBlockchain("127.0.0.1", "5000")
	//callDiffPrivacy()
	//notifyRequester("7d6b9b42-788f-44dd-9d66-83a76f390b2c", []string{"80e74659-3814-42a2-8d9b-94f06a11892d"})
	fmt.Println("End sentinel ...")
}
