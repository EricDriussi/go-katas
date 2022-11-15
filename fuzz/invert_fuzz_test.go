package fuzz_test

import (
	invert "katas/06_invert"
	"testing"
)

func FuzzInvert(f *testing.F) {
	randSentences := []string{"Hola mundo", "Hello World", "SuperCaliFragiListicoSpiraliDoso"}
	for _, sentence := range randSentences {
		f.Add(sentence)
	}

	f.Fuzz(func(t *testing.T, inputSentence string) {
		t.Skip()
		inverted := invert.Run(inputSentence)

		if invert.Run(inverted) != inputSentence {
			t.Fail()
		}
	})
}
