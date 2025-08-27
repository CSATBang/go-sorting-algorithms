package main

import (
	"fmt"
	"github.com/CSATBang/go-sorting-algorithms/algorithms"
	"time"
)

func pileSort(a []int) {
	start := time.Now()

	fmt.Println("pileSort：", time.Since(start))
}
func countSort(a []int) {
	start := time.Now()

	fmt.Println("countSort：", time.Since(start))
}
func bucketSort(a []int) {
	start := time.Now()

	fmt.Println("bucketSort：", time.Since(start))
}
func cardinalSort(a []int) {
	start := time.Now()

	fmt.Println("cardinalSort：", time.Since(start))
}
func main() {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 47, 4, 19, 50, 48}
	algorithms.MergeSort(arr)
	fmt.Println("MergeSort:", arr)
}
