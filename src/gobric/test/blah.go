package main

import (
	"gopkg.in/xmlpath.v1"
	"log"
    "os"
    "fmt"
)

func main() {

file, err := os.Open("../testfiles/woa13_95A4_s00_01.xml")
    checkError("Open woa13_95A4_s00_01.xml file failed", err)

	xpath := "/MI_Metadata"
    path := xmlpath.MustCompile(xpath)
	root, err := xmlpath.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}

}

// Generic error checking function
func checkError(reason string, err error) {
    if err != nil {
        fmt.Printf("%s: %s\n", reason, err)
        os.Exit(1)
    }
}
