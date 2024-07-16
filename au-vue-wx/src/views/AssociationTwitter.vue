<template>
  <div>
    <iframe src="http://10.161.7.87:8080/admin/test"></iframe>
    <button @click="test">11111</button>
  </div>
</template>

<script>

export default {
  name: 'AssociationTwitter',
  data () {
    return {
      url: 'http://10.161.7.87:8080/admin/test'
    }
  },
  created () {
    // this.$router.go(this.url)
  },
  methods: {
    test () {
      // this.url = axios.get('http://10.161.7.87:8080/admin/test')
      window.location.href = 'https://mp.weixin.qq.com/s/ZClnmFzROEIgPUlkPqTAXw'
    }
  }
}
</script>

<style scoped>

</style>
