package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.21))
	fmt.Println(math.Ceil(2.21))
	fmt.Println(strings.Title("go developer!"))
	fmt.Println("kim\nInha\t\"\\") // \n :줄바꿈 \t: tap만큼 띄우기 \": 따옴표 출력 \\: 역슬래시 출력
	fmt.Println('2', '가', '\u18a0')
}
