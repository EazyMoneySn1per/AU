<template>
  <div>
    <router-view />
    我的
    <van-tabbar route v-model="activePage">
      <van-tabbar-item replace to="/" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item replace to="/association" icon="search">社团</van-tabbar-item>
      <van-tabbar-item replace to="/my" icon="friends-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
export default {
  name: 'My',
  data () {
    return {
      activePage: 2
    }
  }
}
</script>

<style scoped>

</style>
