package linearlist

//SinglyNode 单向链表节点
type SinglyNode struct {
	value interface{}
	next  *SinglyNode
}

//Singly 单向链表
type Singly struct {
	head   *SinglyNode
	length int
}

//NewSinglyNode 创建一个单向链表节点
func NewSinglyNode(value interface{}) *SinglyNode {
	return &SinglyNode{
		value: value,
		next:  nil,
	}
}

//NewSingly 创建一个单向链表
func NewSingly() *Singly {
	return &Singly{NewSinglyNode(0), 0}
}

//insertAfter 在某个节点之后插入一个节点
func (s *Singly) insertAfter(p *SinglyNode, value interface{}) bool {
	newNode := NewSinglyNode(value)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	s.length++
	return true
}

//insertBefore 在某个节点之前插入一个节点
func (s *Singly) insertBefore(p *SinglyNode, value interface{}) bool {
	if p == s.head {
		return false
	}
	cur := s.head.next
	pre := s.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewSinglyNode(value)
	pre.next = newNode
	newNode.next = cur
	s.length++
	return true
}

//InsterHead 在头部插入节点
func (s *Singly) InsterHead(value interface{}) {
	s.head = NewSinglyNode(value)
}

//InsertTail 在尾部插入节点
func (s *Singly) InsertTail() {
}
