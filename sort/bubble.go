package sort

/**
冒泡排序 O(n^2)
@Param []int 整形切片
*/
func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
