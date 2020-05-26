//while循环
package main

import "fmt"

func main() {

	//go 里面没有 while 和 do while

	// while
	var i = 1
	for {
		if i > 3 {
			break
		}
		fmt.Println("while")
		i++
	}

	fmt.Println("---------------")

	// do while
	var j = 1
	for {
		fmt.Println("while")
		j++
		if j > 3 {
			break
		}
	}
}
