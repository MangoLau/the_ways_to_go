package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// （1）如何修改字符串中的一个字符：
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	str2 := string(c)
	fmt.Println("alter str:", str2)

	// （2）如何获取字符串的子串：
	substr := str[1:2]
	fmt.Println("substr:", substr)

	// （3）如何使用for或者for-range遍历一个字符串：
	// gives only the bytes:
	for i := 0; i < len(str); i++ {
		fmt.Printf("for: val: %v\n", str[i])
	}
	// gives the Unicode characters:
	for ix, ch := range str {
		fmt.Printf("for-range: id: %d, val: %v\n", ix, ch)
	}

	// （4）如何获取一个字符串的字节数：len(str)
	// 如何获取一个字符串的字符数：
	// 最快速
	length := utf8.RuneCountInString(str);
	fmt.Printf("RuneCountInString: %d\n", length)
	length = len([]rune(str))
	fmt.Printf("len([]rune(str)): %d\n", length)

	// （5）如何连接字符串：
	// 最快速： with a bytes.Buffer
	var b bytes.Buffer
	b.WriteString(str)
	b.WriteString("1243")
	fmt.Println(b.String())

	// 把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。
	n := 3
	b1, b2 := c[:n], c[n:]
	fmt.Printf("pre:%v\nafter:%v\n", b1, b2)
}
