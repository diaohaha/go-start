package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := []int{}
	toTraverse := []*Node{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, child := range root.Children {
		toTraverse = append(toTraverse, child)
	}
	for true {
		if len(toTraverse) == 0 {
			break
		}
		nowRoot := toTraverse[0]
		// 切片操作
		toTraverse = toTraverse[1:]
		res = append(res, nowRoot.Val)
		// ... 语法
		toTraverse = append(nowRoot.Children, toTraverse...)
	}
	return res
}
