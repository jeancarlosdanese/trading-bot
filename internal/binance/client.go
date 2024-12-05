package binance

import (
	"context"
	"log"

	"github.com/adshao/go-binance/v2"
)

func NewBinanceClient(apiKey, apiSecret string) *binance.Client {
	client := binance.NewClient(apiKey, apiSecret)
	client.BaseURL = "https://testnet.binance.vision"

	return client
}

func GetTopPairs(client *binance.Client) []string {
	// Exemplo: Filtrar pares com alta liquidez
	symbols := []string{}
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		log.Fatalf("Erro ao buscar preços: %v", err)
	}

	for _, price := range prices {
		if price.Symbol == "BTCUSDT" || price.Symbol == "ETHUSDT" {
			symbols = append(symbols, price.Symbol)
		}
	}

	return symbols
}
