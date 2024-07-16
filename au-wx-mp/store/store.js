import {
    observable,
    action
} from 'mobx-miniprogram'

export const store = observable({
    // 数据字段
    activeTabBarIndex: 1,
    // 用于前后端身份认证
    token: '',
    // 用户信息
    avatar: 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0',
    nickName: '',
    realName: '',
    studentId: '',
    assName: '',
    // 调用wx.login获取的code
    code: '',
    // actions函数
    updateActiveTabBarIndex: action(function (index) {
        this.activeTabBarIndex = index
    }),
    updateUserInfo: action(function (userInfo) {
        this.avatar = aurl
    })
})