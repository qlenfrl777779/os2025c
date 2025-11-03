package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayBool := [3]bool{true, false, true} // array literal
	var arrayInt [3]int
	fmt.Println(reflect.TypeOf(arrayBool))
	fmt.Printf("%#v\n", arrayBool)
	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])

}
