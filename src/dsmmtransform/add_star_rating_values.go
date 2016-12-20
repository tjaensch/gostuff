package main

func addStarRatingValues(singleRecord DsmmAssessmentRecord) DsmmAssessmentRecord {

	singleRecord.Lightgrey = `<a:srgbClr val="BFBFBF"/>`
	singleRecord.Fullgrey = `<a:schemeClr val="tx1"><a:lumMod val="50000"/><a:lumOff val="50000"/></a:schemeClr>`

	// Preservability
	rating := singleRecord.O

	if rating == "0.5" {
		singleRecord.Star1 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star1 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Fullgrey
		singleRecord.Star4 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Fullgrey
		singleRecord.Star4 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Fullgrey
		singleRecord.Star4 = singleRecord.Fullgrey
		singleRecord.Star5 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star1 = singleRecord.Fullgrey
		singleRecord.Star2 = singleRecord.Fullgrey
		singleRecord.Star3 = singleRecord.Fullgrey
		singleRecord.Star4 = singleRecord.Fullgrey
		singleRecord.Star5 = singleRecord.Fullgrey
	}

	// Accessibility
	rating = singleRecord.P

	if rating == "0.5" {
		singleRecord.Star6 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star6 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Fullgrey
		singleRecord.Star9 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Fullgrey
		singleRecord.Star9 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Fullgrey
		singleRecord.Star9 = singleRecord.Fullgrey
		singleRecord.Star10 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star6 = singleRecord.Fullgrey
		singleRecord.Star7 = singleRecord.Fullgrey
		singleRecord.Star8 = singleRecord.Fullgrey
		singleRecord.Star9 = singleRecord.Fullgrey
		singleRecord.Star10 = singleRecord.Fullgrey
	}

	// Usability
	rating = singleRecord.Q

	if rating == "0.5" {
		singleRecord.Star11 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star11 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Fullgrey
		singleRecord.Star14 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Fullgrey
		singleRecord.Star14 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Fullgrey
		singleRecord.Star14 = singleRecord.Fullgrey
		singleRecord.Star15 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star11 = singleRecord.Fullgrey
		singleRecord.Star12 = singleRecord.Fullgrey
		singleRecord.Star13 = singleRecord.Fullgrey
		singleRecord.Star14 = singleRecord.Fullgrey
		singleRecord.Star15 = singleRecord.Fullgrey
	}

	// Production Sustainability
	rating = singleRecord.R

	if rating == "0.5" {
		singleRecord.Star16 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star16 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Fullgrey
		singleRecord.Star19 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Fullgrey
		singleRecord.Star19 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Fullgrey
		singleRecord.Star19 = singleRecord.Fullgrey
		singleRecord.Star20 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star16 = singleRecord.Fullgrey
		singleRecord.Star17 = singleRecord.Fullgrey
		singleRecord.Star18 = singleRecord.Fullgrey
		singleRecord.Star19 = singleRecord.Fullgrey
		singleRecord.Star20 = singleRecord.Fullgrey
	}

	// Data Quality Assurance
	rating = singleRecord.S

	if rating == "0.5" {
		singleRecord.Star21 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star21 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Fullgrey
		singleRecord.Star24 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Fullgrey
		singleRecord.Star24 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Fullgrey
		singleRecord.Star24 = singleRecord.Fullgrey
		singleRecord.Star25 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star21 = singleRecord.Fullgrey
		singleRecord.Star22 = singleRecord.Fullgrey
		singleRecord.Star23 = singleRecord.Fullgrey
		singleRecord.Star24 = singleRecord.Fullgrey
		singleRecord.Star25 = singleRecord.Fullgrey
	}

	// Data Quality Control/Monitoring
	rating = singleRecord.T

	if rating == "0.5" {
		singleRecord.Star26 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star26 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Fullgrey
		singleRecord.Star29 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Fullgrey
		singleRecord.Star29 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Fullgrey
		singleRecord.Star29 = singleRecord.Fullgrey
		singleRecord.Star30 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star26 = singleRecord.Fullgrey
		singleRecord.Star27 = singleRecord.Fullgrey
		singleRecord.Star28 = singleRecord.Fullgrey
		singleRecord.Star29 = singleRecord.Fullgrey
		singleRecord.Star30 = singleRecord.Fullgrey
	}

	// Data Quality Assessment Rating
	rating = singleRecord.U

	if rating == "0.5" {
		singleRecord.Star31 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star31 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Fullgrey
		singleRecord.Star34 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Fullgrey
		singleRecord.Star34 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Fullgrey
		singleRecord.Star34 = singleRecord.Fullgrey
		singleRecord.Star35 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star31 = singleRecord.Fullgrey
		singleRecord.Star32 = singleRecord.Fullgrey
		singleRecord.Star33 = singleRecord.Fullgrey
		singleRecord.Star34 = singleRecord.Fullgrey
		singleRecord.Star35 = singleRecord.Fullgrey
	}

	// Transparency/Traceability
	rating = singleRecord.V

	if rating == "0.5" {
		singleRecord.Star36 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star36 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Fullgrey
		singleRecord.Star39 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Fullgrey
		singleRecord.Star39 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Fullgrey
		singleRecord.Star39 = singleRecord.Fullgrey
		singleRecord.Star40 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star36 = singleRecord.Fullgrey
		singleRecord.Star37 = singleRecord.Fullgrey
		singleRecord.Star38 = singleRecord.Fullgrey
		singleRecord.Star39 = singleRecord.Fullgrey
		singleRecord.Star40 = singleRecord.Fullgrey
	}

	// Data Integrity Rating
	rating = singleRecord.W

	if rating == "0.5" {
		singleRecord.Star41 = singleRecord.Lightgrey
	} else if rating == "1" {
		singleRecord.Star41 = singleRecord.Fullgrey
	} else if rating == "1.5" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Lightgrey
	} else if rating == "2" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
	} else if rating == "2.5" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Lightgrey
	} else if rating == "3" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Fullgrey
	} else if rating == "3.5" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Fullgrey
		singleRecord.Star44 = singleRecord.Lightgrey
	} else if rating == "4" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Fullgrey
		singleRecord.Star44 = singleRecord.Fullgrey
	} else if rating == "4.5" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Fullgrey
		singleRecord.Star44 = singleRecord.Fullgrey
		singleRecord.Star45 = singleRecord.Lightgrey
	} else if rating == "5" {
		singleRecord.Star41 = singleRecord.Fullgrey
		singleRecord.Star42 = singleRecord.Fullgrey
		singleRecord.Star43 = singleRecord.Fullgrey
		singleRecord.Star44 = singleRecord.Fullgrey
		singleRecord.Star45 = singleRecord.Fullgrey
	}

	//Return updated record
	return singleRecord

}