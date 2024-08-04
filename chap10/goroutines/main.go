package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(5 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)// close(doneChan)
}

func main() {

	// go greet("I hope you're liking the course1!")
	// go greet("I hope you're liking the course2!")
	// done := make(chan bool)
	// go slowGreet("How ... are ... you ...?", done)
	// // <-done// Wait for slowGreet to finish before exiting
	// go greet("I hope you're liking the course3!")// the last greet function call is started only after the slowGreet function finishes and sends a signal to the done channel.
	// <-done// Wait for slowGreet to finish before exiting

	// done := make(chan bool)
	// go greet("I hope you're liking the course1!", done)
	// go greet("I hope you're liking the course2!", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course3!", done)
	// <-done
	// <-done
	// <-done
	// <-done

	// dones := make([] chan bool,4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	// go greet("I hope you're liking the course1!", dones[0])
	// go greet("I hope you're liking the course2!", dones[1])
	// go slowGreet("How ... are ... you ...?", dones[2])
	// go greet("I hope you're liking the course3!", dones[3])
	
	// for _,done := range dones{
	// 	<-done
	// }


	done := make(chan bool)

	go greet("I hope you're liking the course1!", done)
	go greet("I hope you're liking the course2!", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course3!", done)
	
	for doneChan:=range done{
		fmt.Println(doneChan)
	}
	fmt.Println("hello world")

}