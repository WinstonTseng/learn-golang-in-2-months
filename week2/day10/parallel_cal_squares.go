package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers)) // Buffered channel
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}
}
