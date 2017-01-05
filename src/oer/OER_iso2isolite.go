package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

var (
	xmlFilePath string = "/nodc/projects/satdata/OER/Metadata/waf/"
	xslFile    string = "/nodc/users/tjaensch/onestop.git/xsl/oer/XSL/OER_ISO2ISOLite_conversion.xsl"
	//OER collection metadata template file
	//isocofile  string = "/nodc/web/data.nodc/htdocs/nodc/archive/metadata/test/collection/NDBC-COOPS.xml"
	xmlFiles    []string = findXmlFiles(xmlFilePath)
	xmlFileName string
)

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()
	// prepDirs()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of ncFiles slice
	fileSegments := getFileSegments()
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			iso2isolite(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()
	t1 := time.Now()
	log.Printf("The program took %v minutes to run.\n", t1.Sub(t0).Minutes())
}

// Create file directories
func prepDirs() {
	os.Mkdir("./isolite_xml_output", 0777)
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments() [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 1000
	// Add subslices to fileSegments slice
	for i := 0; i < len(xmlFiles)-increaseRate; i +=increaseRate {
	fileSegments = append(fileSegments, xmlFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, xmlFiles[len(xmlFiles)-increaseRate:])
	return fileSegments
}

func iso2isolite(xmlFiles []string) {
	for _, xmlFile := range xmlFiles {
		xsltprocToISO(xmlFile, xslFile)
		// addCollectionMetadata(xmlFile)
	}
}

func findXmlFiles(xmlFilePath string) []string {
	var xmlFiles []string
	var files []byte
	var err error
	cmdName := "find"
	cmdArgs := []string{"-L", xmlFilePath, "-type", "f", "-name", "*.xml"}
	if files, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with finding .xml files in source directory, program exiting.", err)
		os.Exit(1)
	}

	for _, rune := range bytes.Split(files, []byte{'\n'}) {
		xmlFiles = append(xmlFiles, string(rune))
	}
	log.Printf("%d files found in source directory", len(xmlFiles)-1)
	//Last element is blank that's why -1
	return xmlFiles[:len(xmlFiles)-1]
}

// Get just the file name without file extension
func getFileName(xmlFile string) (xmlFileName string) {
	xmlFileBasePath := filepath.Base(xmlFile)
	xmlFileName = strings.TrimSuffix(xmlFileBasePath, ".xml")
	return
}

func xsltprocToISO(xmlFile string, xslFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	//Convert to ISO
	cmdArgs := []string{xslFile, xmlFilePath + getFileName(xmlFile) + ".xml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the XSLT conversion, program exiting.", err)
		os.Exit(1)
	}
	// Write xsltproc conversion to file
	err = ioutil.WriteFile("./isolite_xml_output/"+getFileName(xmlFile)+".xml", isoXML, 0644)
	checkError("write xml file failed", err)
}

/* func addCollectionMetadata(xmlFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	cmdArgs := []string{"--stringparam", "collFile", isocofile, "/nodc/users/tjaensch/onestop.git/xsl/oer/XSL/granule.xsl", "./isolite_xml_output/" + getFileName(xmlFile) + ".xml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the collection metadata addition, program exiting.", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("./isolite_xml_output/"+getFileName(xmlFile)+".xml", isoXML, 0644)
	if err == nil {
		fmt.Printf("%s.xml successfully written to output directory\n", getFileName(xmlFile))
	}
	checkError("write xml file failed, program exiting", err)
} */

func countOutputFiles() {
	files, err := ioutil.ReadDir("./isolite_xml_output/")
	checkError("read isolite_xml_output failed, program exiting", err)
	log.Printf("%d files written to isolite_xml_output directory\n", len(files))
}
