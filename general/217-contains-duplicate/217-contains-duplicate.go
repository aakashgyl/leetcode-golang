func containsDuplicate(nums []int) bool {
    mymap := make(map[int]bool)
    for _, val := range nums{
        if _, ok := mymap[val]; !ok {
            mymap[val] = true
        } else {
            return true    
        }
    }
    return false
}