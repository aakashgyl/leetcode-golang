var climbWays = make(map[int]int)

func climbStairs(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    }
    
    if ways, ok := climbWays[n]; ok {
        return ways
    }
    
    climbWays[n] = climbStairs(n-1) + climbStairs(n-2)
    return climbWays[n]
}