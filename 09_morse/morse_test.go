package morse_test

import (
	morse "katas/09_morse"
	"strings"
	"testing"
)

func TestNaturalToMorse(t *testing.T) {
	result := morse.Run("Hello World")
	if result != ".... . .-.. .-.. ---   .-- --- .-. .-.. -.." {
		t.Fail()
	}
}

func TestMorseToNatural(t *testing.T) {
	result := morse.Run(".... . .-.. .-.. ---   .-- --- .-. .-.. -..")
	if result != "HELLO WORLD" {
		t.Fail()
	}
}

func TestMorseToNaturalAndBack(t *testing.T) {
	morseInput := ".... . .-.. .-.. ---   .-- --- .-. .-.. -.."
	naturalSentence := morse.Run(morseInput)
	morseOutput := morse.Run(naturalSentence)
	if morseInput != morseOutput {
		t.Fail()
	}
}

func TestNaturalToMorseAndBack(t *testing.T) {
	naturalInput := "Hello World"
	morseSentence := morse.Run(naturalInput)
	naturalOutput := morse.Run(morseSentence)
	trimmedUpperCaseInput := strings.Trim(strings.ToUpper(naturalInput), " ")
	if naturalOutput != strings.ToUpper(trimmedUpperCaseInput) {
		t.Fail()
	}
}

func TestNaturalToMorseWhitespace(t *testing.T) {
	result := morse.Run(" ")
	if result != "" {
		t.Fail()
	}
}

func TestNaturalToMorseUnknown(t *testing.T) {
	result := morse.Run("!")
	if result != "!" {
		t.Fail()
	}
}

func TestNaturalToMorseEndsInSpace(t *testing.T) {
	result := morse.Run("0 ")
	if result != "-----" {
		t.Fail()
	}
}
