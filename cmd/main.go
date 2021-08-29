package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Spudymun/LRU/pkg/lrucache"
)

func main() {
	lru := lrucache.New(5)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20)
	fmt.Println("n =", n)
	for i := 0; i <= n; i++ {
		lru.Set(strconv.Itoa(i), i)
	}
	lru.Set("1", 1)
	lru.Set("4", 1)
	for i := 0; i <= n; i++ {
		lru.Get(strconv.Itoa(i))
	}
	fmt.Println(lru)
}
