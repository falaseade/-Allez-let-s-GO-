package main

import (
	"fmt"
)
func main() {
	channel := make(chan int)
	go writeOnly(channel)
	readOnly(channel)

}

func readOnly(read <- chan int){
	fmt.Println("Started reading from channel")
	readValue := <- read
	fmt.Printf("Finished reading from channel %d", readValue)
}

func writeOnly(write chan <- int){
	fmt.Println("Started writing to channel")
	write <- 5
	fmt.Println("Finished writing to channel")
}
