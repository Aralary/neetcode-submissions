func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func findPivot(nums []int, start, end int) int {
	mid := (end - start)/2 + start
	a, b, c := nums[start], nums[mid], nums[end]
    median := a + b + c - min(a, min(b, c)) - max(a, max(b, c))
    switch median {
		case a: return start
		case b: return mid
		default: return end
    }
}

func quickSort(nums[]int, start, end int) {
	if start < end {
		p := partition(nums,start, end)
		quickSort(nums, start, p - 1)
		quickSort(nums, p + 1, end)
	}
}

func partition(nums []int, start, end int)int {
	pivotPos := findPivot(nums, start, end)
	pivot := nums[pivotPos]
	nums[end], nums[pivotPos] = nums[pivotPos], nums[end]
	endPos := end
	end--
	for start < end {
		if nums[start] < pivot {
			start++
			continue
		}
		if nums[end] >= pivot {
			end--
			continue
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	if nums[end] >= pivot { nums[endPos], nums[end] = nums[end], nums[endPos] }
	return end
}