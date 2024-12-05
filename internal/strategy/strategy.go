package strategy

import (
	"context"
	"log"

	"github.com/adshao/go-binance/v2"
)

func ExecuteTrade(client *binance.Client, symbol string, quantity string, isBuy bool) {
	side := binance.SideTypeBuy
	if !isBuy {
		side = binance.SideTypeSell
	}

	order, err := client.NewCreateOrderService().
		Symbol(symbol).
		Side(side).
		Type(binance.OrderTypeMarket).
		Quantity(quantity).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Erro ao executar ordem: %v", err)
	}

	log.Printf("Ordem executada: %+v\n", order)
}
