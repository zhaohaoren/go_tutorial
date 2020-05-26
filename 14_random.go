// 随机数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 需要一个种子
	fmt.Println(rand.Intn(100) + 1)
}
