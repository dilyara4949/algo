func merge(itv [][]int) [][]int {
    arr := Matrix(itv)
    sort.Sort(&arr)

    res := [][]int{}
    res = append(res, arr[0])
    for i := 0; i < len(arr)-1; i++ {
        if res[len(res)-1][1] >= arr[i+1][0] {
            res[len(res)-1][1] = max(arr[i+1][1], res[len(res)-1][1])
        } else {
            res = append(res, arr[i+1])
        }
    }
    return res
}
type Matrix [][]int
func (m Matrix) Len() int { return len(m) }
func (m Matrix) Less(i, j int) bool {
    for x := range m[i] {
        if m[i][x] == m[j][x] {
            continue
        }
        return m[i][x] < m[j][x]
    }
    return false
}

func (m Matrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }