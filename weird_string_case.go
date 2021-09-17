package main

import (
	"fmt"
	"strings"
	"unicode"
)

// my solution
func toWeirdCase(str string) string {

	var myStrArray []string

	for _, v := range strings.Split(str, " ") {
		myStr := ""
		for i, _ := range v {
			if i%2 == 0 {
				myStr += strings.ToUpper(string(v[i]))
			} else {
				myStr += strings.ToLower(string(v[i]))
			}
		}
		myStrArray = append(myStrArray, myStr)
	}
	return strings.Join(myStrArray, " ")
}

// Best Practice
func BestPractice(str string) string {
	cap := true
	var result []rune

	for _, c := range str {
		if c == 32 {
			result = append(result, c)
			cap = true
		} else if cap {
			result = append(result, unicode.ToUpper(c))
			cap = false
		} else {
			result = append(result, unicode.ToLower(c))
			cap = true
		}
	}
	return string(result)
}

func main() {
	got := toWeirdCase("Ab cD")
	expected := "Ab Cd"
	fmt.Println(got == expected)
	fmt.Println(got, expected)

	got = BestPractice("Ab cD")
	fmt.Println(got == expected)
	fmt.Println(got, expected)
}
