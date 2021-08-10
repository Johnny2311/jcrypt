package model

import "fmt"

type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

func (c CryptoResponse) TextOutput() string {
	t := ""
	for _, crypto := range c {
		t += fmt.Sprintf("\nName: %s\nPrice: %s\nRank: %s\nHigh: %s\nCirculation Supply: %s\n",
			crypto.Name, crypto.Price, crypto.Rank, crypto.High, crypto.CirculatingSupply)
	}
	if t == "" {
		return  "There is something wrong, maybe a typo in the names of the currencies"
	}
	return t
}