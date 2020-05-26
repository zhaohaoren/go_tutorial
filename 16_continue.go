// continue
package main

import "fmt"

func main() {

	// 带label的continue 跳到对应的循环继续下一次循环
label:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(j)
			if j == 2 {
				continue label
			}
		}

	}
}
