package cryptoprice

import "fmt"

type AssetData struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"current_price"`
	PriceChange float64 `json:"price_change_24h"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("[ID] - %s | [NAME] - %s | [Price] - $%.2f   [CHANGES FOR 24H ] - $%.2f ", d.ID, d.Name, d.Price, d.PriceChange)
}
