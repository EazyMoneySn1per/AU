// index.js
// 获取应用实例
import { getUserInfo, testNetWork } from '../../api/wechatApi'
import { createStoreBindings } from 'mobx-miniprogram-bindings'
import { store } from '../../store/store'


Page({
  data: {
    userInfo: {},
    uname: '',
    isConnected: true,
    show: false,
  },
  // 事件处理函数
  goto(e) {
    // console.log(e.currentTarget.dataset.url);
    const url = e.currentTarget.dataset.url

    switch (url) {
      case 'department': {
        wx.navigateTo({
          url: '/pages/Au/department/department',
        })
        break
      }
      case 'actvities': {
        wx.navigateTo({
          url: '/pages/Au/activities/activities',
        })
        break
      }
      case 'assoDesc': {
        wx.navigateTo({
          url: '/pages/Asso/assoDescription/assoDescription',
        })
        break
      }
      case 'page1': {
        wx.navigateTo({
          url: '/components/Article/Article?url=https://mp.weixin.qq.com/s/vGCtDelYADJpEd5tpJuC0w',
        })
        break
      }
      case 'page2': {
        wx.navigateTo({
          url: '/components/Article/Article?url=https://mp.weixin.qq.com/s/vGCtDelYADJpEd5tpJuC0w',
        })
        break
      }
      case 'mysztu': {
        wx.navigateToMiniProgram({
          appId: 'wx69dd18a5fbd69ae0',
          path: 'pages/index/index',
        })
      }

    }
  },

  onConfirm() {
    // console.log('onConfirm');
    wx.requestSubscribeMessage({
      tmplIds: ['wBtDklx451_30mFe3tNFmJ6tTlIt1T8rzUYU_qIeS-s']
    })
  },

  onCancel() {
    return
  },

  onClose() {
    this.setData({show: false})
  },

  onLoad(options) {
    this.storeBindings = createStoreBindings(this, {
      store, // 需要绑定的数据仓库
      actions: ['updateActiveTabBarIndex'], // 将 this.setList 绑定为仓库中的 setList action
    })
    if (options.success) {
      this.setData({ show: true })
    }
  },
  async onShow() {
    this.updateActiveTabBarIndex(1)

    testNetWork().catch(err => {
      this.setData({
        isConnected: false,
      })
    })

    setTimeout(async () => {
      console.log(getApp().globalData.isLogin);
      if (!getApp().globalData.isLogin && this.data.isConnected) {
        //TODO 调用微信登录接口获取OpenId
        const { code } = await wx.p.login()
        const { data: res } = await getUserInfo(code)
        if (res.code === 20000) {
          //TODO 当前数据库已有此人
          getApp().globalData.isLogin = true
          const { Token, StudentId, RealName } = res.data
          wx.setStorageSync('token', Token)
          wx.setStorageSync('studentId', StudentId)
          wx.setStorageSync('realName', RealName)
          getApp().globalData.userInfo = res.data
          console.log('login from Index');
          console.log(getApp().globalData.isLogin);
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
    }, 1000)
  },
  onUnload() {},
  reConnected() {
    testNetWork().then(() => {
      this.setData({
        isConnected: true,
      })
    })
  },
})
