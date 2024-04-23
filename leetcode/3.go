func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    l, r := 0, 0
    res := 0
    for ; r < len(s); r++ {
        if mp[s[r]]==0 && (s[0] != s[r]) || mp[s[r]]==-1  || r == 0{
            mp[s[r]] = r
            res = max(res, r-l+1)
        } else {
            for l < mp[s[r]] {
                mp[s[l]] = -1
                l++
            }
            l = mp[s[r]]+1
            mp[s[r]] = r
        }
    }
    return res
}