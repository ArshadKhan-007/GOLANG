package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels in Go provide a way for goroutines to communicate with each other and synchronize their execution.

// Here is a function to avoid deadlock by processing messages from the channel in a separate goroutine.
// Sending Channel
func processMessage(messageChan chan string) {
	fmt.Println("Processing message", <-messageChan)
}

// Function for random delay
// Sending Channel
func RandomDelay(numChannel chan int) {
	for num := range numChannel { // Continuously receive numbers from the channel
		fmt.Println("Processing number", num)
		time.Sleep(time.Second * 2) // Simulate a delay in processing means wait for 2 seconds to process each number
	}
}

// Receiving Channel
func receiveData(resultChannel chan int, val int, val2 int) {
	res := val + val2
	resultChannel <- res
}

func main() {
	// 	messageChan := make(chan string) // chan is short for channel

	// 	messageChan <- "Hello, Channels!" // Send a message to the channel

	// 	msg := <-messageChan // Receive the message from the channel

	// 	fmt.Println(msg)
	// }

	// Note: The above code will result in a deadlock because the main goroutine is trying to send a message to the channel without a corresponding receiver. To fix this, you can use a separate goroutine to receive the message.

	// Corrected version:
	// First we define a function that processes messages from the channel that function will run in a separate goroutine to avoid deadlock.
	// How that function avoid deadlock? By running in a separate goroutine, it can receive messages from the channel independently of the main goroutine, allowing both sending and receiving operations to occur without blocking each other.
	messageChan := make(chan string)
	go processMessage(messageChan)    // Here we call the function in a separate goroutine
	messageChan <- "Hello, Channels!" // Send a message to the channel
	time.Sleep(time.Second * 2)       // Give some time for the goroutine to process the message

	fmt.Println("------------")
	numChannel := make(chan int)
	go RandomDelay(numChannel)
	for {
		numChannel <- rand.Intn(10) // Send random numbers to the channel
	} // Here, for loop is infinite loop that continuously sends random numbers to the channel every 2 seconds.
	// Note: The above code will run indefinitely, sending random numbers to the channel every 2 seconds. You can stop it manually.
	// To run the next part of the code, you may want to comment out the above infinite loop.

	fmt.Println("------------") // This line will not be reached due to the infinite loop above.
	// Receiving Channel
	resultChannel := make(chan int)
	go receiveData(resultChannel, 10, 20)
	result := <-resultChannel
	fmt.Println("Result is:", result)
}
