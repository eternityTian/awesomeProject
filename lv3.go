package main

import (
	"fmt"
)

func main() {

	/*
		func append(slice,elms) []type函数增加切片的容量
	*/

	fmt.Println("input(回车结束一次输入,输入0结束输入)")
	var nums []int
	var num int
	for {
		fmt.Scanln(&num) //得到本次输入的数
		if num == 0 {    //如果输入0跳出循环
			break
		}
		nums = append(nums, num) //允许追加空切片
	}
	//len(slice) 获得切片长度
	for i := 0; i < len(nums)-1; i++ { //冒泡排序，对切片排序
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] { //小数据上浮
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println("output")
	for _, value := range nums { //遍历切片，打印排序后结果
		fmt.Print(value)
		fmt.Print(" ")
	}
}
