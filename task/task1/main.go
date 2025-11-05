package main

import (
	"fmt"
	"os"
	"strconv"
)

var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
var nums1 = []int{2, 7, 11, 15}
var strs = []string{"flower", "flow", "flight"}
var strs1 = []string{"abab", "aba", ""}
var str = "()[]{}"
var str1 = "(])"
var strMap = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

var intervals = [][]int{{2, 3}, {4, 5}, {6, 7}, {1, 10}}

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(os.Args)
	var res = singleNumber(nums)
	//res := plusOne(nums1)
	//res := removeDuplicates(nums)
	// res := longestCommonPrefix(strs1)
	// res := isValid(str1)
	//res := merge(intervals)
	//res := twoSum(nums1, 9)
	//res := isPalindrome(1)
	fmt.Println(res)
	// for s := range nums {
	// 	println(nums[s])
	// }
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	//var temp = [1510]int{}
	rcd := map[int]int{}
	for _, value := range nums {
		switch rcd[value] {
		case 0:
			rcd[value] = 1
		case 1:
			rcd[value] += 1
		}
	}
	var ret int
	for i, v := range rcd {
		if v == 1 {
			ret = i
		}
		
	}

	return ret

}

// 66. 加一
func plusOne(digits []int) []int {
	var res []int
	l := len(digits)
	var t = digits[l-1] + 1
	if l > 1 {
		for i := len(digits) - 2; i >= 0; i-- {
			if i != 0 {
				if t == 10 {
					t = digits[i] + 1 //当前位置加1
					digits[i+1] = 0   //后一位变为0
					continue
				} else {
					digits[i+1] = digits[i+1] + 1
					res = digits
					break
				}
			} else {
				if t == 10 {
					digits[0] = digits[0] + 1
					digits[i+1] = 0
					if digits[0] == 10 {
						s := []int{1}
						digits[0] = 0
						//var newDig  [lens]int
						for j := 0; j < len(digits); j++ {
							s = append(s, digits[j])
						}
						res = s
					} else {
						res = digits
					}
				} else {
					digits[i+1] = digits[i+1] + 1
					res = digits
				}

			}
		}
	} else {
		if t != 10 {
			digits[0] = digits[0] + 1
			res = digits

		} else {
			res = []int{1, 0}
		}
	}

	return res
}

// 26. 删除有序数组中的重复项
// 不相同的整数添加到切片中，最后使用copy到原数组
func removeDuplicates(nums []int) int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		var t int
		if len(res) > 0 {
			t = res[len(res)-1]
		} else {
			t = 0
		}

		if t == nums[i] {
			if t == 0 && i == 0 {
				res = append(res, t)
			}
			continue
		} else {
			res = append(res, nums[i])
			continue
		}
	}
	copy(nums, res)

	return len(res)
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var s = strs[0]
	var slice = []byte{}

	for i := range s {
		t := s[i] //第一个字符串的第一个字符
		var b bool = false
		//以数组第一个字符串为基准，和数组中其他字符串的每个字符顺序比较
		for j := 1; j < len(strs); j++ {

			if len(strs[j]) < (i + 1) {
				b = false
				break
			}

			if t == strs[j][i] {
				b = true
			} else {
				b = false
				break
			}
		}
		if b {
			slice = append(slice, t)
		} else {
			break
		}
	}

	return string(slice)
}

// 20. 有效的括号
func isValid(s string) bool {
	var bt = []string{}
	var bl bool = false
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			t := string(s[i])
			//查询map中是否由对应的值
			v, ok := strMap[t]
			if ok {
				if len(bt) > 0 {
					//若找到对应的值，与切片最后一位比较，相等：出栈，不相等：入栈
					b1 := bt[len(bt)-1]
					if b1 == v {
						// l := len(bt)
						// if l == 1 {

						// }else{
						bt = bt[:len(bt)-1]
						// }

					} else {
						bt = append(bt, t)
						break
					}
				} else { //空栈返回false
					bt = append(bt, t)
					break
				}
			} else {
				bt = append(bt, t)
			}
		}
		if len(bt) > 0 {
			bl = false
		} else {
			bl = true
		}
	} else {
		bl = true
	}

	return bl
}

// 56. 合并区间
// intervals = [[1,3],[2,6],[8,10],[15,18]]
func merge(intervals [][]int) [][]int {
	var res = [][]int{}
	//var t []int = intervals[0]//= [...]int{0,0}
	if len(intervals) > 1 { //数组>1
		for j := 0; j < len(intervals); j++ {
			for i := 0; i < len(intervals)-1; i++ {
				t := intervals[i]
				t1 := intervals[i+1]
				if t[0] > t1[0] {
					intervals[i] = t1
					intervals[i+1] = t
				}
			}
		}
		var t []int
		//从0开始判断前后数组大小
		for i := 0; i < len(intervals)-1; i++ {
			if len(res) > 0 {
				t = res[len(res)-1]
			} else {
				t = intervals[i]
			}
			t1 := intervals[i+1]
			//前后数组有重合时，合并并放入Slice中，之后从切片中查找数据
			if t[1] >= t1[0] && t[1] < t1[1] {
				//[[1,4],[2,3]]
				r := []int{t[0], t1[1]}
				//intervals[i+1] = r
				if len(res) > 0 {
					if t[0] == res[len(res)-1][0] && t[1] == res[len(res)-1][1] {
						res = res[:len(res)-1]
					}

				}
				//else {
				res = append(res, r)
				continue
				//}

			} else if t[1] >= t1[0] && t[1] >= t1[1] {
				if t[1] >= t1[0] {
					r := []int{t[0], t[1]}
					//intervals[i+1] = r
					if len(res) > 0 {
						if t[0] == res[len(res)-1][0] && t[1] == res[len(res)-1][1] {
							continue
						}
					}
					res = append(res, r)
					continue
				}
			} else {
				//切片为空时添加第一个
				if len(res) == 0 {
					//intervals[i + 1] = r
					res = append(res, t)
					res = append(res, t1)
					continue
				}
				res = append(res, t1)
				continue

			}
		}
	} else {
		res = intervals
	}

	return res
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	mapTarget := map[string][]string{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]

		v, ok := mapTarget[strconv.Itoa(nums[i])]
		if ok {
			vi, _ := strconv.Atoi(v[1])
			res = append(res, vi)
			res = append(res, i)
			break
		} else {
			mapTarget[strconv.Itoa(t)] = []string{strconv.Itoa(nums[i]), strconv.Itoa(i)}
			continue
		}
	}

	return res
}

// 9. 回文数
func isPalindrome(x int) bool {
	res := false
	if x < 0 {
		return res
	}
	s := strconv.Itoa(x)
	l := len(s)
	if l > 1 {
		for i := 0; i < l/2; i++ {
			if s[i] == s[l-1-i] {
				res = true
			} else {
				res = false
				break
			}
		}
	} else {
		res = true
	}

	return res
}

