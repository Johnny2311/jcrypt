package model

import (
	"fmt"
	"strconv"
)

type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

func (c CryptoResponse) TextOutput(target string, amount float64) string {
	t := ""
	for _, crypto := range c {
		t += "---------------------------------------------\n"
		t += fmt.Sprintf("Name: %s\n", crypto.Name)
		t += fmt.Sprintf("Price: %s %s\n", crypto.Price, target)
		t += fmt.Sprintf("Rank: %s\n", crypto.Rank)
		t += fmt.Sprintf("High: %s\n", crypto.High)
		t += fmt.Sprintf("Circulation Supply: %s\n", crypto.CirculatingSupply)
		if p, err := strconv.ParseFloat(crypto.Price, 64); amount != 0 && err == nil {
			t += fmt.Sprintf("\nAmount: %f\n", amount * p)
		}
		t += "---------------------------------------------\n"
	}
	if t == "" {
		return  "There is something wrong, maybe a typo in the names of the currencies"
	}
	return t
}