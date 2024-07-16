export const assLogoUrl = 'https://au.sztu.edu.cn/assLogo/'
const baseUrl = 'https://au.sztu.edu.cn:9090/wxapi'

export const request = function ({ path, method, data }) {
  const url = baseUrl + path
  const token = getApp().globalData.userInfo.Token || wx.getStorageSync('token')
  return wx.p
    .request({
      url,
      method,
      data,
      header: { token },
      timeout: 3000,
    })
    .catch(err => {
      wx.hideLoading()
      wx.showToast({
        title: '连接超时',
        icon: 'error',
        duration: 1000,
      })
      wx.switchTab({
        url: '/pages/index/index',
      })
    })
}

export const requestWithErr = function ({ path, method, data }) {
  var app = getApp()
  const url = baseUrl + path
  const token = wx.getStorageSync('token')
  return wx.p.request({
    url,
    method,
    data,
    header: { token },
    timeout: 3000,
  })
}
