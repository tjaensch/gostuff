package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

var pptxSlide1Xml string = "./templates/slide1.xml"

func parseCsvDataToPptxXml(singleRecord DsmmAssessmentRecord) {
	tmpl, err := template.ParseFiles(pptxSlide1Xml)
	checkError("parse template failed", err)

	f, err := os.Create("./Star_rating_template/ppt/slides/slide1.xml")
	checkError("create template file failed", err)

	// Add CSV values to PPTX slide1.xml for Star Rating image
	err = tmpl.ExecuteTemplate(f, "slide1.xml", singleRecord)
	checkError("execute template failed", err)

	// Change into Star_rating_template and zip up files
	os.Chdir("./Star_rating_template")
	cmdName := "zip"
	cmdArgs := []string{"-r", "../" + singleRecord.C + "_Star_rating_template.zip", "_rels", "[Content_Types].xml", "ppt", "docProps" }
	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("Something went wrong with zipping files in source directory, program exiting.", err)
		os.Exit(1)
	}
	os.Chdir("..")

	// Rename .zip file to .pptx
	err =  os.Rename(singleRecord.C + "_Star_rating_template.zip", "./output/" + singleRecord.C + "_Star_rating_template.pptx")
	checkError("renaming .zip file failed", err)
}
