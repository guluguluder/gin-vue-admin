package response

type JobFairList struct {
	ID          string `json:"ID"`
	CompanyName string `json:"companyName"`
	City        string `json:"city"`
	Salary      string `json:"salary"`
	TotalStu    int64  `json:"totalStu"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
type ContentList struct {
	ID         string `json:"ID"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	UpdateTime string `json:"updateTime"`
}
type CommentList struct {
	ID             string `json:"ID"`
	CompanyName    string `json:"companyName"`
	StudentComment string `json:"studentComment"`
	TeacherComment string `json:"teacherComment"`
	CreatedAt      string `json:"createdAt"`
}
