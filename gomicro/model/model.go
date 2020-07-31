package model

// 公共字段
type commonModel struct {
	CreatedTime int `json:"created_time" gorm:"column:createdtime"`
	UpdatedTIme int `json:"updated_time" gorm:"column:updatedtime"`
	DeletedTime int `json:"deleted_time" gorm:"column:deletedtime"`
}
