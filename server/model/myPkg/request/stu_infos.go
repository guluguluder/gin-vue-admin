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

type UpdStudentsInfos struct {
	StudentName string `json:"studentName"`
	CollegeName string `json:"collegeName"`
	ClassNum    string `json:"classNum"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	IsEmployed  string `json:"isEmployed"`
}
