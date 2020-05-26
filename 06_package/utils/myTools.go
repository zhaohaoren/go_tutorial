// 包是作为一个类的，而包中的文件名是没有什么特殊用处的，对外使用包内方法的时候不需要走文件名，而是走包名

package utils

import "fmt"

func PublicFunction() {
	fmt.Println("this is a public function")
}

func privateFunction() {
	fmt.Println("this is a private function")
}

func init() {
	fmt.Println("1 init")
}

func main() {
	fmt.Println("1 main")
}
