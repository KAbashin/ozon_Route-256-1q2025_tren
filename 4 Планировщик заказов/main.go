package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Order struct {
	index    int
	incoming int // время поступления заказа
}

type Machine struct {
	start    int // время начала
	end      int // время окончания
	capacity int // мах свободное место
	index    int
	// currentFreeSpace int // оставшееся свободное место
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// принимаем количество наборов
	var n int
	fmt.Fscanln(in, &n) // количество наборов

	for i := 0; i < n; i++ {
		// прием наборов
		var quantityOffers int

		fmt.Fscanln(in, &quantityOffers) // количество офферов

		arrivalOffers, _ := in.ReadString('\n') // поступление офферов
		elementsOffers := strings.Fields(arrivalOffers)
		arrivalStruct := make([]Order, 0)
		for i, s := range elementsOffers {
			incoming, _ := strconv.Atoi(s)
			arrivalStruct = append(arrivalStruct, Order{index: i + 1, incoming: incoming})
		}

		var mashinsCount int
		trucksStruckt := make([]Machine, mashinsCount)
		fmt.Fscanln(in, &mashinsCount) // количество машин

		for j := 0; j < mashinsCount; j++ {
			var start, end, capacity int
			fmt.Fscan(in, &start, &end, &capacity) // параметры машины
			trucksStruckt = append(trucksStruckt, Machine{start: start, end: end, capacity: capacity, index: j + 1})
		}

		sort.Slice(arrivalStruct, func(i, j int) bool {
			return arrivalStruct[i].incoming < arrivalStruct[j].incoming
		})
		sort.Slice(trucksStruckt, func(i, j int) bool {
			return trucksStruckt[i].start < trucksStruckt[j].start
		})

		// берем в цикле заказы
		// ищем машину со стартом >= закзау
		// проверяем эенд <= заказу
		// проверяем capacity >= 1
		// capacity уменьшаем на 1

		fmt.Fprintln(out, arrivalStruct) // вывод для отладки
		fmt.Fprintln(out, trucksStruckt) // вывод для отладки

		result := make([]int, quantityOffers)
		strSlice := make([]string, len(result))
		for i, v := range result {
			strSlice[i] = fmt.Sprintf("%d", v)
		}
		fmt.Fprintln(out, strings.Join(strSlice, " "))

	}

}
