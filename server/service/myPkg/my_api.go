package myPkg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	r "github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/response"
	"gorm.io/gorm"
)

type MyApiService struct {
}

//获取毕业生信息列表
func (m *MyApiService) GetStudentsListResp(reqInfo request.PageInfo, sysId uint) (list []response.StudentsList, total int64, err error) {

	// TODO:身份权限校验

	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	totalSQL := "select count(*) from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number left join colleges c2 on c.college_number = c2.college_number "
	if err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error; err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}

	sql := "select s.id ID, s.stu_number StuNumber, s.stu_name StuName, s.stu_sex StuSex, s.class_number ClassNumber , s.grade_number GradeNumber , c2.college_name CollegeName, su.phone Phone, su.email Email from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number left join colleges c2 on c.college_number = c2.college_number limit ? ,?"
	if err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error; err != nil {
		return list, total, err
	}
	return list, total, nil
}

// 查看毕业生就业详情
func (m *MyApiService) GetStudentsDetailsResp(reqInfo r.GetStudentsDetails, sysId uint) (info response.StudentDetails, err error) {

	sql := "select s.id ID, s.stu_number StuNumber, s.stu_name StuName, s.stu_sex StuSex, s.class_number ClassNumber , s.grade_number GradeNumber , s.start_time StarTime, s.end_time EndTime, c2.college_name CollegeName, su.phone Phone, su.email Email from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number left join colleges c2 on c.college_number = c2.college_number where s.stu_number = ?"
	err = global.GVA_DB.Raw(sql, reqInfo.StuNumber).Scan(&info).Error
	if err != nil {
		return info, err
	}
	return info, err
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

// 编辑毕业生信息
func (m *MyApiService) UpdStudentsInfosResp(reqInfo r.UpdStudentsInfos, sysId uint) {

	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		return nil
	})
	if err != nil {
		return
	}
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
