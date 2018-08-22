package sort

func HeapSort(arr []int) {
	// 切片长度
	length := len(arr)

	// for par := length/2 -1 // 最后一个非叶子节点

	for i := length/2 - 1; i >= 0; i-- {
		hSort(arr, i, length)
	}
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		hSort(arr, 0, i)
	}
}

func hSort(arr []int, start, end int) {
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
		hSort(arr, son, end)
	}
}
