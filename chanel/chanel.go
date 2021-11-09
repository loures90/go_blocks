package chanel

import "fmt"

func Chanel() {
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("hello")

	// we send a signal to c in order to allow waitAndSay to continue
	c <- true

	// we wait to receive another signal c before we exit
	<-c
}

func waitAndSay(c chan bool, s string) {

	if b := <-c; b { // we wait the value "c", once we get the data we put it in "b"
		fmt.Println(s)
	}
	c <- true
}

// This code lock the go routine till we get data on the C channel
