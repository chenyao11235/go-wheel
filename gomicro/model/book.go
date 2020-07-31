package model

//Book 书籍
type Book struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	PID  int64  `json:"pid" gorm:"column:pid"`
}

//BookKind 书籍种类
type BookKind struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"Id" gorm:"column:name"`
}
