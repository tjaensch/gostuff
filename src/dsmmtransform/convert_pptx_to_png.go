package main

import (
  "fmt"
  "os"
  "os/exec"
)

func convertPptxToPng(singleRecord DsmmAssessmentRecord) {
  os.Chdir("./output")
	cmdName := "libreoffice"
	cmdArgs := []string{"--headless", "--convert-to", "png", singleRecord.C + "_Scoreboard_rating_template.pptx" }
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with converting .pptx to .png: %s_Scoreboard_rating_template.pptx\n", singleRecord.C, err)
	}
  cmdName = "libreoffice"
	cmdArgs = []string{"--headless", "--convert-to", "png", singleRecord.C + "_Star_rating_template.pptx" }
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Printf("something went wrong with converting .pptx to .png: %s_Star_rating_template.pptx\n", singleRecord.C, err)
	}
	os.Chdir("..")
}
