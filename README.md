# GoLang Competition

Este projeto é dedicado ao estudo de concorrência em Go (Golang). Ele explora conceitos fundamentais como:

- **Goroutines**: Execução de funções de forma concorrente.
- **WaitGroup**: Sincronização de goroutines.
- **Canais (Channels)**: Comunicação entre goroutines.
- **Canais Bufferizados (Buffered Channels)**: Canais com capacidade limitada.
- **Select**: Multiplexação de canais.

Cada conceito é implementado em exemplos práticos para facilitar o aprendizado e a compreensão.

## Estrutura do Projeto

- `GoRoutines/`: Exemplos básicos de goroutines.
- `WaitGroup/`: Uso de `sync.WaitGroup` para sincronizar goroutines.
- `Channels/`: Comunicação entre goroutines usando canais.
- `Buffered Channels/`: Exemplos de canais bufferizados.
- `Select/`: Uso do `select` para trabalhar com múltiplos canais.

## Como Executar

1. Certifique-se de ter o Go instalado em sua máquina.
2. Navegue até o diretório do exemplo que deseja executar.
3. Execute o comando:

   ```bash
   go run <nome-do-arquivo>.go
