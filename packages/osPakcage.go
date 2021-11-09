package packages

import (
	"flag"
	"fmt"
)

func Scans() {
	var s string
	var i string
	fmt.Sscanf(" 12345678 ", "%5s%d", &s, &i)
	fmt.Println(s, i)

	x := flag.Int("flagString", 555, "blabla")
	fmt.Println(*x)
}
