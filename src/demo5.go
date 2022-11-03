package main

import (
	"fmt"
)

type USB interface {
	read()
	write()
}

type Computer struct {
}

type Mobile struct {
}

func (c Computer) read() {
	fmt.Println("computer read")
}
func (c Computer) write() {
	fmt.Println("computer write")
}
func (m Mobile) read() {
	fmt.Println("mobile read")
}
func (m Mobile) write() {
	fmt.Println("mobile write")
}

//func main() {
//	var com Computer
//	com.read()
//	com.write()
//	var mob Mobile
//	mob.read()
//	mob.write()
//}
