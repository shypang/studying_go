package main

import (
	"fmt"
	"strconv"
)

func solution8(s string) int {
	answer, _ := strconv.Atoi(s)
	fmt.Println("answer", answer)
	return answer
}

func main() {
	//str := "1234"
	str := "-1234"
	solution8(str)
}
