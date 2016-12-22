package main

import (
	"fmt"
	"os"
	"os/exec"
)

func convertPptxToPng(singleRecord DsmmAssessmentRecord) {
	os.Chdir("./output")
	cmdName := "libreoffice"
	cmdArgs := []string{"--headless", "--convert-to", "png", singleRecord.C + "_" + singleRecord.K + "_Scoreboard_rating_template.pptx"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with converting .pptx to .png: %s_%s_Scoreboard_rating_template.pptx\n", singleRecord.C, singleRecord.K, err)
	}
	cmdName = "libreoffice"
	cmdArgs = []string{"--headless", "--convert-to", "png", singleRecord.C + "_" + singleRecord.K + "_Star_rating_template.pptx"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with converting .pptx to .png: %s_%s_Star_rating_template.pptx\n", singleRecord.C, singleRecord.K, err)
	}
	os.Chdir("..")
}

func updateWordTemplateWithNewPng(singleRecord DsmmAssessmentRecord) {
	os.Chdir("./output")
	cmdName := "cp"
	cmdArgs := []string{singleRecord.C + "_" + singleRecord.K + "_Scoreboard_rating_template.png", "image4.png"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with copying .png file: %s_%s_Scoreboard_rating_template.png\n", singleRecord.C, singleRecord.K, err)
	}
	os.Rename("./image4.png", "../DSMM_WORDDOC_template_unzipped/word/media/image4.png")
	cmdName = "cp"
	cmdArgs = []string{singleRecord.C + "_" + singleRecord.K + "_Star_rating_template.png", "image2.png"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with copying .png file: %s_%s_Star_rating_template.png\n", singleRecord.C, singleRecord.K, err)
	}
	os.Rename("./image2.png", "../DSMM_WORDDOC_template_unzipped/word/media/image2.png")
	// Create new Word template with updated media folder, image4.png and image2.png
	os.Chdir("../DSMM_WORDDOC_template_unzipped")
	cmdName = "zip"
	cmdArgs = []string{"-r", "../zippedWord.zip", "_rels", "customXml", "docProps", "word", "[Content_Types].xml"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with zipping files in DSMM_WORDDOC_template_unzipped directory, program exiting", err)
		os.Exit(1)
	}
	os.Chdir("..")
	os.Rename("./zippedWord.zip", "./DSMM_WORDDOC_template/IRDSMMTemplate_Body_Rev_1.3.docx")
}
