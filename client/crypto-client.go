package client

import (
	"encoding/json"
	"fmt"
	"jcrypt/model"
	"log"
	"net/http"
)

func FetchCrypto(fiat string, crypto string) (string, error) {
	key := "3990ec554a414b59dd85d29b2286dd85"
	URL := fmt.Sprintf("https://api.nomics.com/v1/currencies/ticker?key=%s&interval=1h&ids=%s&convert=%s", key, crypto, fiat)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("oops an error occurred, please try again")
	}
	defer resp.Body.Close()

	var cResp model.CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("oops an error occurred, please try again")
	}
	return cResp.TextOutput(), nil
}
