package main

import "fmt"

type Str struct {
	string
}

type Int struct {
	int64
}

type Stringer interface {
	String() string
}

type Integer interface {
	Double() int64
}

func (str *Str) String() string {
	return str.string
}

func (v *Int) Double() int64 {
	return v.int64<<1
}

func main() {
	var strIntf Stringer
	s := new(Str)
	s.string = "s"
	strIntf = s
	// （1）如何检测一个值v是否实现了接口Stringer
	// strIntf 必须是一个接口变量
	if v, ok := strIntf.(Stringer); ok {
		fmt.Printf("implememts String(): %v\n", v)
	} else {
		fmt.Println("not implememts String()")
	}

	var intIntf Integer
	num := new(Int)
	num.int64 = 10
	intIntf = num
	if v, ok := intIntf.(Stringer); ok {
		fmt.Printf("implememts String(): %v\n", v)
	} else {
		fmt.Println("not implememts String()")
	}
	if v, ok := intIntf.(Integer); ok {
		fmt.Printf("implememts Integer(). double %d is %d\n", v, intIntf.Double())
	} else {
		fmt.Println("not implememts Integer()")
	}

	classifier(1, "str", true)
}

// （2）如何使用接口实现一个类型分类函数：
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}