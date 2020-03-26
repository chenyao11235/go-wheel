package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	Model
	Name string
	Age  int
	Code string
}

func (User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("Code", "random string")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func main() {
	db, err := gorm.Open("mysql", "root:cy.89757@tcp(10.211.55.5:3306)/dispatcher?charset=utf8&parseTime=true")
	if err != nil {
		panic("链接失败")
	}

	defer db.Close()

	db.SingularTable(true)

	db.AutoMigrate(&User{})

	var users []User

	//db.Where("id in(?)", []int{1, 3, 4}).Find(&users)
	db.Where(map[string]interface{}{"name": "eric"}).Find(&users)

	if len(users) != 0 {
		fmt.Println(users[0].ID)
	}
	//user := User{Name: "eric", Age: 0}

	//db.NewRecord(user)

	//db.Create(&user)

	//fmt.Println(db.NewRecord(user))

	//fmt.Println(user.ID)
}
