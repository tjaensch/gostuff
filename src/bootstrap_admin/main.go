package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/schema"
)

var (
	decoder       = schema.NewDecoder()
	ratingsValues = new(RatingsValues)
)

type RatingsValues struct {
	Preservability               string `schema:"preservability"`
	Accessibility                string `schema:"accessibility"`
	Usability                    string `schema:"usability"`
	ProductionSustainability     string `schema:"production_sustainability"`
	DataQualityAssurance         string `schema:"data_quality_assurance"`
	DataQualityControlMonitoring string `schema:"data_quality_control_monitoring"`
	DataQualityAssessment        string `schema:"data_quality_assessment"`
	TransparencyTraceability     string `schema:"transparency_traceability"`
	DataIntegrity                string `schema:"data_integrity"`
}

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", getDsmmRatings)
	http.ListenAndServe("10.90.235.15:1313", nil)
}

func getDsmmRatings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Form submitted
		err := r.ParseForm() // required if no r.FormValue()
		checkError("parse form failed, program exiting", err)

		err = decoder.Decode(ratingsValues, r.PostForm)
		checkError("decode form failed, program exiting", err)

		t := template.Must(template.New("t1").
			Parse("Dot:{{.}}\n"))
		t.Execute(os.Stdout, ratingsValues)
		checkError("parsing template failed, program exiting", err)

	}
	html_template, _ := ioutil.ReadFile("html_template.html")
	w.Write([]byte(html_template))
}
