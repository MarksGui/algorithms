package binary_search_704

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		pivot := left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return -1
}
