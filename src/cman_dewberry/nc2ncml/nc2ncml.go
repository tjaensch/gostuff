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
	ncFilePath string   = "./nc_derived_from_f291/"
	ncFiles    []string = findNcFiles(ncFilePath)
)

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()
	prepDirs()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of ncFiles slice
	fileSegments := getFileSegments(ncFiles)
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			nc2ncml(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()
	t1 := time.Now()
	fmt.Printf("The program took %v minutes to run.\n", t1.Sub(t0).Minutes())
}

// Create file directories
func prepDirs() {
	os.Mkdir("./ncml", 0777)
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments(ncFiles []string) [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 250
	// Add subslices to fileSegments slice
	for i := 0; i < len(ncFiles)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, ncFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, ncFiles[len(ncFiles)-increaseRate:])
	return fileSegments
}

func nc2ncml(ncFiles []string) {
	for _, ncFile := range ncFiles {
		ncdump(ncFile)
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
	fmt.Printf("%d files found in source directory", len(ncFiles)-1)
	//Last element is blank that's why -1
	return ncFiles[:len(ncFiles)-1]
}

func ncdump(ncFile string) {
	// Open log file
	f, err := os.OpenFile("nc2ncml_conversion_failures.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError("open log file failed", err)
	defer f.Close()
	// Assign it to the standard logger
	log.SetOutput(f)
	var ncml []byte
	cmdName := "ncdump"
	cmdArgs := []string{"-x", ncFilePath + getFileName(ncFile) + ".nc"}
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Something went wrong with ncdump on %s, error %v \n", ncFile, stderr.String())
		log.Printf("Something went wrong with ncdump on %s, error %v \n", ncFile, stderr.String())
	}
	if err == nil {
		cmdName = "ncdump"
		cmdArgs = []string{"-x", ncFilePath + getFileName(ncFile) + ".nc"}
		if ncml, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			fmt.Printf("Something went wrong with ncdump on %s, error %v \n", ncFile, stderr.String())
			log.Printf("Something went wrong with ncdump on %s, error %v \n", ncFile, stderr.String())
		}
		// Write ncdump conversion to file
		err = ioutil.WriteFile("./ncml/" + getFileName(ncFile) + ".ncml", ncml, 0644)
		checkError("write ncml file failed, program exiting", err)
		if err == nil {
			fmt.Printf("%s.ncml successfully written to output directory\n", getFileName(ncFile))
		}
	}
}

// Get just the file name without file extension
func getFileName(ncFile string) string {
	ncFileBasePath := filepath.Base(ncFile)
	return strings.TrimSuffix(ncFileBasePath, ".nc")
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./ncml")
	checkError("read ncml directory failed, program exiting", err)
	fmt.Printf("%d files written to ncml directory\n", len(files))
}
