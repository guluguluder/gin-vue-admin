<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="学号">
          <el-input v-model="searchInfo.stuNumber" placeholder="学号" />
        </el-form-item>
        <el-form-item label="是否签约">
          <el-select v-model="searchInfo.isEmployed" clearable placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <!--        <el-table-column align="left" label="ID" min-width="50" prop="ID" />-->
        <el-table-column align="left" label="用户名(学号)" min-width="150" prop="stuNumber" />
        <el-table-column align="left" label="昵称(姓名)" min-width="100" prop="stuName" />
        <el-table-column align="left" label="是否签约" min-width="180" prop="isEmployed" />
        <el-table-column align="left" label="签约公司" min-width="180" prop="companyName" />
        <el-table-column align="left" label="公司城市" min-width="180" prop="jobCity" />
        <el-table-column align="left" label="工作职位" min-width="180" prop="jobTitle" />
        <el-table-column align="left" label="工资薪资" min-width="180" prop="jobSalary" />
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="magic-stick" size="small" @click="openDetails(scope.row)">详情</el-button>
            <el-button type="primary" link icon="edit" size="small" @click="openEdit(scope.row)">编辑</el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <!--   编辑弹窗    -->
    <el-dialog
      v-model="addUserDialog"
      custom-class="user-dialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
          <el-form-item label="学号" prop="stuNumber">
            <el-input v-model="userInfo.stuNumber" />
          </el-form-item>
          <el-form-item label="姓名" prop="stuName">
            <el-input v-model="userInfo.stuName"/>
          </el-form-item>
          <el-form-item label="是否签约" prop="isEmployed">
            <el-input v-model="userInfo.isEmployed" />
          </el-form-item>
          <el-form-item label="签约公司" prop="companyName">
            <el-input v-model="userInfo.companyName" />
          </el-form-item>
          <el-form-item label="公司城市" prop="jobCity">
            <el-input v-model="userInfo.jobCity" />
          </el-form-item>
          <el-form-item label="工作职位" prop="jobTitle">
            <el-input v-model="userInfo.jobTitle" />
          </el-form-item>
          <el-form-item label="工资薪资" prop="jobSalary">
            <el-input v-model="userInfo.jobSalary" />
          </el-form-item>
        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeAddUserDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <!--    -->
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />

    <!--  ////////////  详情弹窗  ////////// -->
    <el-dialog
      v-model="studentDetailDialog"
      custom-class="user-dialog"
      title="详情"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="StudentForm" :rules="rules" :model="employedDetails" label-width="80px">
          <el-form-item label="学号" prop="stuNumber">
            <el-input v-model="employedDetails.stuNumber" />
          </el-form-item>
          <el-form-item label="姓名" prop="stuName">
            <el-input v-model="employedDetails.stuName" />
          </el-form-item>
          <el-form-item label="性别" prop="stuSex">
            <el-input v-model="employedDetails.stuSex" />
          </el-form-item>
          <el-form-item label="所属学院" prop="collegeName">
            <el-input v-model="employedDetails.collegeName" />
          </el-form-item>
          <el-form-item label="班级" prop="classNumber">
            <el-input v-model="employedDetails.classNumber" />
          </el-form-item>
          <el-form-item label="年级" prop="gradeNumber">
            <el-input v-model="employedDetails.gradeNumber" />
          </el-form-item>
          <el-form-item label="是否签约" prop="isEmployed">
            <el-input v-model="employedDetails.isEmployed" />
          </el-form-item>
          <el-form-item label="签约公司" prop="companyName">
            <el-input v-model="employedDetails.companyName" />
          </el-form-item>
          <el-form-item label="工作城市" prop="jobCity">
            <el-input v-model="employedDetails.jobCity" />
          </el-form-item>
          <el-form-item label="工作职位" prop="jobTitle">
            <el-input v-model="employedDetails.jobTitle" />
          </el-form-item>
          <el-form-item label="工作薪资" prop="jobSalary">
            <el-input v-model="employedDetails.jobSalary" />
          </el-form-item>
        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="closeStudentDetailDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'EmploymentPage',
}
</script>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser,
} from '@/api/user'
import {
  deleteStudent,
  getDetailByStuNumber, getEmployedDetails,
  getEmployedList, setEmployedDetails,
  setStudentInfo,
} from '@/api/student'
import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
  AuthorityData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        authorityId: item.authorityId,
        authorityName: item.authorityName,
        children: []
      }
      setAuthorityOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        authorityId: item.authorityId,
        authorityName: item.authorityName
      }
      optionsData.push(option)
    }
  })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  stuNumber: '',
  isEmployed: '',
})

const methodOptions = ref([
  {
    value: 'Y',
    label: '是',
  },
  {
    value: 'N',
    label: '否',
  },
])

const onReset = () => {
  searchInfo.value = {}
}

// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getEmployedList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async(row) => {
  // const res = await deleteUser({ id: row.ID })
  const res = await deleteStudent({ Id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const userInfo = ref({
  ID: '',
  stuNumber: '',
  stuName: '',
  isEmployed: '',
  companyName: '',
  jobCity: '',
  jobTitle: '',
  jobSalary: '',
})
const employedDetails = ref({
  stuNumber: '',
  stuName: '',
  stuSex: '',
  collegeName: '',
  classNumber: '',
  gradeNumber: '',
  isEmployed: '',
  companyName: '',
  jobCity: '',
  jobTitle: '',
  jobSalary: '',
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' },
  ],
  authorityId: [
    { required: true, message: '请选择用户角色', trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setEmployedDetails(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}
const StudentForm = ref(null)
const studentDetailDialog = ref(false)
const closeStudentDetailDialog = () => {
  studentDetailDialog.value = false
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const openDetails = async(row) => {
  dialogFlag.value = 'details'
  const res = await getEmployedDetails({ stuNumber: row.stuNumber })
  employedDetails.value = res.data
  studentDetailDialog.value = true
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  console.log(row.ID)
  addUserDialog.value = true
}

</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>

