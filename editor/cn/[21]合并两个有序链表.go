//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 
//
// 示例 2： 
//
// 
//输入：l1 = [], l2 = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 两个链表的节点数目范围是 [0, 50] 
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 
// 
// Related Topics 递归 链表 
// 👍 1708 👎 0

package cn

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var lastNode, rootNode *ListNode
	isRoot := true
	for {
		listNode := new(ListNode)
		if l1 == nil || l2 == nil {
			if l1 == nil && l2 == nil {
				break
			}
			if l1 == nil {
				listNode = l2
			}
			if l2 == nil {
				listNode = l1
			}
			if lastNode != nil {
				lastNode.Next = listNode
			}
			lastNode = listNode
			if isRoot {
				rootNode = lastNode
				isRoot = false
			}
			break
		} else {
			if l1.Val < l2.Val {
				listNode.Val = l1.Val
				l1 = l1.Next
			} else {
				listNode.Val = l2.Val
				l2 = l2.Next
			}
			if lastNode != nil {
				lastNode.Next = listNode
			}
			lastNode = listNode
			if isRoot {
				rootNode = lastNode
				isRoot = false
			}
		}
	}
	return rootNode
}

//leetcode submit region end(Prohibit modification and deletion)
