package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var chose int
	fmt.Println("Выберите действие")
	fmt.Println("1. Среднее арифмитическое число")
	fmt.Println("2. Сумма чисел")
	fmt.Println("3. Медиана")
	fmt.Scanln(&chose)

	switch {
	case chose == 1:
		avg(ent())
	case chose == 2:
		sum(ent())
	case chose == 3:
		med(ent())
	}
}
func ent() []int {
	k := ""
	slc := make([]int, 0)
	fmt.Println("Вводите числа. Когда закончите нажмите 'e' ")
	for {
		fmt.Scanln(&k)
		if k == "e" {
			break
		} else {
			a, err := strconv.Atoi(k)
			if err == nil {
				slc = append(slc, a)
			}
		}
	}
	return slc
}
func avg(slc []int) {
	var a int
	for _, v := range slc {
		a += v
	}
	fmt.Println(float64(a) / float64(len(slc)))
}
func sum(slc []int) {
	var a int
	for _, v := range slc {
		a += v
	}
	fmt.Println(a)
}
func med(slc []int) {
	sort.Ints(slc)
	if len(slc)%2 == 0 {
		fmt.Println((slc[len(slc)/2-1] + slc[len(slc)/2]) / 2)
	} else {
		fmt.Println(slc[len(slc)/2])
	}

}
