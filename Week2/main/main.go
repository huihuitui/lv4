package main

import (
	"fmt"
	"sort"
)

type newint []coder
type coder struct {
	name   string
	number int64
	code   int
}

var s newint
var key int
var x coder
var tag bool = true
var i string

func main() {
	for tag {
		fmt.Println("-----------学生管理系统------------")
		fmt.Println("1------录入学生------------")
		fmt.Println("2------录出学生------------")
		fmt.Println("3------查找学生------------")
		fmt.Println("4------排序功能------------")
		fmt.Println("5------退出系统------------")
		fmt.Println("6------显示现在------------")
		fmt.Println("-------输入选项-----------")
		fmt.Scanln(&key)
		switch key {
		case 1:
			CHeckinstudent(x)
			fmt.Println("录入成功")
		case 2:
			length := len(s) - 1
			s = append(s[:length])
		case 3:
			var a interface{}

			switch x := a.(type) {
			case int:
				fmt.Scanln(&x)
				typeof(x)
			case string:
				fmt.Scanln(&x)
				typeof(x)
			}

			typeof(a)
		case 4:
			sort.Sort(s)
			fmt.Println("排序成功")
		case 5:
			tag = false
		case 6:
			length := len(s)
			for i := 0; i < length; i++ {
				fmt.Printf("名字：%s\t学号：%d\t代码量：%d\n", s[i].name, s[i].number, s[i].code)
			}
		default:
			fmt.Println("没有此选项")
		}
	}
}
func CHeckinstudent(x coder) {
	fmt.Scanln(&x.name)
	fmt.Scanln(&x.number)
	fmt.Scanln(&x.code)
	s = append(s, x)
}

func typeof(v interface{}) {
	switch x := v.(type) {
	case int:
		length := len(s)
		for i := 0; i < length; i++ {
			if int64(x) == s[i].number {
				fmt.Println("进来了")
				fmt.Println(s[i])
				fmt.Println("录出成功1")
				break
			}
		}

	case string:

		length := len(s)
		for i := 0; i < length; i++ {
			if x == s[i].name {

				fmt.Println(s[i])
				fmt.Println("录出成功2")
				break
			}
		}

	default:
		fmt.Println("输入错误")
	}
}
func (x newint) Len() int {
	return len(x)
}
func (x newint) Less(i, j int) bool {
	return x[i].code < x[j].code
}
func (x newint) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
