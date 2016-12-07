package main

import (
    "github.com/nguyenthenguyen/docx"
)

func main() {
    r, err := docx.ReadDocxFile("./IRDSMMTemplate_Body.docx")
    if err != nil {
        panic(err)
    }
    docx1 := r.Editable()
    docx1.Replace("***DOCTITLE***", "BLAHSUCCESS", -1)
    docx1.WriteToFile("./new_result_1.docx")

    /* docx2 := r.Editable()
    docx2.Replace("old_2_1", "new_2_1", -1)
    docx2.Replace("old_2_2", "new_2_2", -1)
    docx2.WriteToFile("./new_result_2.docx") */

    r.Close()
}