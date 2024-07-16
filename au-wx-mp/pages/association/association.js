// pages/association/association.js
import * as wxapi from '../../api/wechatApi'

Page({
  /**
   * 页面的初始数据
   */
  data: {
    value: '',
    activeKey: 0,
    assoKind: [
      { text: '优秀社团', value: 0 },
      { text: '学术科技类', value: 1 },
      { text: '文化体育类', value: 2 },
      { text: '创新创业类', value: 3 },
      { text: '学生互助类', value: 4 },
      { text: '思想政治类', value: 5 },
      { text: '志愿公益类', value: 6 },
    ],
    currentSelectValue: 0,
    hasSearchInfo: false,
    currentSelectType: '优秀社团',
    assLogoServer: 'http://10.1.20.136/assLogo/',
    totalList: [],
    excellentList: [],
    currentList: [],
    timer: null,
    isSearch: false,
  },
  noSearchInfo() {
    const index = this.data.currentSelectValue
    this.setData({
      currentList: this.data.assoKind[index].data,
      isSearch: false,
    })
  },
  searchChange() {
    this.data.value === '' && this.noSearchInfo()
  },
  getSearchInfo() {
    this.data.value === '' && this.noSearchInfo()
    this.setData({
      isSearch: true,
    })
    wx.showLoading({
      title: '加载中...',
    })
    setTimeout(() => {
      console.log(1)
      //TODO 获取搜索信息
      if (this.data.value === '') {
        const index = this.data.currentSelectValue
        this.setData({
          currentList: this.data.assoKind[index].data,
        })
      } else {
        let list = this.data.totalList.filter(item => {
          return item.assName.includes(this.data.value) === true
        })
        this.setData({
          currentList: list,
        })
      }
      wx.hideLoading()
    }, 300)
  },
  onClick(e) {
    const obj = e.currentTarget.dataset.obj
    // console.log(obj);
    const msg = {
      logo: obj.logo,
      assName: obj.assName,
      presidentName: obj.presidentName,
      desc: obj.assDescription,
    }
    wx.navigateTo({
      url: '/pages/Asso/assoDetail/assoDetail?id=1',
      success: function (res) {
        // 通过 eventChannel 向被打开页面传送数据
        res.eventChannel.emit('acceptDataFromOpenerPage', { data: msg })
      },
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {},
  async onAssTypeChange(event) {
    const index = event.detail
    const type = this.data.assoKind[index].text
    this.setData({
      currentSelectValue: index,
    })
    if (!this.data.assoKind[index].data) {
      const { data: res } = await wxapi.getAssByType(type)
      this.data.assoKind[index].data = res.data
    }
    this.setData({
      currentList: this.data.assoKind[index].data,
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {},

  /**
   * 生命周期函数--监听页面显示
   */
  async onShow() {
    this.data.currentList.length === 0 && this.getList()
  },
  async getList() {
    //TODO 获取优秀社团
    const { data: res } = await wxapi.getExecllentAssociation()

    //TODO 如果社团描述末尾不是句号，加上...c
    console.log(res)
    res.data.forEach(item => {
      if (!item['assDescription'].endsWith('。')) {
        item['assDescription'] += '...'
      }
    })
    //TODO 获取所有社团
    const { data: res1 } = await wxapi.getTotalAssociation()
    //TODO setData
    this.data.assoKind[0].data = res.data
    this.setData({
      currentList: res.data,
      totalList: res1.data,
    })
  },
  openDropDown() {
    const fc = this.selectComponent('#dropdown')
    fc.toggle()
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
