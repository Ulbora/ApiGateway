package handlers

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

func processResponse(resp *http.Response) ([]byte, error) {
	var respbody []byte
	var bdyErr error
	//fmt.Print("Content-Encoding header: ")
	//fmt.Println(resp.Header.Get("Content-Encoding"))
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		//fmt.Println("found body to be gzip")
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Print("gzip error: ")
			fmt.Println(err)
		}
		defer gz.Close()
		resp.Header.Del("Content-Encoding")
		respbody, bdyErr = ioutil.ReadAll(gz)
	case "deflate":
		//fmt.Println("found body to be deflate")
		fz := flate.NewReader(resp.Body)
		defer fz.Close()
		resp.Header.Del("Content-Encoding")
		respbody, bdyErr = ioutil.ReadAll(fz)
	default:
		respbody, bdyErr = ioutil.ReadAll(resp.Body)
	}
	return respbody, bdyErr
}
