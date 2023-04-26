import service from '@/utils/request'

export const getJobFairList = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/get/getJobFairList',
    method: 'post',
    data: data
  })
}
export const deleteJobFair = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/delete/deleteJobFair',
    method: 'delete',
    data: data
  })
}
export const addJobFair = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/add/addJobFair',
    method: 'delete',
    data: data
  })
}
export const setJobFair = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/upd/setJobFair',
    method: 'delete',
    data: data
  })
}
