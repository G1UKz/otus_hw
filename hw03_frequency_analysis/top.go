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
	output := formTop10(outputMap)
	for i := 0; i < len(output); i++ {
		if output[i] == "-" {
			output = append(output[:i], output[i+1:]...)
		}
	}
	if len(output) < 11 {
		return []string{"Error, less than 10 words in input text"}
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

func formTop10(inputMap map[string]int) []string {
	top10 := make([]string, 0)
	top10Sorted := make([]string, 0)
	top10Unsorted := make(map[int][]string)
	maxKey := 0
	words := make([]string, 0)
	for key := range inputMap {
		words = append(words, key)
	}
	sort.SliceStable(words, func(i, j int) bool {
		return inputMap[words[i]] > inputMap[words[j]]
	})
	top10 = append(top10, words...)
	for _, v := range top10 {
		top10Unsorted[inputMap[v]] = append(top10Unsorted[inputMap[v]], v)
	}

	for appearance := range top10Unsorted {
		if appearance > maxKey {
			maxKey = appearance
		}
	}
	for inc := maxKey; inc > 0; inc-- {
		sort.SliceStable(top10Unsorted[inc], func(i, j int) bool {
			return top10Unsorted[inc][i] < top10Unsorted[inc][j]
		})
		top10Sorted = append(top10Sorted, top10Unsorted[inc]...)
	}

	return top10Sorted
}
