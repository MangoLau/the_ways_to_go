package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello World!"
	intArr := []int64{10, 8, 20, 40, 15, 19, 9}
	var intSlice = &intArr
	str1, str2 := StringSlice(str, 3)
	fmt.Println(str1)
	fmt.Println(str2)

	s := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(s)

	fmt.Println(StrReverse(str))
	fmt.Println(StrReverseSingle(str))
	fmt.Println(string(StrDiffLastRune([]byte(str))))
	fmt.Println(BubleSort(intArr))
	fmt.Printf("intArr type: %T\n", intArr)
	fmt.Printf("intSlice type: %T\n", intSlice)
	fmt.Println(mapFunc(tenTimes, intArr))
}

func StringSlice(str string, i int) (str1 string, str2 string) {
	if i > utf8.RuneCountInString(str) {
		return str, ""
	}
	c := []int32(str)
	str1, str2 = string(c[:i]), string(c[i:])
	return
}

func StrReverse(str string) (ret string) {
	arr := []byte(str)
	l := len(arr)
	var retArr = make([]byte, l)
	for k,v := range arr {
		retArr[l - k - 1] = v
	}
	return string(retArr)
}

func StrReverseSingle(str string) (ret string) {
	arr := []byte(str)
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l - i - 1] = arr[l - i - 1], arr[i]
	}
	ret = string(arr)
	return
}

func StrDiffLastRune(arr []byte) (ret []byte) {
	for k,v := range arr {
		if k < 1 || arr[k - 1] != v {
			ret = append(ret, v)
		}
	}
	return ret
}

func BubleSort(arr []int64) (ret []int64) {
	l := len(arr)
	var i, j int
	for i = 0; i < l; i++ {
		for j = i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func mapFunc(fn func(n int64)(ret int64), arr []int64) (ret []int64) {
	for _, v := range arr {
		ret = append(ret, fn(v))
	}
	return ret
}

func tenTimes(num int64) (ret int64) {
	return num * 10
}