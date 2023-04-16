package myPkg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	r "github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/response"
)

type MyApiService struct {
}

//获取毕业生信息列表
func (m *MyApiService) GetStudentsListResp(reqInfo request.PageInfo, sysId uint) (list []response.StudentsList, total int64, err error) {

	// TODO:身份权限校验

	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	totalSQL := "select count(*) from( select ssi.stu_num StudentNum, ssi.stu_name StudentName, scci.college_num CollegeNum, scci.college_name CollegeName, ssi.stu_class_num StudentClassNum, su.phone Phone, su.email Email, ssi.employed Employed from sys_stu_infos ssi left join sys_users su on su.id = ssi.sys_user_id left join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where ssi.deleted_at is null order by ssi.stu_num) t "
	if err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error; err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}

	sql := "select ssi.stu_num StudentNum, ssi.stu_name StudentName, scci.college_num CollegeNum, scci.college_name CollegeName, ssi.stu_class_num StudentClassNum, su.phone Phone, su.email Email, ssi.employed Employed from sys_stu_infos ssi left join sys_users su on su.id = ssi.sys_user_id left join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where ssi.deleted_at is null order by ssi.stu_num limit ? ,?"
	if err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error; err != nil {
		return list, total, err
	}
	return list, total, nil
}

// 根据条件获取毕业生信息列表
func (m *MyApiService) GetStudentsListByConditionsResp(reqInfo r.GetStudentsByConditions, sysId uint) (list []response.StudentsList, total int64, err error) {

	limit := reqInfo.PageInfo.PageSize
	offset := reqInfo.PageInfo.PageSize * (reqInfo.PageInfo.Page - 1)
	whereStr := ""
	if reqInfo.ClassNumber != "" {
		whereStr = fmt.Sprintf("'%v' ssi.stu_class_num = '%v' and ", whereStr, reqInfo.ClassNumber)
	} else if reqInfo.StuNumber != "" {
		whereStr = fmt.Sprintf("'%v' ssi.stu_num = '%v' and ", whereStr, reqInfo.StuNumber)
	} else if reqInfo.IsEmployed != "" {
		whereStr = fmt.Sprintf(" '%v' ssi.employed = '%v' and ", whereStr, reqInfo.IsEmployed)
	} else if reqInfo.IsEmployed != "" {
		whereStr = fmt.Sprintf("'%v' scci.college_name = '%v' and ", whereStr, reqInfo.CollegeName)
	}

	totalSQL := " select count(*) from sys_stu_infos ssi inner join sys_users su on su.id = ssi.sys_user_id inner join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where " + whereStr + " ssi.deleted_at is null  order by ssi.stu_num"
	if err = global.GVA_DB.Raw(totalSQL).Error; err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}
	sql := "select ssi.stu_num StudentNum, ssi.stu_name StudentName, scci.college_num CollegeNum, scci.college_name CollegeName, ssi.stu_class_num StudentClassNum, su.phone Phone, su.email Email, ssi.employed Employed from sys_stu_infos ssi inner join sys_users su on su.id = ssi.sys_user_id inner join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where " + whereStr + " ssi.deleted_at is null  order by ssi.stu_num limit ?,?"
	err = global.GVA_DB.Raw(sql, limit, offset).Scan(&limit).Error
	if err != nil {
		return list, total, err
	}
	fmt.Println("hello World")
	return list, total, nil
}

// 查看毕业生就业详情
func (m *MyApiService) GetStudentsDetailsResp(reqInfo r.GetStudentsDetails, sysId uint) (info response.StudentDetails, err error) {

	sql := "SELECT ssji.stu_num StudentNum, ssi.stu_name StudentName, ssji.company_name CompanyName, ssji.job_city JobCity, ssji.job_title JobTitle, ssji.job_salary JobSalary FROM sys_stu_job_infos ssji INNER JOIN sys_stu_infos ssi ON ssi.stu_num = ssji.stu_num WHERE ssji.stu_num = ? "
	err = global.GVA_DB.Raw(sql, reqInfo.StuNumber).Scan(&info).Error
	if err != nil {
		return info, err
	}
	return info, err
}

// 编辑毕业生信息
func (m *MyApiService) UpdStudentsInfosResp(reqInfo r.UpdStudentsInfos, sysId uint) {

	fmt.Println("hello World")
}

// 删除毕业生信息
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
