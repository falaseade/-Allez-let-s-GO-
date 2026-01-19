package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan string)
	go checkServer("A", time.Second * 1, results)
	go checkServer("B", time.Second * 5, results)
	for range 2 {
		select {
		case v := <- results:
			fmt.Printf("Printing Result %s\n", v)
		case <- time.After(3 * time.Second):
			fmt.Println("Timeout!!!!")
		default:
			fmt.Println("No one is listening, shutting down")
		}
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