package common

import (
	by "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//GetRequest GetRequest
func GetRequest(url string, method string, obj *[]byte) (*http.Request, bool) {
	//fmt.Print("obj: ")
	//fmt.Println(obj)
	var err = false
	var bf *by.Buffer
	var req *http.Request
	var rErr error
	if obj != nil {
		//fmt.Println("in first: ")
		bf = by.NewBuffer(*obj)
		req, rErr = http.NewRequest(method, url, bf)
		if rErr != nil {
			err = true
			fmt.Print("request err: ")
			fmt.Println(rErr)
		}
	} else {
		//fmt.Println("in second: ")
		req, rErr = http.NewRequest(method, url, nil)
		if rErr != nil {
			err = true
			fmt.Print("request err: ")
			fmt.Println(rErr)
		}
	}
	//fmt.Println(bf)
	return req, err
}

//ProcessRespose ProcessRespose
func ProcessRespose(resp *http.Response, obj interface{}) bool {
	var rtn bool
	//fmt.Print("resp in processResponse: ")
	//fmt.Println(resp)
	if resp != nil {
		//fmt.Print("resp body: ")
		//fmt.Println(resp.Body)
		decoder := json.NewDecoder(resp.Body)
		var err error
		if obj != nil {
			err = decoder.Decode(obj)
			//fmt.Print("response obj: ")
			//fmt.Println(obj)
		}
		if err != nil {
			fmt.Print("response err: ")
			fmt.Println(err)
		} else {
			rtn = true
		}
	} else {
		log.Println("response = nil in processResponse")
	}
	return rtn
}

//GetJSONEncode GetJSONEncode
func GetJSONEncode(obj interface{}) *[]byte {
	//fmt.Print("obj in json: ")
	//fmt.Println(obj)
	aJSON, _ := json.Marshal(obj)
	return &aJSON
}

//ProcessServiceCall ProcessCall
func ProcessServiceCall(req *http.Request, obj interface{}) int {
	var code int
	client := &http.Client{}
	resp, cErr := client.Do(req)
	//fmt.Print("resp: ")
	//fmt.Println(resp)
	if cErr != nil {
		fmt.Print("Service err: ")
		fmt.Println(cErr)
		code = http.StatusBadRequest
	} else {
		defer resp.Body.Close()
		ProcessRespose(resp, obj)
		code = resp.StatusCode
	}
	return code
}
