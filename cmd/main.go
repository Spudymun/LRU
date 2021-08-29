package main

import (
	"fmt"
	"strconv"

	"github.com/Spudymun/LRU/pkg/lrucache"
)

func main() {
	lru := lrucache.New(5)
	lru.Set("1", 1)
	lru.Set("2", 2)
	lru.Set("3", 3)
	lru.Set("4", 4)
	lru.Set("5", 5)
	lru.Set("6", 6)
	lru.Set("1", 1)
	lru.Set("7", 7)
	lru.Invalidate("1")
	lru.Set("8", 8)
	for i := 1; i <= 9; i++ {
		fmt.Println(lru.Get(strconv.Itoa(i)))
	}
	fmt.Println(lru)
}
