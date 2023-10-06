package isEmail

import (
	"fmt"
	"testing"
)

func TestIsEmail(t *testing.T) {

	testSlice := []string{"hello", "Tony@gmail.com"}

	if isEmail("Hello World") == false {
		t.Errorf(" %v is not a valid email address\n", "Hello World")
	} else {
		fmt.Printf("%s is an email address\n", "Hello World")
	}

	if isEmail(testSlice[0]) == false {
		t.Errorf(" %v is not a valid email address\n", testSlice[0])
	} else {
		fmt.Printf("%s is an email address\n", testSlice[0])
	}

	if isEmail(testSlice[1]) == false {
		t.Errorf(" %v is not a valid email address\n", testSlice[1])
	} else {
		fmt.Printf("%s is an email address\n", testSlice[1])
	}

}
