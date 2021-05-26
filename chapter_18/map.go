package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	map1["d"] = 4
	map1["e"] = 5
	// （1）如何使用for或者for-range遍历一个映射：
	for k, v := range map1 {
		fmt.Printf("map for-range. k: %s, v: %d\n", k, v)
	}

	map2 := map[string]int{"one":1, "two":2, "three":3, "four":4, "five": 5}
	// （2）如何在一个映射中检测键key1是否存在：
	key := "five"
	if _, isPresent := map2[key]; !isPresent {
		fmt.Printf("map2[%s] is not existed\n", key)
	} else {
		fmt.Printf("map2[%s] is exist, the value is %d\n", key, map2[key])
	}

	// （3）如何在映射中删除一个键：
	delete(map2, "four")
	fmt.Println(map2)
}