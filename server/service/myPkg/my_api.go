package myPkg

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myPkg"
	r "github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myPkg/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"time"
)

type MyApiService struct {
}

//获取毕业生信息列表
func (m *MyApiService) GetStudentsListResp(reqInfo r.SearchStu, sysId uint) (list []response.StudentsList, total int64, err error) {

	// TODO:身份权限校验

	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	whereStr := ""
	if reqInfo.ClassNumber != "" {
		whereStr = fmt.Sprintf("%v s.class_number = '%v' and ", whereStr, reqInfo.ClassNumber)
	}
	if reqInfo.StuNumber != "" {
		whereStr = fmt.Sprintf("%v s.stu_number = '%v' and ", whereStr, reqInfo.StuNumber)
	}
	if reqInfo.CollegeNumber != "" {
		whereStr = fmt.Sprintf("%v s.college_number = '%v' and ", whereStr, reqInfo.CollegeNumber)
	}

	totalSQL := "select count(*) from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number left join colleges c2 on c.college_number = c2.college_number where " + whereStr + " s.deleted_at is null "
	if err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error; err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}

	sql := "select s.id ID, s.stu_number StuNumber, s.stu_name StuName, s.stu_sex StuSex, s.class_number ClassNumber , s.grade_number GradeNumber , c2.college_name CollegeName, su.phone Phone, su.email Email from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number left join colleges c2 on c.college_number = c2.college_number where " + whereStr + " s.deleted_at is null limit ? ,?"
	if err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error; err != nil {
		return list, total, err
	}
	return list, total, nil
}

// 查看毕业生就业详情
func (m *MyApiService) GetStudentsDetailsResp(reqInfo r.GetStudentsDetails, sysId uint) (info response.StudentDetails, err error) {

	sql := "select s.id ID, s.stu_number StuNumber, s.stu_name StuName, s.stu_sex StuSex, s.class_number ClassNumber , s.grade_number GradeNumber , s.start_time StarTime, s.end_time EndTime, c2.college_name CollegeName, su.phone Phone, su.email Email from students s inner join sys_users su on s.sys_stu_id = su.id and su.deleted_at is null left join classes c on c.class_number = s.class_number and c.deleted_at is null left join colleges c2 on c.college_number = c2.college_number and c2.deleted_at is null where s.stu_number = ? and s.deleted_at is null"
	err = global.GVA_DB.Raw(sql, reqInfo.StuNumber).Scan(&info).Error
	if err != nil {
		return info, err
	}
	return info, err
}

// 编辑学生信息
func (m *MyApiService) SetStudentInfoResp(reqInfo r.UpdStudentsReq, sysId uint) error {
	// TODO:身份权限校验

	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var sysStuId string
		sql := "select sys_stu_id  from students s where s.id = ? and deleted_at is null"
		if err := tx.Raw(sql, reqInfo.ID).Scan(&sysStuId).Error; err != nil {
			return err
		}

		//更新sys_user表
		if err := tx.Model(&system.SysUser{}).Where("id = ? and deleted_at is null", sysStuId).Updates(system.SysUser{
			//Username: reqInfo.StuNumber,
			NickName: reqInfo.StuName,
			Phone:    reqInfo.Phone,
			Email:    reqInfo.Email,
		}).Error; err != nil {
			return err
		}
		// 校验学院是否存在
		sql = "select c.college_number  from colleges c where c.college_name = ? and c.deleted_at is null "
		var collegeNumber string
		if err := tx.Raw(sql, reqInfo.CollegeName).Scan(&collegeNumber).Error; err != nil {
			return err
		}
		if collegeNumber == "" {
			return errors.New("该学院不存在")
		}
		// 校验班级是否合法
		sql = "select  c.class_number  from classes c where c.college_number = ? and c.class_number = ? and c.deleted_at is null "
		var classNumber string
		if err := tx.Raw(sql, collegeNumber, reqInfo.ClassNumber).Scan(&classNumber).Error; err != nil {
			return err
		}
		if classNumber == "" {
			return errors.New("该班级不存在")
		}
		// 修改学院或班级需要对应的同步人数
		var data map[string]interface{}
		sql = "select s.class_number ,s.college_number  from students s  where s.id = ? and s.deleted_at is null "
		if err := tx.Raw(sql, reqInfo.ID).Scan(&data).Error; err != nil {
			return err
		}
		if data["class_number"] != reqInfo.ClassNumber {
			err := tx.Exec("update classes set total_stu = total_stu - 1 where class_number = ? and deleted_at is null", data["class_number"]).Error
			if err != nil {
				return err
			}
			err = tx.Exec("update classes set total_stu = total_stu + 1 where class_number = ? and deleted_at is null", reqInfo.ClassNumber).Error
			if err != nil {
				return err
			}
		}
		if data["college_number"] != collegeNumber {
			err := tx.Exec("update colleges set total_stu = total_stu - 1 where college_number = ? and deleted_at is null", data["college_number"]).Error
			if err != nil {
				return err
			}
			err = tx.Exec("update colleges set total_stu = total_stu + 1 where college_number = ? and deleted_at is null", collegeNumber).Error
			if err != nil {
				return err
			}
		}
		// 更新students表
		if err := tx.Model(&myPkg.Students{}).Where("sys_stu_id = ? and deleted_at is null", sysStuId).Updates(myPkg.Students{
			//StuNumber:     reqInfo.StuNumber,
			StuName:       reqInfo.StuName,
			StuSex:        reqInfo.StuSex,
			CollegeNumber: collegeNumber,
			ClassNumber:   reqInfo.ClassNumber,
			GradeNumber:   reqInfo.GradeNumber,
		}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 删除学生
func (m *MyApiService) DeleteStudentResp(reqId r.DelStudentReq, sysId uint) error {

	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// todo: 删除students表数据,同步修改classes 和 colleges 表的总人数
		var data map[string]interface{}
		sql := "select s.class_number ,s.college_number  from students s  where s.id = ? and s.deleted_at is null"
		if err := tx.Raw(sql, reqId.ID).Scan(&data).Error; err != nil {
			return err
		}
		// 更新班级表
		err := tx.Exec("update classes set total_stu = total_stu - 1 where class_number = ? and deleted_at is null", data["class_number"]).Error
		if err != nil {
			return err
		}
		// 更新学院表
		err = tx.Exec("update colleges set total_stu = total_stu - 1 where college_number = ? and deleted_at is null", data["college_number"]).Error
		if err != nil {
			return err
		}
		// 删除学生信息
		err = tx.Model(&myPkg.Students{}).Where("id = ? and deleted_at is null", reqId.ID).Update("deleted_at", time.Now()).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// 获取下拉列表学院
func (m *MyApiService) GetCollegesResp(sysId uint) (colleges []response.CollegeList, err error) {
	sql := "select c.college_number CollegeNumber ,c.college_name CollegeName from colleges c order by c.id"
	err = global.GVA_DB.Raw(sql).Scan(&colleges).Error
	if err != nil {
		return colleges, err
	}
	return colleges, nil
}

// 获取学生就业信息列表
func (m *MyApiService) GetEmployedListResp(reqInfo r.SearchStu, sysId uint) (list []response.EmployedInfos, total int64, err error) {

	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	whereStr := ""
	if reqInfo.StuNumber != "" {
		whereStr = fmt.Sprintf("%v s.stu_number = '%v' and ", whereStr, reqInfo.StuNumber)
	}
	if reqInfo.IsEmployed != "" {
		whereStr = fmt.Sprintf("%v s.employed = '%v' and ", whereStr, reqInfo.IsEmployed)
	}

	totalSQL := "select count(*) from students s left join student_jobs sj on sj.stu_number = s.stu_number where " + whereStr + " s.deleted_at is null"
	err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}

	sql := "select s.id ID, s.stu_number StuNumber , s.stu_name StuName, case when s.employed = 'Y' then '是' else '否' end as  IsEmployed, sj.company_name CompanyName, sj.job_city JobCity, sj.job_title JobTitle, sj.job_salary JobSalary from students s left join student_jobs sj on sj.stu_number = s.stu_number where " + whereStr + " s.deleted_at is null order by s.id limit ?,? "
	err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error
	if err != nil {
		return list, total, err
	}

	return list, total, nil
}

// 查看毕业生就业详情
func (m *MyApiService) GetEmployedDetailsResp(reqInfo r.GetStudentsDetails, sysId uint) (info response.EmployedDetails, err error) {

	sql := "select s.id ID, s.stu_number StuNumber , s.stu_name StuName, s.stu_sex StuSex, s.class_number ClassNumber , s.grade_number GradeNumber, c.college_name CollegeName, s.employed IsEmployed, sj.company_name CompanyName, sj.job_city JobCity, sj.job_title JobTitle, sj.job_salary JobSalary from students s inner join colleges c on c.college_number = s.college_number left join student_jobs sj on sj.stu_number = s.stu_number where s.stu_number = ? and s.deleted_at is null"
	err = global.GVA_DB.Raw(sql, reqInfo.StuNumber).Scan(&info).Error
	if err != nil {
		return info, err
	}
	return info, err
}

// 编辑毕业生就业详情
func (m *MyApiService) SetEmployedDetailsResp(reqInfo r.UpdEmployReq, sysId uint) (err error) {

	if reqInfo.IsEmployed == "是" {
		reqInfo.IsEmployed = "Y"
	}
	if reqInfo.IsEmployed == "否" {
		reqInfo.IsEmployed = "N"
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		var stuNumber string
		err = global.GVA_DB.Raw("select s.stu_number from students s where s.id = ?", reqInfo.ID).Scan(&stuNumber).Error
		if err != nil {
			return err
		}
		err = tx.Model(&myPkg.Students{}).Where("stu_number = ?", stuNumber).Update("employed", reqInfo.IsEmployed).Error
		if err != nil {
			return err
		}
		var stuStr string
		SQL := "select sj.stu_number  from student_jobs sj where sj.stu_number  = ? "
		err = tx.Raw(SQL, stuNumber).Scan(&stuStr).Error
		if err != nil {
			return err
		}
		if stuStr != "" {
			err = tx.Model(&myPkg.StudentJobs{}).Where("stu_number = ?", reqInfo.StuNumber).Updates(myPkg.StudentJobs{
				CompanyName: reqInfo.CompanyName,
				JobCity:     reqInfo.JobCity,
				JobTitle:    reqInfo.JobTitle,
				JobSalary:   reqInfo.JobSalary,
			}).Error
			if err != nil {
				return err
			}
		} else if stuStr == "" {
			err = tx.Create(&myPkg.StudentJobs{
				GVA_MODEL:   global.GVA_MODEL{},
				StuNumber:   stuNumber,
				CompanyName: reqInfo.CompanyName,
				JobCity:     reqInfo.JobCity,
				JobTitle:    reqInfo.JobTitle,
				JobSalary:   reqInfo.JobSalary,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

//获取班级列表
func (m *MyApiService) GetClassListResp(reqInfo r.SearchClass, sysId uint) (list []response.ClassList, total int64, err error) {
	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	whereStr := ""
	if reqInfo.ClassNumber != "" {
		whereStr = fmt.Sprintf("%v c.class_number = '%v' and ", whereStr, reqInfo.ClassNumber)
	}
	if reqInfo.CollegeNumber != "" {
		whereStr = fmt.Sprintf("%v c.college_number = '%v' and ", whereStr, reqInfo.CollegeNumber)
	}

	totalSQL := "select count(*) from classes c inner join colleges c2 on c2.college_number = c.college_number left join( select count(s.id) employed_stu, s.class_number class_number from students s where s.employed = 'Y' group by s.class_number) t on t.class_number = c.class_number where " + whereStr + " c.deleted_at is null order by c.id "
	err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}
	sql := "select c.id ID, c.class_number ClassNumber, c.college_number CollegeNumber, c.total_stu TotalStu, c2.college_name CollegeName, t.employed_stu EmployedSum ,concat( round(t.employed_stu/c.total_stu*100,2) ,'%') EmployedPercent from classes c inner join colleges c2 on c2.college_number = c.college_number left join( select count(s.id) employed_stu, s.class_number class_number from students s where s.employed = 'Y' group by s.class_number) t on t.class_number = c.class_number where " + whereStr + " c.deleted_at is null order by c.id limit ?,?"
	err = global.GVA_DB.Raw(sql, offset, limit).Scan(&list).Error
	if err != nil {
		return list, total, err
	}

	return list, total, nil
}

//获取班级就业信息详情
func (m *MyApiService) GetClassEmployedDetailsResp(reqInfo r.SearchClassDetails, sysId uint) (list []response.ClassDetails, total int64, err error) {
	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	totalSQL := " select count(*) from students s left join student_jobs sj on sj.stu_number = s.stu_number where s.class_number = ? and s.deleted_at is null"
	err = global.GVA_DB.Raw(totalSQL, reqInfo.ClassNumber).Scan(&total).Error
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}
	sql := " select s.stu_number StuNumber, s.stu_name StuName , s.class_number ClassNumber, s.employed Employed, sj.company_name CompanyName from students s left join student_jobs sj on sj.stu_number = s.stu_number where s.class_number = ? and s.deleted_at is null limit ?,? "
	err = global.GVA_DB.Raw(sql, reqInfo.ClassNumber, offset, limit).Scan(&list).Error
	if err != nil {
		return list, total, err
	}
	return list, total, nil
}

//获取招聘会信息列表
func (m *MyApiService) GetJobFairListResp(reqInfo r.SearchJobFairs, sysId uint) (list []response.JobFairList, total int64, err error) {
	limit := reqInfo.PageSize
	offset := reqInfo.PageSize * (reqInfo.Page - 1)

	totalSQL := " select count(*) from job_fairs jf where jf.deleted_at is null"
	err = global.GVA_DB.Raw(totalSQL).Scan(&total).Error
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, nil
	}
	SQL := "select jf.id ID, jf.company_name CompanyName , jf.city City , jf.salary Salary , jf.total_stu TotalStu , jf.telephone Telephone , jf.email Email, jf.address Address from job_fairs jf where jf.deleted_at is null limit ?,?"
	err = global.GVA_DB.Raw(SQL, offset, limit).Scan(&list).Error
	if err != nil {
		return list, total, err
	}
	return list, total, nil
}

//删除招聘会信息
func (m *MyApiService) DeleteJobFair(reqInfo r.SearchJobFairs, sysId uint) (err error) {

	err = global.GVA_DB.Model(&myPkg.JobFairs{}).Where("id = ?", reqInfo.ID).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	return nil

}

func (m *MyApiService) AddJobFairResp(reqInfo r.AddJobFair, sysId uint) (err error) {

	err = global.GVA_DB.Create(&myPkg.JobFairs{
		GVA_MODEL:   global.GVA_MODEL{},
		CompanyName: reqInfo.CompanyName,
		City:        reqInfo.City,
		Salary:      reqInfo.Salary,
		TotalStu:    reqInfo.TotalStu,
		Telephone:   reqInfo.Telephone,
		Email:       reqInfo.Email,
		Address:     reqInfo.Address,
	}).Error
	if err != nil {
		return err
	}
	return nil

}

func (m *MyApiService) SetJobFairResp(reqInfo r.AddJobFair, sysId uint) (err error) {

	err = global.GVA_DB.Model(&myPkg.JobFairs{}).Where("id = ?", reqInfo.ID).Updates(myPkg.JobFairs{
		//NoticeId:    reqInfo.NoticeId,
		CompanyName: reqInfo.CompanyName,
		City:        reqInfo.City,
		Salary:      reqInfo.Salary,
		TotalStu:    reqInfo.TotalStu,
		Telephone:   reqInfo.Telephone,
		Email:       reqInfo.Email,
		Address:     reqInfo.Address,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

/*-----------------------------------------------------------------------*/

// 根据条件获取毕业生信息列表
//func (m *MyApiService) GetStudentsListByConditionsResp(reqInfo r.SearchStu, sysId uint) (list []response.StudentsList, total int64, err error) {
//
//	limit := reqInfo.PageInfo.PageSize
//	offset := reqInfo.PageInfo.PageSize * (reqInfo.PageInfo.Page - 1)
//	whereStr := ""
//	if reqInfo.ClassNumber != "" {
//		whereStr = fmt.Sprintf("'%v' ssi.stu_class_num = '%v' and ", whereStr, reqInfo.ClassNumber)
//	} else if reqInfo.StuNumber != "" {
//		whereStr = fmt.Sprintf("'%v' ssi.stu_num = '%v' and ", whereStr, reqInfo.StuNumber)
//	}
//
//	totalSQL := " select count(*) from sys_stu_infos ssi inner join sys_users su on su.id = ssi.sys_user_id inner join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where " + whereStr + " ssi.deleted_at is null  order by ssi.stu_num"
//	if err = global.GVA_DB.Raw(totalSQL).Error; err != nil {
//		return list, total, err
//	}
//	if total == 0 {
//		return list, total, nil
//	}
//	sql := "select ssi.stu_num StudentNum, ssi.stu_name StudentName, scci.college_num CollegeNum, scci.college_name CollegeName, ssi.stu_class_num StudentClassNum, su.phone Phone, su.email Email, ssi.employed Employed from sys_stu_infos ssi inner join sys_users su on su.id = ssi.sys_user_id inner join sys_class_college_infos scci on scci.class_num = ssi.stu_class_num where " + whereStr + " ssi.deleted_at is null  order by ssi.stu_num limit ?,?"
//	err = global.GVA_DB.Raw(sql, limit, offset).Scan(&limit).Error
//	if err != nil {
//		return list, total, err
//	}
//	fmt.Println("hello World")
//	return list, total, nil
//}

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
