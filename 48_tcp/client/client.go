package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "0.0.0.0:8888")
	if e != nil {
		fmt.Println("fail to get conn...")
		return
	}
	fmt.Println("conn success", conn)
	// 客户端发送单行数据然后退出
	reader := bufio.NewReader(os.Stdin)
	//os.stdin 表示标准输入，即键盘输入，终端输入,
	// 这些其实本质上都是文件linux一切都是文件，socket等硬件交互都是写入文件再读出来交互的
	line, e := reader.ReadString('\n')
	if e != nil {
		fmt.Println("read console fail", e)
	}
	// 将line发送给服务端
	n, e := conn.Write([]byte(line))
	if e != nil {
		fmt.Println("conn write fail", e)
		return
	}
	fmt.Println("client send bytes", n)
}
