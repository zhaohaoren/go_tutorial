// 接口
package main

import "fmt"

// Go的核心编程就是接口编程， 接口是多态的前提：前面的继承可以看到——子类的方法只能自己来去调，没有办法由父类来多态的调用

// 总之：接口的作用：在保证全局统一标准的情况下不使用继承对结构体扩展功能，

// 接口的应用场景：
//我们需要哪些东西，比如我想输入哪些参数，然后你给我输出哪些参数，具体功能你自己去实现的，这时候可以定义接口来约束，接口的作用来源于这里
//这样项目开发，可以便于管理： 所有功能的具体实现都内聚于某个接口内，接口和接口之间是分开的非耦合的
//使用接口可以再不破坏继承的关系上，对结构体的功能进行扩展：
//比如：人继承于猴子，但是要扩展其他的功能，比如飞行，游泳这些，go只要加上这些接口，然后添加让类实现该接口即可。然后该类就可以直接调用了，
// 如果不用接口，那么要么单独为子类添加方法（行为系统内部会不统一，其他类的飞行不一定和这个类一个名），要么就要破坏继承关系了，继承一个飞行的类。

type usb interface {
	// 接口可以定义一组方法，但是不能定义任何变量，方法不能有方法体
	Start()
	Stop()
}

type disk interface {
	Read()
	Write()
}

type Camera struct {
}

// golang 实现接口的方式： 只要实现接口中的方法就算实现该接口了。如果两个接口方法一样，那么默认该结构体实现了这两个接口。
// go实现接口不是基于接口的，而是基于方法的。我写的方法和哪个接口耦合，那就是我实现了哪个接口
// 然后实现了（全部)接口方法的结构体，都可以通过结构体类接受
func (camera *Camera) Start() {
	fmt.Println("camera start")
}
func (camera *Camera) Stop() {
	fmt.Println("camera stop")
}

type Phone struct {
}

func (phone *Phone) Start() {
	fmt.Println("phone start")
}
func (phone *Phone) Stop() {
	fmt.Println("phone stop")
}
func (phone *Phone) Read() {
	fmt.Println("phone read")
}
func (phone *Phone) Write() {
	fmt.Println("phone write")
}

type Computer struct {
}

func (computer *Computer) Work(usb2 usb) {
	usb2.Start()
	usb2.Stop()
}

func main() {
	// 接口使用的简单案例：
	var phone = Phone{}
	var camera = Camera{}
	var computer = Computer{}
	computer.Work(&phone)
	computer.Work(&camera)

	//var usbVal usb = Camera{}: 错误写法：注意1. 结构体是指针类型，camera变量不是指针类型。2，上面实现接口方法的变量类型为*camera不是camera
	// 所以指针类型只能接受指针类型（存放地址），基本类型接受基本类型。找好对应关系就好了

	// 接口的注意事项
	// 接口不能创建对象，但是可以指向一个实现该接口的变量
	var usbVal usb = &Camera{}
	usbVal.Start() // 这种感觉和java不像是接口，更像是java中的父类指向子类
	// 必须实现全部接口方法才叫实现了该接口,接口变量才能引用该变量
	// 所有type自定义的类型都可以实现接口
	// 一个自定义类型可以实现多个接口，接口方法必须都实现
	var diskVal disk = &phone
	diskVal.Read()
	usbVal = &phone
	usbVal.Start()
	// go接口不可以有任何变量
	// 接口默认是一个指针（是引用类型），所以上面传的都是地址
	// 可以有空接口，没有任何方法，默认所有的接口都实现了空接口，可以指向任意的自定义变量
	// 接口可以继承，但是继承的接口不能有相同的方法名，会冲突报错

}
