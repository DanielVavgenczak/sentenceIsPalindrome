package main

import (
	"testing"
)

type SentenceCase struct {
	input string
	result bool
}
type SentenceInvertCase struct {
	input string
	result string
}
func TestIsPalindrome(t *testing.T) {
	testSentence := []SentenceCase{
		{
			input: "A base do teto desaba",
			result: true,
		},
		{
			input: "Anotaram a data da maratona",
			result: true,
		},
		{
			input:  "Para retornar falso",
			result: false,
		},
	}

	for _, testSent := range testSentence {
		if isPalindrome(testSent.input) != testSent.result {
			t.Error("FAIL", testSent)
		}
	}
}
