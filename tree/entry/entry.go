package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets"
	"imooc.com/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(100000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}

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

	/*	fmt.Println()
		myRoot := myTreeNode{&root}
		myRoot.postOrder()
		fmt.Println()

		testSparse()*/
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
