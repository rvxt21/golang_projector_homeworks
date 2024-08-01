package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Numbers struct {
	numbers []int
}

func New() *Numbers {
	return &Numbers{}
}

func main() {
	out := make(chan Numbers)
	in := make(chan int)
	numbers := New()
	var count int
	var max int
	fmt.Print("Enter count of numbers: ")
	fmt.Scan(&count)
	fmt.Print("Enter second num(maximum): ")
	fmt.Scan(&max)
	go GenerateNumbers(count, max, numbers, out)
	go AverageNum(out, in)
	go PrintAverage(in)

	time.Sleep(500 * time.Millisecond)

}

func GenerateNumbers(count int, maxNumber int, numbers *Numbers, out chan Numbers) {
	fmt.Println("Starting generating...")
	for i := 0; i < count; i++ {
		randomNumber := rand.Intn(maxNumber)
		fmt.Println(randomNumber)
		numbers.numbers = append(numbers.numbers, randomNumber)
		out <- *numbers
	}
	fmt.Println("*******ADDED NUMBERS TO NUMBERS STRUCT*******")
	fmt.Println(numbers)
	close(out)
	fmt.Println("Ended generating...")
}

func AverageNum(out chan Numbers, in chan int) {
	fmt.Println("Starting calculating average...")
	var sum int
	var count int
	for num := range out {
		for _, num := range num.numbers {
			sum += num
			count++
		}
	}
	if count == 0 {
		fmt.Println("No numbers found")
	} else {
		average := sum / count
		in <- int(average)
	}
	close(in)
	fmt.Println("Ended calculating average...")

}

func PrintAverage(in chan int) {
	fmt.Println("Average is:", <-in)
}
