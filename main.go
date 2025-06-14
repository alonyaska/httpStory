package main

import (
	"fmt"
	cryptoprice "golang-alonya/cryptoPrices/CryptoPrice"
	"log"
	// путь зависит от твоего go.mod
)

func main() {
	assets, err := cryptoprice.FetchAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}
}

