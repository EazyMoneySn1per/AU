// pages/Asso/myInterview/myInterview.js
import {
  getUserInterviews,
  FindAuStudentInterviewLists,
  studentConfirm,
  studentRefuse,
  getUserInfo,
} from '../../../api/wechatApi'
import Dialog from '../../../miniprogram_npm/@vant/weapp/dialog/dialog'
import Notify from '../../../miniprogram_npm/@vant/weapp/notify/notify'

Page({
  /**
   * 页面的初始数据
   */
  data: {
    assLogoServer: 'http://10.1.20.136/assLogo/',
    assoIsEmpty: true,
    auIsEmpty: true,
    InterviewList: [1, 2],
    activeNames: [],
    auActiveNames: ['1'],
    list1: [],
    list2: [],
    steps: [],
    activeSteps: [],
    refresh: false,
  },

  goto() {
    wx.navigateTo({
      url: '/components/Article/Article?url=https://mp.weixin.qq.com/s/rNT_g5wnlcoV6SSjbJUW3A',
    })
  },
  onChange(event) {
    this.setData({
      activeNames: event.detail,
    })
  },
  auOnChange(event) {
    this.setData({
      auActiveNames: event.detail,
    })
  },
  setInterviewStep(step) {
    switch (step) {
      case '1':
        return 0
      case '2':
        return 1
      case '3':
        return 2
      case '8':
        return 1
      case '9':
        return 0
    }
  },
  studentConfirm(e) {
    const id = e.currentTarget.dataset.id
    console.log(id)
    const that = this
    Dialog.confirm({
      title: '确认加入该社团吗',
      message: '一个人最多只能加入两个社团哦',
    }).then(async () => {
      console.log('confirm')
      const { data: res } = await studentConfirm(id)
      if (res.code === 20000) {
        // 成功通知
        Notify({ type: 'success', message: res.msg, duration: 1500 })
        //TODO 重新获取数据
        that.getList1()
        that.updateUserInfo()
      } else {
        Notify({ type: 'warning', message: '操作失败', duration: 1500 })
      }
      // console.log(res)
    })
  },
  studentCancel(e) {
    const id = e.currentTarget.dataset.id
    const that = this
    Dialog.confirm({
      title: '确定放弃加入该社团吗',
    }).then(async () => {
      // console.log('cancel')
      const { data: res } = await studentRefuse(id)
      // console.log(res)
      if (res.code === 20000) {
        // 成功通知
        Notify({ type: 'success', message: res.msg, duration: 1500 })
        //TODO 重新获取数据
        that.getList1()
      } else {
        Notify({ type: 'warning', message: '操作失败', duration: 1500 })
      }
    })
  },
  openMessageDialog(e) {
    const msg = e.currentTarget.dataset.msg
    // console.log(msg)
    Dialog.alert({
      title: '反馈信息',
      message: msg,
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  async getList1() {
    const studentId = wx.getStorageSync('studentId')
    const { data: res } = await getUserInterviews(studentId)
    console.log(res.data)
    const list = res.data.filter(item => item.interViewStatus !== 5 || item.interViewStatus !== 4)
    if (list && list.length > 0) {
      this.setData({
        list1: list,
        auIsEmpty: false,
      })
    } else if (list && list.length === 0) {
        this.setData({
          list1: list,
          auIsEmpty: true,
        })
    }
  },
  async getList2() {
    const studentId = wx.getStorageSync('studentId')
    const { data: res } = await FindAuStudentInterviewLists(studentId)
    console.log(res.data);
    if (res.data && res.data.length > 0) {
      this.setData({
        list2: res.data,
        assoIsEmpty: false,
      })
    } else if (res.data === null) {
        this.setData({
          list2: res.data,
          assoIsEmpty: true,
        })
    }
  },
  async onLoad(options) {
    if (options.success) {
      //TODO 从消息订阅进入小程序
      Dialog.confirm({
        title: '授权通知',
        message: '我们将向您发起一次消息订阅，以通知您下一次的面试结果',
      }).then()
    }
  },

  onSubscribeMessage() {
    wx.requestSubscribeMessage({
      tmplIds: ['wBtDklx451_30mFe3tNFmJ6tTlIt1T8rzUYU_qIeS-s', 'PHl68j37-9Imvpf_TK0Qo1Bz0aTnA77QlsNimxDjxNk']
    })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {
    const that = this
    wx.getSystemInfo({
      success: res => {
        console.log(res.windowHeight);
        wx.createSelectorQuery().select('.scroll-view').boundingClientRect(function(rect) {
          console.log(rect.top)
          that.setData({
            scrollViewHeight: res.windowHeight - rect.top
          })
        }).exec()
    
      },
    })
  },

  /**
   * 生命周期函数--监听页面显示
   */
  async onShow() {
    wx.showLoading({
      title: '加载中',
    })
    //TODO 获取社团面试
    await this.getList1()
    //TODO 获取社联面试
    await this.getList2()
    wx.hideLoading()
    //TODO 设置 steps
    let steps = []
    let activeSteps = []
    console.log(this.data.list2)
    if (this.data.list2) {
      this.data.list2.forEach(item => {
        const tmp = [{ text: '一面' }, { text: '二面' }, { text: '成功' }]
        item.status === '9' && (tmp[0].text = '一面失败')
        item.status === '8' && (tmp[1].text = '二面失败')
        steps.push(tmp)
        console.log(item.status)
        activeSteps.push(this.setInterviewStep(item.status))
      })
    }
    console.log(activeSteps)
    // console.log(activeSteps)
    this.setData({
      steps,
      activeSteps,
    })
  },

  async updateUserInfo() {
    const { code } = await wx.p.login()
    const { data: res } = await getUserInfo(code)
    if (res.code === 20000) {
      getApp().globalData.userInfo = res.data
    }
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
  onPullDownRefresh() {
    console.log('刷新中');
    this.setData({ refresh: true });
    setTimeout(() => {
      this.setData({
        refresh: false
      })
      this.getList1()
    },1000)
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {},

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {},
})
