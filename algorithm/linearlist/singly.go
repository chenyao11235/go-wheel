package linearlist

import "fmt"

//SinglyNode 单向链表节点
type SinglyNode struct {
	value interface{}
	next  *SinglyNode
}

//GetValue 获取节点的值
func (n *SinglyNode) GetValue() interface{} {
	return n.value
}

//GetNext 获取下一个节点
func (n *SinglyNode) GetNext() *SinglyNode {
	return n.next
}

//Singly 单向链表
type Singly struct {
	head   *SinglyNode // 带头链表，链表中的第一个节点是head.next
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
	// 找到p节点的前一个节点
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 链表中只有一个节点(头节点)
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
func (s *Singly) InsterHead(value interface{}) bool {
	return s.insertAfter(s.head, value)
}

//InsertTail 在尾部插入节点
func (s *Singly) InsertTail(value interface{}) bool {
	curl := s.head
	for curl.next != nil {
		curl = curl.next
	}
	return s.insertAfter(curl, value)
}

//FindByIndex 查找指定值的索引
func (s *Singly) FindByIndex(index int) *SinglyNode {
	if index >= s.length {
		return nil
	}

	cur := s.head.next
	for i := 0; i < index; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}
	return cur
}

//DeleteNode 删除节点
func (s *Singly) DeleteNode(p *SinglyNode) bool {
	if p == nil {
		return false
	}

	pre := s.head
	cur := s.head.next

	for cur != nil {
		if cur == p {
			pre.next = p.next
			s.length--
			return true
		}
		pre = cur
		cur = cur.next
	}
	return false
}

//Print 打印链表
func (s *Singly) Print() {
	cur := s.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("=>%+v", cur.GetValue())
		cur = cur.next
	}
	fmt.Println(format)
}
