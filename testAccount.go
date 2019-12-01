package main

import "fmt"

func main() {
	var loop bool
	var balance float64 //余额
	var money float64   //金额
	var note string     //说明
	var detail string
	detail = "收支\t账户金额\t收支金额\t说明\n"
	loop = true
	for {
		fmt.Println("----------家庭收支明细表---------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出登录")
		fmt.Println("请选择1-4进行对应操作\n")
		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("----------当前收支明细表---------")
			fmt.Println(detail)
		case 2:
			fmt.Println("请输入金额")
			fmt.Scanln(&money)
			fmt.Println("请输入说明")
			fmt.Scanln(&note)
			balance += money
			detail += fmt.Sprintf("收入\t%.2f\t%.2f\t%v\t\n", balance, money, note)
		case 3:
			fmt.Println("请输入金额")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("请输入说明")
			fmt.Scanln(&note)
			balance -= money
			detail += fmt.Sprintf("支出\t%.2f\t%.2f\t%v\t\n", balance, money, note)
		case 4:
			loop = false
			break
		default:
			fmt.Println("非法输入")
		}
		if !loop {
			fmt.Println("一退出")
			break
		}
	}
}
