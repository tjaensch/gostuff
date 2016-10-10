package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	firstFile  = "testfiles/test1.txt"
	secondFile = "testfiles/test2.txt"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
	compareFiles(firstFile, secondFile)
}

func compareFiles(firstFile string, secondFile string) {
	if readFile(firstFile) == readFile(secondFile) {
		fmt.Println("No differences found.")
		return
	}

	// Open log file
	f, err := os.OpenFile("compareFiles_log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError("open log file failed", err)
	defer f.Close()
	// Assign it to the standard logger
	log.SetOutput(f)

	openFile, err := os.Open(firstFile)
	checkError("open file failed", err)

	r := bufio.NewReader(openFile)
	line, err := readFileLine(r)
	for err == nil {
		// Only find lines that have letters in them
		hasLetters := regexp.MustCompile(`[a-zA-Z]`).MatchString
		if hasLetters(line) && !strings.Contains(readFile(secondFile), strings.TrimSpace(line)) {
			fmt.Printf("%s: \"%s\" not found in %s\n", firstFile, strings.TrimSpace(line), secondFile)
			log.Printf("%s: \"%s\" not found in %s\n", firstFile, strings.TrimSpace(line), secondFile)
		}
		line, err = readFileLine(r)
	}
}

func readFile(file string) string {
	fileBytes, err := ioutil.ReadFile(file)
	checkError("open secondFile file failed", err)
	return string(fileBytes)
}

func readFileLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
			ln = append(ln, line...)
	}
	return string(ln), err
}
