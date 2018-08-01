# 80. Remove Duplicates from Sorted Array II

## 题目
Follow up for "Remove Duplicates": What if duplicates are allowed at most twice?

For example, Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.




## 思路

      Slot
j
1   1   1   2   2   3
i

i遍历每个元素
j指向用来比对的值，与nums[i]不同时，更新j=i
找Slot，第一次重复3次的位置，用后续元素不断覆盖



## 注意
count初始值设为1