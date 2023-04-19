package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// Comments 评价表
type Comments struct {
	global.GVA_MODEL
	JobFairsId      int64  `json:"job_fairs_id" gorm:"job_fairs_id" db:"job_fairs_id"`             // 招聘信息ID
	CompanyName     string `json:"company_name" gorm:"company_name" db:"company_name"`             // 公司名称
	StuComments     string `json:"stu_comments" gorm:"stu_comments" db:"stu_comments"`             // 学生评价
	TeacherComments string `json:"teacher_comments" gorm:"teacher_comments" db:"teacher_comments"` // 老师评价
}

// TableName 表名称
func (*Comments) TableName() string {
	return "comments"
}
