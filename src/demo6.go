package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func show(msg int) {
	defer wg.Done()
	fmt.Println(msg)
}
func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a:%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b:%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

//func main() {
//	println(runtime.NumCPU())
//	runtime.GOMAXPROCS(32)
//	go a()
//	go b()
//	time.Sleep(time.Second)
//
//	/*	//waitgroup
//		for i := 0; i < 10; i++ {
//			go show(i)
//			wg.Add(1)
//		}
//		wg.Wait()
//		runtime.Gosched() //让出当前cpu时间
//		fmt.Println("end")*/
//
//	/*	//通道
//		goroutine := make(chan string, 5)
//		goroutine <- "hello" //值放入通道
//		data := <-goroutine  //通道中的值传入
//		fmt.Printf("%v", data)*/
//
//	/*	//协程
//		go show("test1")
//		show("test2")
//		time.Sleep(time.Millisecond * 200)*/
//}
