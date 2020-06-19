package linear_list

type SinglyNode struct {
    value interface{}
    next  *SinglyNode
}

type Singly struct {
    head   *SinglyNode
    length int
}

func NewSinglyNode(value interface{}) *SinglyNode {
    return &SinglyNode{
        value: value,
        next:  nil,
    }
}

func NewSingly() *Singly {
    return &Singly{NewSinglyNode(0), 0}
}

func (s *Singly) insertAfter(p *SinglyNode, value interface{}) bool {
    newNode := NewSinglyNode(value)
    oldNext := p.next
    p.next = newNode
    newNode.next = oldNext
    s.length++
    return true
}

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

func (s *Singly) InterHead(value interface{}) {
    s.head = NewSinglyNode(value)
}

func (s *Singly) InsertTail() {

}
