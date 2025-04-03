func majorityElement(nums []int) int {
    hist := make(map[int]int)
    for _, num := range nums {
        hist[num] += 1 
    }
    times := len(nums) / 2
    maj := 0
    for key, value := range hist {
        if value > times {
            maj = key
        }
    } 
    return maj
}
