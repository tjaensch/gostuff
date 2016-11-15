package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

var (
  protocol string = "http://"
  host string = "localhost:"
  port string = "8097"
  path string = "/onestop/api/metadata/"
  xmlFilePath string = "./testfiles/"
	xmlFiles    []string = findXmlFiles(xmlFilePath)
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

  findXmlFiles(xmlFilePath)

  for _ , xmlFile := range (xmlFiles) {
    postFile(xmlFile)
  }

  t1 := time.Now()
	log.Printf("The program took %v seconds to run.\n", t1.Sub(t0).Seconds())
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

func postFile(xmlFile string)  {
  data, err := os.Open(xmlFile)
  checkError("open file failed", err)
  defer data.Close()

  req, err := http.NewRequest("POST", protocol + host + port + path, data)
  checkError("POST request failed", err)

  req.Header.Set("Content-Type", "application/xml")

  client := &http.Client{}
  res, err := client.Do(req)
  checkError("client request failed", err)
  defer res.Body.Close()

  // Access res.Body to collect server response errors
  bs, err := ioutil.ReadAll(res.Body)
  checkError("reading res.Body tanked", err)

  if res.Status == "201 Created" {
    fmt.Printf("%v successfully posted\n", filepath.Base(xmlFile))
  } else {
    fmt.Printf("ERROR: %v failed with %v------%v\n", filepath.Base(xmlFile), res.Status, string(bs))
  }
}
