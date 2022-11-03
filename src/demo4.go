package main

import "fmt"

//func main() {
//	/*pointer()
//	array()
//	typeDefine()*/
//	testp()
//	var per = Person{
//		1, 2, "qqq", "111",
//	}
//	per.sleep()
//}

func testp() {
	type Person struct {
		id, age     int
		name, email string
	}
	var person = Person{
		id:    0,
		age:   12,
		name:  "tom",
		email: "111@qq.com",
	}
	var p_person = new(Person)
	p_person = &person
	fmt.Println(person)
	fmt.Println(p_person)
	fmt.Println(*p_person)
}

type Person struct {
	id, age     int
	name, email string
}

func (per Person) sleep() {
	fmt.Println("sleep")
}

func typeDefine() {
	type MyInt int
	var a MyInt = 10
	fmt.Printf("%T\n", a)
}

func pointer() {
	var ip *int
	fmt.Printf("ip:%v\n", ip)
	var i int = 100
	ip = &i
	fmt.Printf("ip:%v\n", ip)
	fmt.Printf("ip:%v\n", *ip)
}

func array() {
	a := [3]int{1, 2, 3}
	var pa [3]*int
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Println(pa)
}
