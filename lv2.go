package main

import (
	"fmt"
)

func main() {
	user := "admin"        //提前设置好系统登录账号
	password := "abcdefgh" //账号的密码

	fmt.Println("请输入账号：")
	var id string   //存放输入的账号
	fmt.Scanln(&id) //输入账号
	fmt.Println("请输入密码：")
	var pw string   //存放输入的密码
	fmt.Scanln(&pw) //输入密码

	if user == id && password == pw { //判断输入的账号和密码是否都与设定相等
		fmt.Println("密码正确，成功登录!")
	} else {
		fmt.Println("密码错误!")
	}
}
