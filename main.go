package main

import (
	"fmt"
	"strconv"
)

func main() {
	input()
}
func input() {
	var valChange, valTrans string
	var valCount string
	var valNum float32
	fmt.Println("Введите валюту : ")
	fmt.Println("RUB")
	fmt.Println("USD")
	fmt.Println("EUR")
	for {
		fmt.Scan(&valChange)
		if valChange == "RUB" || valChange == "USD" || valChange == "EUR" {
			break
		} else {
			fmt.Println("Попробуйте ещё раз")
		}
	}
	switch {
	case valChange == "RUB":
		fmt.Println("RUB")
	case valChange == "USD":
		fmt.Println("USD")
	case valChange == "EUR":
		fmt.Println("EUR")
	}

	fmt.Println("Введите количество вашей валюты")
	for {
		fmt.Scan(&valCount)
		valCheck, err := strconv.ParseFloat(valCount, 32)
		if err == nil {
			valNum = float32(valCheck)
			break
		} else {
			fmt.Println("Некорректный ввод")
		}
	}
	fmt.Println("You have: ", valCount, valChange)
	fmt.Println("В какую валюту вы хотите перевести свои деньги?")
	if valChange != "RUB" {
		fmt.Println("RUB")
	}
	if valChange != "USD" {
		fmt.Println("USD")
	}
	if valChange != "EUR" {
		fmt.Println("EUR")
	}

	for {
		fmt.Scan(&valTrans)
		if valTrans == "RUB" || valTrans == "USD" || valTrans == "EUR" {
			break
		} else {
			fmt.Println("Некорректное значение")
		}
	}
	switch {
	case valChange == "RUB" && valTrans == "USD":
		fmt.Println(valNum/80, " USD")
	case valChange == "RUB" && valTrans == "EUR":
		fmt.Println(valNum/90, " EUR")
	case valChange == "USD" && valTrans == "RUB":
		fmt.Println(valNum*80, " RUB")
	case valChange == "USD" && valTrans == "EUR":
		fmt.Println(valNum*0.89, "EUR")
	case valChange == "EUR" && valTrans == "USD":
		fmt.Println(valNum/0.89, "USD")
	case valChange == "EUR" && valTrans == "RUB":
		fmt.Println(valNum*90, "RUB")
	}
}
