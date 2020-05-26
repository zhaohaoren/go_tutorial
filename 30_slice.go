// slice 切片
package main

import "fmt"

func main() {
	// 切片是一个引用类型： 其本质为动态数组

	var intArr = [...]int{1, 3, 5, 7, 9}
	fmt.Println(intArr)
	// 定义切片： 所谓引用就是引用一个数组
	slice := intArr[1:4] // 切出数组索引为[1,3)的数组
	fmt.Println(slice)
	fmt.Println("slice is:", slice)
	fmt.Println("slice len:", len(slice))      // 切片的大小：切片已经使用的容量
	fmt.Println("slice capacity:", cap(slice)) //切片的容量：内存为该切片分配的内存容量，应该是每次空间不足对其扩容，容量是动态变化的
	slice[0] = 0
	// 因为切片是引用类型的，所以slice的修改会直接反应到引用的数组
	fmt.Println(slice)
	fmt.Println(intArr)
	//slice 内存中相当于一个结构体：存放引用的内存起始地址，引用的长度，该切片当前的容量

	// 申明方式：
	// 切片使用方式1: 直接引用定义的数组
	var slice1 = intArr[0:5]
	fmt.Println("slice1:", slice1)

	// 切片使用方式2: make来直接创建切片 make(切片类型，长度，容量)
	// 和用法1的区别，1是真数组，我们程序可直接见的，2则是内部也会创建一个数组对于我们来讲是不可见的
	var slice2 = make([]int, 3, 10)
	slice2[0] = 1
	slice2[1] = 1
	fmt.Println("slice2:", slice2)

	// 切片使用方式3: 直接指明其对应数组，但是数组不需要赋值大小
	var slice3 = []int{1, 2, 3}
	fmt.Println("slice3:", slice3)

	// 切片遍历
	// 1.传统for
	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}
	// 2.for-range
	for i, v := range slice3 {
		fmt.Println(i, v)
	}

	// 添加元素
	slice3 = append(slice3, 4, 5, 6)
	// 添加切片 ... 必须带上
	slice3 = append(slice3, slice1...)
	// append 在底层会创建一个新数组，并将老数据复制给新数组！--- 在append之后就不会再对原引用数组有影响了，【感觉一举两用：既得到了】

	// 切片的拷贝: 有多少拷贝多少，没有就保留原来的0，大拷贝小也可以，有多少考多少
	var a = []int{1, 2, 3, 4, 5}
	var sliceA = make([]int, 10)
	copy(sliceA, a) // 必须都是切片类型的
	fmt.Println(sliceA)

	// string 和 slice
	// string的底层就是一个byte数组，可以切片
	// 但是string是不可变的，通过切片去修改是会报错的
	var str = "zhaohaoren666@gmail.com"
	var sliceStr = str[:6] //var 虽然是自动推导的，但是一定申明的时候定下来了，那么以后其类型也是确定了的
	fmt.Println(sliceStr)
	//str[0] = 'm'
	//sliceStr[0] = 'm' 都是不可修改string的
	//如果一定要修改字符串: 先转成[]byte或[]rune( rune可以处理中文，rune是按照字符进行处理的)切片，在转成string
	sliceStr2 := []byte(str)
	sliceStr2[0] = 'm'
	str = string(sliceStr2)
	fmt.Println(str)

	// 注意事项
	// 和python一样，切片表示
	var slice4 = intArr[:2] /*= intArr[0:2]*/
	slice4 = intArr[1:]     /*= intArr[1:end]*/
	slice4 = intArr[:]      /*= intArr[0:end]*/
	fmt.Println(slice4)
	// 切片还可以继续切片
	var slice5 = slice4[:]
	fmt.Println(slice5)

	testSliceAppend()

	fmt.Println(genFebSlice(0))
}

func testSliceAppend() {
	var arrs = [4]int{1, 3, 45, 6}
	var slice = arrs[:]
	slice[0] = 0
	slice = append(slice, 10)
	slice[1] = 0
	fmt.Println(arrs)
}

func genFebSlice(n int) []uint64 {
	if n < 1 {
		return nil
	}
	var febSlice = make([]uint64, n)
	febSlice[0] = 1
	if n > 1 {
		febSlice[1] = 1
	}
	for i := 2; i < n; i++ {
		febSlice[i] = febSlice[i-1] + febSlice[i-2]
	}
	return febSlice
}
