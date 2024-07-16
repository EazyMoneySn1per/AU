import { request, requestWithErr } from '../utils/http'

// 获取用户信息
export function getUserInfo(code) {
  return request({
    path: `/getUserInfo`,
    method: 'GET',
    data: {
      code,
    },
  })
}

// 进行统一认证
export function addWxUser(data) {
  return request({
    path: '/addWxUser',
    method: 'post',
    data,
  })
}

// 获取所有的社团以及社联组织
export function getTotalAssociation() {
  return request({
    path: '/getTotalAssociation',
    method: 'get',
  })
}

// 获取优秀社团
export function getExecllentAssociation() {
  return request({
    path: '/getExecllentAssociation',
    method: 'get',
  })
}

export function getAssociationsNameMapAssid() {
  return request({
    path: '/getAssociationsNameMapAssid',
    method: 'get',
  })
}

export function getAssociationsByAssid(assId) {
  return request({
    path: '/getAssociationsByAssid',
    method: 'get',
    data: {
      assId,
    },
  })
}

export function getAssByType(type) {
  return request({
    path: `/getAssByType`,
    method: 'get',
    data: {
      type,
    },
  })
}

export function getTwitter(url) {
  return request({
    path: '/getTwitterPic',
    method: 'get',
    data: {
      url,
    },
  })
}

// 面试接口
// 社团面试接口
export function addInterviewUser(data) {
  return request({
    path: '/interview/addInterviewUser',
    method: 'post',
    data,
  })
}

export function getUserInterviews(studentId) {
  return request({
    path: `/interview/getUserInterviews`,
    method: 'get',
    data: {
      studentId,
    },
  })
}

export function exitAss(studentId, assId) {
  return request({
    path: `/exitAss`,
    method: 'get',
    data: {
      studentId,
      assId,
    },
  })
}

export function studentConfirm(id) {
  return request({
    path: '/interview/studentConfirm',
    method: 'get',
    data: {
      id,
    },
  })
}

export function studentRefuse(id) {
  return request({
    path: '/interview/studentRefuse',
    method: 'get',
    data: {
      id,
    },
  })
}

// 社联面试接口
export function FindAuStudentInterviewLists(studentId) {
  return request({
    path: '/interview/goGet/FindStudentInterviewLists',
    method: 'get',
    data: {
      studentId,
    },
  })
}

export function AddInterviewUserAu(data) {
  return request({
    path: '/interview/goPost/AddInterviewUser',
    method: 'post',
    data,
  })
}

export function testNetWork() {
  return requestWithErr({
    path: '/',
    method: 'get',
  })
}
