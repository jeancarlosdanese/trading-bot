# Trading Bot

Este é um robô de trading automatizado desenvolvido em Go, integrado à Binance, projetado para realizar operações de compra e venda de criptomoedas com foco em maximizar o valor da carteira em dólares (USDT). Ele também inclui uma lógica para reservar uma parte do saldo em USDT ao atingir certos limites.

## Funcionalidades

- **Gerenciamento Automático de Reservas**:
  - Ao atingir um valor configurado na carteira, 30% do saldo total é reservado em USDT.
- **Negociações Automatizadas**:
  - Suporte para múltiplos pares de negociação configuráveis (exemplo: `BTCUSDT` e `ETHUSDT`).
  - Quantidade mínima de compra/venda por par configurada no arquivo de configuração.
- **Ajuste Dinâmico de Conversão**:
  - Verifica o saldo disponível antes de tentar converter ativos para USDT.
  - Ajusta automaticamente a quantidade para respeitar o saldo disponível e os limites mínimos exigidos pela Binance.
- **Modo de Teste (Sandbox)**:
  - Permite simular negociações antes de operar na conta real.

## Tecnologias Utilizadas

- **Go**:
  - Linguagem principal do projeto.
- **Binance API**:
  - Para integração com a exchange Binance.
- **Viper**:
  - Para gerenciamento de configurações.
- **Go Modules**:
  - Para gerenciamento de dependências.

## Pré-requisitos

- **Conta na Binance** com acesso à API (em modo de teste ou real).
- **Chaves da API Binance** configuradas para `TRADE` e `USER_DATA`.
- **Go** instalado (versão 1.20 ou superior recomendada).

## Configuração do Projeto

### 1. Clonar o Repositório

```bash
git clone https://github.com/seu-usuario/trading-bot.git
cd trading-bot
```

### 2. Configurar o Arquivo `config.json`

Crie ou edite o arquivo `config.json` no diretório raiz do projeto. Exemplo:

```json
{
  "api_key": "SUA_API_KEY",
  "api_secret": "SUA_API_SECRET",
  "trade_pairs": ["BTCUSDT", "ETHUSDT"],
  "reserve_threshold": 1000,
  "quantity": "0.01",
  "test_mode": true
}
```

**Descrição dos Campos**:

- `api_key` e `api_secret`: Suas chaves da API Binance.
- `trade_pairs`: Lista de pares de negociação.
- `reserve_threshold`: Valor em USD que, ao ser atingido, ativa a reserva de 30% em USDT.
- `quantity`: Quantidade padrão para ordens de compra.
- `test_mode`: Se `true`, utiliza o ambiente de Sandbox da Binance.

### 3. Instalar Dependências

Certifique-se de que o Go está configurado e instale as dependências:

```bash
go mod tidy
```

### 4. Executar o Robô

Para iniciar o robô, execute:

```bash
go run ./cmd/trading/main.go
```

## Estrutura do Projeto

```plaintext
trading-bot/
├── cmd/
│   └── trading/
│       └── main.go      # Arquivo principal do robô
├── internal/
│   ├── binance/         # Integração com a Binance API
│   ├── config/          # Carregamento e gerenciamento de configurações
│   ├── strategy/        # Lógica de trading e reserva
│   └── utils/           # Funções auxiliares (ex.: cálculo de portfólio)
├── config.json          # Arquivo de configuração do robô
├── go.mod               # Gerenciamento de dependências
├── README.md            # Documentação do projeto
```

## Logs do Robô

O robô gera logs detalhados sobre suas operações, incluindo:

- Valor total da carteira em USD.
- Ativos ignorados por não terem preço na Binance.
- Ordens executadas com detalhes de quantidade e preço.
- Ajustes automáticos devido a saldo insuficiente.

Exemplo de saída de log:

```plaintext
INFO: 2024/12/05 00:30:21 Iniciando o robô de trading...
Configuração carregada com sucesso!
Valor total da carteira em USD: 595000.76
Reservando $178500.23 em USDT...
Saldo insuficiente de ETH. Necessário: 46.333609 ETH, Disponível: 0.987100 ETH
Ajustando para converter o máximo possível com margem: 3612.66 USD
Reserva concluída com sucesso!
```

## Funcionalidades Futuras

- Notificações via Telegram ou e-mail.
- Suporte para diversificação automática durante conversões para USDT.
- Estratégias mais avançadas de trading baseadas em análise técnica.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir _issues_ ou enviar _pull requests_.

---

### **Atenção**

Este robô foi projetado para ser usado com cuidado. Recomenda-se utilizar o modo de teste antes de operar em produção. Negociar criptomoedas envolve riscos significativos. Use por sua conta e risco.

---
