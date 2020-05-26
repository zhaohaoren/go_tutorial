// 方法 - 作用在指定的数据类型上的才是方法
package main

import "fmt"

type A1 struct {
	// 时刻注意结构体是值拷贝的，需要改变值得时候需要使用地址传递
	Name string
}

// 和之前理解的还不一样，不在结构体内部的，而是外部指定结构体的
func (a *A1) test() { // (a A1) 表明和A1类型绑定的，只能A1的对象使用, 一般推荐使用*Type来进行函数申明，避免结构体的值拷贝（当然方法内 本质上使用的是*a,调用本质上应为&a）
	fmt.Println(a.Name)
}

type integer int

// 自定义类型都可以有方法： type定义的类型都可以有方法，不仅仅是struct，可以给任何定义类型赋予它方法。
func (i integer) testInt() {
	i++
	fmt.Println(i)
}

func main() {
	var a = A1{"justin"}
	a.test()

	// 调用和传输机制
	// 和函数不同的是：方法会将调用该方法的变量本身作为参数传递到方法中去

	var i integer = 10
	i.testInt()
	fmt.Println(i)

	// 注意点：方法可以传入变量或者&变量，效果是一样的，底层都是值拷贝
	(&i).testInt()
	fmt.Println(i) // 看着好像i传递的是地址，本质上传的还是值，所以这里不会变，依然为10，值拷贝还是引用拷贝在于参数是带不带*

	//本节重点便在于 使用方法的时候注意调用者是引用传递还是值传递，针对结构体这种值类型尤其要注意。
}
