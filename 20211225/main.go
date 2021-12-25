package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var vArr [10]int
	for i := 0; i < len(vArr); i++ {
		vArr[i] = rand.Intn(100)
	}
	fmt.Println(vArr)

	var vRes [10]int
	for i := len(vRes) - 1; i >= 0; i-- {
		vRes[len(vRes)-1-i] = vArr[i]
	}
	fmt.Println(vRes)

	for i := 0; i < len(vArr)/2; i++ {
		vArr[i], vArr[len(vArr)-1-i] = vArr[len(vArr)-1-i], vArr[i]
	}
	fmt.Println(vArr)

	for i, j := 0, len(vArr)-1; j > i; i, j = i+1, j-1 {
		vArr[i], vArr[len(vArr)-1-i] = vArr[len(vArr)-1-i], vArr[i]
	}
	fmt.Println(vArr)

}
