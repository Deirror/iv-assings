func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]bool)
    maxLen, leftMost := 0, 0

    for right := 0; right < len(s); right++ {
        for set[s[right]] {
            delete(set, s[leftMost])
            leftMost++
        }
        set[s[right]] = true
        maxLen = max(maxLen, len(set))
    }
    return maxLen
}
