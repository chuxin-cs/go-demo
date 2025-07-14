package demo

import "fmt"

// 声明 字符串 类型变量
var age int = 28

// 像ts一样会自动推导，不写类型也没关系
var isActive = true

func Demo22() {
	name := "云层上的光"
	fmt.Println("花名==>" + name)
	fmt.Println(age)

}
