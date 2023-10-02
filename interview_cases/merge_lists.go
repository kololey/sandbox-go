package interview_cases

// https://leetcode.com/problems/merge-two-sorted-lists/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(N+M) time and O(1) space
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
