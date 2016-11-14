package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
)

var (
  xmlFile string = "testfiles/WTDL_20150803v10001.xml"
  protocol string = "http://"
  host string = "localhost:"
  port string = "8097"
  path string = "/onestop/api/metadata/"
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
  postFile(xmlFile)
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
