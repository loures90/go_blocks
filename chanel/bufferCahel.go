package chanel

import "fmt"

func BufferChanel() {
	ch := make(chan string, 2) // bufferChannel of size 2

	ch <- "Hello"
	ch <- "Wolrd"
	// ch <- "hello again"  if ch is bigger than buffer size, in this case 2, the buffer lock and throw an error

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
