package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"myfirst/src/bean"
	service1 "myfirst/src/service"
	"strconv"

	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	//var r, r2 float64
	//r = 2
	//r2 = geng(r)
	//println(r2)
	check()
}

func t1() {
	fmt.Println("service")
}

func t2() {
	var sum int
	sum = service1.T3(10)
	println(sum)
}

func t3() {
	//service1.T4()
	//var k int
	//k = service1.T5(10,0)
	//fmt.Println(k)
	service1.T12()
}

func t4() {
	var num = flag.Int("n", 2, "cpu core")
	flag.Parse()
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(*num)
	fmt.Println("t4 begin")
	go service1.Longwait()
	go service1.Shortwait()
	fmt.Println("t4 end")
	time.Sleep(time.Minute)
}

func t5() {
	go service1.Say("t")
	service1.Say("a")
	time.Sleep(10 + time.Second)
}

func t6() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		s := "thread"
		s = fmt.Sprintf("%s%d", s, i)
		go service1.Counter(lock, s)
	}
	time.Sleep(time.Minute)
	//for {
	//	lock.Lock()
	//	c := service1.Count
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if c >= 10 {
	//		break
	//	}
	//}
}

func t7() {
	num := 5
	chs := make([]chan int, num)
	for i := 0; i < num; i++ {
		chs[i] = make(chan int)
		go service1.Channel(chs[i])
	}
	for {
		select {
		case <-chs[0]:
			println(0)
			break
		case <-chs[1]:
			println(1)
			break
		case <-chs[2]:
			println(2)
		case <-chs[3]:
			println(3)
		case <-chs[4]:
			println(4)
		}
	}
	//for _, ch := range chs {
	//	println(<-ch)
	//}
}

func t8() {
	var ch chan int
	service1.T24(ch)
	var v int
	v = <-ch
	println(v)
}

func t9() {
	service1.Doall()
}

func t10() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, " method $s ip is ", os.Args[0])
		os.Exit(1)
	}
	service1.Tip(os.Args[1])
}

func t11() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, " param $s  is ", os.Args[0])
		os.Exit(1)
	}
	service1.Tping(os.Args[1])
}

func insert() {
	var user bean.User
	user.Password = "123456"
	user.UserName = "jim"
	service1.InitDB()
	//service1.InsertUser(user)
	//service1.DeleteUser(44)
	user.Id = 1
	//service1.UpdateUser(user)
	var u bean.User
	u = service1.SelectUserById(45)
	fmt.Printf("%v\n", u)
	users := service1.SelectAllUser()
	fmt.Printf("%v", users)
}

func redis() {
	service1.Redis()
	//service1.RedisGetSet()
	//service1.BatchGet()
	//service1.Rtimeout()
	//service1.Rlist()
	//service1.Rzset()
	//service1.Rinit()
	//service1.Rpool()
	service1.Rtransaction()
}

func mongo() {
	//service1.MoConnecToDB();
	service1.MoRemoveFromMgo()
}

func filewr() {
	//service1.Filew();
	service1.Filer()
}

//根号函数, in , float,  out , float,  level 0.99%
func geng(value float64) float64 {
	var level float64
	level = 0.001
	//2fengfa
	var tmp, begin, end, result, v float64

	tmp = value / 2
	begin = 0
	end = value
	for true {
		result = tmp * tmp
		v = result/value - 1
		if v < 0 {
			v = 0 - v
		}
		// >
		if v <= level {
			return tmp
		} else {
			if result > value {
				end = tmp
				tmp = (begin + end) / 2
			} else {
				begin = tmp
				tmp = (begin + end) / 2
			}
		}
	}
	return 0
}

func check() {
	var tmp float64
	var level float64
	level = 0.001
	var i int
	for i = 0; i < 1000; i++ {
		tmp = rands(0.1, 100.9)
		sqrt := math.Sqrt(tmp)
		f := geng(tmp)
		if math.Abs((sqrt - f)) < level {
			println(i , " ok")
		} else {
			println(i,"no ok")
		}
	}
}

func rands(min, max float32) float64 {
	max = max - min
	rand.Seed(time.Now().UnixNano()) //设置随机种子，使每次结果不一样
	res := Round2(float64(min+max*rand.Float32()), 2)
	fmt.Println(res)
	return res
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}
