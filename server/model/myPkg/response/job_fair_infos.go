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
