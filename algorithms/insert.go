package algorithms

// 插入排序
func InsertSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		//找位置并右移
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}
