package main

import (
	"log"
	"path/filepath"

	"github.com/jeancarlosdanese/trading-bot/internal/binance"
	"github.com/jeancarlosdanese/trading-bot/internal/config"
	"github.com/jeancarlosdanese/trading-bot/internal/strategy"
	"github.com/jeancarlosdanese/trading-bot/internal/utils"
	"github.com/spf13/viper"
)

type Config struct {
	APIKey     string   `mapstructure:"api_key"`
	APISecret  string   `mapstructure:"api_secret"`
	TradePairs []string `mapstructure:"trade_pairs"`
	Quantity   string   `mapstructure:"quantity"`
	TestMode   bool     `mapstructure:"test_mode"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	// Inicializar logger
	utils.InitLogger()
	utils.Info("Iniciando o robô de trading...")

	// Carregar configuração
	configPath := filepath.Join(".", "config")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Configurar cliente Binance
	client := binance.NewBinanceClient(cfg.APIKey, cfg.APISecret)

	// Gerenciar reservas
	strategy.ManageReserves(client, cfg)

	for _, pair := range cfg.TradePairs {
		log.Printf("Executando compra para o par %s\n", pair)
		strategy.ExecuteTrade(client, pair, cfg.Quantity, true)
	}
}
