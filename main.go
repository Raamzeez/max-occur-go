package main

import (
	"fmt"
	"time"
)

var exampleStringSlice = []string{
	"bitcoin", "bitcoin-cash", "ethereum", "ethereum", "solana", "bitcoin", "bitcoin",
	"bitcoin", "litecoin", "solana", "ripple", "ripple", "the-graph", "the-graph", "zrx",
	"tara", "dogecoin", "dogecoin", "dogecoin", "celo", "celo",
	"serum", "zrx", "serum", "serum", "serum", "serum", "celo",
}

func implemenation1() {
	begin := time.Now()

	m := map[string]int{}
	mostMentioned := []string{}

	for _, asset := range exampleStringSlice {
		m[asset] += 1
	}
	fmt.Println(m)

	// loop through to get top 5
	for i := 0; i < 5; i++ {
		maxAsset := ""
		max := 0
		// looping through map
		for key, val := range m {
			if val > max {
				max = val
				maxAsset = key
			}
		}
		mostMentioned = append(mostMentioned, maxAsset)
		delete(m, maxAsset)
	}

	fmt.Println("implementation 1:")
	fmt.Println(mostMentioned)
	fmt.Printf("Time: %s\n", time.Now().Sub(begin))
}

func implemenation2() {
	begin := time.Now()
	m := map[string]int{}

	for _, asset := range exampleStringSlice {
		m[asset] += 1
	}

	max := 0

	for _, v := range m {
		if v > max {
			max = v
		}
	}

	returnSlice := []string{}

	for _, _ = range m {
		for key, value := range m {
			if len(returnSlice) == 5 {
				break
			}
			if value == max {
				returnSlice = append(returnSlice, key)
				delete(m, key)
			}
		}
		max -= 1
	}

	fmt.Println("implementation 2:")
	fmt.Println(returnSlice)
	fmt.Printf("Time: %s\n", time.Now().Sub(begin))

}

func main() {
	implemenation1()
	implemenation2()
}
