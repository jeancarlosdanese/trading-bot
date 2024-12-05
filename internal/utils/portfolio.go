package utils

import (
	"context"
	"log"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

func GetPortfolioValueInUSD(client *binance.Client) (float64, error) {
	// Obter saldo atual da conta
	account, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return 0, err
	}

	totalValueUSD := 0.0

	// Obter preços de mercado para os pares principais
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		return 0, err
	}

	// Criar um mapa de preços para facilitar o acesso
	priceMap := make(map[string]float64)
	for _, price := range prices {
		priceMap[price.Symbol] = parseToFloat(price.Price)
	}

	for _, balance := range account.Balances {
		asset := balance.Asset
		free := parseToFloat(balance.Free)
		locked := parseToFloat(balance.Locked)

		total := free + locked
		if total > 0 {
			if asset == "USDT" || asset == "FDUSD" {
				totalValueUSD += total
			} else if price, ok := priceMap[asset+"USDT"]; ok {
				if total*price < 0.01 {
					log.Printf("Saldo irrelevante para o ativo: %s, valor total: %.2f\n", asset, total*price)
					continue
				}
				totalValueUSD += total * price
			} else {
				log.Printf("Sem preço para o ativo: %s. Ignorando...\n", asset)
				continue
			}
		}
	}

	return totalValueUSD, nil
}

// Função auxiliar para converter string para float64
func parseToFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return v
}
