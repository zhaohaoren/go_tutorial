// channel
package main

import (
	"fmt"
	"strconv"
	"time"
)

//为什么需要channel
//主线程无法确定是否所有的协程都执行完了

// channel
//本质：就是一个数据结构-队列
//队列本身是线程安全的，多个协程访问的时候不需要加锁
//channel是有类型的，string类型只能放string类型数据 -- 废话
// channel是引用类型，必须初始化make后才能写入数据

func main() {

	go recoverDemo()

	// 定义管道
	var intChan chan int // 定义一个int的channel
	intChan = make(chan int, 3)
	fmt.Println(intChan) // 是引用类型

	// 写入管道
	intChan <- 10
	intChan <- 11
	intChan <- 12
	//intChan <- 13 管道不能超出定义的容量
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	// 管道取出
	var num = <-intChan
	fmt.Println(num)
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	num = <-intChan
	//num = <- intChan
	//num = <- intChan  没有数据的时候再取报错 deadlock

	// 关闭管道
	// 关闭后不能再从管道写数据了，但是可以读数据
	// 个人理解：(关闭管道就意味着不会有数据再继续放入了，所以读完就可以退出，但是如果没关就可能有新数据，这时候我们需要阻塞等待新数据，但是如果是在主线程，因为不可阻塞主线程所以直接deadlock）
	intChan <- 10
	intChan <- 11
	close(intChan)
	//intChan <- 13  panic: send on closed channel
	fmt.Println(<-intChan)

	// 遍历管道
	// 使用for-range,不可以for循环，应为队列是在不断变的
	// 遍历如果管道没有close，遍历完后报deadlock错误
	// 遍历如果管道close了，正常遍历后退出遍历
	for v := range intChan {
		fmt.Print(" ", v)
	}
	fmt.Println()

	// 默认管道是 可读可写的
	// 只读管道 -- 函数的参数可以定义为只读或者只写的，别人调用的时候可以控制
	var readOnlyChan <-chan int = make(chan int, 3)
	fmt.Println(<-readOnlyChan)
	//readOnlyChan <- 1 // 不可写
	// 只写管道
	var writeOnlyChan chan<- int = make(chan int, 3)
	writeOnlyChan <- 1
	//<- writeOnlyChan 不可读

	selectDemo()

	// 练习1：一读一写完成后程序退出
	demoChannel()
	// 练习2：计算素数
	primeDemo()

}

func writeChannel(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readChannel(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", v)
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func demoChannel() {
	// 代码一环套一环的： read 套 write的close channel才会退出，main 又套在read 结束后close channel才结束
	var intChan = make(chan int, 50)  // go的读写频率不一致不会有问题，底层有处理
	var exitChan = make(chan bool, 1) // 单独定义一个管道判断是否诚信读取完成， 其实一个标志位也可以啊？
	go writeChannel(intChan)
	go readChannel(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}

	}
}

// --------------------------------------------------------
// 计算素数
func primeDemo() {
	// 需要开启一个管道不断的放入1-8000个数字
	// 开启4个协程不断从管道取出数据判断是否为素数
	// 然后将素数放入新管道
	var intChan = make(chan int, 1000)
	var primeChan = make(chan int, 2000)
	var exitChan = make(chan bool, 4)
	go putNum(intChan)

	for i := 1; i <= 4; i++ {
		go judge(intChan, primeChan, exitChan)
	}

	// 测试后发现：如果不是在主线程中，协程中执行程序，取不到数据的时候会被阻塞住
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 取出数据
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("prime: >", res)
	}

	fmt.Println("main exit")

}

func putNum(intChan chan int) {
	// 放8000个数入channel
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func judge(intChan chan int, primeChan chan int, exiChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		time.Sleep(time.Millisecond * 10)
		if !ok {
			// intChan取不到值了
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("this goroutine can not found prime any more")
	exiChan <- true
}

// --------------------------------------------------------
func selectDemo() {
	//使用select
	//无法确定管道是否关闭的时候取值，使用select防止出现deadlock。
	//方式：如果遇到了select case取不到值，不报deadlock而是执行下一个case
	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + strconv.Itoa(i)
	}
	for {
		select {
		case v := <-intChan2:
			fmt.Println("get int v: ", v)
			time.Sleep(time.Second)
		case v := <-stringChan: // 如果上面取不到数据就执行下面
			fmt.Println("get string v: ", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("all get nothing")
			time.Sleep(time.Second)
			return
		}
	}
}

// --------------------------------------------------------
func recoverDemo() {
	// 解决方法：
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverDemo panic happen", err)
		}
	}()

	// 如果协程中出现panic，会导致程序崩溃
	var myMap map[int]string
	myMap[0] = "hello" // 这个异常会导致程序崩溃！
}
