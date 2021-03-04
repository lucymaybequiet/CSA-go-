package main

import (
	"fmt"
	"sync"
)

var (//全局变量a记录素数的个数
	a = 0
	lock sync.Mutex
	wg sync.WaitGroup
)

func sushu(n int)  {
	defer wg.Done()

	for i := 2; i < n; i++ {
		//判断是否为素数
		if n%i==0 {
			return
		}
	}
	fmt.Println(n)
	//加锁
	lock.Lock()
	a++
	lock.Unlock()
}

func main()  {
	for i := 2; i < 50000; i++ {
		wg.Add(1)
		go sushu(i)
	}
	wg.Wait()
	fmt.Println("素数的数量为",a)
}
