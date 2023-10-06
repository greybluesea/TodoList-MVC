package main

import (
	"fmt"
	"go-course-derek-banas/util"
	"reflect"
)

func main() {
	fmt.Println("Hello,", util.Name)
	intSlice := []int{2, 3, 5, 17, 35}

	fmt.Println(util.IntSliceToStrSlice(intSlice))
	fmt.Println(reflect.TypeOf(util.IntSliceToStrSlice(intSlice)))
}
