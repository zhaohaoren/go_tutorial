// 自定义数据类型
package main

func main() {

	type myInt int // 自定义数据类型，相当于起一个别名： myInt 可以作为int来用了

}

type myFuncType func(int, int) int // 这种自定义的用处

// 函数在go中本质也是变量，可以作为形参传递，（这样可以回调）
func testVar2(myFunction myFuncType, num1 int, num2 int) int {
	return myFunction(num1, num2)
}
