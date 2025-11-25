package main


import(
	"strings"

)


func cleanInput(text string) []string{

	//lower case the input of the string 
	lc_string := strings.ToLower(text)

	// trim the whitespace in the strings 
	tr_string := strings.TrimSpace(lc_string)

	//split at the whitespaces
	split_string := strings.Fields(tr_string)
	
	return split_string
}
