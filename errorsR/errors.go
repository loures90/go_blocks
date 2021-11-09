package errorsR

import "fmt"

func ErrosRecovers() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover", err)
		}
	}()
	throws()
}
func throws() error {
	panic("new Error")
}
