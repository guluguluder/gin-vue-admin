package myPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyApiRouter struct {
}

func (m *MyApiRouter) InitMyApiRouter(Router *gin.RouterGroup) {
	MyRouterGroup := Router.Group("myapp")
	myApi := v1.ApiGroupApp.PkgApiGroup.MyApi
	{
		// 学生管理
		MyRouterGroup.POST("v1/students/list/getAll", myApi.GetStudentsList)                   // 获取毕业生信息列表
		MyRouterGroup.POST("v1/students/detail/getByStuNumber", myApi.GetByStuNumber)          // 获取学生信息详情
		MyRouterGroup.POST("v1/students/update/setStudentInfo", myApi.SetStudentInfo)          // 编辑学生信息
		MyRouterGroup.POST("v1/students/delete/deleteStudent", myApi.DeleteStudent)            // 删除学生
		MyRouterGroup.GET("v1/students/get/getColleges", myApi.GetColleges)                    // 获取下拉列表学院
		MyRouterGroup.POST("v1/students/get/getEmployedList", myApi.GetEmployedList)           // 获取学生就业信息列表
		MyRouterGroup.POST("v1/students/details/getEmployedDetails", myApi.GetEmployedDetails) // 获取学生就业详情
		MyRouterGroup.POST("v1/students/update/setEmployedDetails", myApi.SetEmployedDetails)  // 编辑学生就业信息

		// 班级管理
		MyRouterGroup.POST("v1/class/get/getClassList", myApi.GetClassList)                       // 获取班级就业信息列表
		MyRouterGroup.POST("v1/class/get/getClassEmployedDetails", myApi.GetClassEmployedDetails) // 获取班级就业信息详情
		// 招聘会管理
		MyRouterGroup.POST("v1/jobFairs/get/getJobFairList", myApi.GetJobFairList)    // 获取招聘会信息列表
		MyRouterGroup.DELETE("v1/jobFairs/delete/deleteJobFair", myApi.DeleteJobFair) // 删除招聘会信息
		MyRouterGroup.POST("v1/jobFairs/add/addJobFair", myApi.AddJobFair)            // 添加招聘会信息
		MyRouterGroup.POST("v1/jobFairs/upd/setJobFair", myApi.SetJobFair)            // 编辑招聘会信息
		MyRouterGroup.POST("v1/jobFairs/upd/addCommentInfo", myApi.AddCommentInfo)    // 评价
		//评价管理
		MyRouterGroup.POST("v1/jobFairs/get/getCommentList", myApi.GetCommentList) //评价列表
		//公告管理
		MyRouterGroup.POST("v1/content/list/getAll", myApi.GetContentList)                     // 获取公告列表
		MyRouterGroup.POST("v1/content/upd/setContent", myApi.SetContent)                      // 编辑公告信息
		MyRouterGroup.POST("v1/content/add/addContent", myApi.AddContent)                      // 添加公告信息
		MyRouterGroup.DELETE("v1/content/del/deleteContent", myApi.DeleteContent)              // 添加公告信息
		MyRouterGroup.POST("v1/infos/getEmploymentInfos", myApi.GetEmploymentInfos)            // 获取各学院的就业情况列表
		MyRouterGroup.POST("v1/infos/getCollegeEmployedDetails", myApi.GetCollegeEmployedList) // 获取各学院的就业情况详情列表
		// 就业总结
		MyRouterGroup.POST("v1/infos/addSummary", myApi.AddSummary) // 获取各学院的就业情况详情列表

		/*-----------------------------------------------------------------------------------------*/
		//MyRouterGroup.POST("v1/students/list/getByConditions", myApi.GetStudentsListByConditions) // 根据条件获取毕业生信息列表
		//MyRouterGroup.POST("v1/students/details/get", myApi.GetStudentsDetails)                   // 查看毕业生就业详情
		//MyRouterGroup.POST("v1/students/update", myApi.UpdStudentsInfos)      // 编辑毕业生信息
		MyRouterGroup.DELETE("v1/students/delete", myApi.DeleteStudentsInfos) // 删除毕业生信息
		// 就业信息统计
		MyRouterGroup.POST("v1/college/employment/list/getByConditions", myApi.GetEmploymentInfosListByConditions) // 根据条件获取各学院的就业信息
		MyRouterGroup.POST("v1/college/employment/details/get", myApi.GetEmploymentDetails)                        // 查看学院就业情况详情
		MyRouterGroup.POST("v1/college/employment/update", myApi.UpdEmploymentInfos)                               // 更新学院就业情况
		MyRouterGroup.DELETE("v1/college/employment/delete", myApi.DeleteEmploymentInfos)                          // 删除学院就业信息
		MyRouterGroup.POST("v1/college/students/employed/details/get", myApi.GetEmployedStudentInfos)              // 获取签约学生列表详情
		MyRouterGroup.POST("v1/college/students/unemployed/details/get", myApi.GetUnemployedStudentInfos)          // 获取未签约学生列表详情
	}
}
