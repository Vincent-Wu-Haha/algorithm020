package _05

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	inorderIdx := findIdx(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:inorderIdx+1], inorder[:inorderIdx]),
		Right: buildTree(preorder[inorderIdx+1:], inorder[inorderIdx+1:]),
	}
}

func findIdx(arr []int, item int) int {
	for i, a := range arr {
		if a == item {
			return i
		}
	}
	return 0
}