package sort

/**
希尔排序 O(nlogn)
@Param {[]int} arr
*/
func ShellSort(arr []int) {
	length := len(arr)
	for fraction := length / 2; fraction > 0; fraction /= 2 {
		for i := fraction; i < length; i++ {
			for j := i - fraction; j >= 0 && arr[j] > arr[j+fraction]; j -= fraction {
				arr[j], arr[j+fraction] = arr[j+fraction], arr[j]
			}
		}
	}
}
