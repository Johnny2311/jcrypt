package client

import (
	"encoding/json"
	"fmt"
	"jcrypt/model"
	"log"
	"net/http"
)

// fetch info of crypto and convert from origin to target
func FetchCrypto(origin, target string, amount float64) (string, error) {
	key := "3990ec554a414b59dd85d29b2286dd85"
	URL := fmt.Sprintf("https://api.nomics.com/v1/currencies/ticker?key=%s&interval=1h&ids=%s&convert=%s", key, origin, target)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("oops an error occurred, please try again\n Error: " + err.Error())
	}
	defer resp.Body.Close()

	var cResp model.CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("oops an error occurred, please try again\n Error: " + err.Error())
	}
	return cResp.TextOutput(target, amount), nil
}
