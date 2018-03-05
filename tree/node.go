package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print(){
	fmt.Print(node.Value, " ")
}

/*func (node *treeNode) setValue(Value int){
	node.Value = Value
}*/

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored")
		return  // 这行必须加，否则报“panic: runtime error: invalid memory address or nil pointer dereference”
	}
	node.Value = value
}

