// pages/feedback/feedback.js
import Toast from '@vant/weapp/toast/toast'

Page({
  /**
   * 页面的初始数据
   */
  data: {
    result: [],
    message: '',
    fileList: [],
    showProgress: false,
    percentage: 0,
  },

  onChange(event) {
    this.setData({
      result: event.detail,
    })
    console.log(this.data.result)
  },
  onSubmit() {
    wx.showToast({
      icon: 'success',
      title: '提交成功',
      duration: 1000,
    })
    setTimeout(() => {
      wx.navigateBack()
    }, 1000)
  },
  afterRead(event) {
    const { file } = event.detail
    // 当设置 mutiple 为 true 时, file 为数组格式，否则为对象格式
    console.log(file)
    const that = this
    const uploadTask = wx.uploadFile({
      url: 'https://au.sztu.edu.cn:9090/home/auServerData/images', // 仅为示例，非真实的接口地址
      filePath: file.url,
      name: 'file',
      formData: { user: 'test' },
      success(res) {
        // 上传完成需要更新 fileList
        const { fileList = [] } = that.data
        fileList.push({ ...file, url: res.data })
        that.setData({
          fileList,
          percentage: 0,
        })
        console.log('upload Successfully')
      },
    })
    uploadTask.onProgressUpdate(res => {
      console.log(res)
      this.setData({
        percentage: res.progress,
      })
    })
  },
  controlProgress() {
    let timer = setInterval(() => {
      this.setData({
        percentage: this.data.percentage + 1,
      })
      this.data.percentage === 100 && clearInterval(timer)
    }, 20)
  },
  deleteImg(event) {
    let index = event.detail.index
    const list = [...this.data.fileList]
    list.splice(index, 1)
    this.setData({
      fileList: list,
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {},

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {},

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {},

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
