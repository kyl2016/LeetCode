package main

import "fmt"

func main() {
	var a []int = []int{1,1,1,1,2,3,3,3,4,5,5,5,6,6,7}
	a = []int{1,1,1,2,2,3}

	fmt.Println(a)

	len := RemoveDuplicates(a)
	fmt.Println(len)
	fmt.Println(a)
}

func RemoveDuplicates(nums []int) int {

	j,sameCount,count := 0,1,1
	slot := -1

	for i:=1;i<len(nums);i++{
		if nums[i] != nums[j]{
			j=i
			count++
			sameCount = 1

			if slot >-1 {
				if nums[slot] != nums[i] {
				nums[slot] = nums[i]
				}
				slot++
			}
		} else {
		if sameCount < 2 {
			count++

			if slot >-1{
				if nums[slot] != nums[i]{
					nums[slot] = nums[i]
				}
				slot++
			}
			}else {
				if sameCount == 2 && slot < 0 {
				slot = i
				}
			}

			sameCount++
		}
	}

	return count
}