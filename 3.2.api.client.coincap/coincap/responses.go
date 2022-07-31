package coincap

import (
	"fmt"
)

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUsd"`
	VolumeUSD24h string `json:"volumeUsd24Hr"`
	PriceUSD     string `json:"priceUsd"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s 1 [RANK] %s 1 [SYMBOL] %s 1 [NAME] %s 1 [PRICE] %s",
		d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
