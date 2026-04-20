func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // всегда бинарный поиск по меньшему массиву
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    m, n := len(nums1), len(nums2)
    half := (m + n + 1) / 2 // размер левой половины

    low, high := 0, m
    for low <= high {
        mid1 := (low + high) / 2   // берём mid1 элементов из nums1
        mid2 := half - mid1         // берём mid2 элементов из nums2

        // граничные значения разреза (±inf для краёв массива)
        l1 := math.MinInt64
        if mid1 > 0 { l1 = nums1[mid1-1] }

        r1 := math.MaxInt64
        if mid1 < m { r1 = nums1[mid1] }

        l2 := math.MinInt64
        if mid2 > 0 { l2 = nums2[mid2-1] }

        r2 := math.MaxInt64
        if mid2 < n { r2 = nums2[mid2] }

        if l1 <= r2 && l2 <= r1 {
            // нашли правильный разрез!
            if (m+n)%2 == 1 {
                return float64(max(l1, l2))
            }
            return float64(max(l1,l2)+min(r1,r2)) / 2.0
        } else if l1 > r2 {
            high = mid1 - 1 // взяли слишком много из nums1
        } else {
            low = mid1 + 1  // взяли слишком мало из nums1
        }
    }
    return 0
}