import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'vant/lib/index.css'
import 'lib-flexible'
import TabbarAu from '@/components/TabbarAu'
// import VueRouter from 'vue-router'

import { Icon, Image as VanImage, Button } from 'vant'
// Vue.use(VueRouter)

Vue.use(Icon)
Vue.use(VanImage)
Vue.use(Button)

Vue.config.productionTip = false
Vue.component('TabbarAu', TabbarAu)
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

// 禁止双击放大
var lastTouchEnd = 0
document.documentElement.addEventListener('touchend', function (event) {
  var now = Date.now()
  if (now - lastTouchEnd <= 300) {
    event.preventDefault()
  }
  lastTouchEnd = now
}, false)

// router.beforeEach((to, from, next) => {
//   const studentId = store.getters.studentId
//   if (!studentId) {
//     console.log("测试")
//     console.log(studentId)
//     if (to.path !== '/bindid') {
//       window.location.href = 'http://wxtest.fran6k.live'
//     } else {
//       next()
//     }
//   } else {
//     next()
//   }
// })
