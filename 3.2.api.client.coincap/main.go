package main

import (
	"coincap/coincap"
	"fmt"
	"log"
	"time"
)

func main() {
	coincapCtient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapCtient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	bitcoin, err := coincapCtient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin.Info())
}
