import request from '@/utils/request'

// export function getCode () {
//   axios.get('https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxce51d21a2b4b2e7d&secret=4e62b2d0d5980faa0375765724dc8c34'
//   ).then(res => {
//     console.log(res)
//   })
// }

export function wxlogin (code) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/getUserInfo',
    method: 'get',
    params: { code }
  })
}
export function addWxUser (data) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/addWxUser',
    method: 'post',
    data
  })
}
export function getTotalAssociation () {
  return request({
    url: '/getTotalAssociation',
    method: 'get'
  })
}
export function getExecllentAssociation () {
  return request({
    url: '/getExecllentAssociation',
    method: 'get'
  })
}
export function getAssociationsNameMapAssid () {
  return request({
    url: '/getAssociationsNameMapAssid',
    method: 'get'
  })
}
export function getAssociationsByAssid (assId) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/getAssociationsByAssid',
    method: 'get',
    params: { assId }
  })
}
export function getTwitter (url) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/getTwitterPic',
    method: 'get',
    params: { url }
  })
}

// 面试接口
// 社团面试接口
export function addInterviewUser (data) {
  return request({
    url: '/interview/addInterviewUser',
    method: 'post',
    data
  })
}

export function getUserInterviews (studentId) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/interview/getUserInterviews',
    method: 'get',
    params: { studentId }
  })
}
export function studentConfirm (id) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/interview/studentConfirm',
    method: 'get',
    params: { id }
  })
}
export function studentRefuse (id) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/interview/studentRefuse',
    method: 'get',
    params: { id }
  })
}
// 社联面试接口
export function FindAuStudentInterviewLists (studentId) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/interview/goGet/FindStudentInterviewLists',
    method: 'get',
    params: { studentId }
  })
}
export function AddInterviewUserAu (data) {
  return request({
    url: '/interview/goPost/AddInterviewUser',
    method: 'post',
    data
  })
}
