import service from '@/utils/request'

export const getStudentList = (data) => {
  return service({
    url: '/myapp/v1/students/list/getAll',
    method: 'post',
    data: data
  })
}
export const getDetailByStuNumber = (data) => {
  return service({
    url: '/myapp/v1/students/detail/getByStuNumber',
    method: 'post',
    data: data
  })
}
export const setStudentInfo = (data) => {
  return service({
    url: '/myapp/v1/students/update/setStudentInfo',
    method: 'post',
    data: data
  })
}
export const deleteStudent = (data) => {
  return service({
    url: '/myapp/v1/students/delete/deleteStudent',
    method: 'post',
    data: data
  })
}
export const getColleges = (data) => {
  return service({
    url: '/myapp/v1/students/get/getColleges',
    method: 'get',
    data: data
  })
}
