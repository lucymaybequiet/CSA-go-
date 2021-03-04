package main

import "fmt"

//level1：
//请参照上面的例子，声明鸽子，复读机，柠檬精，真香怪四种接口，然后让人类类型分别实现四种接口。

type dove interface {
	gugugu()
}

type repeater interface {
	repeat(string)
}

type lemon interface {
	suan()
}

type zhenXiangGuai interface {
	zhenxiang()
}

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) gugugu() {
	fmt.Println(p.name,"又鸽了")
}

func (p *person) repeat(word string)  {
	fmt.Println(word)
}

func (p *person) suan()  {
	fmt.Println("也不怎么样嘛")
}

func (p *person) zhenxiang()  {
	fmt.Println("真香")
}

func main()  {

	p := &person{name: "ljy" , age: 19 , gender: "female"}
	p.gugugu()
	p.repeat("针布戳")
	p.suan()
	p.zhenxiang()
}