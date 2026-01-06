package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover", r)
		}
	}()
	oper, slice := getUserEntry()

	fmt.Println("Результат Вашей операции: ", calcUserOper(oper, slice))

}

func getUserEntry() (string, []int) {
	var oper string
	var numberSeries string
OPER:
	fmt.Println("Введите вид операции(AVG/SUM/MED):")
	fmt.Scan(&oper)
	if oper == "AVG" || oper == "SUM" || oper == "MED" {
	NUMBERSSERIES:
		fmt.Println("Введите числовой ряд через запятую:")
		fmt.Scan(&numberSeries)

		slice, err := regx(numberSeries)

		if err != nil {
			fmt.Printf("Ошибка при рбработке числового ряда %s.\n Вы ввели: %s", err, numberSeries)
			goto NUMBERSSERIES
		} else {
			return oper, slice
		}

	} else {
		fmt.Printf("Доступные виды операций: AVG - среднее значение, SUM - сумма значений, MED - медиана.\n Вы ввели: %s\n", oper)
		goto OPER
	}

}

func regx(str string) ([]int, error) {

	//fmt.Println("Test string", str)

	re := regexp.MustCompile("[0-9]+")

	reStr := re.FindAllString(str, -1)
	//fmt.Println(reStr)
	//fmt.Println(len(reStr), cap(reStr), reStr[0])
	slice := make([]int, 0)
	var err error
	for _, v := range reStr {
		var num int
		num, err = strconv.Atoi(v)
		slice = append(slice, num)
		if err != nil {
			//fmt.Println("Ошибка")
			break

		}
	}
	if err != nil {
		//fmt.Println(err, "Метка")
		return nil, err
	} else {
		//fmt.Println(slice, "Метка")
		return slice, nil

	}

}

func calcUserOper(oper string, slice []int) int {
	sum := 0
	switch oper {
	case "AVG":
		sum = calcAvg(slice)
	case "SUM":
		sum = calcSum(slice)
	case "MED":
		sum = calcMed(slice)
	}

	return sum
}

func calcAvg(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum / len(slice)
}

func calcSum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func calcMed(slice []int) int {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	//fmt.Println(slice)
	sum := 0
	if len(slice)%2 == 0 {
		sum = slice[len(slice)/2]
	} else {
		sum = slice[len(slice)/2+1]
	}
	return sum
}
