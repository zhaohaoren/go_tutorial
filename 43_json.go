// json
package main

import (
	"encoding/json"
	"fmt"
)

type Astudent struct {
	Name string `json:"name"` //序列化指定key	-- 反射机制
	Age  int    `json:"age"`
}

func main() {

	serialize()

	deSerialize()
}

func serialize() {

	// 使用json.Marshal
	var stu = Astudent{
		Name: "justin",
		Age:  24,
	}

	bytes, e := json.Marshal(&stu) // 注意序列化的跨包问题，结构体的私有类型需要注意
	if e != nil {
		fmt.Println("serialize error:", e)
	}
	fmt.Println("serialize result:", string(bytes))
}

func deSerialize() {
	// 使用json.Unmarshal 反序列化
	var str = "{\"name\":\"justin\",\"age\":24}"
	var stu Astudent
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("deSerialize error", err)
	}
	fmt.Println(stu)
}
