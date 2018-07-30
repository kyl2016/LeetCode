package main

import "fmt"

func main(){
	var a1 [] int = []int{0,1,4,0,2,0}

	result := RemoveZeros(a1)

	fmt.Println(a1)
	fmt.Println(result)
}

func RemoveZeros (a []int) []int {
	var temp []int = make([]int, len(a))
	var j int = 0;
	for i := 0;i<len(a);i++{
		if a[i] != 0{
			temp[j] = a[i]
			j++
		}
	}

	for ;j<len(a);j++{
		temp[j] = 0
	}

	return temp
}