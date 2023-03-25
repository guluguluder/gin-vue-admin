package myPkg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/response"
)

type MyApiService struct {
}

func (m *MyApiService) GetStudentsListResp(reqInfo request.PageInfo, sysId uint) {

	fmt.Println("hello World")
}
func (m *MyApiService) GetStudentsListByConditionsResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) GetStudentsDetailsResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) UpdStudentsInfosResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) DeleteStudentsInfosResp() {
	fmt.Println("hello World")
}

// 获取各学院的就业情况列表
func (m *MyApiService) GetEmploymentInfosListResp(reqInfo request.PageInfo, sysId uint) (list []response.StuEmploymentInfos, total int64, err error) {

	//TODO: 校验身份权限

	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	totalSQL := "select count(*) from( select scci.college_num , scci.college_name , ifnull(sum(scci.total_students),0) TotalStudents , ifnull(sum(scci.employed_sum),0) EmployedSum, ifnull(sum(scci.unemployed_sum),0) UnemployedSum, case when ifnull(sum(scci.total_students),0) != 0 then concat(sum(scci.employed_sum)/sum(scci.total_students)*100,'%') else '0' end AS Percentage from sys_class_college_infos scci group by scci.college_num, scci.college_name) t "
	err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error
	if err != nil {
		return list, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}

	sql := "select scci.college_num CollegeNum, scci.college_name CollegeName, ifnull(sum(scci.total_students),0) TotalStudents , ifnull(sum(scci.employed_sum),0) EmployedSum, ifnull(sum(scci.unemployed_sum),0) UnemployedSum, case when ifnull(sum(scci.total_students),0) != 0 then concat(sum(scci.employed_sum)/sum(scci.total_students)*100,'%') else '0' end AS Percentage from sys_class_college_infos scci group by scci.college_num, scci.college_name limit ?,?"
	err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error
	if err != nil {
		return list, total, err
	}
	return list, total, nil
}
func (m *MyApiService) GetEmploymentInfosListByConditionsResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) GetEmploymentDetailsResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) UpdEmploymentInfosResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) DeleteEmploymentInfosResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) GetEmployedStudentInfosResp() {
	fmt.Println("hello World")
}
func (m *MyApiService) GetUnemployedStudentInfosResp() {
	fmt.Println("hello World")
}
