package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Notices 公告表
type Notices struct {
	global.GVA_MODEL
	Content      string `json:"content" gorm:"content" db:"content"`                   // 公告内容
	AuthorNumber string `json:"author_number" gorm:"author_number" db:"author_number"` // 作者工号
}

// TableName 表名称
func (*Notices) TableName() string {
	return "notices"
}
