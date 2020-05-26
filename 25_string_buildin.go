// string 自带函数  builtin,strings,strconv包
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.求长度：len - builtin包里面  内建函数
	var str = "hello赵"
	fmt.Println(len(str))

	//2.遍历：切片
	var str2 = []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}
	fmt.Println()

	//3.字符串转整数
	n, err := strconv.Atoi("12a")
	if err != nil {
		fmt.Println("convert failed! ", err)
	} else {
		fmt.Println(n)
	}

	//4.数字转字符串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str=%v, type=%T \n", str3, str3)

	//6.字符串转byte，byte转字符串
	var bytes = []byte("hello")
	fmt.Println(bytes)
	var strBytes = string(bytes)
	fmt.Println(strBytes)

	//7.10进制转其他进制: 返回是字符串
	fmt.Println(strconv.FormatInt(123, 2))

	//8.查找子串,统计子串，子串第一次，最后一次出现的位置
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.Count("abcaabc", "abc"))
	fmt.Println(strings.Index("abcaabc", "abc"))
	fmt.Println(strings.LastIndex("abcaabc", "abc"))
	//8.1 子串替换
	fmt.Println(strings.Replace("abcaabc", "abc", "mmm", 1 /*希望替换几个*/))
	//8.2 字符串分割为数组
	fmt.Println(strings.Split("hello,world,ok", ","))

	//9. 不区分大小写的比较
	fmt.Println("abc" == "ABC")
	fmt.Println(strings.EqualFold("abc", "Abc"))
	//9.1 大小写转换
	fmt.Println(strings.ToLower("ABC"))
	fmt.Println(strings.ToUpper("abc"))

	//10 去除左右两边的空格
	fmt.Println(strings.TrimSpace("  i am justin      "))
	//10.1 去除指定左右两边的字符
	fmt.Println(strings.Trim("! this is justin!!", "!"))
	//10.2 去除左边/右边的指定字符
	fmt.Println(strings.TrimLeft("!hi!", "!"))
	fmt.Println(strings.TrimRight("!hi!", "!"))

	//11 字符串是否以 指定字符开头/结尾
	fmt.Println(strings.HasPrefix("hello", "h"))
	fmt.Println(strings.HasSuffix("hello", "h"))

}
