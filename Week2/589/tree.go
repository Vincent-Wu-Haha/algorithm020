package _89

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (result []int) {
	stack := make([][]*Node, 0, 10)
	for root != nil {
		result = append(result, root.Val)
		if len(root.Children) == 1 {
			root = root.Children[0]
			continue
		} else if len(root.Children) > 1 {
			stack = append(stack, root.Children[1:])
			root = root.Children[0]
			continue
		}
		if len(stack) == 0 {
			break
		}
		level := stack[len(stack)-1]
		root = level[0]
		if len(level) == 1 {
			stack = stack[:len(stack)-1]
		} else {
			stack[len(stack)-1] = level[1:]
		}
	}
	return
}
