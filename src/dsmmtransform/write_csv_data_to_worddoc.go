package main

import (
	"github.com/nguyenthenguyen/docx"
)

func writeCsvDataToWordDoc(singleRecord DsmmAssessmentRecord) {

		r, err := docx.ReadDocxFile("DSMM_WORDDOC_template/IRDSMMTemplate_Body_Rev_1.3.docx")
		checkError("read doc template file failed", err)
		defer r.Close()

		docx := r.Editable()

		docx.Replace("DSMM_A", singleRecord.A, -1)
		docx.Replace("DSMM_B", singleRecord.B, -1)
		docx.Replace("DSMM_C", singleRecord.C, -1)
		docx.Replace("DSMM_D", singleRecord.D, -1)
		docx.Replace("DSMM_E", singleRecord.E, -1)
		docx.Replace("DSMM_F", singleRecord.F, -1)
		docx.Replace("DSMM_G", singleRecord.G, -1)
		docx.Replace("DSMM_H", singleRecord.H, -1)
		docx.Replace("DSMM_I", singleRecord.I, -1)
		docx.Replace("DSMM_J", singleRecord.J, -1)
		docx.Replace("DSMM_K", singleRecord.K, -1)
		docx.Replace("DSMM_L", singleRecord.L, -1)
		docx.Replace("DSMM_M", singleRecord.M, -1)
		docx.Replace("DSMM_N", singleRecord.N, -1)
		docx.Replace("DSMM_O", singleRecord.O, -1)
		docx.Replace("DSMM_P", singleRecord.P, -1)
		docx.Replace("DSMM_Q", singleRecord.Q, -1)
		docx.Replace("DSMM_R", singleRecord.R, -1)
		docx.Replace("DSMM_S", singleRecord.S, -1)
		docx.Replace("DSMM_T", singleRecord.T, -1)
		docx.Replace("DSMM_U", singleRecord.U, -1)
		docx.Replace("DSMM_V", singleRecord.V, -1)
		docx.Replace("DSMM_W", singleRecord.W, -1)
		docx.Replace("DSMM_X", singleRecord.X, -1)
		docx.Replace("DSMM_Y", singleRecord.Y, -1)

		if singleRecord.AA == "" {
			docx.Replace("DSM_AA", "N/A", -1)
		} else {
			docx.Replace("DSM_AA", singleRecord.AA, -1)
		}

		if singleRecord.AB == "" {
			docx.Replace("DSM_AB", "N/A", -1)
		} else {
			docx.Replace("DSM_AB", singleRecord.AB, -1)
		}

		if singleRecord.AC == "" {
			docx.Replace("DSM_AC", "N/A", -1)
		} else {
			docx.Replace("DSM_AC", singleRecord.AC, -1)
		}

		if singleRecord.AD == "" {
			docx.Replace("DSM_AD", "N/A", -1)
		} else {
			docx.Replace("DSM_AD", singleRecord.AD, -1)
		}

		if singleRecord.AE == "" {
			docx.Replace("DSM_AE", "N/A", -1)
		} else {
			docx.Replace("DSM_AE", singleRecord.AE, -1)
		}

		if singleRecord.AF == "" {
			docx.Replace("DSM_AF", "N/A", -1)
		} else {
			docx.Replace("DSM_AF", singleRecord.AF, -1)
		}

		if singleRecord.AG == "" {
			docx.Replace("DSM_AG", "N/A", -1)
		} else {
			docx.Replace("DSM_AG", singleRecord.AG, -1)
		}

		if singleRecord.AH == "" {
			docx.Replace("DSM_AH", "N/A", -1)
		} else {
			docx.Replace("DSM_AH", singleRecord.AH, -1)
		}

		if singleRecord.AI == "" {
			docx.Replace("DSM_AI", "N/A", -1)
		} else {
			docx.Replace("DSM_AI", singleRecord.AI, -1)
		}

		if singleRecord.AJ == "" {
			docx.Replace("DSM_AJ", "N/A", -1)
		} else {
			docx.Replace("DSM_AJ", singleRecord.AJ, -1)
		}

		if singleRecord.AK == "" {
			docx.Replace("DSM_AK", "N/A", -1)
		} else {
			docx.Replace("DSM_AK", singleRecord.AK, -1)
		}

		if singleRecord.AL == "" {
			docx.Replace("DSM_AL", "N/A", -1)
		} else {
			docx.Replace("DSM_AL", singleRecord.AL, -1)
		}

		if singleRecord.AM == "" {
			docx.Replace("DSM_AM", "N/A", -1)
		} else {
			docx.Replace("DSM_AM", singleRecord.AM, -1)
		}

		if singleRecord.AN == "" {
			docx.Replace("DSM_AN", "N/A", -1)
		} else {
			docx.Replace("DSM_AN", singleRecord.AN, -1)
		}

		if singleRecord.AO == "" {
			docx.Replace("DSM_AO", "N/A", -1)
		} else {
			docx.Replace("DSM_AO", singleRecord.AO, -1)
		}

		if singleRecord.AP == "" {
			docx.Replace("DSM_AP", "N/A", -1)
		} else {
			docx.Replace("DSM_AP", singleRecord.AP, -1)
		}

		if singleRecord.AQ == "" {
			docx.Replace("DSM_AQ", "N/A", -1)
		} else {
			docx.Replace("DSM_AQ", singleRecord.AQ, -1)
		}

		if singleRecord.AR == "" {
			docx.Replace("DSM_AR", "N/A", -1)
		} else {
			docx.Replace("DSM_AR", singleRecord.AR, -1)
		}

		if singleRecord.AS == "" {
			docx.Replace("DSM_AS", "N/A", -1)
		} else {
			docx.Replace("DSM_AS", singleRecord.AS, -1)
		}

		if singleRecord.AT == "" {
			docx.Replace("DSM_AT", "N/A", -1)
		} else {
			docx.Replace("DSM_AT", singleRecord.AT, -1)
		}

		if singleRecord.AU == "" {
			docx.Replace("DSM_AU", "N/A", -1)
		} else {
			docx.Replace("DSM_AU", singleRecord.AU, -1)
		}

		if singleRecord.AV == "" {
			docx.Replace("DSM_AV", "N/A", -1)
		} else {
			docx.Replace("DSM_AV", singleRecord.AV, -1)
		}

		if singleRecord.AW == "" {
			docx.Replace("DSM_AW", "N/A", -1)
		} else {
			docx.Replace("DSM_AW", singleRecord.AW, -1)
		}

		if singleRecord.AX == "" {
			docx.Replace("DSM_AX", "N/A", -1)
		} else {
			docx.Replace("DSM_AX", singleRecord.AX, -1)
		}

		if singleRecord.AY == "" {
			docx.Replace("DSM_AY", "N/A", -1)
		} else {
			docx.Replace("DSM_AY", singleRecord.AY, -1)
		}

		if singleRecord.AZ == "" {
			docx.Replace("DSM_AZ", "N/A", -1)
		} else {
			docx.Replace("DSM_AZ", singleRecord.AY, -1)
		}

		if singleRecord.BA == "" {
			docx.Replace("DSM_BA", "N/A", -1)
		} else {
			docx.Replace("DSM_BA", singleRecord.AY, -1)
		}

		if singleRecord.BB == "" {
			docx.Replace("DSM_BB", "N/A", -1)
		} else {
			docx.Replace("DSM_BB", singleRecord.AY, -1)
		}

		if singleRecord.BC == "" {
			docx.Replace("DSM_BC", "N/A", -1)
		} else {
			docx.Replace("DSM_BC", singleRecord.AY, -1)
		}

		if singleRecord.BD == "" {
			docx.Replace("DSM_BD", "N/A", -1)
		} else {
			docx.Replace("DSM_BD", singleRecord.AY, -1)
		}

		if singleRecord.BE == "" {
			docx.Replace("DSM_BE", "N/A", -1)
		} else {
			docx.Replace("DSM_BE", singleRecord.AY, -1)
		}

		if singleRecord.BF == "" {
			docx.Replace("DSM_BF", "N/A", -1)
		} else {
			docx.Replace("DSM_BF", singleRecord.AY, -1)
		}

		if singleRecord.BG == "" {
			docx.Replace("DSM_BG", "N/A", -1)
		} else {
			docx.Replace("DSM_BG", singleRecord.AY, -1)
		}

		if singleRecord.BH == "" {
			docx.Replace("DSM_BH", "N/A", -1)
		} else {
			docx.Replace("DSM_BH", singleRecord.AY, -1)
		}

		// Write to file with Dataset Short Name and version number as doc title
			docx.WriteToFile("./output/" + singleRecord.C + "_" + singleRecord.K + ".docx")
} // end writeCsvDataToWordDoc
