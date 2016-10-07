package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	f291FilePath string   = "/nodc/projects/buoy/F291_CDL/f291/"
	f291Files    []string = findF291Files(f291FilePath)
)

func main() {
	log.Printf("Working digging up files...")
	t0 := time.Now()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of ncFiles slice
	fileSegments := getFileSegments(f291Files)
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			f2912nc(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()
	t1 := time.Now()
	fmt.Printf("The program took %v hours to run.\n", t1.Sub(t0).Hours())
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments(f291Files []string) [][]string {
	// Create a slice of ncFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	increaseRate := 250
	// Add subslices to fileSegments slice
	for i := 0; i < len(f291Files)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, f291Files[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, f291Files[len(f291Files)-increaseRate:])
	return fileSegments
}

func f2912nc(f291Files []string) {
	for _, f291File := range f291Files {
			f29one2nc(f291File)
	}
}

func findF291Files(f291FilePath string) []string {
	var files []string
	filepath.Walk(f291FilePath, func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	//fmt.Printf("%v", files)
	fmt.Printf("%d files found in source directory\n", len(files))
	return files
}

func f29one2nc(f291File string) {
	// Open log file
	f, err := os.OpenFile("f29one2nc_conversion_failures.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError("open log file failed", err)
	defer f.Close()
	// Assign it to the standard logger
	log.SetOutput(f)
	cmd := exec.Command("/nodc/users/tjaensch/workspace/f291tools/f291nc/Debug/f291nc", f291File)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err == nil {
		fmt.Println(f291File + " stdout: " + out.String())
		log.Println(f291File + " stdout: " + out.String())
	}
	if err != nil {
		fmt.Println(f291File + " stderror: " + stderr.String())
		log.Println(f291File + " stderror: " + stderr.String())
	}
}

// Get just the file name without file extension
func getFileName(f291File string) string {
	f291FileBasePath := filepath.Base(f291File)
	return f291FileBasePath
}

func countOutputFiles() {
	files, err := ioutil.ReadDir(".")
	checkError("read nc_f291_Sonny_script directory failed, program exiting", err)
	fmt.Printf("%d files written to nc_f291_Sonny_script directory\n", len(files))
}
