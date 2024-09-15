package main

import (
	"fmt"
)

func main() {
	var v1, v2 int
	fmt.Scanln(&v1, &v2)
	if v1 > v2 {
		//print v1 before v2
		fmt.Println(v2, v1)
		if v1 < v2 {
			//print v2 before v1
			fmt.Println(v1, v2)
		}
	} else {
		//print v2 before v1
		fmt.Println(v1, v2)
	}
}
