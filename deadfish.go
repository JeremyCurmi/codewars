package main

import "fmt"

func Parse(data string) []int {
	var result []int
	var value int
	for _, letter := range data {
		switch string(letter) {
		case "i":
			value += 1
		case "d":
			value += -1
		case "s":
			value *= value
		case "o":
			result = append(result, value)
		}
	}

	if result == nil {
		return nil
	}

	return result
}

func main() {
	tests := []string{"iiisdoso", "ooo", "ioioio", "codewars", "isoisoiso", "ssdizbsnsdkkincuybidds"}
	expected := [][]int{[]int{8, 64}, []int{0, 0, 0}, []int{1, 2, 3}, []int{0}, []int{1, 4, 25}, nil}

	for i, par := range tests {
		got := Parse(par)
		fmt.Println(got, expected[i])
	}
}
