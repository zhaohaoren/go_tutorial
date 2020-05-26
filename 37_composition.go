// 组合
package main

import "fmt"

type AAA struct {
	Name string
}

type CCC struct {
	a AAA // 这种方式不叫继承，叫做组合
	Name string
}

func main()  {

	var c CCC
	c.Name = "CCC"
	c.a.Name = "AAA"
	fmt.Println(c)
}