package main


import (
	"fmt"  
	"time" 
)

func f(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go f("Hello")
	go f("World")
	time.Sleep(1 * time.Second)
}