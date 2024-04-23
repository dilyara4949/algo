func findPeakElement(nums []int) int {
    l, r := 0, len(nums)-1

    for l <= r {
        m := (r+l)/2
        if m-1>= 0 && m+1<len(nums) || l==r{
            if l!=r && nums[m] > nums[m-1] && nums[m] > nums[m+1] || l==r{
              return m
            } else if nums[m-1] > nums[m+1] {
                r = m-1
            } else {
                l = m+1
            }
        } else if m-1 == -1 && nums[m] < nums[m+1] {
            l++
        } else {
            r--
        }
    }
    return l
}