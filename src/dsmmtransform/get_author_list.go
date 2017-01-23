package main

import (
	"fmt"
	"strings"
)

var authorName string

func getAuthorList(singleRecord DsmmAssessmentRecord) DsmmAssessmentRecord {
	s := make([]string, 1)
	authors := strings.Split(singleRecord.M, "; ")
	for _, author := range authors {
		name := strings.Split(author, ",")
		lastNameFirstName := strings.Split(name[0], " ")

		
			// one first name and one last name
			if len(lastNameFirstName) == 2 {
				firstName := lastNameFirstName[0]
				if len(s) == 1 {
					authorName = lastNameFirstName[1] + ", " + string(firstName[0]) + "."
				} else {
					authorName = string(firstName[0]) + ". " + lastNameFirstName[1]
				}
				s = append(s, authorName)
			}
			// one first name and two last names
			if len(lastNameFirstName) == 3 {
				firstName := lastNameFirstName[0]
				if len(s) == 1 {
					authorName = lastNameFirstName[1] + " " + lastNameFirstName[2] + ", " + string(firstName[0]) + ". "
				} else {
					authorName = string(firstName[0]) + ". " + lastNameFirstName[1] + " " + lastNameFirstName[2]
				}
				s = append(s, authorName)
			}
			// two first names and two last names
			if len(lastNameFirstName) == 4 {
				firstName := lastNameFirstName[0]
				secondFirstName := lastNameFirstName[1]
				if len(s) == 1 {
					authorName = lastNameFirstName[2] + " " + lastNameFirstName[3] + ", " + string(firstName[0]) + "."  + string(secondFirstName)
				} else {
					authorName = string(firstName[0]) + "."  + string(secondFirstName) + " " + lastNameFirstName[2] + " " + lastNameFirstName[3]
				}
				s = append(s, authorName)
			}
		

		if len(authors) == 1 {
			singleRecord.Authorlist = authorName
			fmt.Println(singleRecord.Authorlist)
		} else if author == authors[len(authors)-1] {
			s = s[:len(s)-1]
			s = append(s, "and " + authorName)
			Authorlist := strings.Join(s, ", ")
			singleRecord.Authorlist = Authorlist[2:]
			fmt.Println(singleRecord.Authorlist)
		}
	}
	
	return singleRecord
}
