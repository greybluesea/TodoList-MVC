package util

import (
	"strconv"
)

var Name string = "Tony"

func IntSliceToStrSlice(intSlice []int) []string {
	var strSlice = make([]string, len(intSlice))
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = strconv.Itoa(intSlice[i])
	} //		fmt.Println("i:",i," intSlice[i]:",intSlice[i])

	return strSlice
}
