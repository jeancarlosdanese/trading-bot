package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey           string   `mapstructure:"api_key"`
	APISecret        string   `mapstructure:"api_secret"`
	TradePairs       []string `mapstructure:"trade_pairs"`
	ReserveThreshold float64  `mapstructure:"reserve_threshold"`
	Quantity         string   `mapstructure:"quantity"`
	TestMode         bool     `mapstructure:"test_mode"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.SetConfigName("config") // Nome do arquivo sem extensão
	viper.SetConfigType("json")   // Tipo do arquivo
	viper.AddConfigPath(path)     // Caminho do arquivo de configuração

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	log.Println("Configuração carregada com sucesso!")
	return config, nil
}
