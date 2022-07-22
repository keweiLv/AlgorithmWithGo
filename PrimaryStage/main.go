package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

// 存在重复元素Ⅱ
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	set := map[int]bool{}
	for i := 0; i < n; i++ {
		if i > k {
			set[nums[i-k-1]] = false
		}
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}
	return false
}

// 石子游戏Ⅸ
func stoneGameIX(stones []int) bool {
	cnts := make([]int, 3)
	for _, stone := range stones {
		cnts[stone%3]++
	}
	if cnts[0]%2 == 0 {
		return !(cnts[1] == 0 || cnts[2] == 0)
	} else {
		return !(math.Abs(float64(cnts[1]-cnts[2])) <= 2)
	}
}

// 句子中的有效单词数
func countValidWords(sentence string) (ans int) {
	for _, s := range strings.Fields(sentence) {
		if valid(s) {
			ans++
		}
	}
	return
}
func valid(s string) bool {
	hasHyphene := false
	for i, ch := range s {
		if unicode.IsDigit(ch) || strings.ContainsRune("!.,", ch) && i < len(s)-1 {
			return false
		}
		if ch == '-' {
			if hasHyphene || i == 0 || i == len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
				return false
			}
			hasHyphene = true
		}
	}
	return true
}

// 仅仅翻转字母
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !unicode.IsLetter(rune(s[left])) { // 判断左边是否扫描到字母
			left++
		}
		for right > left && !unicode.IsLetter(rune(s[right])) { // 判断右边是否扫描到字母
			right--
		}
		if left >= right {
			break
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}

//leetcode-682 棒球比赛
func calPoints(ops []string) (ans int) {
	points := []int{}
	for _, op := range ops {
		n := len(points)
		switch op[0] {
		case '+':
			ans += points[n-1] + points[n-2]
			points = append(points, points[n-1]+points[n-2])
		case 'D':
			ans += points[n-2] * 2
			points = append(points, 2*points[n-2])
		case 'C':
			ans -= points[n-1]
			points = points[:len(points)-1]
		default:
			pt, _ := strconv.Atoi(op)
			ans += pt
			points = append(points, pt)
		}
	}
	return
}

// 增键字符串匹配
func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	low, hign := 0, n
	for i, ch := range s {
		if ch == 'I' {
			perm[i] = low
			low++
		} else {
			perm[i] = hign
			hign--
		}
	}
	perm[n] = low
	return perm
}

// 搜索插入位置
func searchInsert(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len
	for left < right {
		mid := (right-left)>>1 + len
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 有效的回文
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left,right := 0,len(s)-1
	for left < right{
		for left < right && !isAlnum(s[left]){
			left++
		}
		for left < right && !isAlnum(s[right]){
			right--
		}
		if left < right{
			if s[left] != s[right]{
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isAlnum(ch byte) bool {
	return (ch  >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}