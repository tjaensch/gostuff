package main

import (
    "fmt"
    "net/http"
    "os"
)

var (
  xmlFile string = "testfiles/NDBC-CMANWx.NDBC_51000_201509_D3_v00.xml"
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
}
