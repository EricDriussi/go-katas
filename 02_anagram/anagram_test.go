package anagram_test

import (
	"katas/02_anagram"
	"testing"
)

func TestIdenticalWordsAreNotAnagrams(t *testing.T) {
	result := anagram.Run("testing", "testing")
	if result != false {
		t.Fail()
	}
}

func TestAmorRoma(t *testing.T) {
	result := anagram.Run("amor", "roma")
	if result != true {
		t.Fail()
	}
}

func TestInchChin(t *testing.T) {
	result := anagram.Run("inch", "chin")
	if result != true {
		t.Fail()
	}
}

func TestStressedDesserts(t *testing.T) {
	result := anagram.Run("stressed", "desserts")
	if result != true {
		t.Fail()
	}
}

func TestBragGrab(t *testing.T) {
	result := anagram.Run("brag", "grab")
	if result != true {
		t.Fail()
	}
}

func TestUpperCaseWorks(t *testing.T) {
	result := anagram.Run("Brag", "Grab")
	if result != true {
		t.Fail()
	}
}
