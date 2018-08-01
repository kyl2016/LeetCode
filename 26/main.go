package main

import "fmt"

func main() {
	var a []int = []int{1,1,2,3,3,4,5,5,5,6,6}

	fmt.Println(a)

	len := RemoveDuplicates(a)
	fmt.Println(len)
	fmt.Println(a)
}

func RemoveDuplicates(a []int) int {

	if len(a) == 0{
		return 0
	}

	j, k, count := 1,0,1

	for i := 1; i<len(a);i++{
		if a[i] != a[k]{
			count++
			k = i
			a[j] = a[i] // 将不重复的元素放到前面
			j++
		}
	}

	a = a[0:1]

	return count
}