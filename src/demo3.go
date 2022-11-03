package main

import "fmt"

func test1() {
	var a []int
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("%T\n", a)
	fmt.Println(a1)
}

func test2() {
	var a = map[string]string{"name": "zhangsan", "age": "20"}
	fmt.Println(a)
	a = make(map[string]string)
	a["name"] = "tom"
	a["age"] = "20"
	fmt.Println(a)
	for k, v := range a {
		fmt.Println(k, v)
	}

}

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}

func test(name string, f func(string)) {
	f(name)
}

// 变量初始化->init()->main()
//func init() {
//	fmt.Println("init")
//}

//func main() {
//	fmt.Println("start")
//	defer fmt.Println("step1")
//	defer fmt.Println("step2")
//	defer fmt.Println("step3")
//	fmt.Println("end")
//
//	//test("zhangsan", sayHello)
//
//	//type f1 func(int, int) int
//	//var ff f1
//	//ff = max
//	//fmt.Println(ff(1, 2))
//	//
//	//ff = sum
//	//fmt.Println(ff(1, 2))
//}
