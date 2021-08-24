package main

import (
	"fmt"
)

func main() {
	lru := New(5)
	lru.Set("1", 1)
	lru.Set("2", 2)
	lru.Set("3", 3)
	lru.Set("4", 4)
	lru.Set("5", 5)
	lru.Set("6", 4)
	fmt.Println(lru.items)
}
