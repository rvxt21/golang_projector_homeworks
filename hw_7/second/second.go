package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var min int
	var max int
	fmt.Print("Enter first num(minimum): ")
	fmt.Scan(&min)
	fmt.Print("Enter second num(maximum): ")
	fmt.Scan(&max)
	ch := make(chan int)
	resultCh := make(chan int)

	go GenerateRandomNumsIdRange(min, max, ch)
	go FindMaxMin(ch, resultCh)

	fmt.Println("Min", <-resultCh)
	fmt.Println("Max", <-resultCh)

}

func GenerateRandomNumsIdRange(min int, max int, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 10; i++ {
		random := rand.Intn(max-min+1) + min
		fmt.Println(random)
		ch <- random
	}
	close(ch)
}

func FindMaxMin(ch chan int, resultCh chan int) {
	min := <-ch
	max := min

	for i := range ch {
		if i <= min {
			min = i
		} else {
			if i >= max {
				max = i
			}
		}
	}
	resultCh <- min
	resultCh <- max
}
