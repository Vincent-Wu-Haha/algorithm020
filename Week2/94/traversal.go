package _4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	l := inorderTraversal(root.Left)[:]
	l = append(l, root.Val)
	l = append(l, inorderTraversal(root.Right)...)
	return l
}
