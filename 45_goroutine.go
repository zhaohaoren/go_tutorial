// 协程
package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//go主线程，也可以直接称为线程或者直接叫进程
//一个go主线程可以起多个协程（go可以轻轻松松的开几万个协程） 主线程如结束，协程也就完蛋了，但是不代表主线程输出结束了就表示主线程结束了

//协程的特点
//1、有独立的栈空间
//2、共享进程的堆空间
//3、调度由用户控制
//4、协程是轻量级的线程

func main() {

	// 设置Go运行的cpu数量
	num := runtime.NumCPU() // 获取当前电脑cpu数量
	fmt.Println("cpu num is:", num)
	runtime.GOMAXPROCS(num - 1) // 设置程序可以利用的cpu数量

	go testRoutine() // 开启了一个协程

	// 主线程是一个物理线程，直接作用于cpu是重量级的消耗资源
	// 协程是从主线程开启的，轻量级线程是一个逻辑态，不是真正意义上的线程，所以资源消耗少，可以轻松开启上万个协程，（依据P 创建的一系列的协程队列）
	// 其他语言基于线程的，是内核态的，需要占用cpu的，几千个就可能消耗完了cpu了，所以撑不了那么大
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	for i := 1; i <= 200; i++ {
		go demoRoutine(i)
	}

	// 打印的map 不一定所有协程都执行完成了，后面数据的0是因为数字太大存不下去了
	fmt.Println(myMap)
	fmt.Println(len(myMap))

	//MPG模型 - go协程的调度模型  参考：https://i6448038.github.io/2017/12/04/golang-concurrency-principle/
	//M 操作系统的主线程
	//P 协程执行需要的上下文环境（开启所需要的资源）
	//G 协程
}

func testRoutine() {
	for i := 1; i <= 10; i++ {
		fmt.Println("testRoutine() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

var (
	myMap = make(map[int]int, 10)
	// 申明一个全局互斥锁
	lock sync.Mutex
)

func demoRoutine(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里会报错 -- 解决方法：
	// 1、加全局锁控制写入(适用于低水平程序线程)，-- 缺点：我不知道什么时候所有的协程都执行完了，所以需要channel来实现协程的通信
	// 2、使用channel
	//concurrent map writes  不可以去并发的写 go build -race xx.go 可以检测并发写入问题
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
