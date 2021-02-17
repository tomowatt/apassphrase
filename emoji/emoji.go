package emoji

import (
	"errors"
	"math/rand"
)

//
func GetAMap(dictionary []map[string]string) (string, string, error) {
	if len(dictionary) == 0 {
		return "", "", errors.New("Empty Dictionary")
	}
	for icon, name := range dictionary[rand.Intn(len(dictionary))] {
		return icon, name, nil
	}
	return "", "", nil
}

//
func GetEmojiDictionaries() [][]map[string]string {
	return [][]map[string]string{sports, animals, foods, weathers}
}
