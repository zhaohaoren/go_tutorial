// 接口编程最佳实践 - 结构体排序
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 系统库中很多参数都是要接口传参的，那些接口操作的一般是实现了该接口的对象

// 1. 申明结构体
type student struct {
	Name string
	Age  int
}

// 2. 申明结构体切片
type StudSlice []student

// 3. 实现库排序所需要的接口
func (ss StudSlice) Len() int {
	return len(ss)
}
func (ss StudSlice) Less(i, j int) bool {
	return ss[i].Age > ss[j].Age
}
func (ss StudSlice) Swap(i, j int) {
	tmp := ss[i]
	ss[i] = ss[j]
	ss[j] = tmp

	//ss[i], ss[j] = ss[j], ss[i] 这种写法和上面也是等价的
}

func main() {

	// 初始化切片
	var heroes StudSlice
	for i := 0; i < 10; i++ {
		hero := student{
			Name: fmt.Sprintf("hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	// 使用接口+库方法进行排序
	sort.Sort(heroes) // heroes实现了sort需要接口的所有方法，所以就实现了该接口，接口便可以接受heroes对象
	fmt.Println(heroes)
}
