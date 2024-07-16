
<template>
  <div id="app">
    <head><meta name="referrer" content="never"></head>

    <router-view/>
  </div>
</template>
<script>
export default {
  name: 'App',
  mounted () {
    window.addEventListener('unload', this.saveState)
    window.addEventListener('load', this.getState)
  },
  methods: {
    saveState () {
      console.log("保存到本地存储变量")
      sessionStorage.setItem('state', JSON.stringify(this.$store.state))
    },
    getState () {
      if (sessionStorage.getItem('state')) {
        console.log("获取本地存储变量")
        var json = JSON.parse(sessionStorage.getItem('state'))
        this.$store.dispatch('user/login', json.user).then(res => {
          // sessionStorage.clear()
        })
      } else {
        // window.location.href = 'http://wxtest.fran6k.live'
      }
    }
  }
}
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /*text-align: center;*/
  /*width: 24.38rem;*/
  /*height: 52.75rem;*/
  /*width: 100vw;*/
  /*height: 100vh;*/
  /*table-layout: fixed;*/
  /*background-color: #F1F1F1;*/
}

html,body,#app{
  height: 100%;
  background-color: #F1F1F1;
  width: 100%;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
