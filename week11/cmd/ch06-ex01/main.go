package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectSlice := subjects[:3] // slicing
	subjects[0] = "Java"
	subjectSlice = append(subjectSlice, "Go")
	// subjectSlice = append(subjectSlice, "Go", "DB") // 리스트라면 list.append("Go")

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("=================")
	for i := 0; i < len(subjectSlice); i++ {
		fmt.Println(subjectSlice[i])
	}
}
