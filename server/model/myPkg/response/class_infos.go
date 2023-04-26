package response

type ClassList struct {
	ID              string `json:"ID"`
	CollegeName     string `json:"collegeName"`
	ClassNumber     string `json:"classNumber"`
	TotalStu        int64  `json:"totalStu"`
	EmployedSum     int64  `json:"employedSum"`
	EmployedPercent string `json:"employedPercent"`
}

type ClassDetails struct {
	StuNumber   string `json:"stuNumber"`
	StuName     string `json:"stuName"`
	ClassNumber string `json:"classNumber"`
	Employed    string `json:"employed"`
	CompanyName string `json:"companyName"`
}
