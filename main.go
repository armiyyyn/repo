func twoSum(nums []int, target int) []int {

	var res []int
	for i := 0; i < len(nums); i++ {
		for f := 0; f < len(nums); f++ {
			if i != f && nums[i]+nums[f] == target {
				res = append(res, i)
				break
			}
		}
	}
	return res
}