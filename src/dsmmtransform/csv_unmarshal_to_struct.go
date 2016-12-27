package main

import (
	"encoding/csv"
	//"fmt"
	"os"
	//"reflect"
)

// Struct that represents values of one line of CSV input file
type DsmmAssessmentRecord struct {
	A           string
	B           string
	C           string
	D           string
	E           string
	F           string
	G           string
	H           string
	I           string
	J           string
	K           string
	L           string
	M           string
	N           string
	O           string
	P           string
	Q           string
	R           string
	S           string
	T           string
	U           string
	V           string
	W           string
	X           string
	Y           string
	Z           string
	AA          string
	AB          string
	AC          string
	AD          string
	AE          string
	AF          string
	AG          string
	AH          string
	AI          string
	AJ          string
	AK          string
	AL          string
	AM          string
	AN          string
	AO          string
	AP          string
	AQ          string
	AR          string
	AS          string
	AT          string
	AU          string
	AV          string
	AW          string
	AX          string
	AY          string
	AZ          string
	BA          string
	BB          string
	BC          string
	BD          string
	BE          string
	BF          string
	BG          string
	BH          string
	
	Stardefault string
	Lightgrey   string
	Fullgrey    string
	Star1       string
	Star2       string
	Star3       string
	Star4       string
	Star5       string
	Star6       string
	Star7       string
	Star8       string
	Star9       string
	Star10      string
	Star11      string
	Star12      string
	Star13      string
	Star14      string
	Star15      string
	Star16      string
	Star17      string
	Star18      string
	Star19      string
	Star20      string
	Star21      string
	Star22      string
	Star23      string
	Star24      string
	Star25      string
	Star26      string
	Star27      string
	Star28      string
	Star29      string
	Star30      string
	Star31      string
	Star32      string
	Star33      string
	Star34      string
	Star35      string
	Star36      string
	Star37      string
	Star38      string
	Star39      string
	Star40      string
	Star41      string
	Star42      string
	Star43      string
	Star44      string
	Star45      string

	LevelDarkblue string
	LevelLightblue string
	Level1 string
	Level2 string
	Level3 string
	Level4 string
	Level5 string
	PreservabilityLevel1 string
	PreservabilityLevel2 string
	PreservabilityLevel3 string
	PreservabilityLevel4 string
	PreservabilityLevel5 string
	AccessibilityLevel1 string
	AccessibilityLevel2 string
	AccessibilityLevel3 string
	AccessibilityLevel4 string
	AccessibilityLevel5 string
	UsabilityLevel1 string
	UsabilityLevel2 string
	UsabilityLevel3 string
	UsabilityLevel4 string
	UsabilityLevel5 string
	ProductionSustainabilityLevel1 string
	ProductionSustainabilityLevel2 string
	ProductionSustainabilityLevel3 string
	ProductionSustainabilityLevel4 string
	ProductionSustainabilityLevel5 string
	DataQualityAssuranceLevel1 string
	DataQualityAssuranceLevel2 string
	DataQualityAssuranceLevel3 string
	DataQualityAssuranceLevel4 string
	DataQualityAssuranceLevel5 string
	DataQualityControlMonitoringLevel1 string
	DataQualityControlMonitoringLevel2 string
	DataQualityControlMonitoringLevel3 string
	DataQualityControlMonitoringLevel4 string
	DataQualityControlMonitoringLevel5 string
	DataQualityAssessmentLevel1 string
	DataQualityAssessmentLevel2 string
	DataQualityAssessmentLevel3 string
	DataQualityAssessmentLevel4 string
	DataQualityAssessmentLevel5 string
	TransparencyTraceabilityLevel1 string
	TransparencyTraceabilityLevel2 string
	TransparencyTraceabilityLevel3 string
	TransparencyTraceabilityLevel4 string
	TransparencyTraceabilityLevel5 string
	DataIntegrityLevel1 string
	DataIntegrityLevel2 string
	DataIntegrityLevel3 string
	DataIntegrityLevel4 string
	DataIntegrityLevel5 string
}

// Read CSV data into struct slice
func getCsvData(datafile string) []DsmmAssessmentRecord {

	// Assign default star ratings to singleRecord struct
	singleRecord.Stardefault = `<a:srgbClr val="FFFFFF"/>`
	singleRecord.Star1 =  singleRecord.Stardefault
	singleRecord.Star2 =  singleRecord.Stardefault
	singleRecord.Star3 =  singleRecord.Stardefault
	singleRecord.Star4 =  singleRecord.Stardefault
	singleRecord.Star5 =  singleRecord.Stardefault
	singleRecord.Star6 =  singleRecord.Stardefault
	singleRecord.Star7 =  singleRecord.Stardefault
	singleRecord.Star8 =  singleRecord.Stardefault
	singleRecord.Star9 =  singleRecord.Stardefault
	singleRecord.Star10 = singleRecord.Stardefault
	singleRecord.Star11 = singleRecord.Stardefault
	singleRecord.Star12 = singleRecord.Stardefault
	singleRecord.Star13 = singleRecord.Stardefault
	singleRecord.Star14 = singleRecord.Stardefault
	singleRecord.Star15 = singleRecord.Stardefault
	singleRecord.Star16 = singleRecord.Stardefault
	singleRecord.Star17 = singleRecord.Stardefault
	singleRecord.Star18 = singleRecord.Stardefault
	singleRecord.Star19 = singleRecord.Stardefault
	singleRecord.Star20 = singleRecord.Stardefault
	singleRecord.Star21 = singleRecord.Stardefault
	singleRecord.Star22 = singleRecord.Stardefault
	singleRecord.Star23 = singleRecord.Stardefault
	singleRecord.Star24 = singleRecord.Stardefault
	singleRecord.Star25 = singleRecord.Stardefault
	singleRecord.Star26 = singleRecord.Stardefault
	singleRecord.Star27 = singleRecord.Stardefault
	singleRecord.Star28 = singleRecord.Stardefault
	singleRecord.Star29 = singleRecord.Stardefault
	singleRecord.Star30 = singleRecord.Stardefault
	singleRecord.Star31 = singleRecord.Stardefault
	singleRecord.Star32 = singleRecord.Stardefault
	singleRecord.Star33 = singleRecord.Stardefault
	singleRecord.Star34 = singleRecord.Stardefault
	singleRecord.Star35 = singleRecord.Stardefault
	singleRecord.Star36 = singleRecord.Stardefault
	singleRecord.Star37 = singleRecord.Stardefault
	singleRecord.Star38 = singleRecord.Stardefault
	singleRecord.Star39 = singleRecord.Stardefault
	singleRecord.Star40 = singleRecord.Stardefault
	singleRecord.Star41 = singleRecord.Stardefault
	singleRecord.Star42 = singleRecord.Stardefault
	singleRecord.Star43 = singleRecord.Stardefault
	singleRecord.Star44 = singleRecord.Stardefault
	singleRecord.Star45 = singleRecord.Stardefault

	// Assign default level colors to singleRecord struct
	singleRecord.LevelDarkblue = `<a:srgbClr val="CCD5EA"/>`
	singleRecord.LevelLightblue = `<a:srgbClr val="E7EBF5"/>`
	singleRecord.PreservabilityLevel1 = singleRecord.LevelDarkblue
	singleRecord.PreservabilityLevel2 = singleRecord.LevelLightblue
	singleRecord.PreservabilityLevel3 = singleRecord.LevelDarkblue
	singleRecord.PreservabilityLevel4 = singleRecord.LevelLightblue
	singleRecord.PreservabilityLevel5 = singleRecord.LevelDarkblue
	singleRecord.AccessibilityLevel1 = singleRecord.LevelDarkblue
	singleRecord.AccessibilityLevel2 = singleRecord.LevelLightblue
	singleRecord.AccessibilityLevel3 = singleRecord.LevelDarkblue
	singleRecord.AccessibilityLevel4 = singleRecord.LevelLightblue
	singleRecord.AccessibilityLevel5 = singleRecord.LevelDarkblue
	singleRecord.UsabilityLevel1 = singleRecord.LevelDarkblue
	singleRecord.UsabilityLevel2 = singleRecord.LevelLightblue
	singleRecord.UsabilityLevel3 = singleRecord.LevelDarkblue
	singleRecord.UsabilityLevel4 = singleRecord.LevelLightblue
	singleRecord.UsabilityLevel5 = singleRecord.LevelDarkblue
	singleRecord.ProductionSustainabilityLevel1 = singleRecord.LevelDarkblue
	singleRecord.ProductionSustainabilityLevel2 = singleRecord.LevelLightblue
	singleRecord.ProductionSustainabilityLevel3 = singleRecord.LevelDarkblue
	singleRecord.ProductionSustainabilityLevel4 = singleRecord.LevelLightblue
	singleRecord.ProductionSustainabilityLevel5 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssuranceLevel1 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssuranceLevel2 = singleRecord.LevelLightblue
	singleRecord.DataQualityAssuranceLevel3 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssuranceLevel4 = singleRecord.LevelLightblue
	singleRecord.DataQualityAssuranceLevel5 = singleRecord.LevelDarkblue
	singleRecord.DataQualityControlMonitoringLevel1 = singleRecord.LevelDarkblue
	singleRecord.DataQualityControlMonitoringLevel2 = singleRecord.LevelLightblue
	singleRecord.DataQualityControlMonitoringLevel3 = singleRecord.LevelDarkblue
	singleRecord.DataQualityControlMonitoringLevel4 = singleRecord.LevelLightblue
	singleRecord.DataQualityControlMonitoringLevel5 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssessmentLevel1 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssessmentLevel2 = singleRecord.LevelLightblue
	singleRecord.DataQualityAssessmentLevel3 = singleRecord.LevelDarkblue
	singleRecord.DataQualityAssessmentLevel4 = singleRecord.LevelLightblue
	singleRecord.DataQualityAssessmentLevel5 = singleRecord.LevelDarkblue
	singleRecord.TransparencyTraceabilityLevel1 = singleRecord.LevelDarkblue
	singleRecord.TransparencyTraceabilityLevel2 = singleRecord.LevelLightblue
	singleRecord.TransparencyTraceabilityLevel3 = singleRecord.LevelDarkblue
	singleRecord.TransparencyTraceabilityLevel4 = singleRecord.LevelLightblue
	singleRecord.TransparencyTraceabilityLevel5 = singleRecord.LevelDarkblue
	singleRecord.DataIntegrityLevel1 = singleRecord.LevelDarkblue
	singleRecord.DataIntegrityLevel2 = singleRecord.LevelLightblue
	singleRecord.DataIntegrityLevel3 = singleRecord.LevelDarkblue
	singleRecord.DataIntegrityLevel4 = singleRecord.LevelLightblue
	singleRecord.DataIntegrityLevel5 = singleRecord.LevelDarkblue

	csvfile, err := os.Open(datafile)
	checkError("open csv file failed", err)
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	checkError("reading csv data failed", err)

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
		singleRecord.AZ = cell[51]
		singleRecord.BA = cell[52]
		singleRecord.BB = cell[53]
		singleRecord.BC = cell[54]
		singleRecord.BD = cell[55]
		singleRecord.BE = cell[56]
		singleRecord.BF = cell[57]
		singleRecord.BG = cell[58]

		// Append singleRecord to allRecords struct slice
		allRecords = append(allRecords, singleRecord)
	}


	/* // Print all values from one record to stdout just for quick check that it works
	v := reflect.ValueOf(allRecords[35])
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	fmt.Println(values) */

	return allRecords
} // end getCsvData()
