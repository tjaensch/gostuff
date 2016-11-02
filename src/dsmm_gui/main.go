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
	fmt.Println("Listening on 10.90.235.15:1313")
	http.ListenAndServe("10.90.235.15:1313", nil)
}

func getDsmmRatings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Form submitted
		err := r.ParseForm() // required if no r.FormValue()
		checkError("parsing html form failed, program exiting", err)

		err = decoder.Decode(ratingsValues, r.PostForm)
		checkError("decode form failed, program exiting", err)

		t, err := template.ParseFiles("templates/dsmm.tmpl")
		// Write executed dsmm_template to text file
		f, err := os.Create("dsmm_snippet.txt")
		checkError("write parsed template to file failed, program exiting", err)
		defer f.Close()
		t.ExecuteTemplate(os.Stdout, "dsmm", ratingsValues)
		t.ExecuteTemplate(f, "dsmm", ratingsValues)
		checkError("parsing template failed, program exiting", err)
		fmt.Println()

	}
	html_template, _ := ioutil.ReadFile("templates/ui.html")
	w.Write([]byte(html_template))
}
