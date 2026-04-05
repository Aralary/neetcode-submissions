// func findDuplicate(nums []int) int {
//     n := len(nums) - 1
// 	tmp := 0 << n
// 	for _, el := range nums {
// 		t := 1 << el
// 		if tmp | t == tmp {
// 			return el
// 		}
// 		tmp = tmp | t
// 	}
// 	return 0
// }

func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }

    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}
