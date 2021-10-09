package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println("Input:")

	//var str string //存放输入数据

	//fmt.Scan(&str)

	//func Contains(s, substr string) boll
	//功能：字符串s中是否包含substr,返回bool
	str2 := "ssdasbc"
	fmt.Println(strings.Contains(str2, "abc"))

	//func Join(a []string, sep string) string
	//功能：字符串连接，把slice a 通过sep连接起来
	s := []string{"I","love","go"}
	fmt.Println(strings.Join(s," "))

	//func Index(s, sep string) int
	//功能：在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("hello","e"))
	fmt.Println(strings.Index("hello","a"))

	//func Repeat(s string, count int) string
	//功能：重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na",2))

	//func Replace(s, old, new string, n int) string
	//在s字符串中,把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("I love coding","love","hate",1))
	fmt.Println(strings.Replace("I love love coding","love","hate",-5))

	//func Split(s, sep string) []string
	//把s字符串按照sep分割，返回slice
	arr := strings.Split("1+5-5+4-2","+")
	fmt.Println(arr[2])
	arr2 := strings.Split("5 +4 +1 +2","+")
	fmt.Println(arr2)

	//func Trim(s string, cutset string) string
	//在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim("wuwuwuI love wu coding wuwuwuwu","wu"))//I love wu coding

	//func Fields(s string) []string
	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields("ha ha ha I can't ")[1:3])

	//slice操作
	a := []string {"hello","hhh","wuhu"}
	b := make([]string, len(a))
	copy(b,a)//将slice a复制到slice b
	fmt.Println(b)

	c := "hello"+"world"//字符串拼接
	fmt.Println(c)

	/*
	for循环range格式对slice,map,数组，字符串等进行迭代循环。
	for key, value := range oldMap{
	 newMap[key] = value
	}
	 */
	//44行--->arr2 := strings.Split("5 +4 +1 +2","+")
	//说明空格被忽略了
	for i, s := range arr2{
		fmt.Printf("%d索引上为%s",i,s)
	}


}