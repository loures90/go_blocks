package methodsinterfaces

import "fmt"

type Node interface {
	SetValue(v int)
	GetValue() int
}

type SLLNode struct {
	Next  *SLLNode
	Value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.Value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.Value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode) // new keyword allocates memory
	// constructor -> it returns a pointer to a new sslNode.
}

//powernode
type PowerNode struct {
	next  *PowerNode
	value int
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

func (sNode SLLNode) SetValueWithNoPointers(v int) {
	sNode.Value = v
}

func (sNode SLLNode) GetValueWithNoPointers() int {
	return sNode.Value
}

func NewSLLNodeWithNoPointers() SLLNode {
	return SLLNode{}
}

func Test() {
	var node Node // Node interface implements SetValue and GetValue methods so it can be used by SSLNode and PowerNode, because both also implements this methods
	// node := NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of Value: ", node.GetValue())

	nodeWithNoPointers := NewSLLNodeWithNoPointers()
	nodeWithNoPointers.SetValueWithNoPointers(5)
	fmt.Println("Node with no pointer is of Value: ", nodeWithNoPointers.GetValue())
	// sslNode{} is not pointer so the "setter" does not change the value of the memory where the variable  "nodeWithNoPointers" was created (line 40)

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Node is of Value: ", node.GetValue())

}
