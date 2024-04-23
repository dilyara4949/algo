func search(num []int, t int) int {
    l, r := 0, len(num)-1
    for l < r {
        m :=(r+l)/2
        if m-1>=0 && m+1 < len(num) && num[m] < num[m-1] && num[m] < num[m+1] {
            l = m
            break
        } else if num[m] >= num[l] && num[m] >= num[r] {
            l = m+1
        } else {
            r = m
        }
    }
    fmt.Println(l)
    if num[l] <= t && t <= num[len(num)-1] {
        return binSearch(num, l, len(num)-1, t)
    }
    return binSearch(num, 0, l-1, t)
}

func binSearch(num []int, l, r, t int) int {
    for l <= r {
        m := l + (r-l)/2
        if num[m] == t {
            return m
        } else if num[m] > t {
            r = m-1
        } else {
            l = m+1
        }
    }
    return -1
}