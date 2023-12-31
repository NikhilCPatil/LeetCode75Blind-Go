package main

import "fmt"
//binarry search
func binarrySearch(nums []int, target int) int{
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}

	}

	return -1
}

func main(){
	nums :=	[]int{0, 1, 2, 4, 5, 6, 7}
	target := 7
	result := binarrySearch(nums, target)
	fmt.Println("Target found at index", result)
}