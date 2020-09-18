package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func invalidGetMethod() {
	http.Get("http://localhost:8090/sms")
}

func validPostJSON(os string) {
	payload := map[string]interface{}{
		"deviceOS": os,
		"sms": map[string]string{
			"to":      "987654321",
			"from":    "998877665",
			"message": "SMS message to be sent",
		},
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error occured at the JSON.")
		panic(err)
	}
	fmt.Println(string(reqBody))
	executePost("sms", reqBody)
}
func invalidPostJSON(os string) {
	smsInvalidJSON := []byte(`{"deviceOS":` + os + `,"sms":{"to":"987654321","from":"998877665","message":"SMS }}`)
	executePost("sms", smsInvalidJSON)
}

func invalidPostDataJSON(os string) {
	invalidPayload := map[string]interface{}{
		"deviceOS": os,
		"sms": map[string]string{
			"name":    "NagaVenkatesh",
			"city":    "Madurai",
			"message": "SMS message to be sent",
		},
	}

	reqBody, err := json.Marshal(invalidPayload)
	if err != nil {
		log.Println("Error occured at the JSON.")
		panic(err)
	}
	executePost("sms", reqBody)
}

// func createRemoteDevice(os string) {
// 	reqBody := []byte(`{"deviceOS":"android"}`)
// 	executePost("create", reqBody)
// }

func main() {
	// os := "android"
	// createRemoteDevice("android")
	// 1. Invalid 'Get' method used
	invalidGetMethod()

	// 2. Valid JSON message payload
	// validPostJSON(os)

	// 3. Invalid JSON message payload
	// invalidPostJSON(os)

	// 4. Invalid JSON message payload
	// invalidPostDataJSON(os)

}
