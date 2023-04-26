package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type SearchJobFairs struct {
	request.PageInfo
	ID string `json:"ID"`
}
type AddJobFair struct {
	ID          string `json:"ID"`
	NoticeId    string `json:"noticeId"`
	CompanyName string `json:"companyName"`
	City        string `json:"city"`
	Salary      string `json:"salary"`
	TotalStu    int64  `json:"totalStu"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
