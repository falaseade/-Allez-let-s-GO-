package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Starting the Main Function!!!")
	doWork()
	fmt.Println("Finishing the Main Function")
}

func doWork() {
	fmt.Println("Time to start doWork Function")
	time.Sleep(5 * time.Second)
	fmt.Println("End after 5 seconds")
}