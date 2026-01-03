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

	/*rates := map[string]map[string]float32{
		"RUB": {"USD": 0.0125, "EUR": 0.0111}, // 1/80 и 1/90
		"USD": {"RUB": 80.0, "EUR": 0.89},
		"EUR": {"USD": 1.124, "RUB": 90.0}, // 1/0.89
	}*/

	var ratesP *map[string]map[string]float32
	rates := make(map[string]map[string]float32)
	ratesP = &rates
	(*ratesP)["RUB"] = map[string]float32{
		"USD": 0.0125,
		"EUR": 0.0111,
	}
	(*ratesP)["USD"] = map[string]float32{
		"RUB": 80.0,
		"EUR": 0.89,
	}
	(*ratesP)["EUR"] = map[string]float32{
		"USD": 1.124,
		"RUB": 90.0,
	}

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

	for {
		fmt.Scan(&valTrans)
		if valTrans == "RUB" || valTrans == "USD" || valTrans == "EUR" {
			break
		} else {
			fmt.Println("Некорректное значение")
		}
	}
	result := valNum * (*ratesP)[valChange][valTrans]
	fmt.Println(result, valTrans)

}
