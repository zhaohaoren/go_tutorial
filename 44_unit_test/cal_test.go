package test

import (
	"testing" // 引入test包，testing包提供对go包自动化测试的支持，使用go test命令会自动测试所有的TestXxx()的函数，注意Test后第一个字母大写
	// 工作内容
	// 1. 将所有xxtest.go结尾的文件引入
	// 2. 执行所有的Test方法
	// 3. 运行测试 go test -v v表示输出日志，不要日志就不写v； go test xx.go 测试单个文件，也支持多个文件；测试单个方法 go test xx.go xx方法名
)

// 编写测试用例
// 1. 函数名必须为Test开头的并且紧跟的后面的字母必须是大写
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		//fmt.Println("fail")
		// 输出日志并结束程序,fatalf表示致命的错误
		t.Fatalf("test fail hope:%v actual:%v", 55, res)
	}
	// 输出日志
	t.Logf("success")
}
