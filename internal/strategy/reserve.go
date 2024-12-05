package strategy

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jeancarlosdanese/trading-bot/internal/config"
	"github.com/jeancarlosdanese/trading-bot/internal/utils"

	"github.com/adshao/go-binance/v2"
)

// Estrutura do Config
type Config struct {
	ReserveThreshold float64 `mapstructure:"reserve_threshold"`
}

func ManageReserves(client *binance.Client, cfg config.Config) {
	totalValueUSD, err := utils.GetPortfolioValueInUSD(client)
	if err != nil {
		log.Fatalf("Erro ao calcular valor total da carteira: %v", err)
	}

	log.Printf("Valor total da carteira em USD: %.2f\n", totalValueUSD)

	reserveThreshold := cfg.ReserveThreshold
	if totalValueUSD >= reserveThreshold {
		// Calcular o valor a ser reservado
		reserveAmount := totalValueUSD * 0.30
		log.Printf("Reservando $%.2f em USDT...\n", reserveAmount)

		// Converter ativos para USDT
		err := convertToUSDT(client, reserveAmount)
		if err != nil {
			log.Printf("Erro ao reservar fundos em USDT: %v\n", err)
		} else {
			log.Println("Reserva concluída com sucesso!")
		}
	} else {
		log.Println("A meta de $1000 ainda não foi atingida. Continuando as operações...")
	}
}

// Função para converter ativos em USDT
func convertToUSDT(client *binance.Client, amount float64) error {
	// Obter saldo de ETH
	account, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return fmt.Errorf("erro ao obter informações da conta: %v", err)
	}

	ethBalance := 0.0
	for _, balance := range account.Balances {
		if balance.Asset == "ETH" {
			ethBalance = parseToFloat(balance.Free)
			break
		}
	}

	// Obter o preço atual do ETH
	ethPrice, err := getCurrentPrice(client, "ETHUSDT")
	if err != nil {
		return fmt.Errorf("erro ao obter preço atual de ETH: %v", err)
	}

	ethNeeded := amount / ethPrice
	if ethBalance < ethNeeded {
		log.Printf("Saldo insuficiente de ETH. Necessário: %.6f ETH, Disponível: %.6f ETH\n", ethNeeded, ethBalance)

		// Ajustar o valor para o máximo possível, com margem de segurança de 95% do saldo
		amount = ethBalance * ethPrice * 0.95
		log.Printf("Ajustando para converter o máximo possível com margem: %.2f USD\n", amount)
	}

	// Verificar quantidade mínima para ordem
	if amount < 10 { // Geralmente, a Binance exige ordens mínimas de $10
		log.Printf("Valor ajustado é muito pequeno para conversão: %.2f USD\n", amount)
		return nil
	}

	// Executar ordem de venda
	order, err := client.NewCreateOrderService().
		Symbol("ETHUSDT").
		Side(binance.SideTypeSell).
		Type(binance.OrderTypeMarket).
		QuoteOrderQty(fmt.Sprintf("%.2f", amount)). // Valor em USD
		Do(context.Background())
	if err != nil {
		return err
	}

	log.Printf("Ordem executada: %+v\n", order)
	return nil
}

func getCurrentPrice(client *binance.Client, symbol string) (float64, error) {
	price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return 0, err
	}
	if len(price) == 0 {
		return 0, fmt.Errorf("preço não encontrado para o par %s", symbol)
	}
	return parseToFloat(price[0].Price), nil
}

func parseToFloat(value string) float64 {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("Erro ao converter string para float: %v", err)
	}
	return parsedValue
}
