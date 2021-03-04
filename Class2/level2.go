package main

//写一个接收者函数,该接收者能够判断传入参数的类型，并作出不同的反应

import "fmt"

func Receivere(v interface{})  {
	switch v.(type) {
	case string:
		fmt.Println(v,"是个string")
		break
	case int:
		fmt.Println(v,"是个int")
		break
	case bool:
		fmt.Println(v,"是个bool")
	}
}

func main()  {
	Receivere(123)
	Receivere("ljy")
	Receivere(true)
}