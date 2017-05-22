package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int { //initiales "a" and "b" and returns a func of type int that does ...
	a := 0
	b := 1
	return func() int { // ... the following each time the func called
		temp := a
		a += b
		b = temp
		return a
	}
}

func main() {
	f := fibonacci() //returns function variable f()
	for i := 0; i < 10; i++ {
		fmt.Println(f()) //calls function variable f performing inner algorithm and returning an int
	}
}
