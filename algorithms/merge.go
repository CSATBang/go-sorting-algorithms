package algorithms

// 归并排序
func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	//先递归
	left := append([]int(nil), arr[:mid]...)
	right := append([]int(nil), arr[mid:]...)

	MergeSort(left)
	MergeSort(right)
	//再合并
	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			k++
			i++
		} else {
			arr[k] = right[j]
			k++
			j++
		}
	}
	for i < len(left) {
		arr[k] = left[i]
		k++
		i++
	}
	for j < len(right) {
		arr[k] = right[j]
		k++
		j++
	}
}
