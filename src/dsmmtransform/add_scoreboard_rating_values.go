package main

func addScoreboardRatingValues(singleRecord DsmmAssessmentRecord) DsmmAssessmentRecord {

	singleRecord.Level1 = `<a:srgbClr val="E5F4E0"/>`
	singleRecord.Level2 = `<a:srgbClr val="CBEAC0"/>`
	singleRecord.Level3 = `<a:srgbClr val="B0DFA1"/>`
	singleRecord.Level4 = `<a:srgbClr val="55A839"/>`
	singleRecord.Level5 = `<a:srgbClr val="387026"/>`

	// Preservability
	switch rating := singleRecord.O; rating {
	case "0":
		singleRecord.PreservabilityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.PreservabilityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.PreservabilityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.PreservabilityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.PreservabilityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.PreservabilityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.PreservabilityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.PreservabilityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.PreservabilityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.PreservabilityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.PreservabilityLevel5 = singleRecord.Level5
	}


	// Accessibility
	switch rating := singleRecord.P; rating {
		case "0":
		singleRecord.AccessibilityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.AccessibilityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.AccessibilityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.AccessibilityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.AccessibilityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.AccessibilityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.AccessibilityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.AccessibilityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.AccessibilityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.AccessibilityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.AccessibilityLevel5 = singleRecord.Level5
	}

	// Usability
	switch rating := singleRecord.Q; rating {
		case "0":
		singleRecord.UsabilityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.UsabilityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.UsabilityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.UsabilityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.UsabilityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.UsabilityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.UsabilityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.UsabilityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.UsabilityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.UsabilityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.UsabilityLevel5 = singleRecord.Level5
	}

	// Production Sustainability
	switch rating := singleRecord.R; rating {
		case "0":
		singleRecord.ProductionSustainabilityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.ProductionSustainabilityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.ProductionSustainabilityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.ProductionSustainabilityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.ProductionSustainabilityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.ProductionSustainabilityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.ProductionSustainabilityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.ProductionSustainabilityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.ProductionSustainabilityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.ProductionSustainabilityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.ProductionSustainabilityLevel5 = singleRecord.Level5
	}

	// Data Quality Assurance
	switch rating := singleRecord.S; rating {
		case "0":
		singleRecord.DataQualityAssuranceLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.DataQualityAssuranceLevel1 = singleRecord.Level1
	case "1":
		singleRecord.DataQualityAssuranceLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.DataQualityAssuranceLevel1 = singleRecord.Level1
	case "2":
		singleRecord.DataQualityAssuranceLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.DataQualityAssuranceLevel2 = singleRecord.Level2
	case "3":
		singleRecord.DataQualityAssuranceLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.DataQualityAssuranceLevel3 = singleRecord.Level3
	case "4":
		singleRecord.DataQualityAssuranceLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.DataQualityAssuranceLevel4 = singleRecord.Level4
	case "5":
		singleRecord.DataQualityAssuranceLevel5 = singleRecord.Level5
	}

	// Data Quality Control Monitoring
	switch rating := singleRecord.T; rating {
		case "0":
		singleRecord.DataQualityControlMonitoringLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.DataQualityControlMonitoringLevel1 = singleRecord.Level1
	case "1":
		singleRecord.DataQualityControlMonitoringLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.DataQualityControlMonitoringLevel1 = singleRecord.Level1
	case "2":
		singleRecord.DataQualityControlMonitoringLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.DataQualityControlMonitoringLevel2 = singleRecord.Level2
	case "3":
		singleRecord.DataQualityControlMonitoringLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.DataQualityControlMonitoringLevel3 = singleRecord.Level3
	case "4":
		singleRecord.DataQualityControlMonitoringLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.DataQualityControlMonitoringLevel4 = singleRecord.Level4
	case "5":
		singleRecord.DataQualityControlMonitoringLevel5 = singleRecord.Level5
	}

	// Data Quality Assessment
	switch rating := singleRecord.U; rating {
		case "0":
		singleRecord.DataQualityAssessmentLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.DataQualityAssessmentLevel1 = singleRecord.Level1
	case "1":
		singleRecord.DataQualityAssessmentLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.DataQualityAssessmentLevel1 = singleRecord.Level1
	case "2":
		singleRecord.DataQualityAssessmentLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.DataQualityAssessmentLevel2 = singleRecord.Level2
	case "3":
		singleRecord.DataQualityAssessmentLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.DataQualityAssessmentLevel3 = singleRecord.Level3
	case "4":
		singleRecord.DataQualityAssessmentLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.DataQualityAssessmentLevel4 = singleRecord.Level4
	case "5":
		singleRecord.DataQualityAssessmentLevel5 = singleRecord.Level5
	}

	// Transparency/Traceability
	switch rating := singleRecord.V; rating {
		case "0":
		singleRecord.TransparencyTraceabilityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.TransparencyTraceabilityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.TransparencyTraceabilityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.TransparencyTraceabilityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.TransparencyTraceabilityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.TransparencyTraceabilityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.TransparencyTraceabilityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.TransparencyTraceabilityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.TransparencyTraceabilityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.TransparencyTraceabilityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.TransparencyTraceabilityLevel5 = singleRecord.Level5
	}

	// Data Integrity
	switch rating := singleRecord.W; rating {
		case "0":
		singleRecord.DataIntegrityLevel1 = singleRecord.Level1
	case "0.5":
		singleRecord.DataIntegrityLevel1 = singleRecord.Level1
	case "1":
		singleRecord.DataIntegrityLevel1 = singleRecord.Level1
	case "1.5":
		singleRecord.DataIntegrityLevel1 = singleRecord.Level1
	case "2":
		singleRecord.DataIntegrityLevel2 = singleRecord.Level2
	case "2.5":
		singleRecord.DataIntegrityLevel2 = singleRecord.Level2
	case "3":
		singleRecord.DataIntegrityLevel3 = singleRecord.Level3
	case "3.5":
		singleRecord.DataIntegrityLevel3 = singleRecord.Level3
	case "4":
		singleRecord.DataIntegrityLevel4 = singleRecord.Level4
	case "4.5":
		singleRecord.DataIntegrityLevel4 = singleRecord.Level4
	case "5":
		singleRecord.DataIntegrityLevel5 = singleRecord.Level5
	}

	//Return updated record
	return singleRecord

}
