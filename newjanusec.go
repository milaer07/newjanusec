package main

import (
	"flag"
	"fmt"
	"janusec/utils"
	"os"
	"path/filepath"
	"runtime"

	"github.com/milaer07/newjanusec/data"
)

func main() {
	//注册查看版本参数，查看版本后不运行其他代码，-version=1 或者 true
	ver := flag.Bool("version", false, "Display Version Information")
	flag.Parse()
	if *ver {
		fmt.Println(data.Version)
		os.Exit(0)
	}
	//修改工作目录为当前目录
	dir, _ := os.Executable()
	exePath := filepath.Dir(dir)
	err := os.Chdir(exePath)
	if err != nil {
		utils.CheckError("os.Chdir error", err)
		os.Exit(1)
	}
	//调整并发数
	runtime.GOMAXPROCS(runtime.NumCPU())

}
