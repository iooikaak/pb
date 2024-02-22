package model

// wxbind表
type WxBind struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id"`
	UserID      int64  `gorm:"column:user_id" json:"user_id"`
	OpenID      string `gorm:"column:open_id" json:"open_id"`
	CreatedTime int64  `gorm:"column:created_time" json:"created_time"`
}

func (WxBind) TableName() string {
	return "wxbind"
}
