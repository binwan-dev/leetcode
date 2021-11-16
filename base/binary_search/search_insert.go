package main

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums)
	for {
		dis := end - start
		if dis == 1 {
			if nums[start] >= target {
				return start
			}
			return end
		}

		middle := dis/2 + start
		if nums[middle] > target {
			end = middle
		}
		if nums[middle] < target {
			start = middle
		}
		if nums[middle] == target {
			return middle
		}
	}
}
