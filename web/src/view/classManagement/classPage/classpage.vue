<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="班级">
          <el-input v-model="searchInfo.classNumber" placeholder="班级" />
        </el-form-item>
        <el-form-item label="所属学院">
          <el-select v-model="searchInfo.collegeNumber" clearable placeholder="请选择">
            <el-option
              v-for="item in methodOptions"
              :key="item.collegeNumber"
              :label="item.collegeName"
              :value="item.collegeNumber"
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
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="所属学院" min-width="100" prop="collegeName" />
        <el-table-column align="left" label="班级" min-width="100" prop="classNumber" />
        <el-table-column align="left" label="总人数" min-width="180" prop="totalStu" />
        <el-table-column align="left" label="已签约人数" min-width="180" prop="employedSum" />
        <el-table-column align="left" label="签约率" min-width="180" prop="employedPercent" />
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="magic-stick" size="small" @click="openDetails(scope.row)">详情</el-button>
<!--            <el-button type="primary" link icon="edit" size="small" @click="openEdit(scope.row)">编辑</el-button>-->
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
    <!--  详情   -->
    <el-dialog
      v-model="classDetailsDialog"
      custom-class="user-dialog"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div class="gva-table-box">
        <el-table
          :data="classDetailsData"
          row-key="ID"
        >
          <el-table-column align="left" label="ID" min-width="50" prop="ID" />
          <el-table-column align="left" label="学号" min-width="100" prop="stuNumber" />
          <el-table-column align="left" label="姓名" min-width="100" prop="stuName" />
          <el-table-column align="left" label="班级" min-width="180" prop="classNumber" />
          <el-table-column align="left" label="是否签约" min-width="180" prop="employed" />
          <el-table-column align="left" label="公司名称" min-width="180" prop="companyName" />
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
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="closeClassDetailDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <!--    -->
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
export default {
  name: 'ClassPage',
}
</script>

<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getColleges, getEmployedDetails } from '@/api/student'
import { getClassEmployedDetails, getClassList } from '@/api/class'
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
  classNumber: '',
  collegeNumber: '',
})

const methodOptions = ref([{
  collegeNumber: '',
  collegeName: '',
}])

const getCollegeData = async() => {
  const res = await getColleges({})
  if (res.code === 0) {
    methodOptions.value = res.data
  }
  console.log(res.data)
}
getCollegeData()

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
  const table = await getClassList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
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
  // userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'detail') {
        // const res = await register(req)
        // if (res.code === 0) {
        //   ElMessage({ type: 'success', message: '创建成功' })
        //   await getTableData()
        //   closeAddUserDialog()
        // }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const classDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  classDialog.value = false
}

const classDetailsDialog = ref(false)
const closeClassDetailDialog = () => {
  classDetailsDialog.value = false
}

const dialogFlag = ref()

const classDetailsData = ref([])
const openDetails = async(row) => {
  dialogFlag.value = 'details'
  const res = await getClassEmployedDetails({ page: page.value, pageSize: pageSize.value, classNumber: row.classNumber })
  classDetailsData.value = res.data.list
  total.value = res.data.total
  page.value = res.data.page
  pageSize.value = res.data.pageSize
  // classDetailsData.value = res.data
  classDetailsDialog.value = true
}
const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  classDialog.value = true
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

