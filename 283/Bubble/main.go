package main 

import "fmt"

func main(){
var a []int = []int{1,0,2,0,56,1,0,1,0}

fmt.Println(a)

r := RemoveZeros(a)
fmt.Println(r)
}

func RemoveZeros(a []int) []int{
	i, j := 0,0

	for i<len(a){
if a[i] != 0{
	a[j] = a[i]
	j++
}
		i++
	}

	for j<len(a){
		a[j] = 0
		j++
	}

	return a
}