package _1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0, 10)
	result := make([]int, 0, 10)
	stack = append(stack, root)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, n.Val)
		r := n.Right
		if r != nil {
			stack = append(stack, r)
		}
		l := n.Left
		if l != nil {
			stack = append(stack, l)
		}
	}
	return result
}
