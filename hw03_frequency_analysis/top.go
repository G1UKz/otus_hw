package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(inputString string) []string {
	if inputString == "" {
		return []string{}
	}
	filteredString := regexp.MustCompile(`[^а-яА-Яa-zA-Z\-]`).ReplaceAllString(inputString, " ")
	splitedString := strings.Split(filteredString, " ")
	filteredSlice := FilterString(splitedString)
	outputMap := countStrings(filteredSlice)
	output := sortTop10(outputMap)
	for i := 0; i < len(output); i++ {
		if output[i] == "-" {
			output = append(output[:i], output[i+1:]...)
		}
	}
	if len(output) < 11 {
		return output
	}
	return output[0:10]
}

func FilterString(inputSlice []string) []string {
	var outputSlice []string
	for _, word := range inputSlice {
		if word != "" {
			splitedWord := strings.Split(word, " ")
			for _, word := range splitedWord {
				if word != "" {
					outputSlice = append(outputSlice, word)
				}
			}
		}
	}
	return outputSlice
}

func countStrings(inputSlice []string) map[string]int {
	countedWords := make(map[string]int)
	for _, word := range inputSlice {
		countedWords[strings.ToLower(word)]++
	}

	return countedWords
}

func sortTop10(inputMap map[string]int) []string {
	words := make([]string, 0)
	for key := range inputMap {
		words = append(words, key)
	}
	sort.SliceStable(words, func(i, j int) bool {
		if inputMap[words[i]] == inputMap[words[j]] {
			return words[i] < words[j]
		}
		return inputMap[words[i]] > inputMap[words[j]]
	})
	return words
}
