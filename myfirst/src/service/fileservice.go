package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func Filew() {
	var wireteString = "guoke\n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	n1, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("写入 %d 个字节", n1)
	n1, err1 = f.WriteString(wireteString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("写入 %d 个字节", n1)
	n2, err2 := f.Write([]byte(wireteString)) //写入文件(字节)
	if err2 != nil {
		panic(err1)
	}
	fmt.Println("写入 %d 个字节", n2)
}


func Filer() {
	filepath := "./test.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	file , err = os.Open(filepath);
	buf1 := make([]byte, 1024);
	for{
		n,_:=file.Read(buf1);
		if (n==0){
			break;
		}
		os.Stdout.Write(buf1[:n]);

		//println(buf1);
	}
	println("buf File read ok!")
}