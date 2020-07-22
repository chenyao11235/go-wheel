package binarytree

//BinarySearchTreeNode 节点
type BinarySearchTreeNode struct {
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
	value int
	times int // 值出现的次数
}

func newBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		value: value,
	}
}

func (n *BinarySearchTreeNode) add(value int) {
	// 添加到左子节点
	if value < n.value {
		if n.left == nil {
			n.left = newBinarySearchTreeNode(value)
		} else {
			// 递归
			n.left.add(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = newBinarySearchTreeNode(value)
		} else {
			// 递归
			n.right.add(value)
		}
	} else {
		n.times++
	}
	return
}

//findMinValue 找出最小值
func (n *BinarySearchTreeNode) findMinValue() *BinarySearchTreeNode {
	if n.left == nil {
		return n
	}
	return n.left.findMinValue()
}

//findMaxValue 找出最大值
func (n *BinarySearchTreeNode) findMaxValue() *BinarySearchTreeNode {
	if n.right == nil {
		return n
	}
	return n.right.findMinValue()
}

//BinarySearchTree 二叉查找树
type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

//NewBinarySearchTree 新建
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

//Add 添加
func (t *BinarySearchTree) Add(value int) {
	if t.root == nil {
		t.root = newBinarySearchTreeNode(value)
	} else {
		t.root.add(value)
	}
}

//FindMinValue 找出最小值的节点
func (t *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	return t.root.findMinValue()
}

//FindMaxValue 找出最大值的节点
func (t *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}
	return t.root.findMaxValue()
}
