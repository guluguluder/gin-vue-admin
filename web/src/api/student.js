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
