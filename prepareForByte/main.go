package prepareForByte

func main() {
}

// 打家劫舍
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pre := nums[0]
	cur := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		pre, cur = cur, max(pre+nums[i], cur)
	}
	return cur
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
