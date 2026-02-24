package main

import (
	"strings"
)

func RepeatStr(repetitions int, value string) (string) {
	a := []string{}

	for r := 1; r <= repetitions; r++ {
		a = append(a, value)
	}

	result := strings.Join(a, "")

	return result
}

