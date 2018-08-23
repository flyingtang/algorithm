package sort

/**
快速排序
@Param []int 整形切片
*/
func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	first, last := low, high
	key := arr[first]
	for first < last {
		for first < last && arr[last] >= key {
			last--
		}
		arr[first] = arr[last]
		for first < last && arr[first] <= key {
			first++
		}
		arr[last] = arr[first]
	}
	arr[first] = key
	QuickSort(arr, low, first-1)
	QuickSort(arr, first+1, high)
}
