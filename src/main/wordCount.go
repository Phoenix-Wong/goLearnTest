package main

import (
	"strings"
)





func WordCount(s string)map[string]int {
	result := strings.Fields(s)
	m := make(map[string]int)
	for _,word :=range result{
		m[word]++
	}
	return m
}


func main() {
}
