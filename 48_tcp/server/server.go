package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server begin to listen...")
	// 开启监听 ：参数1 表示使用tcp协议，参数2 表示监听的ip和端口
	listener, e := net.Listen("tcp", "0.0.0.0:8888") //0.0.0.0 ipv4和v6都支持，127那个只支持ipv4
	if e != nil {
		fmt.Println("listen error:", e)
		return
	}
	defer listener.Close()
	for {
		//等待客户端连接,获取连接
		fmt.Println("wait to accept...")
		conn, e := listener.Accept() // 这里会阻塞等待着,可以使用telnet进行测试
		if e != nil {
			fmt.Println("fail to get conn")
		}
		fmt.Println(conn)
		fmt.Println("remote host conn:", conn.RemoteAddr().String())
		go process(conn)
	}

	fmt.Println(listener)

}

//协程处理客户端信息
func process(conn net.Conn) { // 一个连接对应一个处理
	defer conn.Close() // 一定不要忘了关闭连接
	for { // 循环接受消息

		//
		buf := make([]byte, 1024)
		// read会等待客户端发送请求信息，如果客户端没有write，则会阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read fail", err)
			return
		}
		// 输出接受到消息，消息会放在buffer中是一个字节数组，需要转为string，n表示我们获取的字节数组长度，收多少转多少。
		fmt.Print("server get:", string(buf[:n]))
	}
}
