package main

import (
	//"io/ioutil"
	//"os"
    "reflect"
	"testing"
)

var testFile string = "AGE00147710"

func TestReadInStationsFileInfo(t *testing.T) {
	_, latMap, lonMap, stationLongNameMap := readInStationsFileInfo()
	if len(latMap) != len(lonMap) {
		t.Error("len(latMap) doesn't match len(lonMap)")
	}
	if len(stationLongNameMap) < 103000 {
		t.Error("len(stationLongNameMap) not right, got", len(stationLongNameMap))
	}
}

func TestGetIndividualDataFile(t *testing.T) {
	expected := getIndividualDataFile(testFile)
    if reflect.TypeOf(expected).Kind() != reflect.String {
    	t.Error("expected string type, got", reflect.TypeOf(expected))
    }
}