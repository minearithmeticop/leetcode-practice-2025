package problem1480

// RunningSum calculates the running sum of a 1D array
func RunningSum(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    result := make([]int, len(nums))
    result[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        result[i] = result[i-1] + nums[i]
    }
    
    return result
}
