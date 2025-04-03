func hIndex(citations []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(citations)))
    h := 0
    for i, c := range citations {
        if c >= i + 1 {
            h = i + 1
        } else {
            break
        }
    }
    return h
}
