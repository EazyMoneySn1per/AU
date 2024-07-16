import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog'
import { getAssociationsNameMapAssid, exitAss, getUserInfo } from '../../api/wechatApi'

Page({
  options: {
    // styleIsolation: 'apply-shared'
  },
  /**
   * 页面的初始数据
   */
  data: {
    id: 1,
    avatarUrl: '',
    studentId: '',
    leftAss: {},
    rightAss: {},
    logoSrc: '',
    realName: '',
    show: false,
  },
  goto(event) {
    const url = event.currentTarget.dataset.url
    url === 'feedback' &&
      wx.navigateTo({
        url: `/pages/feedback/feedback`,
      })
    wx.navigateTo({
      url: `/pages/Asso/${url}/${url}`,
    })
  },
  exitAss(e) {
    const ass = e.currentTarget.dataset.obj
    const studentId = wx.getStorageSync('studentId')
    let message = '' //弹窗信息
    let assId = '' //社团id
    if (ass === 'left') {
      // data.assId =
      message = `确定要退出${this.data.leftAss.AssName}吗`
      assId = this.data.leftAss.id
    } else {
      message = `确定要退出${this.data.rightAss.AssName}吗`
      assId = this.data.rightAss.id
    }
    // console.log(assId)
    const that = this
    Dialog.confirm({
      title: '退出社团',
      message,
    }).then(async () => {
      // on confirm
      console.log('onConfirm')
      const { data: res } = await exitAss(studentId, assId)
      console.log(res)
      let title = '操作成功'
      let icon = 'success'
      res.code !== 20000 && ((title = '操作失败'), (icon = 'error'))
      wx.showToast({
        title,
        icon,
        duration: 1200,
      })
      await that.updateUserInfo()
      that.getList()
    }).catch(() => {})
  },
  NoPermisson() {
    Dialog.alert({
      title: 'No Permisson',
      message: '您没有权限进入该页面哦',
    })
  },
  onClose() {
    this.setData({
      show: false
    })
  },
  onChooseAvatar(e) {
    const { avatarUrl } = e.detail
    wx.setStorageSync('avatar', avatarUrl)
    this.setData({
      avatarUrl,
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad(options) {
    //TODO 获取头像
    const avatarUrl = wx.getStorageSync('avatar')
    const studentId = wx.getStorageSync('studentId')
    const realName = wx.getStorageSync('realName')
    this.setData({
      avatarUrl,
      studentId,
      logoSrc: getApp().globalData.assLogoServer,
      realName,
    })
    // console.log(userInfo)
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {},

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    this.getList()
  },

  async updateUserInfo() {
    const { code } = await wx.p.login()
    const { data: res } = await getUserInfo(code)
    if (res.code === 20000) {
      getApp().globalData.userInfo = res.data
    }
  },

  async getList() {
    //TODO 获取已加入社团列表
    const { AssName: assList } = getApp().globalData.userInfo
    const { data: res } = await getAssociationsNameMapAssid()
    const assMap = res.data
    for (let item of assList) {
      if (item.AssName) {
        for (let map of assMap) {
          if (item.AssName === map.Assname) {
            item.id = map.Assid
          }
        }
      }
    }

    // console.log(assList)
    this.setData({
      leftAss: assList[0],
      rightAss: assList[1],
    })
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
