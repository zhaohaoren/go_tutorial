package model

// go中 首字母小写被认为是私有变量
var private_val int = 10
// 首字母大写认为是public变量
var Public_val int = 11

// 不同包无法访问-使用工厂模式
type student struct {
	// 结构体小写其他包也不能访问
	Name string
	age  int // 该字段为小写，其他包也不能直接访问
}

// 提供外包创建结构体接口
func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		age:  age,
	}
}

// 提供外包获取私有属性接口
func (stu *student) GetAge() int {
	return stu.age
}
