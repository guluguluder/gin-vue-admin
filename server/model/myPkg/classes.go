package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Classes 班级表
type Classes struct {
	global.GVA_MODEL
	ClassNumber   string `json:"class_number" gorm:"class_number" db:"class_number"`       // 班级编号
	CollegeNumber string `json:"college_number" gorm:"college_number" db:"college_number"` // 学院编号
	TotalStu      int64  `json:"total_stu" gorm:"total_stu" db:"total_stu"`                // 班级人数
}

// TableName 表名称
func (*Classes) TableName() string {
	return "classes"
}
