package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Max(numbers []int) (max int) {
	max = numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(numbers []int) (min int) {
	min = numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

func HighAndLow(in string) string {
	var numbers []int

	numbersText := strings.Fields(in)

	for _, num := range numbersText {
		val, _ := strconv.Atoi(num)
		numbers = append(numbers, val)
	}

	result := []string{strconv.Itoa(Max(numbers))}
	result = append(result, strconv.Itoa(Min(numbers)))

	return strings.Join(result, " ")
}

func main() {
	test := "1 2 3 4 5"
	expected := "5 1"
	got := HighAndLow(test)
	fmt.Println("got", got, "expected", expected)
	fmt.Println(got == expected)
}
