package countwords

import (
	"regexp"
	"sort"
	"strings"
)

func Run(sentence string) map[string]int {
	cleanSentence := cleanInput(sentence)
	sortedWords := sortedArrayFrom(cleanSentence)

	count := map[string]int{}
	for _, v := range sortedWords {
		if count[v] != 0 {
			count[v] = count[v] + 1
		} else {
			count[v] = 1
		}
	}
	return count
}

func cleanInput(input string) string {
	lowerCase := strings.ToLower(input)

	reg := regexp.MustCompile(`[^\p{L}\p{N} ]+`)
	cleanedInput := reg.ReplaceAllString(lowerCase, "")
	return cleanedInput
}

func sortedArrayFrom(input string) []string {
	wordsArray := strings.Split(input, " ")
	sort.Strings(wordsArray)
	return wordsArray
}
