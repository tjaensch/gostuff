package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	csvfile string = "dsmm_assessments.csv"
)

// For character values see CSV input file
type DsmmAssessmentRecord struct {
	A  string
	B  string
	C  string
	D  string
	E  string
	F  string
	G  string
	H  string
	I  string
	J  string
	K  string
	L  string
	M  string
	N  string
	O  float32
	P  float32
	Q  float32
	R  float32
	S  float32
	T  float32
	U  float32
	V  float32
	W  float32
	X  string
	Y  string
	Z  string
	AA string
	AB string
	AC string
	AD string
	AE string
	AF string
	AG string
	AH string
	AI string
	AJ string
	AK string
	AL string
	AM string
	AN string
	AO string
	AP string
	AQ string
	AR string
	AS string
	AT string
	AU string
	AV string
	AW string
	AX string
	AY string
	AZ string
}

func main() {

	csvfile, err := os.Open(csvfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord DsmmAssessmentRecord

	var allRecords []DsmmAssessmentRecord

	for _, cell := range rawCSVdata {
		oneRecord.A = cell[0]
		oneRecord.B = cell[1]
		oneRecord.C = cell[2]
		oneRecord.D = cell[3]
		allRecords = append(allRecords, oneRecord)
	}

	fmt.Printf("Timestamp: %s,\nusername: %s,\ndatasetshortname: %s,\ndatasettitle: %s\n", allRecords[2].A, allRecords[2].B, allRecords[2].C, allRecords[2].D)

}
