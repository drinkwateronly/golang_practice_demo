package main

import (
	"fmt"
	"os"
	"time"
)

func basicInfo() {
	//获取主机名
	fmt.Println(os.Hostname())
	//获取当前目录
	fmt.Println(os.Getwd())
	//获取用户ID
	fmt.Println(os.Getuid())
	//获取有效用户ID
	fmt.Println(os.Geteuid())
	//获取组ID
	fmt.Println(os.Getgid())
	//获取有效组ID
	fmt.Println(os.Getegid())
	//获取进程ID
	fmt.Println(os.Getpid())
	//获取父进程ID
	fmt.Println(os.Getppid())
	//获取环境变量的值
	//fmt.Println(os.Getenv("GOPATH"))
	//设置环境变量的值
	//os.Setenv("NAME", "test")

	//清除所有环境变量（慎用）
	//os.Clearenv()
	//改变当前工作目录
	//os.Chdir("/")
	fmt.Println(os.Getwd())
}

func fileOp() {
	//创建目录
	os.Mkdir("abc", os.ModePerm) // 0777
	//创建多级目录
	os.MkdirAll("xxxx/x", os.ModePerm)
	//删除文件或目录
	os.Remove("xxx/x")
	//删除指定目录下所有文件
	os.RemoveAll("xxxxx")
	//重命名文件
	os.Rename("./test1.txt", "./twst1_new.txt")
	// 获取文件信息-->FileInfo接口类型，有很多方法，如文件大小，是否是文件夹等等
	info, _ := os.Stat("./lqz")
	fmt.Println(info.Size())
	//创建文件
	f1, _ := os.Create("./test.txt")
	defer f1.Close()
	//修改文件权限
	os.Chmod("./test.txt", 0777)
	//修改文件所有者
	os.Chown("./test.txt", 0, 0)
	//修改文件的访问时间和修改时间
	os.Chtimes("./test.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour))
	//打开文件
	//func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile("./test2.txt", os.O_RDWR|os.O_CREATE, 0766)

}

func main() {
	basicInfo()

}
