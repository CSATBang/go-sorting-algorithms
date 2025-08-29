package main

import (
	"fmt"
	"github.com/CSATBang/go-sorting-algorithms/algorithms"
)

func pileSort(a []int) {

}
func countSort(a []int) {

}
func bucketSort(a []int) {

}
func cardinalSort(a []int) {

}
func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 47, 4, 19, 50, 48}
	bubbleArr := append([]int(nil), arr...)
	hillArr := append([]int(nil), arr...)
	insertArr := append([]int(nil), arr...)
	quickArr := append([]int(nil), arr...)
	mergeArr := append([]int(nil), arr...)
	fmt.Println("BubbleSort:", algorithms.BubbleSort(bubbleArr))
	fmt.Println("HillSort:", algorithms.HillSort(hillArr))
	fmt.Println("InsertSort:", algorithms.InsertSort(insertArr))
	fmt.Println("QuickSort:", algorithms.QuickSort(quickArr, 0, len(quickArr)-1))
	algorithms.MergeSort(mergeArr)
	fmt.Println("MergeSort:", mergeArr)
	fmt.Println("arr:", arr)

}
