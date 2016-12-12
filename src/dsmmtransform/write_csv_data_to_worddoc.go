package main

import (
	"github.com/nguyenthenguyen/docx"
)

func writeCsvDataToWordDoc(allRecords []DsmmAssessmentRecord) {

	for i, _ := range allRecords {
		r, err := docx.ReadDocxFile("DSMM_WORDDOC_template/IRDSMMTemplate_Body.docx")
		checkError("read doc template file failed", err)
		defer r.Close()

		docx := r.Editable()

		docx.Replace("DSMM_A", allRecords[i].A, -1)
		docx.Replace("DSMM_B", allRecords[i].B, -1)
		docx.Replace("DSMM_C", allRecords[i].C, -1)
		docx.Replace("DSMM_D", allRecords[i].D, -1)
		docx.Replace("DSMM_E", allRecords[i].E, -1)
		docx.Replace("DSMM_F", allRecords[i].F, -1)
		docx.Replace("DSMM_G", allRecords[i].G, -1)
		docx.Replace("DSMM_H", allRecords[i].H, -1)
		docx.Replace("DSMM_I", allRecords[i].I, -1)
		docx.Replace("DSMM_J", allRecords[i].J, -1)
		docx.Replace("DSMM_K", allRecords[i].K, -1)
		docx.Replace("DSMM_L", allRecords[i].L, -1)
		docx.Replace("DSMM_M", allRecords[i].M, -1)
		docx.Replace("DSMM_N", allRecords[i].N, -1)
		docx.Replace("DSMM_O", allRecords[i].O, -1)
		docx.Replace("DSMM_P", allRecords[i].P, -1)
		docx.Replace("DSMM_Q", allRecords[i].Q, -1)
		docx.Replace("DSMM_R", allRecords[i].R, -1)
		docx.Replace("DSMM_S", allRecords[i].S, -1)
		docx.Replace("DSMM_T", allRecords[i].T, -1)
		docx.Replace("DSMM_U", allRecords[i].U, -1)
		docx.Replace("DSMM_V", allRecords[i].V, -1)
		docx.Replace("DSMM_W", allRecords[i].W, -1)
		docx.Replace("DSMM_X", allRecords[i].X, -1)
		docx.Replace("DSMM_Y", allRecords[i].Y, -1)

		if allRecords[i].AA == "" {
			docx.Replace("DSM_AA", "N/A", -1)
		} else {
			docx.Replace("DSM_AA", allRecords[i].AA, -1)
		}

		if allRecords[i].AB == "" {
			docx.Replace("DSM_AB", "N/A", -1)
		} else {
			docx.Replace("DSM_AB", allRecords[i].AB, -1)
		}

		if allRecords[i].AC == "" {
			docx.Replace("DSM_AC", "N/A", -1)
		} else {
			docx.Replace("DSM_AC", allRecords[i].AC, -1)
		}

		if allRecords[i].AD == "" {
			docx.Replace("DSM_AD", "N/A", -1)
		} else {
			docx.Replace("DSM_AD", allRecords[i].AD, -1)
		}

		if allRecords[i].AE == "" {
			docx.Replace("DSM_AE", "N/A", -1)
		} else {
			docx.Replace("DSM_AE", allRecords[i].AE, -1)
		}

		if allRecords[i].AF == "" {
			docx.Replace("DSM_AF", "N/A", -1)
		} else {
			docx.Replace("DSM_AF", allRecords[i].AF, -1)
		}

		if allRecords[i].AG == "" {
			docx.Replace("DSM_AG", "N/A", -1)
		} else {
			docx.Replace("DSM_AG", allRecords[i].AG, -1)
		}

		if allRecords[i].AH == "" {
			docx.Replace("DSM_AH", "N/A", -1)
		} else {
			docx.Replace("DSM_AH", allRecords[i].AH, -1)
		}

		if allRecords[i].AI == "" {
			docx.Replace("DSM_AI", "N/A", -1)
		} else {
			docx.Replace("DSM_AI", allRecords[i].AI, -1)
		}

		if allRecords[i].AJ == "" {
			docx.Replace("DSM_AJ", "N/A", -1)
		} else {
			docx.Replace("DSM_AJ", allRecords[i].AJ, -1)
		}

		if allRecords[i].AK == "" {
			docx.Replace("DSM_AK", "N/A", -1)
		} else {
			docx.Replace("DSM_AK", allRecords[i].AK, -1)
		}

		if allRecords[i].AL == "" {
			docx.Replace("DSM_AL", "N/A", -1)
		} else {
			docx.Replace("DSM_AL", allRecords[i].AL, -1)
		}

		if allRecords[i].AM == "" {
			docx.Replace("DSM_AM", "N/A", -1)
		} else {
			docx.Replace("DSM_AM", allRecords[i].AM, -1)
		}

		if allRecords[i].AN == "" {
			docx.Replace("DSM_AN", "N/A", -1)
		} else {
			docx.Replace("DSM_AN", allRecords[i].AN, -1)
		}

		if allRecords[i].AO == "" {
			docx.Replace("DSM_AO", "N/A", -1)
		} else {
			docx.Replace("DSM_AO", allRecords[i].AO, -1)
		}

		if allRecords[i].AP == "" {
			docx.Replace("DSM_AP", "N/A", -1)
		} else {
			docx.Replace("DSM_AP", allRecords[i].AP, -1)
		}

		if allRecords[i].AQ == "" {
			docx.Replace("DSM_AQ", "N/A", -1)
		} else {
			docx.Replace("DSM_AQ", allRecords[i].AQ, -1)
		}

		if allRecords[i].AR == "" {
			docx.Replace("DSM_AR", "N/A", -1)
		} else {
			docx.Replace("DSM_AR", allRecords[i].AR, -1)
		}

		if allRecords[i].AS == "" {
			docx.Replace("DSM_AS", "N/A", -1)
		} else {
			docx.Replace("DSM_AS", allRecords[i].AS, -1)
		}

		if allRecords[i].AT == "" {
			docx.Replace("DSM_AT", "N/A", -1)
		} else {
			docx.Replace("DSM_AT", allRecords[i].AT, -1)
		}

		if allRecords[i].AU == "" {
			docx.Replace("DSM_AU", "N/A", -1)
		} else {
			docx.Replace("DSM_AU", allRecords[i].AU, -1)
		}

		if allRecords[i].AV == "" {
			docx.Replace("DSM_AV", "N/A", -1)
		} else {
			docx.Replace("DSM_AV", allRecords[i].AV, -1)
		}

		if allRecords[i].AW == "" {
			docx.Replace("DSM_AW", "N/A", -1)
		} else {
			docx.Replace("DSM_AW", allRecords[i].AW, -1)
		}

		if allRecords[i].AX == "" {
			docx.Replace("DSM_AX", "N/A", -1)
		} else {
			docx.Replace("DSM_AX", allRecords[i].AX, -1)
		}

		if allRecords[i].AY == "" {
			docx.Replace("DSM_AY", "N/A", -1)
		} else {
			docx.Replace("DSM_AY", allRecords[i].AY, -1)
		}
		// Write to file with Dataset Short Name and version number as doc title
			docx.WriteToFile("./output/" + allRecords[i].C + "_" + allRecords[i].K + ".docx")

	} // end writeCsvDataToWordDoc
}
