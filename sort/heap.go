package sort

/**
堆排序
@Param  {[]int} arr 整形切片
*/
func HeapSort(arr []int) {
	// 切片长度
	length := len(arr)

	// for par := length/2 -1 // 最后一个非叶子节点

	for i := length/2 - 1; i >= 0; i-- {
		Sort(arr, i, length)
	}
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Sort(arr, 0, i)
	}
}

func Sort(arr []int, start, end int) {
	dad := start
	son := dad*2 + 1
	if son >= end {
		return
	}
	if son+1 < end && arr[son] < arr[son+1] {
		son++
	}
	if arr[son] > arr[dad] {
		arr[son], arr[dad] = arr[dad], arr[son]
		Sort(arr, son, end)
	}
}
