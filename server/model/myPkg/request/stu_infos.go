package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type GetStudentsByConditions struct {
	PageInfo    request.PageInfo
	ClassNumber string `json:"classNumber"`
	StuNumber   string `json:"stuNumber"`
	IsEmployed  string `json:"isEmployed"`
	CollegeName string `json:"collegeName"`
}

type GetStudentsDetails struct {
	PageInfo  request.PageInfo
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
type UpdStudentsInfos struct {
	StudentName string `json:"studentName"`
	CollegeName string `json:"collegeName"`
	ClassNum    string `json:"classNum"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	IsEmployed  string `json:"isEmployed"`
}
