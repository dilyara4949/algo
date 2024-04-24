func trap(h []int) int {
    leftMax, rightMax := make([]int, len(h)), make([]int, len(h))
    if len(h) > 0 {
        leftMax[0] = 0
        rightMax[len(h)-1] = 0
    }
    for i := 1; i < len(h); i++ {
        leftMax[i] = max(leftMax[i-1], h[i-1])
        rightMax[len(h)-1-i] = max(rightMax[len(h)-1-i+1], h[len(h)-1-i+1])
    }
    res := 0
    for i := 1; i < len(h)-1; i++ {
        res += max(min(leftMax[i], rightMax[i]) - h[i], 0)
    }
    return res
}