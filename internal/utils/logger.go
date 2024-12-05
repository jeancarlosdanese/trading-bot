package utils

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

// InitLogger inicializa os loggers para diferentes níveis
func InitLogger() {
	logFile, err := os.OpenFile("trading-bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}

	infoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	log.SetOutput(logFile)
}

// Info registra logs de informações
func Info(message string) {
	infoLogger.Println(message)
}

// Error registra logs de erros
func Error(message string) {
	errorLogger.Println(message)
}
