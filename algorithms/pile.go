package algorithms

/*
堆的基本结构：是一个完全二叉树，也就是除了最后一层节点外，上层节点都是非空的,最后一层节点是从左到右依次排布的
堆满足一个性质，以小根堆为例，每一个节点都小于等于它的左右儿子,根节点是整个数据结构的最小值。

如何手写一个小根堆
堆的存储：用一个一维数组，根节点坐标设为从x=1开始 节点的左儿子index=2x,右儿子index=2x+1
若下标0开始，则左儿子2x+1，右儿子 2x+2

堆的操作(小根堆为例)：
1.插入一个数				heap[++size]=x;up(size)
2.求集合中的最小值			heap[1]
3.删除最小值				heap[1]=heap[size] ;size--;down(1);  用最后一个堆的元素来覆盖堆顶的元素再把最后一个元素删除，然后down(1)一遍
4.删除任意一个元素			heap[k]=heap[size] ;size--;down(k);up(k); 与上同理
5.修改任意一个元素			heap[k]=x;down(k);up(k);

都可以由两个函数实现：
down(x){
	节点往下调整
}
up(x){
	节点往上调整
}

堆排序也就参考上面的down操作
这里排序使用大根堆
*/

func PileSort(arr []int) []int {
	heapSize := len(arr)
	//建立大根堆，因为大根堆有节点大于等于左右儿子的性质，根节点的值就是整个数据结构最大的
	buildMaxHeap(arr, heapSize)
	// 于是我们可以把根节点与最后一个节点互换，把最大值固定到数组末尾
	// 经过多轮操作，数组会从前到后变为升序排列
	for i := heapSize - 1; i > 0; i-- {
		//根节点与最后一个节点互换
		swap(arr, 0, i)
		//范围缩小由于
		heapSize--
		//整个树下沉排序
		down(arr, 0, heapSize)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	// 从最后一个非叶子节点开始进行建堆
	// 如果下标从 0 开始，最后一个非叶子节点是 arrLen/2 - 1
	// 如果下标从 1 开始，最后一个非叶子节点是 arrLen/2

	for i := arrLen/2 - 1; i >= 0; i-- {
		down(arr, i, arrLen)
	}
}

// 下沉操作
func down(arr []int, i, arrLen int) {
	left := i*2 + 1
	right := i*2 + 2
	//largest是用来记录某节点与它的左右孩子的最大值
	largest := i
	// 若左儿子存在，且左儿子大于当前最大值节点，则更新 largest = 左儿子下标
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	//若右儿子存在，且右儿子大于当前最大值节点，则更新 largest = 右儿子下标
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	//largest存的是三个节点的最大的节点标号，若不等于i,说明这里根节点就不是最大值
	if largest != i {
		swap(arr, i, largest)
		down(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
