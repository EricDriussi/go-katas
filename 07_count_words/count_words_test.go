package countwords_test

import (
	"katas/07_count_words"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	result := countwords.Run("Hi my name is peter, full name is Peter Parker (PeterMan).")
	expected := map[string]int{
		"full":     1,
		"hi":       1,
		"is":       2,
		"my":       1,
		"name":     2,
		"parker":   1,
		"peter":    2,
		"peterman": 1,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
