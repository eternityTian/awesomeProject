package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("input:（暂时支持包含同一种运算符的多数运算）")
	var str string   //存放计算式
	fmt.Scanln(&str) //得到一行简单计算式
	fmt.Println("output:")
	if strings.Contains(str, "+") {
		fmt.Println(calculate(str, "+"))
	} else if strings.Contains(str, "-") {
		fmt.Println(calculate(str, "-"))
	} else if strings.Contains(str, "*") || strings.Contains(str, "x") {
		fmt.Println(calculate(str, "*"))
	} else if strings.Contains(str, "/") {
		fmt.Println(calculate(str, "/"))
	} else {
		fmt.Println("不存在运算符号")
	}

}

func calculate(str, operator string) float64 { //传入计算式和运算符，返回答案
	arr := strings.Split(str, operator)
	sum := 0.0 //存放结果
	if operator == "*" {
		sum = 1.0
	}
	for key, num := range arr { //计算切片后的各个数组运算结果
		i, _ := strconv.ParseFloat(num, 64)
		switch operator { //go中case后自带break,若需要执行后面case可使用fallthroug
		case "+":
			sum += i
		case "-":
			if key == 0 {
				sum = i
			} else {
				sum -= i
			}
		case "*":
			sum *= i
		case "x":
			sum *= i
		case "/":
			if key == 0 {
				sum = i
			} else {
				sum /= i
			}
		}
	}
	return sum
}
