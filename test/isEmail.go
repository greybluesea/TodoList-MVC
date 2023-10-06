package isEmail

import "regexp"

func isEmail(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(s)
}
