package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go func() {
		number := make([]int, 10)
		for i := 0; i < 10; i++ {
			number[i] = rand.Intn(101)
		}
		for _, num := range number {
			channel1 <- num
		}
		defer close(channel1)
	}()
	go func() {
		for num := range channel1 {
			channel2 <- num * num
		}
		defer close(channel2)
	}()
	var result []int
	for vl := range channel2 {
		result = append(result, vl)
	}
	fmt.Println("Рузультат", result)
}
