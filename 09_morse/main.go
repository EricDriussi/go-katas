package morse

import (
	"regexp"
	"strings"
)

func Run(sentence string) string {
	reg := regexp.MustCompile(`^[-.\s]*$`)
	isMorse := reg.MatchString(sentence)

	if isMorse {
		return turnMorseToNatural(sentence)
	} else {
		return turnNaturalToMorse(strings.ToUpper(sentence))
	}
}

func turnNaturalToMorse(sentence string) string {
	morseSentence := ""
	for _, rune := range sentence {
		char := string(rune)
		if char == " " {
			morseSentence += "  "
		} else {
			value, ok := naturalToMorse[char]
			if ok {
				morseSentence += value + " "
			} else {
				morseSentence += char + " "
			}
		}
	}
	return strings.Trim(morseSentence, " ")
}

func turnMorseToNatural(morse string) string {
	naturalSentence := ""
	for _, morseSeq := range strings.Split(morse, " ") {
		if morseSeq == "" {
			naturalSentence += " "
		} else {
			naturalSentence += morseToNatural[morseSeq]
		}
	}
	return removeDoubleSpaces(naturalSentence)
}

func removeDoubleSpaces(sentence string) string {
	return strings.Join(strings.Fields(sentence), " ")
}

// from https://es.wikipedia.org/wiki/Código_morse
var naturalToMorse = map[string]string{
	"A": ".-", "Ñ": "--.--", "1": ".----",
	"B": "-...", "O": "---", "2": "..---",
	"C": "-.-.", "P": ".--.", "3": "...--",
	"D": "-..", "Q": "--.-", "4": "....-",
	"E": ".", "R": ".-.", "5": ".....",
	"F": "..-.", "S": "...", "6": "-....",
	"G": "--.", "T": "-", "7": "--...",
	"H": "....", "U": "..-", "8": "---..",
	"I": "..", "V": "...-", "9": "----.",
	"J": ".---", "W": ".--", ".": ".-.-.-",
	"K": "-.-", "X": "-..-", ",": "--..--",
	"L": ".-..", "Y": "-.--", "?": "..--..",
	"M": "--", "Z": "--..", "\"": ".-..-.",
	"N": "-.", "0": "-----", "/": "-..-.",
}

var morseToNatural = map[string]string{
	".-": "A", "--.--": "Ñ", ".----": "1",
	"-...": "B", "---": "O", "..---": "2",
	"-.-.": "C", ".--.": "P", "...--": "3",
	"-..": "D", "--.-": "Q", "....-": "4",
	".": "E", ".-.": "R", ".....": "5",
	"..-.": "F", "...": "S", "-....": "6",
	"--.": "G", "-": "T", "--...": "7",
	"....": "H", "..-": "U", "---..": "8",
	"..": "I", "...-": "V", "----.": "9",
	".---": "J", ".--": "W", ".-.-.-": ".",
	"-.-": "K", "-..-": "X", "--..--": ",",
	".-..": "L", "-.--": "Y", "..--..": "?",
	"--": "M", "--..": "Z", ".-..-.": "\"",
	"-.": "N", "-----": "0", "-..-.": "/",
}
