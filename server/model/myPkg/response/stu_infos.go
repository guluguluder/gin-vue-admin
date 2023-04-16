package response

type StudentsList struct {
	StudentNum      string `json:"studentNum"`
	StudentName     string `json:"studentName"`
	CollegeNum      string `json:"collegeNum"`
	CollegeName     string `json:"collegeName"`
	StudentClassNum string `json:"studentClassNum"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Employed        string `json:"employed"`
}

type StudentDetails struct {
	StudentName string `json:"studentName"`
	StudentNum  string `json:"studentNum"`
	CompanyName string `json:"companyName"`
	JobCity     string `json:"jobCity"`
	JobTitle    string `json:"jobTitle"`  // 工作职位
	JobSalary   string `json:"jobSalary"` // 工作薪资
}
