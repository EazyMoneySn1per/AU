import {
  addInterviewUser,
  AddInterviewUserAu,
  getUserInterviews,
  getAssociationsNameMapAssid,
  FindAuStudentInterviewLists,
} from '../../../api/wechatApi'
import Notify from '../../../miniprogram_npm/@vant/weapp/notify/notify'
import Dialog from '../../../miniprogram_npm/@vant/weapp/dialog/dialog'

const chooseList = {
  社团: [],
  社联: ['策划部', '宣传部', '外联部', '财务部', '秘书部'],
}

Page({
  /**
   * 页面的初始数据
   */
  data: {
    gender: '男',
    phone: '',
    wx: '',
    ownInfo: '',
    show: false,
    columns: [],
    pickerValue0: '社团',
    pickerValue1: '',
    tempData: {},
    assMap: [],
    firstOpen: false,
    showDialog: true,
  },
  showPicker() {
    this.setData({
      show: true,
      firstOpen: true,
    })
    setTimeout(() => {
      this.setData({
        pickerValue0: '社团',
        pickerValue1:
          this.data.pickerValue1 === ''
            ? chooseList['社团'][0]
            : this.data.pickerValue1,
      })
    }, 300)
  },
  onClose() {
    this.setData({
      show: false,
    })
  },

  onChange(event) {
    const { picker, value, index } = event.detail
    if (index === 0) {
      this.setData({
        tempData: {
          pickerValue0: value[0],
          pickerValue1: chooseList[value[0]][0],
        },
      })
    } else {
      this.setData({
        tempData: {
          pickerValue0: value[0],
          pickerValue1: value[1],
        },
      })
    }
    picker.setColumnValues(1, chooseList[value[0]])
  },
  onConfirm() {
    this.setData({
      show: false,
      pickerValue0: this.data.tempData.pickerValue0,
      pickerValue1: this.data.tempData.pickerValue1,
    })
  },
  onCancel() {
    this.setData({
      show: false,
    })
  },
  onLogOut() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  },
  async submit() {
    //TODO 获取学号和姓名
    const { StudentId, RealName } = getApp().globalData.userInfo
    //TODO 完善表单信息
    this.data.firstOpen &&
      !this.data.pickerValue1 &&
      this.setData({
        pickerValue0: '社团',
        pickerValue1: chooseList['社团'][0],
      })
    let assId
    this.data.assMap.some(item => {
      if (this.data.pickerValue1 === item.Assname) {
        assId = item.Assid
        return true
      }
    })
    const formData = {
      name: RealName,
      studentId: StudentId,
      sex: this.data.gender,
      phoneNum: this.data.phone,
      wxNum: this.data.wx,
      description: this.data.ownInfo,
      department: this.data.pickerValue1,
      assId,
    }
    console.log(formData)
    //TODO 校验表单信息
    for (const key in formData) {
      if (formData[key] === '') {
        return wx.showToast({
          title: '请完善个人信息',
          icon: 'error',
          duration: 1500,
        })
      }
    }
    let res = {}
    let tmplIds
    if (this.data.pickerValue0 === '社团') {
      //TODO 社团面试
      tmplIds = ['u_Dg5h_Fuq-3LniqEWtM08ZjmYElNJXtva5_kUgI1og']
      console.log('社团面试')
      const { data } = await addInterviewUser(formData)
      res = data
    } else {
      //TODO 社联面试
      tmplIds = ['wBtDklx451_30mFe3tNFmJ6tTlIt1T8rzUYU_qIeS-s', 'PHl68j37-9Imvpf_TK0Qo1Bz0aTnA77QlsNimxDjxNk']
      console.log('社联面试')
      const { data } = await AddInterviewUserAu(formData)
      res = data
    }
    console.log(res)
    
    if (res.code === 20000) {
      Notify({ type: 'primary', message: res.msg, duration: 1200 })
      wx.requestSubscribeMessage({
        tmplIds,
        complete: () => {
          wx.setStorageSync('infoDraft', {
            sex: this.data.gender,
            phoneNum: this.data.phone,
            wxNum: this.data.wx,
            description: this.data.ownInfo,  
          })
          wx.navigateBack()
        }
      })
    } else {
      if (res.msg.includes('报名时间为:')) {
        Dialog.alert({
          title: '报名时间',
          message: res.msg.slice(6)
        })
      } else {
        Dialog.alert({
          title: '提示内容',
          message: res.msg,
        })  
      }
    }
  },
  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad(options) {
    if (wx.getStorageSync('infoDraft')) {
      const {sex, phoneNum, wxNum, description} = wx.getStorageSync('infoDraft')
      this.setData({
        gender: sex,
        phone: phoneNum,
        wx: wxNum,
        ownInfo: description
      })
    }

    setTimeout(async () => {
      const { data: res } = await getAssociationsNameMapAssid()
      chooseList['社团'] = res.data.map(item => item.Assname)
      this.setData({
        assMap: res.data,
        columns: [
          {
            values: Object.keys(chooseList),
            className: 'column1',
          },
          {
            values: chooseList['社团'],
            className: 'column2',
            defaultIndex: 0,
          },
        ],
      })
      console.log(this.data.assMap)
    }, 2000)
  },

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
