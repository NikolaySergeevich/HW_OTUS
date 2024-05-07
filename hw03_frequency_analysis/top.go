package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const maxWord = 10

var reg = regexp.MustCompile(`(?:\p{L}+(?:-\p{L}+)*)|(?:\p{L}-\p{L})`)

var taskWithAsteriskIsCompleted = false

type kv struct {
	key   string
	value int
}

var sortWord []kv

func Top10(text string) []string {
	res := make([]string, 0, maxWord)
	indiv := make(map[string]int)
	parsText := parsText(taskWithAsteriskIsCompleted, text)
	for _, v := range parsText {
		indiv[v]++
	}
	for k, v := range indiv {
		sortWord = append(sortWord, kv{k, v})
	}
	sort.Slice(sortWord, func(i, j int) bool {
		if sortWord[i].value == sortWord[j].value {
			return sortWord[i].key < sortWord[j].key
		}
		return sortWord[i].value > sortWord[j].value
	})
	for i := 0; i < maxWord && i < len(sortWord); i++ {
		res = append(res, sortWord[i].key)
	}

	return res
}

func parsText(flag bool, text string) []string {
	if flag {
		text = strings.ToLower(text)
		return reg.FindAllString(text, -1)
	}
	return strings.Fields(text)
}
