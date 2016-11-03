package main

import (
	//"bytes"
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

func DsmmFormRoute(w http.ResponseWriter, r *http.Request) {
	dsmm_form := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/dsmm_form.html"))
	if err := dsmm_form.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DsmmResultsRoute(w http.ResponseWriter, r *http.Request) {

	// Form submitted
	err := r.ParseForm() // required if no r.FormValue()
	checkError("parsing html form failed, program exiting", err)

	err = decoder.Decode(ratingsValues, r.PostForm)
	checkError("decode form failed, program exiting", err)

	t, err := template.ParseFiles("templates/dsmm/dsmm.tmpl")
	checkError("execute template failed, program exiting", err)
	t.ExecuteTemplate(os.Stdout, "dsmm", ratingsValues)
	//t.ExecuteTemplate(w, "dsmm", ratingsValues)

	dsmm_form := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/dsmm_results.html"))
	if err := dsmm_form.Execute(w, ratingsValues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
