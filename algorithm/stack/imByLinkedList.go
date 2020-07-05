package stack

//Node 节点
type Node struct {
	value interface{}
	next  *Node
}

//LinkedList 链表
type LinkedList struct {
	head   *Node
	length int
}

//BaseOnLinkedList 基于链表实现的栈
type BaseOnLinkedList struct {
}

//Push 压栈
func (s *BaseOnLinkedList) Push(v interface{}) {
	panic("not implemented") // TODO: Implement
}

//Pop 出栈
func (s *BaseOnLinkedList) Pop() interface{} {
	panic("not implemented") // TODO: Implement
}

//IsEmpty 判断是否为空
func (s *BaseOnLinkedList) IsEmpty() bool {
	panic("not implemented") // TODO: Implement
}

//Top 取栈顶
func (s *BaseOnLinkedList) Top() interface{} {
	panic("not implemented") // TODO: Implement
}

//Flush 清空
func (s *BaseOnLinkedList) Flush() {
	panic("not implemented") // TODO: Implement
}
