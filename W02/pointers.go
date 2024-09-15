package main

import "fmt"

func spisminbror() {
	fmt.Println("Hej min bror")
}

func main() {
	for i := 0; i < 10; i++ {
		spisminbror()
		go spisminbror()
	}
}
