package main

import (
	"fmt"
)

type income struct {
	Amount    float64 `json:"amount"`
	Timestamp float64 `json:"timestamp"`
}

// var (
//
//	incomes []income
//
// )
// func homePage(w http.ResponseWriter,r *http.Request)
func main() {
	var item string
	str := "yatheesha"
	words := []string{"yatheesha"}
	// the remaining number of matches
	matches := len(words)
	// start of the sliding window,
	// not the start of the current item
	var start int
	freqs := make(map[string]int)
	for _, word := range words {
		freqs[word]++
		fmt.Println(freqs[word])
	}
	wordSize := len(words[0])
	var res []int
	for stop := wordSize - 1; stop < len(str); stop++ {
		itemStart := stop + 1 - wordSize
		item := str[itemStart : stop+1]
		if freq, exists := freqs[item]; exists {
			if freq > 0 {
				matches--
				fmt.Println(matches)
			}
			freqs[item]--
			fmt.Println(freqs[item])
		}

		if matches == 0 {
			// shrink from the left
			start = stop - len(words)*wordSize + 1
			res = append(res, start)
		}

		// have not yet filled the window,
		// no need to drop the leftmost item
		if stop < len(words)*wordSize-1 {
			continue
		}

		// prepare for the next iteration,
		// i.e. accommodate space for the next item.
		left := str[start : start+wordSize]
		start++
		if freq, exists := freqs[left]; exists {
			if freq >= 0 {
				matches++
			}
			freqs[left]++
		}
	}
	if freq, exists := freqs[item]; exists {
		if freq > 0 {
			matches--
		}
		freqs[item]--
	}
	sum := matches + len(words)
	fmt.Println(sum)
	fmt.Println(item)
	fmt.Println(words)
}
