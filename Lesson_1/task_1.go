package main

import "fmt"

func main() {
	const a float64 = (63.6)
	var amount float64
	fmt.Println("Привет, это программа конвертации валюты.\nКакую сумму в рублях нужно перевести в доллары?")
	fmt.Scanln(&amount)
	fmt.Println("В долларах это будет", amount/a, "USD")
}
