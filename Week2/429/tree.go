package _29

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	cur := []*Node{root}
	var next []*Node
	level := 0
	for len(cur) != 0 {
		for _, node := range cur {
			if len(result) < level+1 {
				result = append(result, []int{})
			}
			result[level] = append(result[level], node.Val)
			next = append(next, node.Children...)
		}
		level++
		cur = next
		next = nil
	}
	return
}
