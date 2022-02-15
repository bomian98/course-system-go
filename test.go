package main

import (
	"fmt"
	"reflect"
	"time"
)

//func main101() {
//
//	fmt.Println("元素的个数为", len(ch), "缓存能力为", cap(ch))
//	ch <- 123
//	t := <-ch
//	fmt.Println("读取到数据", t)
//	ch <- 123
//	fmt.Println("元素的个数为", len(ch), "缓存能力为", cap(ch))
//	ch <- 123
//	fmt.Println("元素的个数为", len(ch), "缓存能力为", cap(ch))
//	ch <- 123
//	fmt.Println("元素的个数为", len(ch), "缓存能力为", cap(ch))
//	ch <- 123
//	fmt.Println("元素的个数为", len(ch), "缓存能力为", cap(ch))
//}

func printA(ch1, ch2 chan string) {
	for i := 0; i < 100; i++ {
		<-ch2
		fmt.Println(i, "A")
		ch1 <- "PrintA"
	}
}

func printB(ch1, ch2, ch3 chan string) {
	ch2 <- "begin"
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Println(i, "B")
		if i != 99 {
			ch2 <- "print B"
		} else {
			ch3 <- "end"
		}
	}
}

func ch1out() {
	for {
		fmt.Println(<-ch1)
		time.Sleep(2 * time.Second)
	}
}

var ch1 = make(chan string, 3)

func main2() {
	age := 3
	fmt.Println(&age)
	fmt.Println(*(&age))
	fmt.Println(reflect.TypeOf(&age))
}
