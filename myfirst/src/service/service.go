package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"myfirst/src/bean"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func M1() {
	fmt.Println("service m1")
}

func T1() {
	var nums = [10]int{1, 2, 3, 4, 5}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum=", sum)
	for i, num := range nums {
		fmt.Println("i=", i, " num=", num)
	}
	kvs := map[string]string{"a": "a1", "b": "b1", "c": "c1"}
	for k, v := range kvs {
		fmt.Printf("%s=%s  ", k, v)
	}
	fmt.Println()
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func T2() {
	var cou map[string]string
	cou = make(map[string]string)
	cou["china"] = "beijing"
	cou["russia"] = "moshike"
	cou["baji"] = "bajishoudou"
	for key, value := range cou {
		fmt.Println(key, value)
	}

	k, v := cou["china"]
	fmt.Println(k, v)
	if v {
		fmt.Println("exist")
	} else {
		fmt.Println("not ")
	}
	k, v = cou["japan"]
	fmt.Println(k, v)
}

func T3(i int) int {
	var sum int = 0
	sum += i
	if i > 1 {
		sum += T3(i - 1)
	}
	return sum
}

type phone interface {
	call()
}

type hwphone struct {
}
type nphone struct {
}

func (hwphone) call() {
	fmt.Println("i am hwphone")
}

func (nphone) call() {
	fmt.Println("i am nphone")
}

func T4() {
	var p phone
	p = new(hwphone)
	p.call()
	p = new(nphone)
	p.call()
}

func T5(i int, j int) int {
	var k int
	k = i / j
	return k
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func T6() {
	// 正常情况
	result, errorMsg := Divide(100, 10)
	if errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	result, errorMsg = Divide(100, 0)
	if errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func T7() {
	go say("hello")
	say("world")
}

func sum(s []int, n int, c chan int) {
	var sum int
	fmt.Println(n, s)
	for _, num := range s {
		sum += num
	}
	c <- sum
}

func T8() {
	s := []int{2, 4, 5, 7, 1, -9, 8}
	c := make(chan int)
	c1 := make(chan int)
	fmt.Println(s[3])
	fmt.Println(s[0:3])
	fmt.Println(s[:3])
	fmt.Println(s[1:3])
	fmt.Println(s[2:3])
	fmt.Println(s[3:3])
	go sum(s[:3], 1, c)
	go sum(s[3:], 2, c1)
	x := <-c
	y := <-c1
	fmt.Println(x)
	fmt.Println(y)
}

func T9() {
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	// 获取这两个数据
	ch <- 22222
	ch <- 1222
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	fmt.Println(n)
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println(x, y)
		x, y = y, x+y
	}
	close(c)
}

func T10() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
func T12() {
	a := make(chan int, 1024)
	b := make(chan int, 1024)
	for i := 0; i < 10; i++ {
		fmt.Printf("the %d num\n", i)
		a <- 1
		b <- 1
		select {
		case <-a:
			fmt.Printf("from a %d\n", a)
		case <-b:
			fmt.Println("from b %d", b)
		}
	}
}

func Longwait() {
	fmt.Println("long begin")
	time.Sleep(5 * time.Second)
	fmt.Println("long end")
}
func Shortwait() {
	fmt.Println("short begin")
	time.Sleep(2 * time.Second)
	fmt.Println("short end")
}

func Say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s, i)
	}
}

var Count int = 0

func Counter(lock *sync.Mutex, s string) {
	lock.Lock()
	Count++
	println(s, Count)
	lock.Unlock()
}

var l int = 0

func Channel(ch chan int) {
	l++
	ch <- l
	println("count", ch)
}

var timeout = make(chan bool, 1)

func T23() {
	time.Sleep(time.Minute)
	timeout <- true
}
func T24(ch chan int) {
	select {
	case <-ch:
		println(<-ch)
	case <-timeout:
		println("timeout")
	}
}

type Vector [1000000]int

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u[i]
	}
	c <- 1
}

func Doall() {
	num := runtime.NumCPU()
	println(num)
	u := Vector{}
	l := len(u)
	println(l)
	var index int = 0
	for j := 0; j < l; j++ {
		u[j] = index
		index++
	}
	//for j := 0; j < l; j++ {
	//	print(u[j])
	//}
	//c := make(chan int, num)
	//for i := 0; i < num; i++ {
	//	go v.DoSome(i*l/num, (i+1)*l/num, u, c)
	//}
	//
	//for i := 0; i < num; i++ {
	//	println(<-c)
	//}

}

func Tip(ip string) {
	addr := net.ParseIP(ip)
	if addr == nil {
		println("invalid ip")
	} else {
		println("ip is :", addr.String())
	}
	os.Exit(0)
}

func Tping(service string) {
	con, err := net.Dial("ip4:icmp", service)
	checkError(err)
	println(con)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, " error:%s", err.Error())
		os.Exit(0)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xfff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

//func Sserver() {
//	address := net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8000}
//	listener, err := net.ListenTCP("tcp4", &address)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		con, err := listener.AcceptTCP()
//		if err != nil {
//			log.Fatal(err)
//		}
//		println("add ", con.RemoteAddr())
//		//go echo(con)
//		handleConnection(con)
//	}
//}
//
//var index int
//
//func echo(con *net.TCPConn) {
//	index++
//	//tick := time.Tick(5 * time.Second)
//	//for now := range tick{
//	b := []byte(string(index))
//	n, err:=fmt.Fprintln(con, b)
//	//n, err := con.Write(b)
//	if err != nil {
//		log.Println(err)
//		con.Close()
//		return
//	}
//	fmt.Printf("send %d byte to %s \n", n, con.RemoteAddr())
//	//}
//}

func S1() {
	//监听端口
	port := ":7777"
	//构建一个基于tcp4的 端口为7777的监听地址对象
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkErr(err)
	//采用tcp连接
	listener, err := net.ListenTCP("tcp", tcpAddr)
	//for循环一个监听客户端的连接
	for {
		//接受客户端的连接，如果没有连接，就堵塞
		con, err := listener.Accept()
		checkErr(err)
		//输出客户端的地址
		log.Println("客户端连接成功:", con.RemoteAddr().String())
		handleConnection(con)
	}
}

//输出消息给客户端
func handleConnection(conn net.Conn) {
	for {
		//格式化当前时间，很抱歉，好像go格式化时间 不支持 YYYYMMdd 这样的写法
		format := time.Now().Format("2006/01/02 15:04:05")
		//输出到客户端
		fmt.Fprintln(conn, format)
		//线程睡眠一秒
		time.Sleep(time.Second * 1)
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func UDP() {
	//监听端口
	port := ":7777"
	//构建一个基于tcp4的 端口为7777的监听地址对象
	tcpAddr, err := net.ResolveUDPAddr("udp4", port)
	checkErr(err)
	//采用tcp连接
	con, err := net.ListenUDP("udp", tcpAddr)
	//for循环一个监听客户端的连接
	for {
		var buf [512]byte
		//接受客户端的连接，如果没有连接，就堵塞
		_, addr, err := con.ReadFromUDP(buf[0:])
		checkErr(err)
		daytime := time.Now().String()
		//输出客户端的地址
		println(string([]byte(daytime)))
		//log.Println("客户端连接成功:", con.RemoteAddr().String())
		con.WriteToUDP([]byte(daytime), addr)
	}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, que *Quotient) error {
	if args.B == 0 {
		return errors.New("divide num should not empty")
	}
	que.Quo = args.A / args.B
	que.Rem = args.A % args.B
	return nil
}

func Rpchttp() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		println(err.Error())
	}
}

func Rpctcp() {
	arith := new(Arith)
	rpc.Register(arith)
	addr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func Rpcjson() {
	arith := new(Arith)
	rpc.Register(arith)
	addr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "jim"
)

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertUser(user bean.User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO sys_user (`name`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.UserName, user.Password)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUser(id int) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM sys_user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func UpdateUser(user bean.User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE sys_user SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName, user.Password, user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func SelectUserById(id int) bean.User {
	var user bean.User
	//var id int
	//var name string
	//var password string

	err := DB.QueryRow("SELECT id,name,password FROM sys_user WHERE id = ?", id).Scan(&user.Id, &user.UserName, &user.Password)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return user
}

func SelectAllUser() []bean.User {
	//执行查询语句
	rows, err := DB.Query("SELECT id,name,password from sys_user")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var users []bean.User
	//循环读取结果
	for rows.Next() {
		var user bean.User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

var c redis.Conn
var err error

func Redis() {
	server := "127.0.0.1:6379"
	option := redis.DialPassword("123456")
	c, err = redis.Dial("tcp", server, option)
	if err != nil {
		log.Println("connect server failed:", err)
		return
	}
	log.Println("connect redis successfully")
}

func RedisGetSet() {
	_, err = c.Do("Set", "10", 100)
	if err != nil {
		fmt.Println("set int failed", err)
		return
	}

	_, err = c.Do("Set", "name", "jimyy")
	if err != nil {
		fmt.Println("set string failed", err)
		return
	}
	res, err := redis.Int(c.Do("Get", "10"))
	if err != nil {
		fmt.Println("get string failed,", err)
		return
	}
	fmt.Println(res)
	// redigo 通过redis.String()函数来获取字符串
	res1, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("get string failed,", err)
		return
	}
	fmt.Println(res1)
}

func BatchGet() {
	res3, err := redis.Ints(c.Do("MGet", 1, 2, 3))
	if err != nil {
		fmt.Println("get int failed,", err)
		return
	}
	fmt.Println(res3)
	res, err := redis.Strings(c.Do("MGet", "a", "b", "c"))
	if err != nil {
		fmt.Println("get string failed,", err)
		return
	}
	fmt.Println(res)
}

func Rtimeout() {
	_, err = c.Do("setex", "a", 2, "10")
	if err != nil {
		fmt.Println("set string failed", err)
		return
	}

	res2, err := redis.String(c.Do("Get", "a"))
	if err != nil {
		fmt.Println("get string failed,", err)
		return
	}
	fmt.Println(res2)
	time.Sleep(5 * time.Second)
	fmt.Println("5秒后")
	res2, err = redis.String(c.Do("Get", "a"))
	if err != nil {
		fmt.Println("get string failed,", err)
		return
	}

}

func Rlist() {
	// 从左边放入元素
	_, err = c.Do("lpush", "NBAplayer", "Jordon", "Kobe", "Lebron")
	if err != nil {
		fmt.Println("push element failed")
		return
	}
	// 取出所有元素
	res4, err := redis.Strings(c.Do("lrange", "NBAplayer", "0", "-1"))
	if err != nil {
		fmt.Println("get element failed")
		return
	}
	for _, v := range res4 {
		fmt.Print(v + " ")
	}
}

func Rhash() {
	// 单个插入
	_, err = c.Do("HSet", "DSB", "name", "SB")
	if err != nil {
		fmt.Println("hset failed")
		return
	}
	_, err = c.Do("HmSet", "DSB", "age", "18", "addr", "usa")
	if err != nil {
		fmt.Println("hmset failed")
		return
	}
	// 单个获取
	res5, err := redis.String(c.Do("Hget", "DSB", "name"))
	if err != nil {
		fmt.Println("hget element failed")
		return
	}
	fmt.Println(res5)
	// 批量获取
	res6, err := redis.Strings(c.Do("HmGet", "DSB", "name", "age", "addr"))
	if err != nil {
		fmt.Println("hmget failed")
		return
	}
	fmt.Println(res6)
}

func Rset() {
	_, err = c.Do("sadd", "qin", "18", "male", "handsome")
	if err != nil {
		fmt.Println("sadd failed, err:", err)
		return
	}
	_, err = c.Do("sadd", "yi", "19", "male", "handsome")
	if err != nil {
		fmt.Println("sadd failed, err:", err)
		return
	}
	res, err := redis.Strings(c.Do("smembers", "qin"))
	if err != nil {
		fmt.Println("smembers failed, err:", err)
		return
	}
	fmt.Println(res)
	res, err = redis.Strings(c.Do("smembers", "yi"))
	if err != nil {
		fmt.Println("smembers failed, err:", err)
		return
	}
	fmt.Println(res)
	res, err = redis.Strings(c.Do("sinter", "qin", "yi"))
	if err != nil {
		fmt.Println("sinter failed, err:", err)
		return
	}
	fmt.Println(res)
	res, err = redis.Strings(c.Do("sdiff", "qin", "yi"))
	if err != nil {
		fmt.Println("sinter failed, err:", err)
		return
	}
	fmt.Println(res)
	res, err = redis.Strings(c.Do("sdiff", "yi", "qin"))
	if err != nil {
		fmt.Println("sinter failed, err:", err)
		return
	}
	fmt.Println(res)
	res1, err := redis.Bool(c.Do("sismember", "qin", "20"))
	if err != nil {
		fmt.Println("sismember failed, err:", err)
	}
	fmt.Println(res1)
	res1, err = redis.Bool(c.Do("sismember", "qin", "18"))
	if err != nil {
		fmt.Println("sismember failed, err:", err)
	}
	fmt.Println(res1)
	res2, err := redis.Int(c.Do("scard", "qin"))
	if err != nil {
		fmt.Println("scard failed, err:", err)
	}
	fmt.Println(res2)
	res3, err := redis.Int(c.Do("srem", "qin", 18))
	if err != nil {
		fmt.Println("srem failed, err:", err)
	}
	fmt.Println(res3)
	res, err = redis.Strings(c.Do("smembers", "qin"))
	if err != nil {
		fmt.Println("smembers failed, err:", err)
		return
	}
	fmt.Println(res)
}

func Rzset() {
	res, err := c.Do("zadd", "RankOfNBAplayer", 1, "Jordon", 2, "LeBron", 3, "kobe", 4, "YaoMing", 5, "JRSmith")
	if err != nil {
		fmt.Println("zadd failed, err:", err)
		return
	}
	fmt.Println("zadd ", res)
	resq, err := redis.Strings(c.Do("zrange", "RankOfNBAplayer", 0, -1))
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("zrange", resq)
	res, err = c.Do("zrem", "RankOfNBAplayer", "JRSmith")
	if err != nil {
		fmt.Println("zrem failed, err:", err)
		return
	}
	fmt.Println(res)
	resq, err = redis.Strings(c.Do("zrange", "RankOfNBAplayer", 0, -1))
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("zrange", resq)
	res, err = redis.Int(c.Do("zscore", "RankOfNBAplayer", "Jordon"))
	if err != nil {
		fmt.Println("zscore failed, err:", err)
		return
	}
	fmt.Println("zscore ", res)
	res, err = redis.Int(c.Do("zrank", "RankOfNBAplayer", "Jordon"))
	if err != nil {
		fmt.Println("zrank failed, err:", err)
		return
	}
	fmt.Println("zrank ", res)
	res, err = redis.Int(c.Do("zcard", "RankOfNBAplayer"))
	if err != nil {
		fmt.Println("zcard failed, err:", err)
		return
	}
	fmt.Println("zcard ", res)
	res, err = redis.Int(c.Do("zincrby", "RankOfNBAplayer", 10, "JRSmith"))
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("zincrby ", res)
	resq, err = redis.Strings(c.Do("zrange", "RankOfNBAplayer", 0, -1))
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("zrange", resq)
	resq, err = redis.Strings(c.Do("zrangebyscore", "RankOfNBAplayer", 0, 3))
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("zrangebyscore", resq)
}

// 创建redis连接池
var pool *redis.Pool

func Rinit() {
	pool = &redis.Pool{
		MaxIdle:     16,  // 连接数量
		MaxActive:   0,   // 最大连接数量，不确定用0
		IdleTimeout: 300, //连接关闭时间
		Dial: func() (redis.Conn, error) { //需要连接的数据库
			server := "127.0.0.1:6379"
			option := redis.DialPassword("123456")
			c, err = redis.Dial("tcp", server, option)
			if err != nil {
				log.Println("connect server failed:", err)
			}
			return c, err
		},
	}
}
func Rpool() {
	// 从连接池中获取连接
	client := pool.Get()
	// 把连接放回连接池
	defer client.Close()
	_, err := client.Do("Set", "CSDN", "good")
	if err != nil {
		fmt.Println("set string failed, err:", err)
		return
	}

	res, err := redis.String(client.Do("Get", "CSDN"))
	if err != nil {
		fmt.Println("set string failed, err:", err)
		return
	}
	fmt.Println(res)
	// 关闭连接池
	pool.Close()
}

func Rtransaction() {
	// 开启事务
	c.Send("MULTI")
	c.Send("set", "k1", "v1")
	c.Send("set", "k2", "v2")
	// 执行事务
	res, err := c.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("exec ", res)
	v, err := redis.String(c.Do("get", "k1"))
	fmt.Println("get k1", v)
	v, err = redis.String(c.Do("get", "k2"))
	fmt.Println("get k2", v)
}



