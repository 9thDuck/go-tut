package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// func greet(phrase string, doneChan chan bool) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

func greetWithRandomErr(phrase string, doneChan chan bool, errorChan chan error) {
	randomNum := rand.Float64()

	if randomNum < 0.5 {
		errorChan <- errors.New("Error")
		return
	}

	doneChan <- true
}

// func greet(phrase int, doneChan chan bool) {
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true
// }

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
}

func main() {
	const name1 = "Sai"
	const name2 = "Sarath"

	const greeting = "How are you?"
	const response = "I am good. Thanks. How are you?"

	// done := make(chan bool)
	dones := make([]chan bool, 4)
	errorChans := make([]chan error, len(dones))

	for idx := range dones {
		dones[idx] = make(chan bool)
		errorChans[idx] = make(chan error)

		go greetWithRandomErr("", dones[idx], errorChans[idx])
	}

	// for idx := range dones {
	// 	go greet(i, dones[i])
	// 	dones[idx] = make(chan bool)
	// }

	// go greet(name1, dones[0])
	// go slowGreet(greeting, dones[1])
	// go greet(name2, dones[2])
	// go slowGreet(response, dones[3])

	// go greet(name1, done)
	// go greet(greeting, done)
	// go slowGreet(name2, done)
	// go greet(response, done)

	// for _, done := range dones {
	// 	<-done
	// }

	for idx := range dones {
		select {
		case err := <-errorChans[idx]:
			fmt.Println(err)
		case <-dones[idx]:
			fmt.Println("Done")
		}

	}

	// for range done {
	// }
}
