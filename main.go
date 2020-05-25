package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var staticDir string
var port int

func main() {

	// 处理参数
	flag.StringVar(&staticDir, "d", "", "-d=静态文件所在目录")
	flag.IntVar(&port, "p", 8888, "-p=静态服务的端口号")

	flag.Parse()

	// 静态文件目录
	path := getStaticDir()

	// 验证目录是否存在
	if isNotExist(path) {
		printErr("%s 不存在或不是一个目录， 请检查参数是否正确", path)
	}

	fmt.Printf("\n启动静态文件服务\n监听的目录是: %s\n", path)
	fmt.Printf("可以通过浏览器上的 http://localhost:%d 访问\n\n", port)

	http.Handle("/", http.FileServer(http.Dir(path)))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		printErr("%v", err)
	}
}

// 日志打印助手函数
func printErr(format string, messages ...interface{}) {
	fmt.Printf("\n服务启动失败：%s\n\n", fmt.Sprintf(format, messages...))
	os.Exit(1)
}

// 检查目录是否真实存在
func isNotExist(path string) bool {
	info, err := os.Stat(path)

	// 如果文件状态出现错误， 直接返回
	if err != nil || os.IsNotExist(err) {
		return true
	}

	// 如果不是目录
	if !info.IsDir() {
		return true
	}

	return false
}

// 获取静态文件的目录
func getStaticDir() string {
	execDir, _ := os.Getwd()

	// 如果没有指定静态文件路径，直接返回可执行文件执行时候的目录
	if staticDir == "" {
		return execDir
	}

	// 如果是绝对路径， 只要是 / 开头就当绝对路径了
	if staticDir[0] == '/' {
		return "/" + strings.TrimLeft(staticDir, "/")
	}

	return execDir + "/" + staticDir

}
