package main

import "strings"

func reverseWords(s string) string {
	var word string
	var words []string

	for i := range s {
		word += string(s[i])

		if string(s[i]) == " " {
			words = append(words, word)
			word = ""
		}
		if i == len(s)-1 {
			words = append(words, word)
		}
	}

	var revWords string
	for i := 0; i < len(words); i++ {
		revWords += string(reverseString(words[i]))
		if i == len(words)-2 {
			revWords += " "
		}
	}
	revWords = strings.TrimSpace(revWords)
	return revWords
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
