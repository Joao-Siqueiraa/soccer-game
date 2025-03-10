package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var t1 string
	var t2 string
	var gt1 int
	var gt2 int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("escolha o primeiro time: ")
	fmt.Scan(&t1)

	fmt.Print("escolha o segundo time: ")
	fmt.Scan(&t2)

	gt1 = rand.Intn(10) //gera um numero aleatorio
	gt2 = rand.Intn(10)

	if gt1 > gt2 {
		fmt.Println("O", t1, "ganhou de", gt1, "x", gt2)

	} else if gt1 == gt2 {
		fmt.Println(" deu empate o jogo foi", gt1, "x", gt2)
	} else if gt2 > gt1 {
		fmt.Println("O", t2, "ganhou de", gt2, "x", gt1)
	}

}
