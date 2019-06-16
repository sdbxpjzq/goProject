package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func incCounter() {
	for count := 0; count < 2; count++ {
		// 安全的对counter +1
		// 强制同一时刻只能有一个 goroutine 运行并完成这个加法操作
		atomic.AddInt64(&counter, 1) // atmoic 包的 AddInt64 ,另外两个有用的原子函数是 LoadInt64 和 StoreInt64

		// 当前goroutine从线程退出, 并放回队列
		runtime.Gosched()
	}
	wg.Done()
}

// 消费者
func consumer(data chan int, done chan bool) {
	for value := range data { // 接收数据, 直到通道关闭
		fmt.Println(value)

	}
	done <- true
	close(done)
}

// 生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data) // 数据生产结束, 关闭通道
}

//
//type Test struct {
//	name string
//}
//
//func (t *Test) IntSum(a int, b int) (sum int) {
//	sum = a + b
//	return
//}
//
//func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
//	inputs := make([]reflect.Value, len(args))
//	for i, _ := range args {
//		inputs[i] = reflect.ValueOf(args[i])
//	}
//
//	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
//		return nil
//	} else {
//		return v.Call(inputs)
//	}
//
//}

type Inter interface {
	Ping()
	Pang()
}

type Test struct {
}

func (t *Test) Ping() {
	fmt.Println("ping")
}
func (t *Test) Pang() {
	fmt.Println("pang")
}

func typeFunc(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("sting", v)
	}
}

func testInt(seed ...int) {
	fmt.Println(len(seed))
	for key, value := range seed {
		fmt.Println(key, value)
	}
}

func main() {
	testInt()

	//typeFunc(10)
	//typeFunc("23")
	//os.Exit(1)
	//
	//a := &Test{}
	//var i interface{} = a
	//
	//comm,ok := i.(Inter)
	//if ok {
	//	comm.Ping()
	//	comm.Pang()
	//}
	//
	//
	//fmt.Println(ok)

	//v := reflect.ValueOf(Test{
	//	name: "z",
	//})
	//
	//fmt.Println(v.FieldByName("name"))           // 获取值
	//fmt.Println(v.FieldByName("name").IsValid()) // 检测是否存在该 字段
	//fmt.Println(v.FieldByName("age").IsValid())  // 检测是否存在该 字段
	//
	//sum := callReflect(&Test{}, "IntSum", 1,2)
	//
	//for key, value := range sum {
	//	fmt.Println(key, value)
	//}
	//
	//fmt.Println(sum[0].Int())

}
