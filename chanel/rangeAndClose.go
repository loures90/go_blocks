package chanel

import "fmt"

func RangeAndClose() {
	c := make(chan string)
	go sayHello(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c // check ig the chanel close
	fmt.Println("chanel close?", !ok, "value ", v)
}

func sayHello(c chan string, n int) {

	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}
