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
