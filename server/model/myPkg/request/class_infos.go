package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type SearchClass struct {
	request.PageInfo
	ClassNumber   string `json:"classNumber"`
	CollegeNumber string `json:"collegeNumber"`
}
type SearchClassDetails struct {
	request.PageInfo
	ClassNumber string `json:"classNumber"`
}
type SearchCollegeDetails struct {
	request.PageInfo
	CollegeNumber string `json:"collegeNumber"`
}
