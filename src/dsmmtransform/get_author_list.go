package main

import (
	"strings"
)

func getAuthorList(singleRecord DsmmAssessmentRecord) DsmmAssessmentRecord {
	s := make([]string, 1)
	authors := strings.Split(singleRecord.M, "; ")
	for _, author := range authors {
		name := strings.Split(author, ",")
		lastNameFirstName := strings.Split(name[0], " ")
		s = append(s, lastNameFirstName[1]+", "+lastNameFirstName[0]+", ")
		if len(authors) == 1 {
			Authorlist := strings.Join(s, "")
			singleRecord.Authorlist = Authorlist[:len(Authorlist)-2]
		} else if author == authors[len(authors)-1] {
			s = s[:len(s)-1]
			s = append(s, "and "+lastNameFirstName[1]+", "+lastNameFirstName[0])
			Authorlist := strings.Join(s, "")
			singleRecord.Authorlist = Authorlist
		}
	}
	return singleRecord
}
