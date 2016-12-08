package main

import (
	"github.com/nguyenthenguyen/docx"
)

func writeCsvDataToWordDoc(allRecords []DsmmAssessmentRecord) {
	r, err := docx.ReadDocxFile("DSMM_WORDDOC_template/IRDSMMTemplate_Body.docx")
	checkError("read doc template file failed", err)
	defer r.Close()

	docx := r.Editable()
	docx.Replace("DSMM_A", allRecords[421].A, -1)
	docx.Replace("DSMM_B", allRecords[421].B, -1)
	docx.Replace("DSMM_C", allRecords[421].C, -1)
	docx.Replace("DSMM_D", allRecords[421].D, -1)
	docx.Replace("DSMM_E", allRecords[421].E, -1)
	docx.Replace("DSMM_F", allRecords[421].F, -1)
	docx.Replace("DSMM_G", allRecords[421].G, -1)
	docx.Replace("DSMM_H", allRecords[421].H, -1)
	docx.Replace("DSMM_I", allRecords[421].I, -1)
	docx.Replace("DSMM_J", allRecords[421].J, -1)
	docx.Replace("DSMM_K", allRecords[421].K, -1)
	docx.Replace("DSMM_L", allRecords[421].L, -1)
	docx.Replace("DSMM_M", allRecords[421].M, -1)
	docx.Replace("DSMM_N", allRecords[421].N, -1)
	docx.Replace("DSMM_O", allRecords[421].O, -1)
	docx.Replace("DSMM_P", allRecords[421].P, -1)
	docx.Replace("DSMM_Q", allRecords[421].Q, -1)
	docx.Replace("DSMM_R", allRecords[421].R, -1)
	docx.Replace("DSMM_S", allRecords[421].S, -1)
	docx.Replace("DSMM_T", allRecords[421].T, -1)
	docx.Replace("DSMM_U", allRecords[421].U, -1)
	docx.Replace("DSMM_V", allRecords[421].V, -1)
	docx.Replace("DSMM_W", allRecords[421].W, -1)
	docx.Replace("DSMM_X", allRecords[421].X, -1)
	docx.Replace("DSMM_Y", allRecords[421].Y, -1)

	if allRecords[421].AA == "" {
		docx.Replace("DSM_AA", "N/A", -1)
	} else {
		docx.Replace("DSM_AA", allRecords[421].AA, -1)
	}

	if allRecords[421].AB == "" {
		docx.Replace("DSM_AB", "N/A", -1)
	} else {
		docx.Replace("DSM_AB", allRecords[421].AB, -1)
	}

	if allRecords[421].AC == "" {
		docx.Replace("DSM_AC", "N/A", -1)
	} else {
		docx.Replace("DSM_AC", allRecords[421].AC, -1)
	}

	if allRecords[421].AD == "" {
		docx.Replace("DSM_AD", "N/A", -1)
	} else {
		docx.Replace("DSM_AD", allRecords[421].AD, -1)
	}

	if allRecords[421].AE == "" {
		docx.Replace("DSM_AE", "N/A", -1)
	} else {
		docx.Replace("DSM_AE", allRecords[421].AE, -1)
	}

	if allRecords[421].AF == "" {
		docx.Replace("DSM_AF", "N/A", -1)
	} else {
		docx.Replace("DSM_AF", allRecords[421].AF, -1)
	}

	if allRecords[421].AG == "" {
		docx.Replace("DSM_AG", "N/A", -1)
	} else {
		docx.Replace("DSM_AG", allRecords[421].AG, -1)
	}

	if allRecords[421].AH == "" {
		docx.Replace("DSM_AH", "N/A", -1)
	} else {
		docx.Replace("DSM_AH", allRecords[421].AH, -1)
	}

	if allRecords[421].AI == "" {
		docx.Replace("DSM_AI", "N/A", -1)
	} else {
		docx.Replace("DSM_AI", allRecords[421].AI, -1)
	}

	if allRecords[421].AJ == "" {
		docx.Replace("DSM_AJ", "N/A", -1)
	} else {
		docx.Replace("DSM_AJ", allRecords[421].AJ, -1)
	}

	if allRecords[421].AK == "" {
		docx.Replace("DSM_AK", "N/A", -1)
	} else {
		docx.Replace("DSM_AK", allRecords[421].AK, -1)
	}

	if allRecords[421].AL == "" {
		docx.Replace("DSM_AL", "N/A", -1)
	} else {
		docx.Replace("DSM_AL", allRecords[421].AL, -1)
	}

	if allRecords[421].AM == "" {
		docx.Replace("DSM_AM", "N/A", -1)
	} else {
		docx.Replace("DSM_AM", allRecords[421].AM, -1)
	}

	if allRecords[421].AN == "" {
		docx.Replace("DSM_AN", "N/A", -1)
	} else {
		docx.Replace("DSM_AN", allRecords[421].AN, -1)
	}

	if allRecords[421].AO == "" {
		docx.Replace("DSM_AO", "N/A", -1)
	} else {
		docx.Replace("DSM_AO", allRecords[421].AO, -1)
	}

	if allRecords[421].AP == "" {
		docx.Replace("DSM_AP", "N/A", -1)
	} else {
		docx.Replace("DSM_AP", allRecords[421].AP, -1)
	}

	if allRecords[421].AQ == "" {
		docx.Replace("DSM_AQ", "N/A", -1)
	} else {
		docx.Replace("DSM_AQ", allRecords[421].AQ, -1)
	}

	if allRecords[421].AR == "" {
		docx.Replace("DSM_AR", "N/A", -1)
	} else {
		docx.Replace("DSM_AR", allRecords[421].AR, -1)
	}

	if allRecords[421].AS == "" {
		docx.Replace("DSM_AS", "N/A", -1)
	} else {
		docx.Replace("DSM_AS", allRecords[421].AS, -1)
	}

	if allRecords[421].AT == "" {
		docx.Replace("DSM_AT", "N/A", -1)
	} else {
		docx.Replace("DSM_AT", allRecords[421].AT, -1)
	}

	if allRecords[421].AU == "" {
		docx.Replace("DSM_AU", "N/A", -1)
	} else {
		docx.Replace("DSM_AU", allRecords[421].AU, -1)
	}

	if allRecords[421].AV == "" {
		docx.Replace("DSM_AV", "N/A", -1)
	} else {
		docx.Replace("DSM_AV", allRecords[421].AV, -1)
	}

	if allRecords[421].AW == "" {
		docx.Replace("DSM_AW", "N/A", -1)
	} else {
		docx.Replace("DSM_AW", allRecords[421].AW, -1)
	}

	if allRecords[421].AX == "" {
		docx.Replace("DSM_AX", "N/A", -1)
	} else {
		docx.Replace("DSM_AX", allRecords[421].AX, -1)
	}

	if allRecords[421].AY == "" {
		docx.Replace("DSM_AY", "N/A", -1)
	} else {
		docx.Replace("DSM_AY", allRecords[421].AY, -1)
	}
	// Write to file with Dataset Short Name as doc title
	docx.WriteToFile("./" + allRecords[421].C + ".docx")
}
