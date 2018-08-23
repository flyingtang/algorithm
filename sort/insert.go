package sort

/**
插入排序 O(n^2)
@Param []int 整形切片
*/
func InsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
