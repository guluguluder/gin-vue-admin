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
    method: 'post',
    data: data
  })
}
export const setJobFair = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/upd/setJobFair',
    method: 'post',
    data: data
  })
}
export const addCommentInfo = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/upd/addCommentInfo',
    method: 'post',
    data: data
  })
}
export const getCommentList = (data) => {
  return service({
    url: '/myapp/v1/jobFairs/get/getCommentList',
    method: 'post',
    data: data
  })
}
