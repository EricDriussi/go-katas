package fuzz_test

import (
	"katas/09_morse"
	"strings"
	"testing"
)

func FuzzMorse(f *testing.F) {
	randSentences := []string{"Hola mundo", "Hello World", "SuperCaliFragiListicoSpiraliDoso", "A long, random test sentence"}
	for _, sentence := range randSentences {
		f.Add(sentence)
	}

	f.Fuzz(func(t *testing.T, naturalSentence string) {
		morseSentence := morse.Run(naturalSentence)
		naturalOutput := morse.Run(morseSentence)

		trimmedUpperCaseInput := strings.Trim(strings.ToUpper(naturalSentence), " ")
		if naturalOutput != trimmedUpperCaseInput {
			t.Fail()
		}
	})
}
