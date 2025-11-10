package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectSlice := subjects[:3] // slicing
	// subjects[0] = "Java"
	subjectSlice[0] = "Database"
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("=================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
