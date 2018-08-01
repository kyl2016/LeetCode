package main

import "fmt"

func main(){
	nums := []int{3,2,2,3,1,3,4,11,2,3,1,1,3}
	fmt.Println(nums)

	len := Remove(nums, 3)
	fmt.Println(nums)
	fmt.Println(len)
}

func Remove(nums []int, k int) int{

	if len(nums) == 0{
		return 0
	}

	count := 0
	j:= 0
	for i:=0;i<len(nums);i++{
		if nums[i] != k {
			if i != j{
				nums[j] = nums[i]
				}
			j++
			count++
		}
	}

	return count
}