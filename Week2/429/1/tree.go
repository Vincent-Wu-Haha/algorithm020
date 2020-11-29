package _1

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for l := len(q); len(q) != 0; l = len(q) {
		var level []int
		for i := 0; i < l; i++ {
			n := q[i]
			level = append(level, n.Val)
			q = append(q, n.Children...)
		}
		q = q[l:]
		result = append(result, level)
	}
	return
}
