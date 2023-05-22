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
export const getEmployedList = (data) => {
  return service({
    url: '/myapp/v1/students/get/getEmployedList',
    method: 'post',
    data: data
  })
}
export const getEmployedDetails = (data) => {
  return service({
    url: '/myapp/v1/students/details/getEmployedDetails',
    method: 'post',
    data: data
  })
}
export const setEmployedDetails = (data) => {
  return service({
    url: '/myapp/v1/students/update/setEmployedDetails',
    method: 'post',
    data: data
  })
}

export const getEmploymentInfos = (data) => {
  return service({
    url: '/myapp/v1/infos/getEmploymentInfos',
    method: 'post',
    data: data
  })
}
export const getCollegeEmployedDetails = (data) => {
  return service({
    url: '/myapp/v1/infos/getCollegeEmployedDetails',
    method: 'post',
    data: data
  })
}
export const addSummary = (data) => {
  return service({
    url: '/myapp/v1/infos/addSummary',
    method: 'post',
    data: data
  })
}
