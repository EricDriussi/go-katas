package fuzz_test

import (
	"katas/02_anagram"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func FuzzAnagram(f *testing.F) {
	randPairs := map[string]string{"testing": "testing", "amor": "roma", "brag": "grab"}
	for k, v := range randPairs {
		f.Add(k, v)
	}

	f.Fuzz(func(t *testing.T, w1 string, w2 string) {
		areAnagrams := anagram.Run(w1, w2)

		w1, w2 = strings.ToLower(w1), strings.ToLower(w2)
		if len(w1) != len(w2) && areAnagrams {
			t.Fail()
		}
		if w1 == w2 && areAnagrams {
			t.Fail()
		}
		a1, a2 := strings.Split(w1, ""), strings.Split(w2, "")
		sort.Strings(a1)
		sort.Strings(a2)
		if !reflect.DeepEqual(a1, a2) && areAnagrams {
			t.Fail()
		}
	})
}
