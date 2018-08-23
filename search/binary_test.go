package search

import (
	"fmt"
	"go-sort/sort"
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	rand.Seed(time.Now().Unix())
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		n := rand.Intn(100000)
		arr[i] = n
	}
	sort.QuickSort(arr, 0, len(arr)-1)
	index := BinarySearch(arr, 100)

	fmt.Println("index = ", index)
	fmt.Println(arr[:index+1], len(arr[:index+1]))
}
