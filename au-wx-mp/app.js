// app.js

// 将wx中的异步请求方法 转化为 promise
import { promisifyAll } from 'miniprogram-api-promise'
import { getUserInfo, testNetWork } from './api/wechatApi'

const wxp = (wx.p = {})
promisifyAll(wx, wxp)
// 将wx中的异步请求方法 转化为 promise

import './utils/http'


App({
  globalData: {
    userInfo: {},
    assLogoServer: 'http://au.sztu.edu.cn/assLogo/',
    isLogin: false,
  },
  onLaunch(options) {
    //TODO 监听网络状态
    this.netWorkManage()
    //TODO 检查是否登录
    this.login()
  },
  
  netWorkManage() {
    wx.onNetworkStatusChange(res => {
      !res.isConnected &&
        wx.showToast({
          title: '网络似乎不太流畅',
          icon: 'error',
          duration: 2000,
        })
    })
  },

  async login() {
    let isConnected = true
    testNetWork().catch(err => {
      this.setData({
        isConnected: false,
      })
    })
    if (!this.globalData.isLogin && isConnected) {
      //TODO 调用微信登录接口获取OpenId
      const { code } = await wx.p.login()
      const { data: res } = await getUserInfo(code)
      if (res.code === 20000) {
        //TODO 当前数据库已有此人
        this.isLogin = true
        console.log(res);
        console.log('login from App');
        const { Token, StudentId, RealName } = res.data
        wx.setStorageSync('token', Token)
        wx.setStorageSync('studentId', StudentId)
        wx.setStorageSync('realName', RealName)
        this.globalData.userInfo = res.data
        this.globalData.isLogin = true
      } else { 
          //TODO 前往用户认证页面
          wx.setStorageSync('openId', res.data.MpOpenId)
          wx.showModal({
            title: '提示',
            content: '请前往用户认证界面',
            success(res) {
              res.confirm &&
                wx.navigateTo({
                  url: '/pages/bindId/bindId',  
                })
            },
          })
        } 
    }
  }
})
