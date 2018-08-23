package sort

/**
选择排序 O(n^2)
@Param []int 整形切片
*/
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		if index != i {
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}
