func search(nums []int, target int) int {
	start := findMin(nums)
	var end int = len(nums) - 1
	if start != 0 {
		if target > nums[end] {
			// ищем в левой части
			end = start - 1
			start = 0
		}
	}
	var mid int = -1
	for start <= end {
		mid = (end - start) / 2 + start
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func findMin(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := (right-left)/2 + left

        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left
}
