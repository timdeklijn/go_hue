package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// url to communicate with light 6 in the house.
const url = "https://192.168.178.22/api/73rtx0fU6CMNysLyU1QkyRf7pvGcEupIL38i982-/lights/6"

// State describes the state of a light
type State struct {
	On     bool       `json:"on"`
	Bri    int        `json:"bri"`
	Hue    int        `json:"hue"`
	Sat    int        `json:"sat"`
	Effect string     `json:"effect"`
	Xy     [2]float64 `json:"xy"`
}

// Light is the main response which requesting info on a light.
type Light struct {
	State State `json:"state"`
}

// OnState is used to turn on or of the light.
type OnState struct {
	On bool `json:"on"`
}

func createCLient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	return client
}

func parseLightResponse(resp http.Response) Light {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var light Light
	json.Unmarshal(body, &light)
	return light
}

func getLightState(client http.Client) State {
	// Request the state of the light
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the json to a Light struct.
	var light Light
	json.Unmarshal(body, &light)
	return light.State
}

func switchOnState(client http.Client) {
	// Request the current state of the light and create a json with the opposite of
	// that state.
	currentState := getLightState(client)
	onState := OnState{On: !currentState.On}
	r, err := json.Marshal(onState)
	if err != nil {
		panic(err)
	}

	// Create a http request.
	req, err := http.NewRequest(http.MethodPut, url+"/state", bytes.NewBuffer(r))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json and send out the request.
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}

func listLightInfo(client http.Client) {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Parsing Body")
	light := parseLightResponse(*resp)
	fmt.Println(light)
}

func main() {
	client := createCLient()
	listLightInfo(client)
	switchOnState(client)
}
