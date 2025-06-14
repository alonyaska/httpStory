package cryptoprice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchAssets() ([]AssetData, error) {
	client := &http.Client{
		Transport: &LoggingRoundTripper{
			Logger: os.Stdout,
			Next:   http.DefaultTransport,
		},
	}

	resp, err := client.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []AssetData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
