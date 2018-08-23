package sort

/**
归并排序  O(n log n)
@Param  {[]int} arr 整形切片
*/
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])
	return Merge(l, r)
}

func Merge(left, right []int) (result []int) {
	l, r := 0, 0 // 左右两切片的下标
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
