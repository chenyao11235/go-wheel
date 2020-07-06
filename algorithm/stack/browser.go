package stack

//Browser 浏览器
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

//NewBrowser 新建一个浏览器
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewStackBaseOnLinkedList(),
		backStack:    NewStackBaseOnArray(),
	}
}

func (b *Browser) canForward() bool {
	return true
}

func (b *Browser) canBack() bool {
	return true
}

//Forward 前进
func (b *Browser) Forward() {

}

//Back 后退
func (b *Browser) Back() {

}
