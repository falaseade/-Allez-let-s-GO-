package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	dataChannel := make(chan int)
	done := make(chan struct{})
	for num := range 3 {
		go printNumbers(fmt.Sprintf("Consumer %d", num), dataChannel, done)
	}
	go produceRandomNumbers(dataChannel)
	for range 3 {
		<- done
	}
}

func printNumbers(name string, ch <- chan int, done chan <- struct{}){
	for num := range ch {
		fmt.Printf("%s is printing a random number %d \n", name, num)
	}
	done <- struct{}{}
}

func produceRandomNumbers(ch chan <- int){
	defer close(ch)
	for range 100 {
		ch <- rand.Int()
	}
}

func checkServer(name string, delay time.Duration, ch chan string){
	fmt.Printf("Checking Server %s \n", name)
	time.Sleep(delay)
	message := fmt.Sprintf("Server %s is UP!!!", name)
	ch <- message
}

func readOnly(read <-chan int) int {
	readValue := <-read
	return readValue
}

func writeOnly(write chan<- int, value int) {
	write <- value
}

func readIncrementWrite(channel chan int) {
	readValue := readOnly(channel)
	readValue = readValue + 1
	writeOnly(channel, readValue)
}

func printValue(print <-chan int) {
	printValue := <-print
	fmt.Printf("Printing Value Read %d", printValue)
}

func player(name string, table chan int){
	for ball := range table {
		fmt.Printf("Player %s hit ball %d\n", name, ball)
		ball++
		if ball >= 100 {
			fmt.Println("Finished playing!!!")
			close(table)
			time.Sleep(2 * time.Second)
		} else {
			table <- ball
		}
		
	}
}

