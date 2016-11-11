package main

import (
	"net/http"
	"os"
	"os/exec"
	"testing"
)

var testFile string = "testfiles/NDBC-CMANWx.NDBC_51000_201509_D3_v00.xml"

func TestMain(m *testing.M) {
	code := m.Run()
	// curl -XDELETE http://localhost:9200/staging_v1/granule/NDBC-CMANWx.NDBC_51000_201509_D3_v00
	exec.Command("curl", "-XDELETE", "http://localhost:8097/onestop/api/metadata/NDBC-CMANWx.NDBC_51000_201509_D3_v00/").Output()
	os.Exit(code)
}

func TestPostFile(t *testing.T) {
	postFile(testFile)
	_, err := http.Get("http://localhost:8097/onestop/api/metadata/NDBC-CMANWx.NDBC_51000_201509_D3_v00/")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
}
