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
		lru.Set(strconv.Itoa(i), rand.Int31n(20))
	}
	lru.Set(strconv.Itoa(5), rand.Int31n(20))
	fmt.Println(lru)
}
