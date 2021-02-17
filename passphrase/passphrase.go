package passphrase

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tomowatt/passphrase/dictionary"
	"github.com/tomowatt/passphrase/emoji"
)

type passphraseStruct struct {
	Passphrase string `json:"passphrase"`
}

type emojiphraseStruct struct {
	Emojiphrase string `json:"emojiphrase"`
	Emojis      string `json:"emojis"`
}

func getPassphraseSlice(seed int64) []string {
	rand.Seed(seed)
	dictionaries := dictionary.GetDefaultDictionaries()
	var passphraseSlice []string
	for _, dict := range dictionaries {
		word, err := dictionary.GetAWord(dict)
		if err != nil {
			log.Fatal(err)
		}
		passphraseSlice = append(passphraseSlice, word)
	}
	return passphraseSlice
}

func getEmojiphraseMap(seed int64) map[string]string {
	rand.Seed(seed)
	dictionaries := emoji.GetEmojiDictionaries()
	emojiphraseMap := make(map[string]string)
	for _, dict := range dictionaries {
		name, icon, err := emoji.GetAMap(dict)
		if err != nil {
			log.Fatal(err)
		}
		emojiphraseMap[name] = icon
	}
	return emojiphraseMap
}

func createPassphraseJSON(passphraseSlice []string) []byte {
	passphraseString := strings.Join(passphraseSlice, "-")
	passphraseJSON, err := json.Marshal(passphraseStruct{passphraseString})
	if err != nil {
		log.Fatal(err)
	}
	return passphraseJSON
}

func createEmojiphraseJSON(emojiphraseMap map[string]string) []byte {
	var iconSlice []string
	var nameSlice []string
	for name, icon := range emojiphraseMap {
		nameSlice = append(nameSlice, name)
		iconSlice = append(iconSlice, icon)
	}

	nameString := strings.Join(nameSlice, "-")
	iconString := strings.Join(iconSlice, "-")

	emojiphraseJSON, err := json.Marshal(emojiphraseStruct{nameString, iconString})
	if err != nil {
		log.Fatal(err)
	}
	return emojiphraseJSON
}

//
func GetPassphrase(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	passphraseSlice := getPassphraseSlice(time.Now().Unix())
	writer.Write(createPassphraseJSON(passphraseSlice))
}

//
func GetEmojiphrase(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	emojiphraseMap := getEmojiphraseMap(time.Now().Unix())
	writer.Write(createEmojiphraseJSON(emojiphraseMap))
}
