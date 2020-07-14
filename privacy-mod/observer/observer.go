package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func observeBlockchain(ip string, port string) {
	//cli := httputil.NewClientConn()
	listen := true
	for listen {
		lastObseration := time.Now().Format(time.RFC3339)
		reqBody, err := json.Marshal(map[string]string{
			"holderID":  "282802",
			"timestamp": lastObseration,
		})
		if err != nil {
			log.Fatalln(err)
		}
		//var cli http.Clien
		http.Get(fmt.Sprintf("http://%s:%s/api/observer", ip, port), "application/json", bytes.NewBuffer(reqBody))
		time.Sleep(5)
		//lastObseration = time.Now().Format(time.RFC3339)
	}

}

func notifyRequester() {

}

func main() {
	//fmt.Println(time.Now().Format(time.RFC3339))
}
