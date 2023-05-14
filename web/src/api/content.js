import service from '@/utils/request'

export const getContentList = (data) => {
  return service({
    url: '/myapp/v1/content/list/getAll',
    method: 'post',
    data: data
  })
}
export const setContentInfo = (data) => {
  return service({
    url: '/myapp/v1/content/upd/setContent',
    method: 'post',
    data: data
  })
}
export const addContentInfo = (data) => {
  return service({
    url: '/myapp/v1/content/add/addContent',
    method: 'post',
    data: data
  })
}
export const deleteContent = (data) => {
  return service({
    url: '/myapp/v1/content/del/deleteContent',
    method: 'delete',
    data: data
  })
}
