package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ncmlAdditions struct {
	ncFileName string
	fileSize   int
	dataPath   string
	englishTitle string
}

//ncmlAdditions struct method that adds information missing from .ncml to same
func (p *ncmlAdditions) encode(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "<title>%s</title>\n<filesize>%d</filesize>\n<path>%s</path>\n<englishtitle>%s</englishtitle>\n</netcdf>",
		p.ncFileName, p.fileSize, p.dataPath, p.englishTitle)
}

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

var (
	ncFilePath string = "/nodc/web/data.nodc/htdocs/ndbc/co-ops"
	xslFile    string = "/nodc/users/tjaensch/xsl/coops/XSL/ncml2iso_modified_from_UnidataDD2MI_COOPS_Thomas_edits.xsl"
	//CO-OPS collection metadata template file
	isocofile  string = "/nodc/web/data.nodc/htdocs/nodc/archive/metadata/test/collection/NDBC-COOPS.xml"
	ncFiles    []string = findNcFiles(ncFilePath)
	ncFileName string
	fileSize   int
	dataPath   string
	englishTitle string
	additions  = ncmlAdditions{ncFileName, fileSize, dataPath, englishTitle}
)

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()
	prepDirs()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of ncFiles slice
	fileSegments := getFileSegments()
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			nc2iso(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()
	t1 := time.Now()
	log.Printf("The program took %v hours to run.\n", t1.Sub(t0).Hours())
}

// Create file directories
func prepDirs() {
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./xml_output", 0777)
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments() [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 50
	// Add subslices to fileSegments slice
	for i := 0; i < len(ncFiles)-increaseRate; i +=increaseRate {
	fileSegments = append(fileSegments, ncFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, ncFiles[len(ncFiles)-increaseRate:])
	return fileSegments
}

func nc2iso(ncFiles []string) {
	for _, ncFile := range ncFiles {
		var (
			ncFileName string = getFileName(ncFile)
			fileSize   int    = getFileSize(ncFile)
			dataPath   string = getFilePath(ncFile)
			englishTitle string = getEnglishTitle(ncFile)
			additions         = ncmlAdditions{ncFileName, fileSize, dataPath, englishTitle}
		)
		ncdump(ncFile)
		appendToNcml(ncFile, additions)
		xsltprocToISO(ncFile, xslFile)
		addCollectionMetadata(ncFile)
	}
}

func findNcFiles(ncFilePath string) []string {
	var ncFiles []string
	var files []byte
	var err error
	cmdName := "find"
	cmdArgs := []string{"-L", ncFilePath, "-type", "f", "-name", "*.nc"}
	if files, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with finding .nc files in source directory, program exiting.", err)
		os.Exit(1)
	}

	for _, rune := range bytes.Split(files, []byte{'\n'}) {
		ncFiles = append(ncFiles, string(rune))
	}
	log.Printf("%d files found in source directory", len(ncFiles)-1)
	//Last element is blank that's why -1
	return ncFiles[:len(ncFiles)-1]
}

func ncdump(ncFile string) {
	var ncml []byte
	var err error
	cmdName := "ncdump"
	cmdArgs := []string{"-x", ncFile}
	if ncml, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with ncdump, program exiting.", err)
		os.Exit(1)
	}
	// Write ncdump conversion to file
	err = ioutil.WriteFile("./ncml/"+getFileName(ncFile)+".ncml", ncml, 0644)
	checkError("write ncml file failed, program exiting", err)
}

// Get just the file name without file extension
func getFileName(ncFile string) (ncFileName string) {
	ncFileBasePath := filepath.Base(ncFile)
	ncFileName = strings.TrimSuffix(ncFileBasePath, ".nc")
	return
}

// Get "English" title for ISO XML
func getEnglishTitle(ncFile string) string {
	re := regexp.MustCompile("_D\\d+")
	return "NDBC-COOPS_" + strings.TrimPrefix(getFileName(ncFile), "NOS_") + " - CO-OPS buoy " + getFileName(ncFile)[4:11] + " for " + getFileName(ncFile)[12:18] + ", deployment " + re.FindString(getFileName(ncFile))[2:]
}

// Get .nc file size in KB
func getFileSize(ncFile string) (fileSize int) {
	file, err := os.Open(ncFile)
	checkError("Open .nc file failed", err)
	defer file.Close()
	stat, err := file.Stat()
	checkError("Get file size failed", err)
	return (int)(stat.Size() / 1024)
}

// Get .nc file path on WAF
func getFilePath(ncFile string) string {
	trimmedPath := ncFile[27:len(ncFile)]
	return strings.Replace(trimmedPath, filepath.Base(ncFile), "", -1)
}

// Append struct fields to .ncml file so XSLT can grab the added values
func appendToNcml(ncFile string, additions ncmlAdditions) {
	// Mess with <netcdf> tag for XSL wo work later on
	input, err := ioutil.ReadFile("./ncml/" + getFileName(ncFile) + ".ncml")
	checkError("read ncml file failed", err)
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "<netcdf") {
			lines[i] = "<netcdf>"
		}
		if strings.Contains(line, "</netcdf>") {
			lines[i] = ""
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile("./ncml/"+getFileName(ncFile)+".ncml", []byte(output), 0644)
		checkError("write ncml file failed", err)
	}
	// Append extra lines to .ncml
	file, err := os.OpenFile("./ncml/"+getFileName(ncFile)+".ncml", os.O_RDWR|os.O_APPEND, 0)
	checkError("open ncml file failed", err)
	if _, err := additions.encode(file); err != nil {
		fmt.Printf("append to ncml file failed: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
}

func xsltprocToISO(ncFile string, xslFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	//Convert to ISO
	cmdArgs := []string{xslFile, "./ncml/" + getFileName(ncFile) + ".ncml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the XSLT conversion, program exiting.", err)
		os.Exit(1)
	}
	// Write xsltproc conversion to file
	err = ioutil.WriteFile("./xml_output/"+getFileName(ncFile)+".xml", isoXML, 0644)
	checkError("write xml file failed", err)
}

func addCollectionMetadata(ncFile string) {
	var isoXML []byte
	var err error
	cmdName := "xsltproc"
	cmdArgs := []string{"--stringparam", "collFile", isocofile, "/nodc/users/tjaensch/xsl/coops/XSL/granule.xsl", "./xml_output/" + getFileName(ncFile) + ".xml"}
	if isoXML, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with the collection metadata addition, program exiting.", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("./xml_output/"+getFileName(ncFile)+".xml", isoXML, 0644)
	if err == nil {
		fmt.Printf("%s.xml successfully written to output directory\n", getFileName(ncFile))
	}
	checkError("write xml file failed, program exiting", err)
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./xml_output")
	checkError("read xml_output failed, program exiting", err)
	log.Printf("%d files written to xml_output directory\n", len(files))
}
