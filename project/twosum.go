package main

import "fmt"

func main() {
	hitung := twosum([]int{3, 3, 2, 2}, 6)
	fmt.Println(hitung)
}

func twosum(nums []int, target int) []int {
	var hasil []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				hasil = append(hasil, i)
				hasil = append(hasil, j)
			}
		}
	}
	return hasil
}
