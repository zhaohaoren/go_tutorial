// go命令行参数
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 使用os.Args 这个切片可以获取到命令行的参数， 缺点：严格按照规定的顺序来写命令行参数
	fmt.Println("params len is", len(os.Args))

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	fmt.Println("use flag-------------------------------------")

	// 使用flag包来解析
	var user string
	var pwd string
	var host string
	var port int
	// &user: 接受保存参数值， "u" 表示接受参数中-u的参数，""默认值，注释说明
	flag.StringVar(&user, "u", "", "这里写说明")
	flag.StringVar(&pwd, "pwd", "", "这里写说明")
	flag.StringVar(&host, "h", "", "这里写说明")
	flag.IntVar(&port, "p", 8080, "这里写说明")
	flag.Parse() // 必须要有该操作进行转换，才会将参数的内容赋给那些变量
	fmt.Println(user, pwd, host, port)
	// go run 42_cmd_params.go -u zhao -pwd 123 -h 127.0.0.1 -p 2013
}
