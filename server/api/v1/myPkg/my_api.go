package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	r "github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

// 获取毕业生信息列表
func (m *MyApi) GetStudentsList(c *gin.Context) {
	var reqInfo r.SearchStu
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	AuthorityId := utils.GetUserAuthorityId(c)
	lists, total, err := myApiService.GetStudentsListResp(reqInfo, Id, AuthorityId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if total == 0 {
		response.OkWithMessage("there is no data", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     lists,
			Total:    total,
			Page:     reqInfo.PageInfo.Page,
			PageSize: reqInfo.PageInfo.PageSize,
		}, "ok", c)
	}
}

// 查看毕业生就业详情
func (m *MyApi) GetByStuNumber(c *gin.Context) {
	var reqInfo r.GetStudentsDetails
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	list, err := myApiService.GetStudentsDetailsResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(list, "ok", c)
}

// 编辑学生信息
func (m *MyApi) SetStudentInfo(c *gin.Context) {
	var reqInfo r.UpdStudentsReq
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	//TODO:身份权限校验
	Id := utils.GetUserID(c)
	err = myApiService.SetStudentInfoResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("操作成功", c)
	}

}

// 删除学生
func (m *MyApi) DeleteStudent(c *gin.Context) {
	var reqId r.DelStudentReq
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.DeleteStudentResp(reqId, Id)
	if err != nil {
		response.FailWithMessage("操作失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// 获取下拉列表学院
func (m *MyApi) GetColleges(c *gin.Context) {
	Id := utils.GetUserID(c)
	list, err := myApiService.GetCollegesResp(Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(list, "ok", c)
	}
}

// 获取学生就业信息列表
func (m *MyApi) GetEmployedList(c *gin.Context) {

	var reqInfo r.SearchStu
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}

	Id := utils.GetUserID(c)
	AuthorityId := utils.GetUserAuthorityId(c)
	list, total, err := myApiService.GetEmployedListResp(reqInfo, Id, AuthorityId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

// 查看毕业生就业详情
func (m *MyApi) GetEmployedDetails(c *gin.Context) {
	var reqInfo r.GetStudentsDetails
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	list, err := myApiService.GetEmployedDetailsResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(list, "ok", c)
}

// 编辑学生就业信息
func (m *MyApi) SetEmployedDetails(c *gin.Context) {
	var reqInfo r.UpdEmployReq
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	if reqInfo.IsEmployed == "" {
		response.OkWithMessage("ok", c)
		return
	}
	err = myApiService.SetEmployedDetailsResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// 获取班级列表
func (m *MyApi) GetClassList(c *gin.Context) {
	var reqInfo r.SearchClass
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	list, total, err := myApiService.GetClassListResp(reqInfo, Id)
	if err != nil {
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

//获取班级就业信息详情
func (m *MyApi) GetClassEmployedDetails(c *gin.Context) {

	var reqInfo r.SearchClassDetails
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	list, total, err := myApiService.GetClassEmployedDetailsResp(reqInfo, Id)
	if err != nil {
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

//获取招聘会信息列表
func (m *MyApi) GetJobFairList(c *gin.Context) {
	var reqInfo r.SearchJobFairs
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	list, total, err := myApiService.GetJobFairListResp(reqInfo, Id)
	if err != nil {
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

//删除招聘会信息
func (m *MyApi) DeleteJobFair(c *gin.Context) {
	var reqInfo r.SearchJobFairs
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.DeleteJobFair(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

//添加招聘会信息
func (m *MyApi) AddJobFair(c *gin.Context) {
	var reqInfo r.AddJobFair
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.AddJobFairResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

//编辑招聘会信息
func (m *MyApi) SetJobFair(c *gin.Context) {
	var reqInfo r.AddJobFair
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.SetJobFairResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("编辑成功", c)
	}

}

//评价
func (m *MyApi) AddCommentInfo(c *gin.Context) {
	var reqInfo r.AddComment
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	AuthorityId := utils.GetUserAuthorityId(c)
	err = myApiService.AddCommentInfoResp(reqInfo, Id, AuthorityId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

//评价列表
func (m *MyApi) GetCommentList(c *gin.Context) {
	var reqInfo r.SearchComment
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	AuthorityId := utils.GetUserAuthorityId(c)
	list, total, err := myApiService.GetCommentListResp(reqInfo, Id, AuthorityId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

//获取公告列表
func (m *MyApi) GetContentList(c *gin.Context) {
	var reqInfo r.SearchContent
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	list, total, err := myApiService.GetContentListResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if total == 0 {
		response.OkWithMessage("暂无数据", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}

}

func (m *MyApi) SetContent(c *gin.Context) {
	var reqInfo r.UpdContent
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.SetContentResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}
func (m *MyApi) AddContent(c *gin.Context) {
	var reqInfo r.AddContent
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.AddContentResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}
func (m *MyApi) DeleteContent(c *gin.Context) {
	var reqInfo r.DelContent
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	err = myApiService.DelContentResp(reqInfo, Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// 获取各学院的就业情况列表
func (m *MyApi) GetEmploymentInfos(c *gin.Context) {
	var reqInfo r.SearchCollegeDetails
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	infos, total, err := myApiService.GetEmploymentInfosResp(reqInfo, Id)
	if err != nil {
		return
	} else if total == 0 {
		response.OkWithMessage("there is no data", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
			Total:    total,
			List:     infos,
		}, "ok", c)
	}
}

// 获取各学院的就业情况列表
func (m *MyApi) GetCollegeEmployedList(c *gin.Context) {
	var reqInfo r.SearchCollegeDetails
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	infos, total, err := myApiService.GetCollegeEmployedListResp(reqInfo, Id)
	if err != nil {
		return
	} else if total == 0 {
		response.OkWithMessage("there is no data", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
			Total:    total,
			List:     infos,
		}, "ok", c)
	}
}
func (m *MyApi) AddSummary(c *gin.Context) {
	var reqInfo r.AddSummary
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Id := utils.GetUserID(c)
	AuthorityId := utils.GetUserAuthorityId(c)
	err = myApiService.AddSummaryResp(reqInfo, Id, AuthorityId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

/*-----------------------------------------------------------------------*/
