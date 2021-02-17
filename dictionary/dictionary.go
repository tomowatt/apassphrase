package dictionary

import (
	"errors"
	"math/rand"
	"strings"
)

func removeWhiteSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

//
func GetAWord(dictionary []string) (string, error) {
	if len(dictionary) == 0 {
		return "", errors.New("Empty Dictionary")
	}
	word := dictionary[rand.Intn(len(dictionary))]

	return removeWhiteSpace(word), nil
}

//
func GetDefaultDictionaries() [][]string {
	return [][]string{colours, animals, verbs, adverbs}
}
