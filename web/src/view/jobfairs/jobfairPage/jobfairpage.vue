<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addJob">新增招聘信息</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="公司名称" min-width="150" prop="companyName" />
        <el-table-column align="left" label="城市" min-width="150" prop="city" />
        <el-table-column align="left" label="薪资" min-width="100" prop="salary" />
        <el-table-column align="left" label="招聘人数" min-width="100" prop="totalStu" />
        <el-table-column align="left" label="手机号" min-width="180" prop="telephone" />
        <el-table-column align="left" label="邮箱" min-width="180" prop="email" />
        <el-table-column align="left" label="招聘会地址" min-width="180" prop="address" />
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" @click="openEdit(scope.row)">编辑</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="deleteUserFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" size="small">删除</el-button>
              </template>
            </el-popover>
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
    <!--    -->
    <el-dialog
      v-model="addUserDialog"
      custom-class="user-dialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="100px">
          <el-form-item label="公司名称" prop="companyName">
            <el-input v-model="userInfo.companyName" />
          </el-form-item>
          <el-form-item label="城市" prop="city">
            <el-input v-model="userInfo.city" />
          </el-form-item>
          <el-form-item label="薪资" prop="salary">
            <el-input v-model="userInfo.salary" />
          </el-form-item>
          <el-form-item label="招聘人数" prop="totalStu">
            <el-input v-model="userInfo.totalStu" />
          </el-form-item>
          <el-form-item label="手机号" prop="telephone">
            <el-input v-model="userInfo.telephone" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item label="招聘会地址" prop="address">
            <el-input v-model="userInfo.address" />
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
  </div>
</template>

<script>
export default {
  name: 'JobFairPage',
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
import { addJobFair, deleteJobFair, getJobFairList, setJobFair } from '@/api/jobFairs'
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
const searchInfo = ref({})

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
  const table = await getJobFairList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
//
// watch(() => tableData.value, () => {
//   setAuthorityIds()
// })

const initPage = async() => {
  getTableData()
}

initPage()

// const setAuthorityIds = () => {
//   tableData.value && tableData.value.forEach((user) => {
//     user.authorityIds = user.authorities && user.authorities.map(i => {
//       return i.authorityId
//     })
//   })
// }

const deleteUserFunc = async(row) => {
  const res = await deleteJobFair({ Id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const userInfo = ref({
  ID: '',
  companyName: '',
  city: '',
  salary: '',
  totalStu: '',
  telephone: '',
  email: '',
  address: '',
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
        const res = await addJobFair(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setJobFair(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addJob = () => {
  dialogFlag.value = 'add'
  userInfo.value = {}
  addUserDialog.value = true
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
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

