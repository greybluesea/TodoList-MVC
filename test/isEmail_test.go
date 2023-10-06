package isEmail

import (
	"fmt"
	"testing"
)

func TestIsEmail(t *testing.T) {

	testSlice := []string{"hello", "Tony@gmail.com"}

	for _, str := range testSlice {

		if isEmail(str) == false {
			t.Errorf(" %s is not a valid email address\n", str)
		} else {
			fmt.Printf("%s is an email address\n", str)
		}

	}

}
