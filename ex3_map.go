package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	 m  :=make(map[string]int)
	for i:=0;i<len(str);i++{
		m[str[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
	fmt.Println(WordCount("toi la la cong cong 123"))
	str :="cong cong cong cong 1 2 3"
	
}
