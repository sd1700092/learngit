package main

import (
	"fmt"

	"imooc.com/learngo/tree"
)

func main() {
	var root tree.Node
	//fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{} //root的左节点指向一个空的treeNode
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.SetValue(10)
	root.Print()
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	/*	pRoot := &root
		pRoot.print()
		pRoot.setValue(200)
		pRoot.print()*/

	/*	var pRoot *treeNode
		pRoot.setValue(200)*/

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
}
