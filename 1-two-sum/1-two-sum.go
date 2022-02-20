func twoSum(nums []int, target int) []int {
    valMap := make(map[int]int)
    for i, val := range nums {
        if j, ok := valMap[target-val]; ok {
            return []int{i,j}
        }
        valMap[val] = i
    }
    return []int{}
}