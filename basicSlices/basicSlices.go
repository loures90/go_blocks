package basicslices

import "fmt"

func Basicslices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = append(a[:2], a[4:]...)
	fmt.Println(a)
	fmt.Println("capacity of a:", cap(a))

}
