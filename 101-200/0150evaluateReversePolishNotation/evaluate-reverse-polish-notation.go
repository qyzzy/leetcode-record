package evaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	nums := []int{}
	for _, v := range tokens {
		switch v {
		case "+":
			{
				sum := nums[len(nums)-1] + nums[len(nums)-2]
				nums = nums[:len(nums)-1]
				nums[len(nums)-1] = sum
			}
		case "-":
			{
				sum := nums[len(nums)-2] - nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				nums[len(nums)-1] = sum
			}
		case "*":
			{
				sum := nums[len(nums)-1] * nums[len(nums)-2]
				nums = nums[:len(nums)-1]
				nums[len(nums)-1] = sum
			}
		case "/":
			{
				sum := nums[len(nums)-2] / nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				nums[len(nums)-1] = sum
			}
		default:
			{
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
		}
	}
	return nums[0]
}
