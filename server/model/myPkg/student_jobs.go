package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// StudentJobs 学生就业工作表
type StudentJobs struct {
	global.GVA_MODEL
	StuNumber   string `json:"stu_number" gorm:"stu_number" db:"stu_number"`       // 学号
	CompanyName string `json:"company_name" gorm:"company_name" db:"company_name"` // 公司名称
	JobCity     string `json:"job_city" gorm:"job_city" db:"job_city"`             // 工作城市
	JobTitle    string `json:"job_title" gorm:"job_title" db:"job_title"`          // 工作职位
	JobSalary   string `json:"job_salary" gorm:"job_salary" db:"job_salary"`       // 工资薪水
}

// TableName 表名称
func (*StudentJobs) TableName() string {
	return "student_jobs"
}
