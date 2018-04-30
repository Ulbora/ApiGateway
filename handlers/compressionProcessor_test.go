package handlers

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"time"
	//"net/http"
	"net/http/httptest"
	//"net/http"
	//"net/http/httptest"
	//"reflect"
	"compress/gzip"
	"testing"
)

func Test_processResponse_gzip(t *testing.T) {
	w := httptest.NewRecorder()

	w.Header().Set("Content-Encoding", "gzip")
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Name = "atest.txt"
	zw.Comment = "A test"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("this is a test"))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	w.Body = &buf

	bdy, err := processResponse(w.Result())
	fmt.Print("bdy: ")
	fmt.Println(bdy)
	fmt.Print("err: ")
	fmt.Println(err)
	if bdy == nil || err != nil {
		t.Fail()
	}
}

func recoverZip() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func Test_processResponse_gzipBody(t *testing.T) {
	w := httptest.NewRecorder()
	defer recoverZip()
	w.Header().Set("Content-Encoding", "gzip")
	var buf bytes.Buffer
	w.Body = &buf
	bdy, err := processResponse(w.Result())
	fmt.Print("bdy: ")
	fmt.Println(bdy)
	fmt.Print("err: ")
	fmt.Println(err)
	if bdy == nil || err != nil {
		t.Fail()
	}
}

func Test_processResponse_deflate(t *testing.T) {
	w := httptest.NewRecorder()

	w.Header().Set("Content-Encoding", "deflate")

	var b bytes.Buffer

	// Compress the data using the specially crafted dictionary.
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte("this is a test"))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	w.Body = &b
	bdy, err := processResponse(w.Result())
	fmt.Print("bdy: ")
	fmt.Println(bdy)
	fmt.Print("err: ")
	fmt.Println(err)
	if bdy == nil || err != nil {
		t.Fail()
	}
}

func Test_processResponse(t *testing.T) {
	w := httptest.NewRecorder()
	w.WriteString("test")
	bdy, err := processResponse(w.Result())
	fmt.Print("bdy: ")
	fmt.Println(bdy)
	fmt.Print("err: ")
	fmt.Println(err)
	if bdy == nil || err != nil {
		t.Fail()
	}
}
