package builtin

import "unsafe"


func AarraySum(arr [10]int) int{
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func SliceSum(s []int) int {
	sum := 0
	for _,v := range s{
		sum += v
	}

	return sum
}

func Array2Slice(arr *[10]int) []int {
	return unsafe.Slice(&arr[0], len(arr))
}