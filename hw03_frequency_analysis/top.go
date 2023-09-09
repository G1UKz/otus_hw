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
	splitedString := strings.Split(inputString, " ")
	filteredSlice := FilterString(splitedString)
	outputMap := countStrings(filteredSlice)
	output := formTop10(outputMap)
	for i := 0; i < len(output); i++ {
		if output[i] == "-" {
			output = append(output[:i], output[i+1:]...)
		}
	}
	return output[0:10]
}

func FilterString(inputSlice []string) []string {
	var outputSlice []string
	for _, word := range inputSlice {
		if word != "" {
			filteredWord := regexp.MustCompile(`[^а-яА-Яa-zA-Z\-]`).ReplaceAllString(word, " ")
			splitedWord := strings.Split(filteredWord, " ")
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
		var exist bool
		count, exist := countedWords[strings.ToLower(word)]
		if exist {
			countedWords[strings.ToLower(word)] = count + 1
			exist = false
		} else {
			countedWords[strings.ToLower(word)] = 1
		}
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
