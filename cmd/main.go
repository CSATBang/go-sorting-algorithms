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
	algorithms.MergeSort(arr)
	fmt.Println("MergeSort:", arr)
}
