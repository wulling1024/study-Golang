package LeetCode

import (
	"fmt"
	"reflect"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func GenerateList(arr []int) *ListNode{
	//var node, result *ListNode
	node := new(ListNode)
	result := node
	fmt.Println(reflect.TypeOf(result))
	for _, tmp := range arr {
		tag := new(ListNode)
		tag.Val = tmp
		node.Next = tag
		node = tag
	}
	return result
}

func Test_main(t *testing.T){

	// 创建三个链表用于测试
	list  := GenerateList([]int{1, 4, 5})
	list2 := GenerateList([]int{1, 3, 4})
	list3 := GenerateList([]int{2, 6})
	lists := make([]*ListNode, 3)
	lists[0] = list.Next
	lists[1] = list2.Next
	lists[2] = list3.Next
	result := mergeKLists(lists)
	for {
		if result == nil {
			break
		}
		fmt.Printf("%v ",result.Val)
		result = result.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0{
		return nil
	}
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	middle := start + (end - start)/2
	left := merge(lists, start, middle)
	right := merge(lists, middle + 1, end)

	return merge2List(left, right)
}

func merge2List(left, right *ListNode) *ListNode{
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val > right.Val {
		right.Next = merge2List(left, right.Next)
		return right
	}else {
		left.Next = merge2List(left.Next, right)
		return left
	}
}
