package main

import (
  "os"
	"text/template"
)

var pptxSlide1Xml string = "./templates/slide1.xml"

func parseCsvDataToPptxXml(singleRecord DsmmAssessmentRecord) {
tmpl, err := template.ParseFiles(pptxSlide1Xml)
checkError("parse template failed", err)

f, err := os.Create("./Star_rating_template/ppt/slides/slide1.xml")
checkError("create template file failed", err)

err = tmpl.ExecuteTemplate(f, "slide1.xml", singleRecord)
checkError("execute template failed", err)
}
