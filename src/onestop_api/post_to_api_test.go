package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

var testFiles []string = []string{
	"testfiles/NDBC-CMANWx.NDBC_51000_201509_D3_v00.xml",
	"testfiles/41010_199101.xml",
	"testfiles/46013_198908.xml",
	"testfiles/CO-OPS.NOS_8447387_201309_D1_v00.xml",
	"testfiles/NDBC_42036_201608_D3_v00.xml",
	"testfiles/NOS_9076070_201312_D1_v00.xml",
	"testfiles/woa13_8594_t05_04.xml",
	"testfiles/WTDL_20150803v10001.xml",
	"testfiles/WTEE_20150928v20001.xml",
	"testfiles/WTEY_20070816v20001.xml",
}

func TestPostFile(t *testing.T) {
	for _, testFile := range testFiles {
		postFile(testFile)
	}

	resp, _ := http.Get("http://localhost:9200/_search?pretty")
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	for _, testFile := range testFiles {
		if !strings.Contains(string(data), strings.TrimSuffix(filepath.Base(testFile), ".xml")) {
			t.Fatalf("Expected http://localhost:9200/_search?pretty to contain %s", strings.TrimSuffix(filepath.Base(testFile), ".xml"))
		}
	}
}
