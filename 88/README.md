# 88. Merge Sorted Array

## 题目
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

- The number of elements initialized in nums1 and nums2 are m and n respectively.
- You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

Example:

```
**Input:**
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

**Output:** [1,2,2,3,5,6]
```
## 思路
从尾部开始合并

## 注意

为什么不能 nums1 = nums2 ？
if m==0 {
		i := n-1
		for i >= 0{
			nums1[i] = nums2[i]
			i--
		}
	}