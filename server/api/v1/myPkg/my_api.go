package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	r "github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

// 获取毕业生信息列表
func (m *MyApi) GetStudentsList(c *gin.Context) {
	var reqInfo request.PageInfo
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	lists, total, err := myApiService.GetStudentsListResp(reqInfo, Id)
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
			Page:     reqInfo.Page,
			PageSize: reqInfo.PageSize,
		}, "ok", c)
	}
}

// 根据条件获取毕业生信息列表
func (m *MyApi) GetStudentsListByConditions(c *gin.Context) {
	var reqInfo r.GetStudentsByConditions
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	lists, total, err := myApiService.GetStudentsListByConditionsResp(reqInfo, Id)
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
func (m *MyApi) GetStudentsDetails(c *gin.Context) {
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

// 编辑毕业生信息
func (m *MyApi) UpdStudentsInfos(c *gin.Context) {

	var reqInfo r.UpdStudentsInfos
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	myApiService.UpdStudentsInfosResp(reqInfo, Id)
	response.Ok(c)
}

// 删除毕业生信息
func (m *MyApi) DeleteStudentsInfos(c *gin.Context) {
	myApiService.DeleteStudentsInfosResp()
	response.Ok(c)
}

// 获取各学院的就业情况列表
func (m *MyApi) GetEmploymentInfosList(c *gin.Context) {
	var reqInfo request.PageInfo
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		return
	}
	Id := utils.GetUserID(c)
	infos, total, err := myApiService.GetEmploymentInfosListResp(reqInfo, Id)
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

// 根据条件获取各学院的就业信息
func (m *MyApi) GetEmploymentInfosListByConditions(c *gin.Context) {
	myApiService.GetEmploymentInfosListByConditionsResp()
	response.Ok(c)
}

// 查看学院就业情况详情
func (m *MyApi) GetEmploymentDetails(c *gin.Context) {
	myApiService.GetEmploymentDetailsResp()
	response.Ok(c)
}

// 更新学院就业情况
func (m *MyApi) UpdEmploymentInfos(c *gin.Context) {
	myApiService.UpdEmploymentInfosResp()
	response.Ok(c)
}

// 删除学院就业信息
func (m *MyApi) DeleteEmploymentInfos(c *gin.Context) {
	myApiService.DeleteEmploymentInfosResp()
	response.Ok(c)
}

// 获取签约学生列表详情
func (m *MyApi) GetEmployedStudentInfos(c *gin.Context) {
	myApiService.GetEmployedStudentInfosResp()
	response.Ok(c)
}

// 获取未签约学生列表详情
func (m *MyApi) GetUnemployedStudentInfos(c *gin.Context) {
	myApiService.GetUnemployedStudentInfosResp()
	response.Ok(c)
}
