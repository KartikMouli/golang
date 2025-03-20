/*

Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in the previous lesson to iterate over the characters of a string.

Assignment
Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only.

*/

package main

import (
	"fmt"

)


func isValidPassword(password string) bool {
	
	len := len(password)

	if len < 5 || len >12{
		return false
	}

	uppercaseLetter,digit:= false,false

	for _,c := range password{
		if c >='A' && c<='Z'{
			uppercaseLetter = true
		}
		if c >= '0' && c <= '9'{
			digit=true
		}
	}

	return uppercaseLetter && digit

}



func main() {
	type testCase struct {
		password string
		isValid  bool
	}

	testCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
		{"VeryLongPassword1", false},
		{"Short", false},
		{"1234short", false},
		{"Short5", true},
		{"P4ssword", true},
		{"AA0Z9", true},
	}

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		result := isValidPassword(test.password)
		if result != test.isValid {
			failCount++
			fmt.Printf(`---------------------------------
Test Case %d
Password:  "%s"
Expecting: %v
Actual:    %v
Result:    Fail
`, i+1, test.password, test.isValid, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Case %d
Password:  "%s"
Expecting: %v
Actual:    %v
Result:    Pass
`, i+1, test.password, test.isValid, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}