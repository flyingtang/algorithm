package search

func BinarySearch(arr []int, key int) (index int) {
	index = -1
	low, high := 0, len(arr)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid
		} else {
			high = mid
		}
	}
	return
}
