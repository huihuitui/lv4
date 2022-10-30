package main

import "fmt"

func Receive(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("这个是%T", v)
	case bool:
		fmt.Printf("这个是%T", v)
	case string:
		fmt.Printf("这个是%T", v)
	default:
		fmt.Printf("这个是%T", v)
	}
}
