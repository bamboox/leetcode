/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(soultion("wsawd1dad"))
	fmt.Println(soultion3("wsawd1dad"))

}

// O(n^3)
func soultion(str string) int {
	maxLength := 0
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if unique(str, i, j) {
				if maxLength < j-i {
					//j需要+1为才能遍历到
					maxLength = j - i
					fmt.Println(j, i)
				}
			}
		}
	}
	return maxLength
}

func unique(str string, start, end int) bool {
	strSet := make(map[string]int)
	for i := start; i < end; i++ {
		str := string(str[i])

		if _, ok := strSet[str]; ok {
			fmt.Println("|-")
			return false
		} else {
			strSet[str] = 1
			fmt.Print(str)
		}
	}
	fmt.Println("|")
	return true
}

// o(2n)
func soultion2(str string) int {
	maxLength, i, j := 0, 0, 0
	n := len(str)
	strSet := make(map[string]int)

	for i < n && j < n {
		if _, ok := strSet[string(str[j])]; !ok {
			strSet[string(str[j])] = 0
			if maxLength < j-i {
				maxLength = j - i
			}
			j++
		} else {
			delete(strSet, string(str[i]))
			i++
		}
	}
	return maxLength + 1
}

// o(n)
func soultion3(str string) int {
	maxLength := 0
	n := len(str)
	strMap := make(map[string]int)

	for i, j := 0, 0; j < n; j++ {
		if index, ok := strMap[string(str[j])]; ok {
			i = index
		}
		strMap[string(str[j])] = j + 1
		if maxLength < j-i {
			maxLength = j - i
			fmt.Println(i, j)
		}
	}
	return maxLength + 1
}

// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length(), ans = 0;
//         int[] index = new int[128]; // current index of character
//         // try to extend the range [i, j]
//         for (int j = 0, i = 0; j < n; j++) {
//             i = Math.max(index[s.charAt(j)], i);
//             ans = Math.max(ans, j - i + 1);
//             index[s.charAt(j)] = j + 1;
//         }
//         return ans;
//     }
// }
