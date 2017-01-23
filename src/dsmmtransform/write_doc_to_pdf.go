package main 

import (
	"fmt"
	"os"
	"os/exec"
)

func writeDocToPdf(singleRecord DsmmAssessmentRecord) {
	os.Chdir("./output")
	cmdName := "libreoffice"
	cmdArgs := []string{"--headless", "--convert-to", "pdf", singleRecord.C + "_" + singleRecord.K + ".docx"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with converting .docx to .pdf: %s_%s.docx\n", singleRecord.C, singleRecord.K, err)
	}
	os.Chdir("..")
}