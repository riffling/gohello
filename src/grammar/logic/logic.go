package logic

import (
	"fmt"
)

func For_Array() {
	// for condition
	var i1 int = 3
	for true {
		fmt.Printf("for true; {break} i1=%d \n", i1)
		break
	}

	// for init; cond; exp
	var array1 [3]int
	array1 = [3]int{10, 20, 30}

	for i := 0; i < len(array1); i++ {
		fmt.Printf("for init; cond; exp {}: idx:%d = val:%d \n", i, array1[i])
	}

	//for array
	var array2 [3]int = [3]int{100, 200, 300}
	for idx, val := range array2 {
		fmt.Printf("for idx, val=range array {}: idx:%d = val:%d  \n", idx, val)
	}

	array3 := [...]string{"red", "blue", 7: "eighth element"} //set value for eighth element
	for idx, val := range array3 {
		fmt.Printf("for idx, val=range array {}: idx:%d = val:%s  \n", idx, val)
	}

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int
	var xy int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	xy = 78
	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
		fmt.Printf("x=%d", xy)
	}

}
