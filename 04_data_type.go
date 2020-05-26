// go的数据类型
// 包使用看文档就行

package main

import (
	"fmt"
	"strconv"
	_ "unicode" //不使用这个包，但是又想保留他就是用_ 表示忽略它
	"unsafe"
)

func main() {
	// 基本数据类型
	var t = 100
	// 获取变量类型和字节大小
	fmt.Printf("t's type is: %T; \nt's size is: %d; \n", t, unsafe.Sizeof(t))

	// 整形---------------------------------------------
	//有符号整数
	var i int = 1
	var i1 int8 = -128 // -128~127  1byte=8位
	var i2 int16 = 2   // 2byte
	var i3 int32 = 2   // 4byte
	var i4 int64 = 2   // 8byte
	fmt.Println(i, i1, i2, i3, i4)

	//无符号整数
	var j uint = 1
	var j1 uint8 = 1 // 0-255 没有符号
	var j2 uint16 = 1
	var j3 uint32 = 1
	var j4 uint64 = 1
	var j5 byte = 1
	fmt.Println(j, j1, j2, j3, j4, j5)

	var r rune = 12        //等价int32 表示一个Unicode码
	var ch rune = '\u8d75' // 好像也不能表示中文
	fmt.Println(ch)
	fmt.Println(r)

	//存字符用byte，等价unit8，表示的数字范围也一样，只是他还可以接收字符而已
	var c byte = 'c'
	fmt.Println("c=", c)         // 直接输出就输出的是ASCII码值
	fmt.Printf("c = %c \n", c)   // 需要输出字符就需要按照格式输出
	var ch1 int = '北'            // 只要类型能表示的范围在该符号对应的编码值范围就可以接受中文
	fmt.Printf("中文: %c \n", ch1) // 输出格式化就行

	// 浮点型---------------------------------------------
	// 浮点数都是有符号的，所以第一位都是符号位
	var k float32 = -123.00009  //单精度
	var k1 float64 = -123.00009 //双精度 - 推荐直接使用
	fmt.Println(k, k1)          // 两者所能表示的精度不一样，小的精度可能丢失
	var k2 = .1                 // 等价于 0.1
	fmt.Printf("%T", k2)        // 给一个小数，go会默认为float64，一个整数则默认为int64
	//支持科学计数法
	var k3 = 1.2345e2
	fmt.Println(k3)

	var b bool = true // 只占一个字节 只可取true和false
	fmt.Println(b)

	var s string = "abc你好啊" // 只要设置的是utf8 就不可能出现乱码问题
	//s[0] = 'm' go的字符串不可以被改变
	fmt.Println(s)
	var s1 string = `*\反引号会原样输出字符串,包括格式\n`
	fmt.Println(s1)
	//字符串拼接
	var s3 string = "str1 " + // 多行拼接 "+" 必须放在后面
		"str2 "
	s3 += "str3 "
	fmt.Println(s3)

	/*
	Go 基本数据默认值，如果没有对变量赋值go自动为其赋默认值
	整形：0
	浮点：0
	字符串：""
	布尔：false
	*/

	base_type_convert()
	//复杂数据类型

}

//基本数据类型的转换
func base_type_convert() {
	/*
		Go的类型在不同类型变量不可以自动转换，只能显式转换
	 */
	var conv int = 100
	//var conv2 float64 = conv 不可以直接转换
	var conv2 float64 = float64(conv) //转换本质：将conv的值转成一个值然后再给conv2，本身conv是不变的。
	fmt.Println(conv2)

	//var num1 int32 = 10
	//var num2 int64 = 1+num1 //数值可以转换但是不可以直接赋值给不同类型，当和其他数字相加的时候，以相加的那个数有类型的为准
	//fmt.Println(num2)

	/*
		基本数据类型和string的转换
	 */
	var num int = 99
	var str string
	var str2 string
	// 方法1
	str = fmt.Sprintf("%d", num)
	// 方法2
	str2 = strconv.FormatInt(int64(num), 10)
	str2 = strconv.Itoa(num) //方法2.2
	fmt.Printf("%T \n", str)
	fmt.Printf("%T \n", str2)
	// 将string转为基本数据类型
	var st string = "abc"
	var num2 int64
	num2, _ = strconv.ParseInt(st, 10, 64) // 字符串，进制，转成多少位的类型 返回值：转换后的值，错误信息（使用_表示忽略该错误）
	fmt.Printf("num type: %T", num2)
}
