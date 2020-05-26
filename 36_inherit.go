// 继承 - 代码复用
package main

import "fmt"

type IStudent struct {
	Name  string
	age   int
	Score float64
}

type Pupil struct {
	IStudent      // 嵌入了Student的匿名结构体 - 继承
	Score float64 // 当结构体和匿名结构体 含有相同的字段的时候，我们将采用就近原则（这里的就近只是说取当前结构体的，如果结构体没有参考下面）
}

type Graduate struct {
	IStudent
}

func (s *IStudent) GetScore() {
	fmt.Println("score is", s.Score)
}

func (s *IStudent) SetScore(score float64) {
	s.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("pupil is testing")
}

func (g *Graduate) testing() {
	fmt.Println("graduate is testing")
}

//  --------继承2---------
type AA struct {
	Name string
}
type BB struct {
	Name string
}


//  --------多重继承--------- 尽量不要使用多重继承
type C struct {
	AA // 当父类都有一个属性，而自己却没有这个属性，那么访问的时候必须带上匿名结构体名
	BB
}

type D struct {
	int // 匿名结构体可以是基础数据类型的, 其实和非匿名使用没啥区别，一般也不会这么用
}

func main() {
	var pupil = &Pupil{}
	pupil.IStudent.Name = "jack"
	pupil.IStudent.age = 20 // 子类可以使用父类所有的属性，不存在共有和私有的问题
	pupil.age = 22          // 也可以直接使用，不需要在通过父类来调用： 他会依次从对应的类去找对应的属性
	pupil.testing()
	pupil.IStudent.SetScore(99.98)
	pupil.IStudent.GetScore()
	fmt.Println(pupil.IStudent.Score)
	fmt.Println(pupil.Score)

	var c C
	c.AA.Name = "lucy"
	fmt.Println(c.AA.Name)

	var c2 C = C{ //创建结构体同时初始化继承结构体
		AA{Name: "AA"},
		BB{Name: "BB"},
	}
	fmt.Println(c2)

	var d D
	d.int = 20
	fmt.Println(d)
}
