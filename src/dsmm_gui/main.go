package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	templates     map[string]*template.Template
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

func init() {
	loadTemplates()
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/", DsmmFormRoute).Methods("GET")
	router.HandleFunc("/dsmm_results", DsmmResultsRoute).Methods("POST")

	fmt.Println("Listening on 10.90.235.15:1313")
	if err := http.ListenAndServe("10.90.235.15:1313", router); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func DsmmFormRoute(res http.ResponseWriter, req *http.Request) {

	if err := templates["dsmm_form"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func DsmmResultsRoute(res http.ResponseWriter, req *http.Request) {

	// Form submitted
	err := req.ParseForm() // required if no r.FormValue()
	checkError("parsing html form failed, program exiting", err)

	err = decoder.Decode(ratingsValues, req.PostForm)
	checkError("decode form failed, program exiting", err)

	t, err := template.ParseFiles("templates/dsmm.tmpl")
	t.ExecuteTemplate(os.Stdout, "dsmm", ratingsValues)

	if err := templates["dsmm_results"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func loadTemplates() {
	var baseTemplate = "templates/layout/_base.html"
	templates = make(map[string]*template.Template)

	templates["dsmm_form"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/dsmm_form.html"))
	templates["dsmm_results"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/dsmm_results.html"))
}
