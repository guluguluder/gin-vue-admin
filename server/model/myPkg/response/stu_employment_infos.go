package response

type StuEmploymentInfos struct {
	CollegeNumber string `json:"collegeNumber"`
	CollegeName   string `json:"collegeName"`
	TotalStudents int64  `json:"totalStudents"`
	EmployedSum   int64  `json:"employedSum"`
	UnemployedSum int64  `json:"unemployedSum"`
	Percentage    string `json:"percentage"`
}
