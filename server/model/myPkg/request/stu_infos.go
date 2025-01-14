package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type SearchStu struct {
	request.PageInfo
	ClassNumber   string `json:"classNumber"`
	StuNumber     string `json:"stuNumber"`
	CollegeNumber string `json:"collegeNumber"`
	IsEmployed    string `json:"isEmployed"`
}

type GetStudentsDetails struct {
	request.PageInfo
	StuNumber string `json:"stuNumber"`
}

type UpdStudentsReq struct {
	ID          string `json:"ID"`
	StuNumber   string `json:"stuNumber"`
	StuName     string `json:"stuName"`
	StuSex      string `json:"stuSex"`
	ClassNumber string `json:"classNumber"`
	GradeNumber string `json:"gradeNumber"`
	CollegeName string `json:"collegeName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type DelStudentReq struct {
	ID string `json:"ID" form:"id"` // 主键
}

type UpdEmployReq struct {
	ID          string `json:"ID"`
	StuNumber   string `json:"stuNumber"`
	StuName     string `json:"stuName"`
	IsEmployed  string `json:"isEmployed"`
	CompanyName string `json:"companyName"`
	JobCity     string `json:"jobCity"`
	JobTitle    string `json:"jobTitle"`
	JobSalary   string `json:"jobSalary"`
}

/*================================================================*/
