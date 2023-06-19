package main

import (
	"fmt"
	"strings"
)

func main(){
	//teste := "Anotaram a data da maratona"
	//teste := "Ovo"
	teste := "Never odd or even"
	//teste := "Frase não palindrome"
	text := isPalindrome(teste)
	fmt.Println(text)
}
func isPalindrome(text string) bool {
	if len(text) == 0 {
		return false
	}
	input := normalizeInput(text)
	return invertSentence(input) == input
}

func invertSentence (text string) string {
	invertedText := []byte{}
	for i := len(text) - 1; i >= 0; i--{
		invertedText = append(invertedText, text[i])
	}
	invert := string(invertedText)
	return invert
}
func normalizeInput(input string) string {
	normalize := strings.ToLower(input)
	normalize = strings.ReplaceAll(normalize, " ", "")
	return normalize
}