<template>
  <div style="z-index: 99">
    <van-tabbar route v-model="active" active-color="#1E70C1">
      <van-tabbar-item replace to="/" icon="wap-home-o">首页</van-tabbar-item>
      <van-tabbar-item replace to="/association" icon="search">社团</van-tabbar-item>
      <van-tabbar-item replace to="/my" icon="friends-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
import Vue from 'vue'
import { Tabbar, TabbarItem } from 'vant'

Vue.use(Tabbar)
Vue.use(TabbarItem)
export default {
  name: 'TabbarAu',
  data () {
    return {
      active: this.activePage
    }
  },
  props: {
    activePage: String
  }
}
</script>

<style scoped>

</style>
