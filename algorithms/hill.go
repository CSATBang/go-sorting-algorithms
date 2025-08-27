package algorithms

//希尔排序  插入排序的进阶版 将序列分组再插入

func HillSort(arr []int) []int {
	gap := 1
	n := len(arr)
	if gap < n {
		gap = n / 2
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap /= 2
	}
	return arr
}
