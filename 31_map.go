// map
package main

import (
	"fmt"
	"sort"
)

func main() {

	// map 就是kv结构： key可以支持很多类型，通常为int 或者string ，也可以指针chanel，bool 结构体数组等，map是引用类型的。
	// slice map function 是不可以作为key的 因为这些不支持==进行比较
	var map1 map[string]string
	// map在使用前需要对其分配空间后才能使用，否则报错
	map1 = make(map[string]string, 1)
	map1["no1"] = "justin"
	map1["no2"] = "justin2" // map是可以自动扩容的
	map1["no1"] = "zhaohaoren"
	fmt.Println(map1)

	// 初始化方式2：不需要指定大小的
	var map2 = make(map[string]string)
	map2["no1"] = "justin"
	map2["no2"] = "jack"
	fmt.Println(map2)

	// 第三种方式：申明并初始化
	var map3 = map[string]string{
		"no1": "a",
		"no2": "b", // 这里逗号都不能省
	}
	fmt.Println(map3)

	var map4 = map[string]map[string]string{
		"a": { //map-map结构 更加适合使用map-struct
			"name": "zhao",
			"sex":  "man",
		},
	}
	fmt.Println(map4)

	//map的增删改查
	// 删： 如果key不存在不报错，也不会执行删除操作
	delete(map1, "no1")
	fmt.Println(map1)
	// 没有一次性方法删除所有的key，可以遍历所有key一个个的去删除，或者直接make一个新的给这个map
	// 查： 可以返回两个值，一个是表示是否有该key的
	val, hasVal := map1["no1"]
	if hasVal {
		fmt.Println(val)
	} else {
		fmt.Println("no key")
	}
	// 增 和 改就没啥说的了

	// map的遍历 -- 用for range
	for k, v := range map1 {
		fmt.Printf("key:%v val:%v\n", k, v)
	}

	// map的kv对数
	fmt.Println(len(map1))

	// map切片： map的个数可以动态的变化， 本质上还是切片，只是切片中存的是map
	var map5 []map[string]string        // 先申明
	map5 = make([]map[string]string, 1) // 切片初始化
	if map5[0] == nil {
		map5[0] = make(map[string]string, 2) // map初始化
		map5[0]["name"] = "zhao"
		map5[0]["age"] = "20s"
	}
	// 后续添加使用append
	var newMap = map[string]string{
		"name": "jack",
		"age":  "22",
	}
	map5 = append(map5, newMap)
	fmt.Println(map5)

	// map排序
	// go没有排序的map，只能对key先进行排序，然后按照顺序取出value
	var keys []string
	for k, _ := range map1 {
		keys = append(keys, k) // 先将所有key放入切片
	}
	sort.Strings(keys) // 使用切片进行排序
	for _, v := range keys {
		fmt.Println("Key:", v, "Value", map1[v]) // 按照排序的切片输出
	}

}
