package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	results := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	exchangeNames := []string{"Exchange 1", "Exchange 2", "Exchange 3"}
	for _, exchange := range exchangeNames {
		wg.Add(1)
		go getDataFromExchange(ctx, exchange, results, &wg)
	}
	go func(){
		wg.Wait()
		close(results)
	}()

	for value := range results {
		fmt.Println(value)
	}
}

func getDataFromExchange(ctx context.Context, nameOfExchange string, results chan <- int, wg *sync.WaitGroup){
	defer wg.Done()
	jitterTimeout := rand.IntN(3) + 1
	select {
	case <- ctx.Done():
		fmt.Printf("Context Cancelled: %s for %s exchange\n", ctx.Err(), nameOfExchange)
	case <- time.After(time.Duration(jitterTimeout) * time.Second):
		results <- jitterTimeout
	}
}

func searchDatabase(nameOfDatabase string, results chan <- string, wg *sync.WaitGroup){
	defer wg.Done()
	jitterTime := rand.IntN(3) + 1
	time.Sleep(time.Duration(jitterTime) * time.Second)
	results <- fmt.Sprintf("Received response from: %s", nameOfDatabase)
}

func downloadFiles(wg *sync.WaitGroup, name string, maxDownloadTime int){
	defer wg.Done()
	jitter := rand.IntN(maxDownloadTime)
	time.Sleep(time.Duration(jitter) * time.Second)
	fmt.Printf("%s has finished downloading!!!\n", name)
}

func workFunction(jitterTimeout int, ch chan <- int, done <- chan struct{}){
	for {
		select {
		case <- done :
			fmt.Println("Work Generator: Received stop signal. Exiting.")
		case ch <- 3:
			jitter := rand.IntN(jitterTimeout)
			time.Sleep(time.Duration(jitter) * time.Second)
		}
	}
}

func workerFunction(waitPeriod int, ch <- chan int, done chan <- struct{}){
	for {
		select {
		case value := <- ch :
			fmt.Printf("Worker is processing: %d\n", value)
		case <- time.After(time.Duration(waitPeriod) * time.Second):
			close(done)
			return
		}
	}
}

func sleepForRandomTimeThenMessage(name string, ch chan <- string){
	randomSeconds := rand.IntN(5)
	time.Sleep(time.Duration(randomSeconds))
	ch <- fmt.Sprintf("%s is done\n", name)
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

