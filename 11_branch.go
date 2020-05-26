// 顺序 分支 循环 之 分支

package main

import "fmt"

func main() {

	//单分支： 必须带大括号，条件可以不加括号
	var age int = 19
	if age > 18 {
		fmt.Println("over 18")
	}

	//双分支： 格式有要求
	// 还支持这种写法：在if中直接定义变量
	if high := 10; high > 1 {
		fmt.Println("over 1")
	} else {
		fmt.Println("low 1")
	}

	//多分支
	if num := 1; num > 1 {
		fmt.Println(">1")
	} else if num > 0 {
		fmt.Println("1=> num >0")
	} else {
		fmt.Println("<= 0")
	}

	// switch case
	// 不再需要break！ case 可以加多个条件
	var key byte
	var character byte = 'f'
	fmt.Println("请输入一个字符：a,b,c,d,e,f")
	_, _ = fmt.Scanf("%c", &key)
	switch key { // 表达式不拘泥于常量，一个有返回值的函数都可以
	case 'a', 'b', 'c': //case 需要和switch表达式的类型一致，也不需要是常量可以用变量表示，但是所有case表达式常量不能重复,变量不会去区分
		fmt.Println("less d")
		fallthrough //向下穿透，以前不加break会自动穿透，这个需要加这个才穿透，不加默认为break的（很有意思的设计，确实break使用的更多点，所以使用穿透的可能小点）
	case 'd':
		fmt.Println("d")
	case 'e', character:
		fmt.Println("over d")
	default:
		fmt.Println("input error")
	}

	//switch case 用法2 -- 可以当做if else来使用
	var age2 int = 20
	switch {
	case age2 == 20:
		fmt.Println("20")
	case age2 == 30:
		fmt.Println("30")
	}

	//type switch  需要接口的知识
}
