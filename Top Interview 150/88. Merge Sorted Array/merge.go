func merge(nums1 []int, m int, nums2 []int, n int)  {
    result := make([]int, m + n)
    var i, j, k int
    for i < m && j < n {
        if nums1[i] < nums2[j] {
            result[k] = nums1[i]
            i++ 
        } else {
            result[k] = nums2[j]
            j++
        }
        k++
    }

    for i < m {
        result[k] = nums1[i]
        k++ 
        i++ 
    }

    for j < n {
        result[k] = nums2[j]
        k++
        j++ 
    }

    for t := range m + n {
        nums1[t] = result[t]
    }
}
