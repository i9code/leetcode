package main

import (
	"fmt"
)

func main()  {
	fmt.Println( permute([]int{1,2,3}))
}

func permute(nums []int) [][]int {
    return subNumberSlice(nums)
}

func subNumberSlice(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }
    if len(nums) == 2 {
        return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
    }

    result := [][]int{}
    for index, value := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		numsSubOne := append(numsCopy[:index],numsCopy[index+1:]...)
		fmt.Println("------------",numsSubOne)
		valueSlice := []int{value}
		fmt.Println("------------",valueSlice)
		newSubSlice := subNumberSlice(numsSubOne)
		fmt.Println(newSubSlice)
        for _, newValue := range newSubSlice {
            result = append(result, append(valueSlice, newValue...))
        }
    }
    return result
}
