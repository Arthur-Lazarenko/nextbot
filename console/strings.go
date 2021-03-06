package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"../core"
	"../core/brain"
)

// LaunchingDialog - функция для выполнения последовательности диалога
func LaunchingDialog(network *brain.NeuralNetwork, dictionary map[string]float64, events chan<- string) {

	// определение считывателя
	reader := bufio.NewReader(os.Stdin)

	for {

		// запрос ввода строки
		fmt.Print("YOU: ")

		// считывание строки
		question, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}

		fmt.Print("BOT: ")
		fmt.Println(core.Input(question, network, dictionary, events))

		// задержка
		time.Sleep(time.Second * 2)

	}

}
