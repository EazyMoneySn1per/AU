import { addWxUser } from '../../api/wechatApi'

Page({
  /**
   * 页面的初始数据
   */
  data: {
    realName: '',
    studentId: '',
    password: '',
    phoneNum: '',
    wechatId: '',
    nickName: '-',
    avatar: '-',
    show: true,
  },
  async submit() {
    //TODO 校验表单内容
    for (const key in this.data) {
      if (this.data[key] === '') {
        return wx.showToast({
          title: '请完善表单',
          icon: 'error',
          duration: 1200
        })
      }
    }
    const MpOpenId = wx.getStorageSync('openId')
    const data = {
      ...this.data,
      MpOpenId,
    }
    data.__webviewId__ && delete data.__webviewId__
    delete data.show
    wx.showLoading({
      title: '正在进行认证',
    })
    console.log(data)
    const { data: res } = await addWxUser(data)
    console.log(res)
    // 认证成功
    if (res.code === 20000) {
      wx.hideLoading()
      wx.showToast({
        title: '学号认证成功',
        icon: 'success',
        duration: 1000,
      })
      setTimeout(() => {
        wx.switchTab({
          url: '/pages/index/index',
        })
      }, 1000)
    } else {
      wx.showToast({
        title: '学号认证失败',
        icon: 'error',
        duration: 1000,
      })
    }
  },

  onCancel() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad(options) {},

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {},

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {},

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {},

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {},

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {},

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {},
})
