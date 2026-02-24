package main

import (
	"fmt"
)

func MultiTable(number int) (string) {
	var table string
	for i := 1; i <= 10; i++ {
		result := i * number
		s := fmt.Sprintf("%d * %d = %d\n", i, number, result)
		table += s
	}

	return table
}

func main() {
	fmt.Println(MultiTable(2))
}