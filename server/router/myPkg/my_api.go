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
		MyRouterGroup.POST("v1/students/list/getAll", myApi.GetStudentsList)          // 获取毕业生信息列表
		MyRouterGroup.POST("v1/students/detail/getByStuNumber", myApi.GetByStuNumber) // 获取学生信息详情
		MyRouterGroup.POST("v1/students/update/setStudentInfo", myApi.SetStudentInfo) // 编辑学生信息
		/*-----------------------------------------------------------------------------------------*/
		MyRouterGroup.POST("v1/students/list/getByConditions", myApi.GetStudentsListByConditions) // 根据条件获取毕业生信息列表
		//MyRouterGroup.POST("v1/students/details/get", myApi.GetStudentsDetails)                   // 查看毕业生就业详情
		MyRouterGroup.POST("v1/students/update", myApi.UpdStudentsInfos)      // 编辑毕业生信息
		MyRouterGroup.DELETE("v1/students/delete", myApi.DeleteStudentsInfos) // 删除毕业生信息
		// 就业信息统计
		MyRouterGroup.POST("v1/college/employment/list/getAll", myApi.GetEmploymentInfosList)                      // 获取各学院的就业情况列表
		MyRouterGroup.POST("v1/college/employment/list/getByConditions", myApi.GetEmploymentInfosListByConditions) // 根据条件获取各学院的就业信息
		MyRouterGroup.POST("v1/college/employment/details/get", myApi.GetEmploymentDetails)                        // 查看学院就业情况详情
		MyRouterGroup.POST("v1/college/employment/update", myApi.UpdEmploymentInfos)                               // 更新学院就业情况
		MyRouterGroup.DELETE("v1/college/employment/delete", myApi.DeleteEmploymentInfos)                          // 删除学院就业信息
		MyRouterGroup.POST("v1/college/students/employed/details/get", myApi.GetEmployedStudentInfos)              // 获取签约学生列表详情
		MyRouterGroup.POST("v1/college/students/unemployed/details/get", myApi.GetUnemployedStudentInfos)          // 获取未签约学生列表详情
	}
}
