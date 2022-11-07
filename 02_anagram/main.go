package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func Run(w1 string, w2 string) bool {
	w1, w2 = strings.ToLower(w1), strings.ToLower(w2)
	if len(w1) != len(w2) {
		return false
	} else if w1 == w2 {
		return false
	} else {
		a1, a2 := strings.Split(w1, ""), strings.Split(w2, "")
		sort.Strings(a1)
		sort.Strings(a2)
		return reflect.DeepEqual(a1, a2)
	}
}
