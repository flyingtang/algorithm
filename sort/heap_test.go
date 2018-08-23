package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		n := rand.Intn(100000)
		arr[i] = n
	}

	// fmt.Printf("heap排序前：%v \n", arr)
	HeapSort(arr)

	//  最好用bce 方式比较两个slice,以下反射不推荐
	// okRes := []int{0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 8}
	// fmt.Printf("heap排序后：%v \n", arr)
	// ok := reflect.DeepEqual(arr, okRes)

	// if ok {
	// 	t.Log("测试通过")
	// } else {
	// 	t.Error("测试失败")
	// }
}
