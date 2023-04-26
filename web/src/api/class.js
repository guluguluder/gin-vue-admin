import service from '@/utils/request'

export const getClassList = (data) => {
  return service({
    url: '/myapp/v1/class/get/getClassList',
    method: 'post',
    data: data
  })
}
export const getClassEmployedDetails = (data) => {
  return service({
    url: '/myapp/v1/class/get/getClassEmployedDetails',
    method: 'post',
    data: data
  })
}
