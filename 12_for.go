// for 循环
package main

import "fmt"

func main() {
	// 第一种写法
	for i := 1; i <= 10; i++ {
		fmt.Println("for ", i)
	}

	// 第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("for2 ", j)
		j++
	}

	// 第三种方式
	// 等价于 for ;; {}
	k := 1
	for {
		fmt.Println("for3 ", k)
		if k == 10 {
			break
		}
		k++
	}

	// 第四种方式
	// for-range()  和python的range()很像 -- 可以用来遍历数组和字符串，中文不会出问题
	strs := "abcdefg中文"
	for idx, val := range strs {
		fmt.Printf("%d,%c \n", idx, val)
	}

	//也可以使用传统方式来遍历字符串
	// 注意：中文会出错，因为len是按照字节来量长度的，中文utf8 会占用三个字节。
	// 解决方式：使用切片
	strsWithCh := "我是justin"
	for i := 1; i < len(strsWithCh); i++ {
		fmt.Printf("%c ", strsWithCh[i])
	}

}
