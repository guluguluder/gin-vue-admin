package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Teachers 教师表
type Teachers struct {
	global.GVA_MODEL
	SysUserId        int64  `json:"sys_user_id" gorm:"sys_user_id" db:"sys_user_id"`                   // 用户id
	WorkerNumber     string `json:"worker_number" gorm:"worker_number" db:"worker_number"`             // 工号
	WorkerName       string `json:"worker_name" gorm:"worker_name" db:"worker_name"`                   // 教师姓名
	WorkerSex        string `json:"worker_sex" gorm:"worker_sex" db:"worker_sex"`                      // 性别
	WorkerEducation  string `json:"worker_education" gorm:"worker_education" db:"worker_education"`    // 学历
	WorkerUniversity string `json:"worker_university" gorm:"worker_university" db:"worker_university"` // 毕业学校
	CollegeNumber    string `json:"college_number" gorm:"college_number" db:"college_number"`          // 学院编号
}

// TableName 表名称
func (*Teachers) TableName() string {
	return "teachers"
}
