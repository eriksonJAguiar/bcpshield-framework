package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomtag"
	shell "github.com/ipfs/go-ipfs-api"
)

var pathTo string
var pathFrom string
var cids []string

func get(w http.ResponseWriter, r *http.Request) {
	sh := shell.NewShell("35.233.252.12:5001")
	//cid, err := sh.Add(strings.NewReader("hello world!"))
	cid := r.URL.Query().Get("id")
	fmt.Println("[Log] Get Element")
	if _, err := os.Stat(pathTo); err != nil {
		os.Mkdir(pathTo, os.ModePerm.Perm())
	}
	err := sh.Get(cid, pathTo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}

	fmt.Sprintf("[Log] sucess to get %s \n", pathTo)
	//insertTokens()
	cids = append(cids, cid)
	removeAll()
	w.WriteHeader(200)
	w.Write([]byte(`OK`))
}

func add(w http.ResponseWriter, r *http.Request) {
	sh := shell.NewShell("35.233.252.12:5001")
	//cid, err := sh.Add(strings.NewReader("hello world!"))
	fmt.Println("[Log] Add Element")
	cid, err := sh.AddDir(pathFrom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return
	}
	fmt.Printf("[Log] added in IPFS, with CID = %s", cid)
	//cids = append(cids, cid)
	//return cid
	file, _ := os.OpenFile("result.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	file.WriteString(cid + "\n")
	defer file.Close()
	w.WriteHeader(200)
	//byteCid, _ := json.Marshal(cid)
	w.Write([]byte(`Ok`))
}

func removeAll() {
	os.RemoveAll(pathTo)
}

func getFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var fileName []string

	for _, f := range files {
		fileName = append(fileName, f.Name())
	}

	return fileName
}

func insertTokens() {
	files := getFiles(pathTo)
	for _, f := range files {
		aux, _ := uuid.NewRandom()
		token := aux.String()
		ds, err := dicom.ReadDataSetFromFile(fmt.Sprintf("%s/%s", pathTo, f), dicom.ReadOptions{})
		if err != nil {
			fmt.Sprintf("[ERROR] %s \n", err.Error())
			continue
		}
		// dicomtag.
		tokenID, err := ds.FindElementByTag(dicomtag.PatientID)
		if err != nil {
			fmt.Sprintf("[ERROR] %s \n", err.Error())
			continue
		}
		tokenID.Value = []interface{}{token}

		buf := bytes.Buffer{}
		if err := dicom.WriteDataSet(&buf, ds); err != nil {
			panic(err)
		}

		ds2, err := dicom.ReadDataSet(&buf, dicom.ReadOptions{})
		if err != nil {
			panic(err)
		}
		tokenID, err = ds2.FindElementByTag(dicomtag.PatientID)
		if err != nil {
			panic(err)
		}
		fmt.Println("[Log] Token: " + tokenID.String())

	}
}

func main() {
	home, _ := os.UserHomeDir()
	pathFrom = filepath.Join(home, "dicom-files")
	//fmt.Sprintf("%s/dicom-files", home)
	//pathTo = fmt.Sprintf("%s/dicom-share", home)
	pathTo = filepath.Join(home, "dicom-share")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ipfs/add", add)
	router.HandleFunc("/ipfs/get", get)
	//router.HandleFunc("/ipfs/run", experiments)

	fmt.Println("Stating Server ...")
	log.Fatal(http.ListenAndServe(":7001", router))
}
