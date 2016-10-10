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
	cdlFilePath string   = "/nodc/projects/buoy/dewberry_cdl/Allf291Files/"
	cdlFiles    []string = findCdlFiles(cdlFilePath)
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
			cdl2nc(fileSegment)
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
	os.Mkdir("./nc", 0777)
	os.Mkdir("./ncml", 0777)
	os.Mkdir("./cdl2ncml_conversion_failures", 0777)
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments() [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 250
	// Add subslices to fileSegments slice
	for i := 0; i < len(cdlFiles)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, cdlFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, cdlFiles[len(cdlFiles)-increaseRate:])
	return fileSegments
}

func cdl2nc(cdlFiles []string) {
	for _, cdlFile := range cdlFiles {
		if _, err := os.Stat("./nc/" + getFileName(cdlFile) + ".nc"); os.IsNotExist(err) {
		ncgen(cdlFile)
		ncdump(cdlFile)
		}
	}
}

func findCdlFiles(cdlFilePath string) []string {
	var cdlFiles []string
	var files []byte
	var err error
	cmdName := "find"
	cmdArgs := []string{"-L", cdlFilePath, "-type", "f", "-name", "*.cdl"}
	if files, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with finding .cdl files in source directory, program exiting.", err)
		os.Exit(1)
	}
	for _, rune := range bytes.Split(files, []byte{'\n'}) {
		cdlFiles = append(cdlFiles, string(rune))
	}
	log.Printf("%d files found in source directory", len(cdlFiles)-1)
	//Last element is blank that's why -1
	return cdlFiles[:len(cdlFiles)-1]
}

func ncgen(cdlFile string) {
	// Open log file
	f, err := os.OpenFile("cdl2ncml_conversion_failures.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError("open log file failed", err)
	defer f.Close()
	// Assign it to the standard logger
	log.SetOutput(f)
	cmdName := "ncgen"
	cmdArgs := []string{"-b", cdlFile}
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(cdlFile + " error: " + stderr.String())
		log.Println(cdlFile + " error: " + stderr.String())
		_ = os.Rename(getFileName(cdlFile) + ".nc", "./cdl2ncml_conversion_failures/" + getFileName(cdlFile) + ".nc")
		return
	}
	// Compress file with nccopy
	cmdName = "nccopy"
	cmdArgs = []string{"-d9", getFileName(cdlFile) + ".nc", getFileName(cdlFile) + "_compressed.nc"}
	cmd = exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(cdlFile + " error: " + stderr.String())
		log.Println(cdlFile + " error: " + stderr.String())
		return
	}
	// If success move file to ./nc directory
	if _, err = os.Stat(getFileName(cdlFile) + "_compressed.nc"); err == nil {
		_ = os.Rename(getFileName(cdlFile)+"_compressed.nc", "./nc/"+getFileName(cdlFile)+".nc")
	}
	// Clean up uncompressed .nc file
	os.Remove(getFileName(cdlFile) + ".nc")
}

func ncdump(cdlFile string) {
	var ncml []byte
	var err error
	cmdName := "ncdump"
	cmdArgs := []string{"-x", "./nc/" + getFileName(cdlFile) + ".nc"}
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return
	}
	if err == nil {
		cmdName = "ncdump"
		cmdArgs = []string{"-x", "./nc/" + getFileName(cdlFile) + ".nc"}
		if ncml, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			fmt.Printf("Something went wrong with ncdump on %s, error %v \n", cdlFile, stderr.String())
			log.Printf("Something went wrong with ncdump on %s, error %v \n", cdlFile, stderr.String())
		}
		// Write ncdump conversion to file
		err = ioutil.WriteFile("./ncml/" + getFileName(cdlFile) + ".ncml", ncml, 0644)
		checkError("write ncml file failed, program exiting", err)
		if err == nil {
			fmt.Printf("%s.ncml successfully written to output directory\n", getFileName(cdlFile))
		}
	}
}

// Get just the file name without file extension
func getFileName(cdlFile string) string {
	cdlFileBasePath := filepath.Base(cdlFile)
	return strings.TrimSuffix(cdlFileBasePath, ".cdl")
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./nc")
	checkError("read nc directory failed, program exiting", err)
	log.Printf("%d files written to nc directory\n", len(files))
}
