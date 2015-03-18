package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	b := 0
	for _, x := range a {
		b += x
	}
	res <- b
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	c1 := <-ch
	c2 := <-ch
	c1 += c2
	fmt.Println(c1)

}
