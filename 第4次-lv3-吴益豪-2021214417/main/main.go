package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	//var tag = false
	//var key, value, skillname string
	//var res bool
	//m := make(map[string]string)
	//for !tag {
	//	fmt.Println("请输入技能名字：")
	//	fmt.Scan(&key)
	//	res = Monit(key)
	//	if res {
	//		fmt.Println("涉嫌违法：请重新输入")
	//		continue
	//	}
	//	fmt.Println("请输入技能模板：")
	//	fmt.Scan(&value)
	//	res = Monit(value)
	//	if res {
	//		fmt.Println("涉嫌违法：请重新输入")
	//		continue
	//	}
	//	m[key] = value
	//	fmt.Println("结束请输入‘true’")
	//	fmt.Scan(&tag)
	//}
	//var tag1 bool = false
	//for !tag1 {
	//	fmt.Println("请输入将要释放技能的名字")
	//	fmt.Scan(&skillname)
	//	ReleaseSkill(m, skillname)
	//	fmt.Println("结束请输入‘true’")
	//	fmt.Scan(&tag1)
	//}
	var i, j, k int = 0, 0, 0
	ticker := time.NewTicker(time.Second)
	mu := sync.WaitGroup{}
	tag := true
	mu.Add(2)
	go func() {
		t1 := time.NewTicker(time.Second)
		tag := true
		for tag {
			select {
			case <-t1.C:
				fmt.Println("王焱落地！")
				i++
			default:
				if i == 5 {
					tag = false
				}
			}
		}
		mu.Done()
	}()
	go func() {
		t2 := time.NewTicker(time.Second)
		tag := true
		for tag {
			select {
			case <-t2.C:
				fmt.Println("赵祚钦落地！")
				j++
			default:
				if j == 5 {
					tag = false
				}
			}
		}
		mu.Done()
	}()
	for tag {
		select {
		case <-ticker.C:
			fmt.Println("芜湖起飞！")
			k++
		default:
			if k == 5 {
				tag = false
			}
		}
	}

	mu.Wait()
}
func ReleaseSkill(m map[string]string, skillname string) {
	v, ok := m[skillname]
	if !ok {
		fmt.Println("未找到该技能")
		return
	}
	fmt.Printf("%v", v)

}
func Monit(s string) bool {
	var res bool
	var i int = 0
	Words := []string{"黄色", "暴力", "色情", "坤坤"}
	for _, s1 := range Words {
		res = strings.Contains(s, s1)
		if res {
			i++
		}
	}
	if i > 0 {
		return true
	}
	return false
}
