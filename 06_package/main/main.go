package main

import (
	"../model"      // 相对路径或者GoPath相对路径或者绝对路径都可以
	util "../utils" // 可以对引入的包起一个别名
	"fmt"
)

func main() {

	fmt.Println(model.Public_val)
	//fmt.Println(model.private_val) 无法访问私有变量

	util.PublicFunction()
	//util.Main()
}
