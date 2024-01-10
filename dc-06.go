package main

import (
	"fmt"
	"sync"
)

// Define a simple struct
type MyStruct struct {
	ID   int
	Name string
}

func main() {
	// Create a channel that carries MyStruct
	myStructChannel := make(chan MyStruct)

	// Use a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Start a goroutine to send a MyStruct instance through the channel
	wg.Add(1)
	go sendData(myStructChannel, &wg)

	// Start another goroutine to receive the MyStruct instance from the channel
	wg.Add(1)
	go receiveData(myStructChannel, &wg)

	// Wait for both goroutines to finish
	wg.Wait()
}

// sendData sends a MyStruct instance through the channel
func sendData(ch chan<- MyStruct, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create an instance of MyStruct
	dataToSend := MyStruct{ID: 1, Name: "Example"}

	// Send the struct through the channel
	ch <- dataToSend

	fmt.Println("Data sent through channel.")
}

// receiveData receives a MyStruct instance from the channel
func receiveData(ch <-chan MyStruct, wg *sync.WaitGroup) {
	defer wg.Done()

	// Receive the struct from the channel
	receivedData := <-ch

	// Print the received data
	fmt.Printf("Received Data: %+v\n", receivedData)
}
