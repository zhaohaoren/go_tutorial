// 面向对象
package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("message is: name[%v] gender[%v] age[%v] id[%v] score[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

func main() {

	var stu1 = Student{
		name:   "justin",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}

	fmt.Println(stu1.say())

}
