package ranhash_test

import (
	"testing"
	"unicode"

	"github.com/newtoallofthis123/ranhash"
)

func hasNoNumbers(input string) bool {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func hasNoAlphabets(input string) bool {
	for _, char := range input {
		if unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func TestStringHash(t *testing.T) {
	hash := ranhash.GenerateRandomString(8)

	if !hasNoNumbers(hash) {
		t.Errorf("Numbers found in only string hash")
	}
}

func TestNumberHash(t *testing.T) {
	hash := ranhash.GenerateRandomNumber(8)

	if !hasNoAlphabets(hash) {
		t.Errorf("Alphabets found in only number hash")
	}
}
