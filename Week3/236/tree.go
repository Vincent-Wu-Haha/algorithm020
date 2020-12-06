package _36

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Left != nil {
		if r := lowestCommonAncestor(root.Left, p, q); r != nil {
			return r
		}
	}
	if root.Right != nil {
		if r := lowestCommonAncestor(root.Right, p, q); r != nil {
			return r
		}
	}

	if commonAncestor(root, p, q) {
		return root
	}
	return nil
}

func commonAncestor(root, p, q *TreeNode) bool {
	if ancestor(root, p) && ancestor(root, q) {
		return true
	}
	return false
}

func ancestor(root, p *TreeNode) bool {
	if root == p {
		return true
	}
	if root.Left != nil && ancestor(root.Left, p) {
		return true
	}
	if root.Right != nil && ancestor(root.Right, p) {
		return true
	}
	return false
}
