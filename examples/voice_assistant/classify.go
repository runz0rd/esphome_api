package main

import "strings"

func match(text string, words []string) string {
	scoreByWord := make(map[string]int)
	for _, input := range strings.Split(text, " ") {
		input = strings.Trim(strings.ToLower(input), " ,.!")
		for _, word := range words {
			word := strings.ToLower(word)
			if strings.Contains(word, input) || strings.Contains(input, word) {
				scoreByWord[word] += 1
			}
		}
	}
	var bestScoreWord string
	for word, score := range scoreByWord {
		if bestScoreWord == "" {
			bestScoreWord = word
		}
		if scoreByWord[bestScoreWord] < score {
			bestScoreWord = word
		}
	}
	return bestScoreWord
}
