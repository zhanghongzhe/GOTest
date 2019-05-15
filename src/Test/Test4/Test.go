package Test4

import "fmt"

func Count(ch chan int)  {
	ch <- 1 //将一个数据写入（发送）至channel
	fmt.Println("Counting")
}
func Test()  {
	chs := make([]chan int, 10)
	for i := 0; i<10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<- ch //从channel中读取数据的语法是：value := <-ch
	}
}