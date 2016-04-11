package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return branchNum(root, 0)
}

func branchNum(root *TreeNode, prev int) int {
	if root == nil {
		return 0
	}
	val := prev*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return val
	} else {
		return branchNum(root.Left, val) + branchNum(root.Right, val)
	}
}

func main() {
}
