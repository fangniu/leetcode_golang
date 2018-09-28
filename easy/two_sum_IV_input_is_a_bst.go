package main

//Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their
// sum is equal to the given target.
//
//Example 1:
//Input:
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//Target = 9
//
//Output: True
//Example 2:
//Input:
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//Target = 28
//
//Output: False


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}
	return find(root, k, m)
}

func find(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok:= m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = true
	return find(root.Left, k, m) || find(root.Right, k, m)
}

//------------------先中序遍历得到有序数组-------------------

type TwoSumIV struct {
	list []int
}

func findTarget(root *TreeNode, k int) bool {
	t := TwoSumIV{}
	t.inOrder(root)
	i := 0
	j := len(t.list) - 1
	for i < j {
		sum := t.list[i] + t.list[j]
		if sum == k {
			return true
		} else if sum < k {
			i ++
		} else {
			j --
		}
	}
	return false
}

func (t *TwoSumIV) inOrder(root *TreeNode)  {
	if root == nil {
		return
	}
	t.inOrder(root.Left)
	t.list = append(t.list, root.Val)
	t.inOrder(root.Right)
}



