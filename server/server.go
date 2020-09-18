package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type payload struct {
	DeviceOS  string    `json:"deviceOS" validate:"required"`
	SmsSchema smsSchema `json:"sms"`
}
type smsSchema struct {
	From    string `json:"from" validate:"required"`
	To      string `json:"to" validate:"required"`
	Message string `json:"message"`
}

func sms(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		fmt.Fprintf(w, "Method not allowed to use")
		return
	}

	// Read request body.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, err.Error())
		return
	}

	// Parse body as json.
	var payloadData payload
	if err = json.Unmarshal(body, &payloadData); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		fmt.Fprintf(w, err.Error())
		return
	}

	// validate the JSON data fields
	validate := validator.New()
	if err = validate.Struct(payloadData); err != nil {
		log.Printf("All the required fields are not available, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		fmt.Fprintf(w, err.Error())
		return
	}
	log.Println(payloadData)

	// Create the device type
	currentDevice, err := CreateDevice(payloadData.DeviceOS)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Send the sms
	smsContext := payloadData.SmsSchema
	_, err = currentDevice.SendSms(smsContext)
	if err != nil {
		log.Printf("Issue with the SMS module, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SMS Sent successfully.")
}

var currentDevice Device

func main() {

	http.HandleFunc("/sms", sms)

	http.ListenAndServe(":8090", nil)
}
