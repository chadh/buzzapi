package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

var jm = make(map[string]string)

func createJSONRequest() {
	// Build request parameter JSON map
	jm["api_app_id"] = env.APIUser
	jm["api_app_password"] = base64.StdEncoding.EncodeToString([]byte(env.APIPasswd))

	if *async {
		jm["api_request_mode"] = "async"
	} else {
		jm["api_request_mode"] = "sync"
	}

	jm["api_receive_timeout"] = fmt.Sprintf("%d", *timeout*1000)

	var myHandle string
	if *clienthandle != "" || *handle != "" {
		if *clienthandle != "" && *handle != "" {
			log.Warning("Ignoring clienthandle arg [%s]", *clienthandle)
		}
		if *handle != "" {
			myHandle = *handle
		} else {
			myHandle = *clienthandle
		}
	} else {
		var h, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		rand.Seed(time.Now().UTC().UnixNano())
		myHandle = fmt.Sprintf("from-%d@%s-rand%d", os.Getpid(), h, rand.Intn(32768))
	}
	jm["api_client_request_handle"] = myHandle

	jm["api_log_level"] = *loglevel

	if params != nil {
		for k, v := range *params {
			jm[k] = v
		}
	}
}

func postRequest() {
	var URL *url.URL
	URL, err := url.Parse(env.APIURL + fmt.Sprintf("/apiv3/%s/%s", res, op))
	if err != nil {
		log.Critical("Bad API URL")
		os.Exit(1)
	}

	o, err := json.Marshal(&jm)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("POST to %s, JSON Request: %s\n", URL, string(o))

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(o))
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Critical("%s", err)
	}
	defer resp.Body.Close()

	log.Info("response Status: %s", resp.Status)
	log.Info("response Headers: %s", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Critical("%s", err)
	}
	log.Info("response Body: %s", string(body))

	// FIXME now process the response body and find the api_result_data and print it out
	var g interface{}
	json.Unmarshal(body, &g)
	for k, v := range g.(map[string]interface{}) {
		switch k {
		case "api_result_data":
			switch vv := v.(type) {
			case string:
				fmt.Printf("%s\n", vv)
			case int:
				fmt.Printf("%d\n", vv)
			case []interface{}:
				m, err := json.MarshalIndent(vv, "", "  ")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", m)
			case map[string]interface{}:
				m, err := json.MarshalIndent(vv, "", "  ")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", m)
			}
		case "api_error_info":
			ei := v.(map[string]interface{})
			log.Error("%s", ei["message"])
		case "api_buzzapi_logs":
			//log.Info("api_buzzapi_logs: %v", v)
		case "api_provider_logs":
			//log.Info("api_provider_logs: %v", v)
		default:
			/*
				switch vv := v.(type) {
				case string:
					log.Debug("%s is string: %s", k, vv)
				case int:
					log.Debug("%s is int: %s", k, vv)
				case float64:
					log.Debug("%s is float64: %f", k, vv)
				case []interface{}:
					log.Debug("%s is an array:", k)
					for i, u := range vv {
						log.Debug("%d, %s", i, u)
					}
				case map[string]interface{}:
					log.Debug("%s is a map:", k)
					for key, val := range vv {
						log.Debug("%s, %v", key, val)
					}
				default:
					log.Debug("%s is of type %s: %v", k, reflect.TypeOf(v), vv)
				}
			*/
		}
	}
}
