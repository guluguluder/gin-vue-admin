package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// StudentJobs 学生就业工作表
type StudentJobs struct {
	global.GVA_MODEL
	StuNumber   string `json:"stuNumber" gorm:"stu_number" db:"stu_number"`       // 学号
	CompanyName string `json:"companyName" gorm:"company_name" db:"company_name"` // 公司名称
	JobCity     string `json:"jobCity" gorm:"job_city" db:"job_city"`             // 工作城市
	JobTitle    string `json:"jobTitle" gorm:"job_title" db:"job_title"`          // 工作职位
	JobSalary   string `json:"jobSalary" gorm:"job_salary" db:"job_salary"`       // 工资薪水
}

// TableName 表名称
func (*StudentJobs) TableName() string {
	return "student_jobs"
}
