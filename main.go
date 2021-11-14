package main

import (
	"flag"
	"fmt"
	"jcrypt/client"
	"log"
	"strings"
)

func main() {
	var origin, target string
	var amount float64

	// bind flags to variables
	flag.StringVar(
		&origin, "p", "BTC", "Input the name of the crypto you would like to know the price of",
	)
	flag.StringVar(
		&target, "conv", "USD", "Input the name of the currency you would like to convert the crypto",
	)
	flag.Float64Var(
		&amount, "calc", 0, "Input the amount of the crypto to be converted",
	)

	flag.Parse()

	target = strings.ToUpper(target)
	origin = strings.ToUpper(origin)

	info, err := client.FetchCrypto(origin, target, amount)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(info)
}
