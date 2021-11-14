package main

import (
	"flag"
	"fmt"
	"jcrypt/client"
	"log"
	"strings"
)

func main() {
	var fiat, crypto string

	// bind flags to variables
	flag.StringVar(
		&fiat, "fiat", "USD", "The name of the fiat currency you would like to know the price of your crypto in",
	)
	flag.StringVar(
		&crypto, "crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)

	flag.Parse()

	fiat = strings.ToUpper(fiat)
	crypto = strings.ToUpper(crypto)

	info, err := client.FetchCrypto(fiat, crypto)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(info)
}
