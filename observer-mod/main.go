package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"
)

var USER string
var lastObservation string
var SERVER string

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
	url := "http://35.211.104.239:3000/api/observerRequests"

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

	var reqPayload request
	reqPayload.User = USER
	reqPayload.Timestamp = lastObservation
	aux, _ := json.Marshal(reqPayload)
	payload := bytes.NewReader(aux)
	//payload := strings.NewReader(fmt.Sprintf("{\n\t\"user\": \"%s\", \n\t\"timestamp\": \"%s\"\n}", USER, lastObservation ))
	req, _ := http.NewRequest("GET", url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if len(body) == 0 {
		return
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
		// home, _ := os.UserHomeDir()
		// getIPFS(path.Join(home, "shared-dicom"), resp.IpfsReference)

		lastObservation = time.Now().Format(time.RFC3339)
	}

}

func getIPFS(path string, ref string) {
	sh := shell.NewShell("35.233.252.12:5001")

	if _, err := os.Stat(path); err != nil {
		os.Mkdir(path, os.ModePerm.Perm())
	}
	err := sh.Get(ref, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}

	fmt.Sprintf("[Log] sucess to get %s \n", path)
}

func callPrivacyKanonymity() []string {

	url := fmt.Sprintf("http://%s:5000/api/getPriv", SERVER)
	fmt.Println(url)
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	var resBody []dicom
	json.Unmarshal(body, &resBody)

	var dcmID []string

	url = "http://35.211.104.239:3000/api/addAsset"
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

	url := fmt.Sprintf("http://%s:5000/api/getPrivDiff", SERVER)
	fmt.Println(url)
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

	url = "http:/35.211.104.239:3000/api/addAssetDiff"
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
	url := "http://35.211.104.239:3000/api/notify"
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

func getObserve(w http.ResponseWriter, r *http.Request) {
	observeBlockchain(SERVER, "5000")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`Ok`))
}

func main() {

	USER = "982572dc-822d-4318-bc84-c4e4ad8d9c31"
	lastObservation = time.Now().Local().Format("2006-01-02")
	SERVER = "localhost"
	fmt.Println("Starting sentinel ...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/getObserve", getObserve)

	fmt.Println("Stating Server ...")

	log.Fatal(http.ListenAndServe(":6000", router))

	fmt.Println("End sentinel ...")
}
