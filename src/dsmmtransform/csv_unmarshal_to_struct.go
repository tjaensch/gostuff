package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

var (
	csvfile string = "DSMM_CSV_input_file/dsmm_assessments.csv"
)

// Struct that represents values of one line of CSV input file
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
	O  string
	P  string
	Q  string
	R  string
	S  string
	T  string
	U  string
	V  string
	W  string
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
}

// Read CSV data into struct slice
func getCsvData() []DsmmAssessmentRecord {
	csvfile, err := os.Open(csvfile)
	checkError("open csv file failed", err)
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	checkError("reading csv data failed", err)

	// Initialize struct with default values
	singleRecord := DsmmAssessmentRecord{A: "N/A", B: "N/A", AA: "N/A", AD: "N/A"}

	var allRecords []DsmmAssessmentRecord

	for _, cell := range rawCSVdata {
		singleRecord.A = cell[0]
		singleRecord.B = cell[1]
		singleRecord.C = cell[2]
		singleRecord.D = cell[3]
		singleRecord.E = cell[4]
		singleRecord.F = cell[5]
		singleRecord.G = cell[6]
		singleRecord.H = cell[7]
		singleRecord.I = cell[8]
		singleRecord.J = cell[9]
		singleRecord.K = cell[10]
		singleRecord.L = cell[11]
		singleRecord.M = cell[12]
		singleRecord.N = cell[13]
		singleRecord.O = cell[14]
		singleRecord.P = cell[15]
		singleRecord.Q = cell[16]
		singleRecord.R = cell[17]
		singleRecord.S = cell[18]
		singleRecord.T = cell[19]
		singleRecord.U = cell[20]
		singleRecord.V = cell[21]
		singleRecord.W = cell[22]
		singleRecord.X = cell[23]
		singleRecord.Y = cell[24]
		singleRecord.Z = cell[25]
		singleRecord.AA = cell[26]
		singleRecord.AB = cell[27]
		singleRecord.AC = cell[28]
		singleRecord.AD = cell[29]
		singleRecord.AE = cell[30]
		singleRecord.AF = cell[31]
		singleRecord.AG = cell[32]
		singleRecord.AH = cell[33]
		singleRecord.AI = cell[34]
		singleRecord.AJ = cell[35]
		singleRecord.AK = cell[36]
		singleRecord.AL = cell[37]
		singleRecord.AM = cell[38]
		singleRecord.AN = cell[39]
		singleRecord.AO = cell[40]
		singleRecord.AP = cell[41]
		singleRecord.AQ = cell[42]
		singleRecord.AR = cell[43]
		singleRecord.AS = cell[44]
		singleRecord.AT = cell[45]
		singleRecord.AU = cell[46]
		singleRecord.AV = cell[47]
		singleRecord.AW = cell[48]
		singleRecord.AX = cell[49]
		singleRecord.AY = cell[50]

		// Append singleRecord to allRecords struct slice
		allRecords = append(allRecords, singleRecord)
	}

	// Print all values from one record to stdout just for quick check that it works
	v := reflect.ValueOf(allRecords[763])
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	fmt.Println(values)

	return allRecords
} // end getCsvData()
