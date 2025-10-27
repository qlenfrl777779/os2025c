package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func main() {
	fmt.Print("정수입력 :")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	passFail := ""
	if score >= 60 {
		passFail = "합격!"
	} else {
		passFail = "불합격"
	}
	fmt.Printf("%.2f점은 %v\n", score, passFail)
}
