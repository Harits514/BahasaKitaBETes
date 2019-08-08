package main

import (
	"fmt"
)

func merge(arr []int, l int, m int, r int) {
	leftLimit := m - l + 1
	rightLimit := r - m
	leftSlc := make([]int, leftLimit)
	copy(leftSlc, arr[l:m+1])
	rightSlc := make([]int, rightLimit)
	copy(rightSlc, arr[m+1:r+1])

	fmt.Println("merging:", l, m, r)

	i := 0
	j := 0
	k := l
	fmt.Println(leftSlc, rightSlc)

	for i < leftLimit && j < rightLimit {
		if leftSlc[i] <= rightSlc[j] {
			arr[k] = leftSlc[i]
			i++
		} else {
			arr[k] = rightSlc[j]
			j++
		}
		k++
	}
	fmt.Println("hey")
	for i < leftLimit {
		arr[k] = leftSlc[i]
		i++
		k++
	}
	for j < rightLimit {
		arr[k] = rightSlc[j]
		j++
		k++
	}
}

func mergesort(arr []int, l int, r int) {
	if l < r {
		m := int(l + (r-l)/2)
		fmt.Println("l, m, r processed:", l, m, r)
		mergesort(arr, l, m)
		mergesort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}

func main() {
	fmt.Println("input any amount of non-negative number separated with enter,\nif done, input any negative number")
	var num int
	var nums []int
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		nums = append(nums, num)
	}
	lenSlice := len(nums)
	mergesort(nums, 0, lenSlice-1)
	fmt.Println(nums)
}
