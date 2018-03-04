package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print(){
	fmt.Print(node.value, " ")
}

/*func (node *treeNode) setValue(value int){
	node.value = value
}*/

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return  // 这行必须加，否则报“panic: runtime error: invalid memory address or nil pointer dereference”
	}
	node.value = value
}

func (node *treeNode) traverse(){
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	//fmt.Println(root)
	root = treeNode{value: 3}
	root.left = &treeNode{} //root的左节点指向一个空的treeNode
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.setValue(10)
	root.print()
	nodes := []treeNode{
		{value: 3},
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

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.traverse()
}