// 工厂模式
package main

import (
	"./06_package/model"
	"fmt"
)

// go中没有构造函数，当结构体小写为私有类型的时候，其他包需要该结构体的时候使用工厂模式
// 这里其实不算是工厂模式，只能算是一个get方法

func main() {

	student := model.NewStudent("justin", 22)
	fmt.Println(student)
	fmt.Println(student.Name)
	fmt.Println((*student).Name)
	fmt.Println(student.GetAge())
}
