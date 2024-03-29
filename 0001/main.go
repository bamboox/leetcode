/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	nums := []int{3, 4, 6, 7, 1, 2}
	target := 9
	fmt.Println(twoSum(nums, target))
}

// o(n)
func twoSum(nums []int, target int) []interface{} {
	hashMap := make(map[int]int, len(nums))
	for i, n := range nums {
		hashMap[n] = i
	}
	result := make([]interface{}, 0)
	resultFilter := make(map[int]int, len(nums))
	for j, n := range nums {
		if hIndex, ok := hashMap[target-n]; ok {
			//filter 0 ,2 ; 2, 0
			if _, ok := resultFilter[j+hIndex]; !ok {
				resultFilter[j+hIndex] = 0
				fmt.Println(j, hIndex)
				result = append(result, []int{j, hIndex})
			}
		}
	}
	return result
}
