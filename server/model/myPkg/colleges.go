package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Colleges 学院表
type Colleges struct {
	global.GVA_MODEL
	CollegeNumber string `json:"college_number" gorm:"college_number" db:"college_number"` // 学院编号
	CollegeName   string `json:"college_name" gorm:"college_name" db:"college_name"`       // 学院名称
	TotalStu      int64  `json:"total_stu" gorm:"total_stu" db:"total_stu"`                // 学院总人数
}

// TableName 表名称
func (*Colleges) TableName() string {
	return "colleges"
}
