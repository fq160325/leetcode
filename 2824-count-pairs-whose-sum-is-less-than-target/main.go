package CountPairsWhoseSumIsLessThanTarget

import (
	"fmt"
	"sync"
)

func Main() {
	var arr = []int{-5, 0, -7, -1, 9, 8, -9, 9}
	fmt.Println(countPairs(arr, -14))
	fmt.Println(arr)
}

/*
1、从数列中挑出一个元素，称为 "基准"（pivot）;

2、重新排序数列，所有元素比基准值小的摆放在基准前面，
所有元素比基准值大的摆在基准的后面（相同的数可以
到任一边）。在这个分区退出之后，该基准就处于数列
的中间位置。这个称为分区（partition）操作；

3、递归地（recursive）把小于基准值元素的子数列和大于
基准值元素的子数列排序；
*/
func quickSort(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 针对特殊情况,数组lens为0或者1时
	if len(nums) <= 1 {
		return
	}
	var pivot, pivotIndex = nums[len(nums)-1], len(nums) - 1
	for i, j := 0, pivotIndex; i < j; {
		for pivot >= nums[i] && i < j {
			i++
		}
		if i >= j {
			break
		}
		nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex]
		pivotIndex = i
		for pivot < nums[j] && i < j {
			j--
		}
		if i >= j {
			break
		}
		nums[pivotIndex], nums[j] = nums[j], nums[pivotIndex]
		pivotIndex = j
	}
	nums[pivotIndex] = pivot
	wg.Add(2)
	quickSort(nums[:pivotIndex], wg)
	quickSort(nums[pivotIndex+1:], wg)
}

func countPairs(nums []int,
	target int) int {
	var wg sync.WaitGroup
	wg.Add(1)
	// 首先进行排序
	quickSort(nums, &wg)
	wg.Wait()
	var res int
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] >= target {
			j--
			continue
		}
		res += j - i
		i++
	}
	return res
}
