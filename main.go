package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Code      int    `json:"code"`
	Success   bool   `json:"success"`
	Jwt       string `json:"jwt"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	var username string
	var password string
	var apiHost string

	flag.StringVar(&username, "username", "", "Username (required)")
	flag.StringVar(&password, "password", "", "Password (required)")
	flag.StringVar(&apiHost, "api_host", "", "API host (required)")
	flag.Parse()

	if username == "" || password == "" || apiHost == "" {
		flag.PrintDefaults()
		return
	}

	client := &http.Client{}
	reqBody := fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
	}`, username, password)

	req, _ := http.NewRequest("POST", apiHost, strings.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode == 200 {
		var data Data
		err := json.Unmarshal(bodyBytes, &data)
		if err != nil {
			log.Fatal(err)
		}
		if data.Success {
			fmt.Println(data.Jwt)
		} else {
			fmt.Println("Response data.success is not true")
			fmt.Println("Response data:", bodyString)
		}
	} else {
		fmt.Println("Response status code is not 200")
		fmt.Println("Response status:", resp.StatusCode)
		fmt.Println("Response data:", bodyString)
	}

	return
}
