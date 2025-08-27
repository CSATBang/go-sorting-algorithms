package algorithms

// QuickSort 快速排序 理念是分治，选一个值，以它为标准,然后划分区间，小于它的在左侧，大于它的再右侧，递归实现
// 选取两个指针，一个从左侧left向右遍历，一个从右侧right向左遍历，遍历的过程与标准值比较
// 若如果所有数都一样的话，时间复杂度会退化到 𝑂(n*2)
// _______<x(这一侧是小于X的值)__________|x(标准值)|____________(这一侧是大于X的值)>x___________
func QuickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}
	midValue := arr[l]
	i, j := l-1, r+1

	for i < j {
		//模仿C语言里的do while实现
		for {
			i++
			if arr[i] >= midValue {
				break
			}
		}
		for {
			j--
			if arr[j] <= midValue {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}

	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
	return arr
}
