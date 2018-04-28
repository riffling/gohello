package main

import (
	"fmt"

	"grammar/hello"
)

const (
	i = 1 << iota //1
	j = 3 << iota //6
	k             //12
	l             //24
)

var i1 int = 12
var x int

func testMap() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

}

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var array1 [2]int
	var array2 []int = []int{4, 2, 3}
	fmt.Println("i1=", i1)

	fmt.Println(array1)
	fmt.Println(array2)

	hello.HelloWorld()
	TestPrimitive()

}
