package passphrase

import (
	"dictionary"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type passphraseStruct struct {
	Passphrase string `json:"passphrase"`
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

func createPassphraseJSON(passphraseSlice []string) []byte {
	passphraseString := strings.Join(passphraseSlice, "-")
	passphraseJSON, err := json.Marshal(passphraseStruct{passphraseString})
	if err != nil {
		log.Fatal(err)
	}
	return passphraseJSON
}

//
func GetPassphrase(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	passphraseSlice := getPassphraseSlice(time.Now().Unix())
	writer.Write(createPassphraseJSON(passphraseSlice))
}
