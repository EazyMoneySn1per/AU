import { storeBindingsBehavior } from 'mobx-miniprogram-bindings'
import { store } from '../store/store'

Component({
  options: {
    styleIsolation: 'shared'
  },
  behaviors: [storeBindingsBehavior],
  storeBindings: {
    store,
    fields: {
      active: 'activeTabBarIndex',
    },
    actions: {
      updateActive: 'updateActiveTabBarIndex'
    }
  },
  data: {
      list: [
          {index: 1, text: '首页', icon: 'wap-home-o', url: '/pages/index/index'},
          {index: 2, text: '社团', icon: 'search', url: '/pages/association/association'},
          {index: 3, text: '我的', icon: 'friends-o', url: '/pages/my/my'}
      ]
    // list: [{
    //   "url": "/pages/index/index",
    //   "text": "首页",
    //   "iconPath": "/images/tabbar/home.png",
    //   "selectedIconPath": "/images/tabbar/home-active.png",
    //   "icon": 'home-o'
    // }, {
    //   "url": "/pages/association/association",
    //   "text": "社团",
    //   "iconPath": "/images/tabbar/search.png",
    //   "selectedIconPath": "/images/tabbar/search-active.png",
    //   "icon": 'search'
    // }, {
    //   "url": "/pages/my/my",
    //   "text": "我的",
    //   "iconPath": "/images/tabbar/friends.png",
    //   "selectedIconPath": "/images/tabbar/friends-active.png",
    //   "icon": 'friends-o'
    // }]
  },
  methods: {
    onChange(e) {
      // event.detail 的值为当前选中项的索引
      // this.updateActive(event.detail)
      // wx.switchTab({
      //   url: this.data.list[event.detail].url
      // });
      const index = e.currentTarget.dataset.index
      // console.log(index);
      // this.data.list.forEach(item => {
      //   if (item.index === index) {
      //     item.active = true
      //   } else item.active = false
      // })
      // console.log(this.data.list);
      this.updateActive(index)
      const url = this.data.list[index-1].url
    //   console.log(url)
    //   console.log(1);
      wx.switchTab({
        url
      })
    }
  }
});