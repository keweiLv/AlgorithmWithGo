package main

import (
	"math"
	"math/bits"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

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

// 寻找比目标字母大的最小字母,sort.Search:一般用于从一个已经排序的数组中找到某个值所对应的索引
func nextGreatestLetter(letters []byte, target byte) byte {
	return letters[sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})%len(letters)]
}

// 二进制表示中质数个计算置位
func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return ans
}
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 单词规律
func wordPattern(pattern string, s string) bool {
	word2ch := map[string]byte{}
	ch2word := map[byte]string{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}

// 分糖果
func distributeCandies(candyType []int) int {
	set := map[int]struct{}{}
	for _, t := range candyType {
		set[t] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}

// 最富有客户的资产总量
func maximumWealth(accounts [][]int) (ans int) {
	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

// 只出现一次的数字
// x & (-x) 可以获得ax最低的非0位
// nums 中的所有元素分成两类.其中一类包含所有二进制表示的第 ll 位为 00 的数，另一类包含所有二进制表示的第 ll 位为 11 的数
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	mask := xorSum & (-xorSum)
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&mask > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}

// 最小差值
func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	ans := maxNum - minNum - 2*k
	if ans > 0 {
		return ans
	}
	return 0
}

// 删列造序
func minDeletionSize(strs []string) (ans int) {
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				ans++
				break
			}
		}
	}
	return ans
}

// leetcode-682 棒球比赛
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
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isAlnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// 数组序号转换
func arrayRankTransform(arr []int) []int {
	copy := append([]int{}, arr...)
	sort.Ints(copy)
	ranks := map[int]int{}
	for _, v := range copy {
		if _, ok := ranks[v]; !ok {
			ranks[v] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	minValue := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}

// 数组中两元素的最大乘积
func maxProduct(nums []int) int {
	p, q := 0, 0
	for _, num := range nums {
		if num > q {
			p, q = q, num
		} else if num > p {
			p = num
		}
	}
	return (p - 1) * (q - 1)
}

// 数组元素积的符号
func arraySign(nums []int) int {
	ans := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			ans *= -1
		}
	}
	return ans
}

// 至少在两个数组中出现的值
func twoOutOfThree(nums1, nums2, nums3 []int) (ans []int) {
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	for x, m := range mask {
		if m&(m-1) > 0 {
			ans = append(ans, x)
		}
	}
	return
}

// 按摩师
func massage(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i%2] = Max(dp[(i-1)%2], dp[(i-2)%2]+nums[i])
	}
	return dp[(l-1)%2]
}

func Max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

// 检查句子中的数字是否递增
func areNumbersAscending(s string) bool {
	pre := 0
	for _, t := range strings.Split(s, " ") {
		if t[0] <= '9' {
			cur, _ := strconv.Atoi(t)
			if pre >= cur {
				return false
			}
			pre = cur
		}
	}
	return true
}

// 统计各数字之和为偶数的整数个数
func countEven(num int) (ans int) {
	for i := 1; i <= num; i++ {
		s := 0
		for x := i; x > 0; x /= 10 {
			s += x % 10
		}
		if s%2 == 0 {
			ans++
		}
	}
	return
}

// 统计包含给定前缀的字符串
func prefixCount(words []string, pref string) (ans int) {
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			ans++
		}
	}
	return
}

// 易混淆数
func confusingNumber(n int) bool {
	m := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	origin := n
	reverse := 0
	for n > 0 {
		if _, ok := m[n%10]; !ok {
			return false
		}
		reverse = reverse*10 + m[n%10]
		n /= 10
	}
	return origin != reverse
}

// 检查相同字母间的距离
func checkDistances(s string, distance []int) bool {
	d := [26]int{}
	for i, c := range s {
		c -= 'a'
		if d[c] > 0 && i-d[c] != distance[c] {
			return false
		}
		d[c] = i + 1
	}
	return true
}

// 老鼠和奶酪
func miceAndCheese(reward1 []int, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x
		reward1[i] -= x
	}
	sort.Ints(reward1)
	n := len(reward1)
	for i := 0; i < k; i++ {
		ans += reward1[n-i-1]
	}
	return
}

// 有序数组的平方
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}
