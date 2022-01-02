package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("Chute um número")
	secretNumber := rand.Intn(101)
	reader := bufio.NewReader(os.Stdin)
	for {
		println("Coloque um número de 0 á 100")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		inputedNumber, err := strconv.Atoi(text)
		if err != nil {
			println("Coloque um número de verdade")
			os.Exit(2)
		}
		if secretNumber > inputedNumber {
			println("Menor que o secreto")

		}
		if secretNumber == inputedNumber {
			println("Acertou")
			break
		}
		if secretNumber < inputedNumber {
			println("Maior que o secreto")

		}
	}
}
