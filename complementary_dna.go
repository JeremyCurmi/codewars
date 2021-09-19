package main

import (
	"fmt"
	"strings"
)

func DNAStrand(dna string) string {
	result := ""
	pairs := map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}

	for _, letter := range dna {
		result += pairs[string(letter)]
	}
	return result
}

func BestPractice(dna string) string {
	replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	return replacer.Replace(dna)
}

func main() {

	test := make(map[string]string, 3)
	test["AAAA"] = "TTTT"
	test["ATTGC"] = "TAACG"
	test["GTAT"] = "CATA"

	for args, expected := range test {
		fmt.Println(DNAStrand(args) == expected)
	}

	for args, expected := range test {
		fmt.Println(BestPractice(args) == expected)
	}

}
