package main

import(
"testing"
)


func TestCleanInput(t *testing.T){

	cases := []struct{
		input string 
		expected []string
	}{
		{

		input: " hello world ",
		expected: []string{"hello", "world"},
		},
		{

		input: "mY name iS John  ",
		expected: []string{"my", "name", "is", "john"},
		},
		{

		input: "THIS WATER IS YELLOW",
		expected: []string{"this", "water", "is", "yellow"},
		},
		{

		input: "my Littlebrother is youNger",
		expected: []string{"my", "littlebrother", "is", "younger"},
		},
		{

		input: "we r going to BarcElona       ",
		expected: []string{"we", "r", "going", "to", "barcelona"},
		},
	}


	for _,c := range cases{
		actual := cleanInput(c.input)
		//Check the length of the actual slice against the expected slice 
		// if they do not match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected){
			t.Errorf("length of %v and %v is not the same", actual, c.expected)
		}

		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice 
			// if they do not match, use t.Errorf to print an error 
			//fail the test 
			if word != expectedWord{
				t.Errorf("%v does not equal the expected word %v", word, expectedWord)
			}	
		}
	}
}
