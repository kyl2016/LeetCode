package main

import "fmt"

func main(){
	nums1 := []int{1,3,4,6,0,0}
	nums2 := []int{2,5}

	merge(nums1, 0, nums2, 2)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m==0 {
		i := n-1
		for i >= 0{
			nums1[i] = nums2[i]
			i--
		}
	}

	if n == 0 {
		return
	}

	i,j := m-1,n-1
	k := m+n-1

	for i>=0 && j>=0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		}else{
			nums1[k] = nums2[j]
			j--
		}

		k--
	}

	for i >= 0{
		nums1[k] = nums1[i]
		k--
		i--
	}

	for j >= 0{
		nums1[k] = nums2[j]
		k--
		j--
	}
}