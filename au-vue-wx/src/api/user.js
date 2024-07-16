import request from '@/utils/request'

export function login (data) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/login',
    url: '/auth/login',
    method: 'post',
    data
  })
}

export function getInfo (token) {
  return request({
    // baseURL: 'http://localhost:8080',
    // url: '/vue-admin-template/user/info',
    url: '/auth/info',
    method: 'get',
    params: { token }
  })
}

export function logout () {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}
