package linearlist

import (
	"fmt"
)

/*
判断一个链表是不是会文链表，这里提供另种方法
*/

// 自己先实现一种最笨的方法
func isPalindrome(l *Singly) bool {
	cur := l.head.next
	m := l.length / 2
	// 链表的前半部分
	firstHalf := make([]interface{}, m)
	// 链表的后半部分
	secondHalf := make([]interface{}, m+1)
	for i := 0; i < l.length; i++ {
		if i < m {
			firstHalf[i] = cur.GetValue()
		} else {
			secondHalf[i] = cur.GetValue()
		}
		cur = cur.next
	}

	fmt.Println(firstHalf, secondHalf)
	for i := 0; i < m; i++ {
		if len(firstHalf) == len(secondHalf) {
			if firstHalf[i] != secondHalf[m-i] {
				return false
			}
		} else {
			if firstHalf[i] != secondHalf[i+1] {
				return false
			}
		}
	}
	return true
}

// 方法1:开一个栈存放链表示链表的前半段
func isPalindrome1(l *Singly) {

}

// 方法2:
func isPalindrome2(l *Singly) {

}
