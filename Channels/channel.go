package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go write("Olá Mundo!", ch)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range ch {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func write(text string, channel chan string) {
	for range 5 {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
