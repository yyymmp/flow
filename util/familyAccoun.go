package util

import "fmt"

type FamilyAccoun struct {
	Loop     bool    //是否退出
	Balance  float64 //余额
	Money    float64 //金额
	Note     string  //说明
	Detail   string  //明细
	IsHaving bool    //当前是否有收支明细
}

//工厂模式返回此对象
func NewFamilyAccoun() *FamilyAccoun {
	return &FamilyAccoun{
		Loop:     false,
		Balance:  0,
		Money:    0,
		Note:     "",
		Detail:   "",
		IsHaving: false,
	}
}

//主菜单
func (f *FamilyAccoun) Menu() {
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
			f.Mingxi()
		case 2:
			f.income()
		case 3:
			if ok := f.expend(); !ok {
				break
			}
		case 4:
			f.logout()
		default:
			fmt.Println("非法输入")
		}

	}
}

//收支明细
func (f *FamilyAccoun) Mingxi() {
	if f.IsHaving {
		fmt.Println("----------当前收支明细表---------")
		fmt.Println(f.Detail)
	} else {
		fmt.Println("当前暂无收支情况，来一笔吗")
	}
}

//收入
func (f *FamilyAccoun) income() {
	fmt.Println("请输入金额")
	fmt.Scanln(&f.Money)
	fmt.Println("请输入说明")
	fmt.Scanln(&f.Note)
	f.Balance += f.Money
	f.Detail += fmt.Sprintf("收入\t%.2f\t%.2f\t%v\t\n", f.Balance, f.Money, f.Note)
	f.IsHaving = true
}

//支出
func (f *FamilyAccoun) expend() bool {
	fmt.Println("请输入金额")
	fmt.Scanln(&f.Money)
	if f.Money > f.Balance {
		fmt.Println("余额不足")
		return false
	}
	fmt.Println("请输入说明")
	fmt.Scanln(&f.Note)
	f.Balance -= f.Money
	f.Detail += fmt.Sprintf("支出\t%.2f\t%.2f\t%v\t\n", f.Balance, f.Money, f.Note)
	f.IsHaving = true
	return true
}
func (f *FamilyAccoun) logout() {
	flag := ""
	for {
		fmt.Println("确定要退出吗y or n")
		fmt.Scanln(&flag)
		if flag == "y" {
			f.Loop = false
			break
		} else if flag == "n" {
			break
		} else {
			fmt.Println("非法输入，请重新输入")
		}
	}
}
