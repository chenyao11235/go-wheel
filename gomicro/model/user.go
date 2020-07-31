package model

//User 用户
type User struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Password string `json:"password" gorm:"column:password"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	commonModel
}
