package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type counter struct {
	word  string
	count int
}

func Top10(input string) []string {
	lowedInput := strings.ToLower(input)
	expression := regexp.MustCompile(`\p{L}+\-\p{L}+|[\p{L}]+`)
	slicedInput := expression.FindAllString(lowedInput, -1)

	wordMap := make(map[string]int)
	for _, word := range slicedInput {
		wordMap[word]++
	}

	wordCounter := make([]counter, 0)
	for i := range wordMap {
		wordCounter = append(wordCounter, counter{
			word:  i,
			count: wordMap[i],
		})
	}

	sort.Slice(wordCounter, func(i, j int) bool {
		return wordCounter[i].count > wordCounter[j].count ||
			wordCounter[i].count == wordCounter[j].count &&
				wordCounter[i].word < wordCounter[j].word
	})

	if len(wordCounter) > 10 {
		wordCounter = wordCounter[:10]
	}

	result := make([]string, 0)
	for _, word := range wordCounter {
		result = append(result, word.word)
	}

	return result
}
