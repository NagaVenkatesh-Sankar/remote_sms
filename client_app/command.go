package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url         = "http://localhost:8090/"
	contentType = "application/json"
)

func executePost(command string, payload []byte) {

	// http.Post("http://localhost:8090/sms", "application/json", bytes.NewBuffer(reqBody))
	resp, err := http.Post(url+command, contentType, bytes.NewBuffer(payload))

	if err != nil {
		log.Println("Error occured. Please check the connection or the Server is not reachable.")
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	//fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if err != nil {
		panic(err)
	}
}
