package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Students 学生表
type Students struct {
	global.GVA_MODEL
	SysStuId      int64  `json:"sys_stu_id" gorm:"sys_stu_id" db:"sys_stu_id"`             // 用户id
	StuNumber     string `json:"stu_number" gorm:"stu_number" db:"stu_number"`             // 学号
	StuName       string `json:"stu_name" gorm:"stu_name" db:"stu_name"`                   // 姓名
	StuSex        string `json:"stu_sex" gorm:"stu_sex" db:"stu_sex"`                      // 性别
	ClassNumber   string `json:"class_number" gorm:"class_number" db:"class_number"`       // 班级
	GradeNumber   string `json:"grade_number" gorm:"grade_number" db:"grade_number"`       // 年级
	CollegeNumber string `json:"college_number" gorm:"college_number" db:"college_number"` // 学院编号
	StartTime     string `json:"start_time" gorm:"start_time" db:"start_time"`             // 入学日期
	EndTime       string `json:"end_time" gorm:"end_time" db:"end_time"`                   // 毕业日期
	Employed      string `json:"employed" gorm:"employed" db:"employed"`                   // 是否签约 （y:已签约，n:未签约）
}

// TableName 表名称
func (*Students) TableName() string {
	return "students"
}
