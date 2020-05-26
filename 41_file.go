// 文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 程序中文件是以流的形式来操作的
// 写入内存叫做输入流，内存写入文件叫做输出流
// go操作文件使用os.File，其中File是一个结构体

func main() {
/*
	//#1 打开文件
	// open 会返回一个file对象，和一个错误信息
	// file 还可以叫file指针，文件句柄等
	file, e := os.Open("./41_file")
	if e != nil {
		fmt.Println("open file error:", e)
	}
	fmt.Println(file)
	defer file.Close()
	//#2 读取文件(带缓存方式) - 返回一个reader
	reader := bufio.NewReader(file)
	//交由reader去读取文件
	for {
		s, e := reader.ReadString('\n') // 读取到换行就结束
		fmt.Print(s)
		if e == io.EOF { // 表示文件读取到末尾了
			break
		}

	}
	fmt.Println("\n----------------一次性读取文件----------------")
	// 一次性读取文件, 这种没有显示的open操作就不要我们去close文件
	bytes, e := ioutil.ReadFile("./41_file")
	if e != nil {
		fmt.Println("read all error:", e)
	}
	fmt.Printf("%v", string(bytes))
*/
	//#3 关闭文件

	//e = file.Close()
	//if e != nil {
	//	fmt.Println("file close error:", e)
	//}
	//
	//writeFile()
	//
	//copyFile()
	fmt.Println(copyFile2())

	//fmt.Println(PathExist("./41_file"))
}

func writeFile() {
	// 写文件  文件路径，打开方式，linux权限
	file, e := os.OpenFile("41_writeFile", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("write file error:", e)
		return
	}
	defer file.Close()
	str := "justin write this"
	writer := bufio.NewWriter(file) // 带缓存的写入
	writer.WriteString(str)         // 这一步其实是先将文件写入缓存的
	writer.Flush()                  // 将缓存刷新到文件中
}

// 使用ioutil来完成文件的拷贝
func copyFile() {
	file1 := "./41_file"
	file2 := "./41_file_copy"

	data, e := ioutil.ReadFile(file1)
	if e != nil {
		fmt.Println("read file error:", e)
	}
	err := ioutil.WriteFile(file2, data, 0666)
	if err != nil {
		fmt.Println("write file error", err)
	}
}

// 使用流进行读写
func copyFile2() (written int64, err error) {
	file1 := "./41_file"
	file2 := "./41_file_copy.txt"

	srcFile, e := os.Open(file1)
	if e != nil {
		fmt.Println(e)
	}
	defer srcFile.Close()
	//reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(file2, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	//writer := bufio.NewWriter(destFile)
	defer destFile.Close()

	//return io.Copy(writer, reader)  //这种方法不可行-拷贝文件没内容！(视频讲错了，这copy最后还是需要flush刷入文件中)
	return io.Copy(destFile, srcFile) // 可以直接使用这样拷贝
}

// 判断文件是否存在
func PathExist(path string) (bool, error) {
	_, e := os.Stat(path)
	if e == nil {
		// 文件存在无异常
		return true, nil
	}
	// 异常为文件不存在返回
	if os.IsNotExist(e) {
		return false, nil
	}
	// 判断出现异常，无法判断是否存在文件返回
	return false, e
}
