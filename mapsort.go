package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"yathee":    32,
		"swathi":    23,
		"teja":      43,
		"prabhakar": 45,
		"sujatha":   34,
	}
	keys := make([]string, 0, len(m))
	for v := range m {
		keys = append(keys, v)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
