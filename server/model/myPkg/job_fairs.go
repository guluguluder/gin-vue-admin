package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// JobFairs 招聘信息表
type JobFairs struct {
	global.GVA_MODEL
	NoticeId    uint   `json:"notice_id" gorm:"notice_id" db:"notice_id"`          // 公告Id
	CompanyName string `json:"company_name" gorm:"company_name" db:"company_name"` // 公司名称
	City        string `json:"city" gorm:"city" db:"city"`                         // 工作城市
	Salary      string `json:"salary" gorm:"salary" db:"salary"`                   // 薪水
	TotalStu    int64  `json:"total_stu" gorm:"total_stu" db:"total_stu"`          // 招聘人数
	Telephone   string `json:"telephone" gorm:"telephone" db:"telephone"`          // 电话
	Email       string `json:"email" gorm:"email" db:"email"`                      // 邮箱
	Address     string `json:"address" gorm:"address" db:"address"`                // 招聘会地址
}

// TableName 表名称
func (*JobFairs) TableName() string {
	return "job_fairs"
}
