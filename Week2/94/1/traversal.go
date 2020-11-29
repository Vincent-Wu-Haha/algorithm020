package _1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0, 10)
	stack = push(stack, root)
	result := make([]int, 0, 10)
	for n := root; len(stack) != 0; {
		if n.Left != nil {
			stack = push(stack, n.Left)
			temp := n
			n = n.Left
			temp.Left = nil
			continue
		}
		stack, n = pop(stack)
		result = append(result, n.Val)
		if n.Right != nil {
			stack = push(stack, n.Right)
			temp := n
			n = n.Right
			temp.Right = nil
			continue
		}
	}
	return result
}

func push(stack []*TreeNode, node *TreeNode) []*TreeNode {
	stack = append(stack, node)
	return stack
}

func pop(stack []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(stack) == 0 {
		return nil, nil
	}
	temp := stack[len(stack)-1]
	return stack[:len(stack)-1], temp
}
