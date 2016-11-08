package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	decoder       = schema.NewDecoder()
	ratingsValues = new(RatingsValues)
	dsmmSnippet   bytes.Buffer
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
	router.HandleFunc("/", DsmmForm).Methods("GET")
	router.HandleFunc("/dsmm_results", DsmmResults).Methods("POST")
	router.HandleFunc("/dsmm_xml", DsmmWriteSnippetToBrowser).Methods("POST")
	router.HandleFunc("/upload_xml", UploadXmlFile).Methods("POST")

	fmt.Println("Listening on 10.90.235.15:1313")
	if err := http.ListenAndServe("10.90.235.15:1313", router); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func DsmmForm(w http.ResponseWriter, r *http.Request) {
	dsmm_form := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/dsmm_form.html"))
	if err := dsmm_form.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DsmmResults(w http.ResponseWriter, r *http.Request) {
	// Form submitted
	err := r.ParseForm() // required if no r.FormValue()
	checkError("parsing html form failed, program exiting", err)

	err = decoder.Decode(ratingsValues, r.PostForm)
	checkError("decode form failed, program exiting", err)

	dsmm_form := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/dsmm_results.html"))
	if err := dsmm_form.Execute(w, ratingsValues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DsmmWriteSnippetToBrowser(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/dsmm/dsmm.tmpl")
	checkError("execute template failed, program exiting", err)
	t.ExecuteTemplate(os.Stdout, "dsmm", ratingsValues)
	t.ExecuteTemplate(w, "dsmm", ratingsValues)
}

func UploadXmlFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/dsmm/dsmm.tmpl")
	checkError("execute template failed, program exiting", err)
	t.ExecuteTemplate(&dsmmSnippet, "dsmm", ratingsValues)

	file, _, err := r.FormFile("file")
	checkError("UploadXmlFile failed, program exiting", err)
	defer file.Close()

	out, err := os.Create("/tmp/uploadedfile")
	checkError("create file folder for upload failed, program exiting", err)
	defer out.Close()

	// write content from POST to file
	_, err = io.Copy(out, file)
	checkError("write content from POST to file failed, program exiting", err)

	dat, err := ioutil.ReadFile("/tmp/uploadedfile")
	checkError("read uploaded file failed, program exiting", err)

	if !strings.Contains(string(dat), "</gmi:MI_Metadata>") {
		fail := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/download_fail.html"))
		if err := fail.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if strings.Contains(string(dat), "<gmd:metadataMaintenance>") {
		strings.Replace(string(dat), "</gmd:metadataMaintenance>", "</gmd:metadataMaintenance>" + dsmmSnippet.String(), 1)
		success := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/download_success.html"))
		if err := success.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		strings.Replace(string(dat), "</gmi:MI_Metadata>", dsmmSnippet.String() + "</gmi:MI_Metadata>", 1)
		success := template.Must(template.ParseFiles("templates/layout/_base.html", "templates/dsmm/download_success.html"))
		if err := success.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	err = os.Remove("/tmp/uploadedfile")
	checkError("remove uploaded file failed, program exiting", err)
}
