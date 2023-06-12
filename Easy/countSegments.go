package main

import (
	"strings"
)

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	s = strings.TrimSpace(s)

	wordsB := []byte(s)

	var word string
	var words []string

	for i := range wordsB {
		word += string(wordsB[i])

		if string(wordsB[i]) == " " {
			words = append(words, word)
			word = ""
		}
		if i == len(wordsB)-1 {
			words = append(words, word)
		}
	}

	var trimmedWords []string
	for word := range words {
		if words[word] == " " {
			continue
		} else {
			trimmedWords = append(trimmedWords, words[word])
		}
	}
	return len(trimmedWords)
}
