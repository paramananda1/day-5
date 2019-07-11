package main

import (
	"fmt"
)

func ChannelTest() {
	var ic chan string
	ic = make(chan string,1)
	ic<-"qeqew"

	fmt.Println("Channel: " ,<-ic )

	var MAP map[int]string
	str := []string{"m","n","b"}
	MAP = make(map[int]string)
	for i,s := range str{
		MAP[i] = s

	}
	fmt.Println("Map : ",MAP)


}
