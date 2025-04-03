func removeElement(nums []int, val int) int {
    var k int
    for i := range len(nums) {
        if nums[i] != val {
            nums[i], nums[k] = nums[k], nums[i]
            k++
        }
    }
    return k
}
