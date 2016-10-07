package main

import (
	"fmt"
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

var (
	xmlFiles    []string = findXmlFiles()
	lenXmlFiles          = len(xmlFiles)
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
	t0 := time.Now()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of xmlFiles slice
	fileSegments := getFileSegments()
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			processXmlFiles(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	t1 := time.Now()
	fmt.Printf("\nThe program took %v minutes to run.\n", t1.Sub(t0).Minutes())
}

func findXmlFiles() []string {
	pathS, err := os.Getwd()
	checkError("get path failed", err)
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".xml", f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments() [][]string {
	var increaseRate int
	// Create a slice of xmlFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	if lenXmlFiles == 0 {
		fmt.Println("No XML files found in current working directory, program exiting.")
		os.Exit(1)
	}
	if lenXmlFiles <= 5 {
		increaseRate = 1
	} else {
		increaseRate = lenXmlFiles / 5
	}

	// Add subslices to fileSegments slice
	for i := 0; i < len(xmlFiles)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, xmlFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, xmlFiles[len(xmlFiles)-increaseRate:])
	return fileSegments
}

func processXmlFiles(xmlFiles []string) {
	for _, xmlFile := range xmlFiles {
		findBrokenHttpLinks(xmlFile)
	}
}

func findBrokenHttpLinks(xmlFile string) {
	// Create slice for links found
	var httpLinks []string
	badLinkCounter := 0
	goodLinkCounter := 0

	// Regular expression to match links
	re := regexp.MustCompile(`("|>|\s)https?:.*?("|<|\s)`)
	// Open log file
	f, err := os.OpenFile("linkchecker_bad_links_log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError("open log file failed", err)
	defer f.Close()
	// Assign it to the standard logger
	log.SetOutput(f)
	input, err := ioutil.ReadFile(xmlFile)
	checkError("read file failed", err)
	linksFound := re.FindAllString(string(input), -1)
	for _, linkFound := range linksFound {
		httpLinks = append(httpLinks, linkFound[1:len(linkFound)-1])
	}
	for _, httpLink := range httpLinks {
		var httpResponse []byte
		var err error
		// Replace "&amp;" with "&" in found links for curl to work better
		httpLink = strings.Replace(httpLink, "&amp;", "&", -1) 
		// Remove the image streaming blah from links for curl to work better
		httpLink = strings.Replace(httpLink, "&stream=true&stream_ID=plot_image", "", -1)
		cmdName := "curl"
		cmdArgs := []string{"--head", httpLink, "--max-time", "10"}
		if httpResponse, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			log.Printf("\n:-( Something went wrong with reaching the server, check link!\n", err)
			log.Println("\n")
			fmt.Printf("\n:-( Something went wrong with reaching the server, check link!\n", err)
			fmt.Println("\n")
		}
		if !strings.Contains(string(httpResponse), "200 OK") {
			badLinkCounter++
			log.Printf("\n--- BAD LINK ---\nFile: %v \nLink: %v \nServer Response: %v \n", filepath.Base(xmlFile), httpLink, string(httpResponse))
			fmt.Printf("\n--- BAD LINK ---\nFile: %v \nLink: %v \nServer Response: %v \n", filepath.Base(xmlFile), httpLink, string(httpResponse))
		}
		if strings.Contains(string(httpResponse), "200 OK") {
			goodLinkCounter++
			fmt.Printf("\n--- WORKING LINK ---\nFile: %v \nLink: %v \n", filepath.Base(xmlFile), httpLink)
		}
	}
	log.Printf("\nNumber of questionable links found in %v: %d \n", filepath.Base(xmlFile), badLinkCounter)
	log.Printf("\nNumber of working links found in %v: %d \n", filepath.Base(xmlFile), goodLinkCounter)
	log.Println("\n")
	log.Printf("\nNumber of all links found in %v: %v \n", filepath.Base(xmlFile), len(httpLinks))
	fmt.Printf("\nNumber of questionable links found in %v: %d \n", filepath.Base(xmlFile), badLinkCounter)
	fmt.Printf("\nNumber of working links found in %v: %d \n", filepath.Base(xmlFile), goodLinkCounter)
	fmt.Printf("\nNumber of all links found in %v: %v \n", filepath.Base(xmlFile), len(httpLinks))
}
