package main

import (
	"fmt"
	"sort"
)

type animal interface {
	name() string
	action() string
}
type Cat struct{}

func (cat Cat) name() string {
	return "小猫"
}
func (cat Cat) action() string {
	return "喵喵"
}

func main() {
	kill := intList{
		"1. First Blood",
		"3. Triple Kill",
		"4. wazhua KIll",
		"5. Penda  Kill",
		"2. Double Kill",
	}
	sort.Sort(kill)
	for _, v := range kill {
		fmt.Printf(" %v\n", v)
	}

}
