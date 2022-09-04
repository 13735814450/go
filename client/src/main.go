package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main()  {
	rpcJsonclient()
	fmt.Println(111)
}

//func sclient(){
//	tcpaddr, err := net.ResolveTCPAddr("tcp4","127.0.0.1:8000")
//	checkError(err)
//	con, err := net.DialTCP("tcp",nil, tcpaddr)
//	checkError(err)
//	n, err := con.Write([]byte("HEAD / HTTP/1.0\r\n\n\n "))
//	checkError(err)
//	println(n)
//	var b []byte
//	r, err :=con.Read(b)
//	checkError(err)
//	println(r)
//	result, err := ioutil.ReadAll(con)
//	println(111)
//	checkError(err)
//	println(string(result))
//	os.Exit(0)
//}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, " error:%s", err.Error())
		os.Exit(0)
	}
}

func startClient()  {
	log.Print("start client ....")
	//连接端口
	port := "127.0.0.1:7777"
	//基于tcp4 端口为7777的地址对象
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)
	// 基于tcp的连接，
	// 第一个参数  是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
	// 第二个参数  表示本机地址，一般设置为nil
	// 第三个参数 表示远程的服务地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//go 提供的流拷贝，将conn的值拷贝到stdout里面，标准输出，也就是控制台
	io.Copy(os.Stdout,conn)
}
func startUdp()  {
	log.Print("start client ....")
	//连接端口
	port := "127.0.0.1:7777"
	//基于tcp4 端口为7777的地址对象
	tcpAddr, err := net.ResolveUDPAddr("udp4", port)
	checkError(err)
	// 基于tcp的连接，
	// 第一个参数  是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
	// 第二个参数  表示本机地址，一般设置为nil
	// 第三个参数 表示远程的服务地址
	conn, err := net.DialUDP("udp", nil, tcpAddr)
	checkError(err)
	_,err =conn.Write([]byte("anything"))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	println(string(buf[0:n]))
	os.Exit(0)
}

type Args struct {
	A,B int
}

type Quotient struct {
	Quo, Rem int
}
func rpcclient(){
	serverAddress := "127.0.0.1:1234"
	client, err := rpc.Dial("tcp", serverAddress)
	//client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dial error:" , err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith : %d * %d = %d\n",args.A,args.B,reply )

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: " , err)
	}
	fmt.Printf("arith: %d/%d=%d 余 %d\n",args.A,args.B,quot.Quo,quot.Rem )
}
func rpcJsonclient(){
	serverAddress := "127.0.0.1:1234"
	client, err := jsonrpc.Dial("tcp", serverAddress)
	//client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("dial error:" , err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith : %d * %d = %d\n",args.A,args.B,reply )

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: " , err)
	}
	fmt.Printf("arith: %d/%d=%d 余 %d\n",args.A,args.B,quot.Quo,quot.Rem )
}
