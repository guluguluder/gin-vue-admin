package response

type StudentsList struct {
	ID          string `json:"ID"`
	StuNumber   string `json:"stuNumber"`
	StuName     string `json:"stuName"`
	StuSex      string `json:"stuSex"`
	ClassNumber string `json:"classNumber"`
	GradeNumber string `json:"gradeNumber"`
	CollegeName string `json:"collegeName"`
	StarTime    string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type StudentDetails struct {
	ID          string `json:"ID"`
	StuNumber   string `json:"stuNumber"`
	StuName     string `json:"stuName"`
	StuSex      string `json:"stuSex"`
	ClassNumber string `json:"classNumber"`
	GradeNumber string `json:"gradeNumber"`
	CollegeName string `json:"collegeName"`
	StarTime    string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type CollegeList struct {
	CollegeNumber string `json:"collegeNumber"`
	CollegeName   string `json:"collegeName"`
}
