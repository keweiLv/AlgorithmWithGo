package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("5555")
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

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
