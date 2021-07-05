package main

import (
	"fmt"
	"golang-algo-datastructures/structs"
)

func bubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func insertSlice(a []int, index int, value int) []int {
	a = append(a, 0)
	copy(a[index+1:], a[index:])
	a[index] = value
	return a
}

func reverseArray(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func binarySearch(a []int, value int) int {
	n := len(a)
	low, high, mid := 0, n, 0
	for low < high {
		mid = (low + high) / 2
		if a[mid] == value {
			return mid
		} else if a[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

var head *structs.Node = new(structs.Node)

func main() {
	//var scores = []int{90,70,50,80,60,85}
	//bubbleSort(scores)
	//fmt.Println(scores)
	//var index = binarySearch(scores, 600)
	////scores = insertSlice(scores, 1, 100)
	////reverseArray(scores)
	//fmt.Println(index)
	head.Data = "San"
	head.Next = nil
	var nodeOakland *structs.Node = &structs.Node{
		Data: "Oakland",
		Next: nil,
	}
	head.Next = nodeOakland
	fmt.Println(head.Data)
}
