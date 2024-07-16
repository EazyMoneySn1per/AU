<template>
  <div class="home">
    <!--轮播图-->
<!--    <div class="top-box">-->
<!--      <van-swipe :autoplay="4000" indicator-color="white">-->
<!--        <van-swipe-item class="top-box-image" @click="goTwitter()">-->
<!--          <van-image fit="cover" lazy-load :src="require('../assets/banner.png')"></van-image>-->
<!--        </van-swipe-item>-->
<!--      </van-swipe>-->
<!--    </div>-->
    <!--部门介绍和社联活动-->
    <div @click="goTwitter('https://mp.weixin.qq.com/s/6mAIAKAUmmOfnKCfEuqoew')">
      <van-image style="display:block;" fit="cover" lazy-load :src="require('../assets/banner.png')"></van-image>
    </div>
    <div @click="goTwitter('https://mp.weixin.qq.com/s/mK-LWFA6-bOPUVGS39J3ow')">
      <van-image style="display:block;" src="http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TibX6YP0TnhVurkSlKhPNZ8hlWjqVrp7yrdQ44wcBPrmhFVsQq1T127YpegD14jXtG9tfOibREwDLA/0?wx_fmt=jpeg"></van-image>
    </div>
    <div class="mid-box">
      <div class="mid-box-content" @click="route('AuDepartment')">
        <van-image style="margin-bottom: 0.5rem" fit="contain" width="4.3rem" :src="require('../assets/home-left.png')"></van-image>
        <div style="font-size: 1.06rem;color: rgba(76, 76, 76, 100);font-weight: 700;">社联部门</div>
      </div>
      <div class="mid-box-line">
      </div>
      <div class="mid-box-content" @click="route('AuActivities')">
        <van-image style="margin-bottom: 0.5rem" fit="contain" width="4.3rem" :src="require('../assets/home-right.png')"></van-image>
        <div style="font-size: 1.06rem;color: rgba(76, 76, 76, 100);font-weight: 700;">社联活动</div>
      </div>
    </div>
    <!--推文-->
    <div class="bottom-box" @click="route('AssociationDescription')">
      <div class="bottom-box-content">
          <!--图片遮罩层-->
        <div class="bottom-box-content-cover">
          <van-image style="z-index: 2;opacity:0.7;" lazy-load fit="cover" width="100%" height="100%" radius="10px" :src="require('../assets/bottomBackground.jpg')"></van-image>
        </div>
        <van-icon style="z-index: 4;position: absolute;left: 20rem;top: 0.5rem;" :name="require('../assets/route.png')" size="1.75rem" color="#ffffff"/>
        <div style="z-index: 4;width: 10rem;font-size: 1.2rem;font-weight: 700;color: #ffffff;
        position: absolute;left: 50%;margin-left: -4rem;top: 4.4rem">深技大社团介绍</div>
      </div>
    </div>
<!--    <div class="bottom-box">-->
<!--      <div class="bottom-box-content">-->
    <div style="width: 100%;text-align: center">
      <video src="http://10.1.20.136/assLogo/xcsp.mp4" poster="http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77R0fUAaccCG1Egu0MS694l7KpbibG0FTHQ7pKllbznexRouWK2nzxcpRZKCIa0YWfvn5DbRRNwz70A/0?wx_fmt=jpeg" webkit-playsinline="isiPhoneShowPlaysinline" playsinline="isiPhoneShowPlaysinline" preload="metadata"  controls="controls" class="video_fill">
        您的浏览器不支持 video 标签        </video>
    </div>

<!--      </div>-->
<!--    </div>-->
    <!--底部tab-->
    <div style="z-index: 99;position: relative;height: 3.2rem;">
      <TabbarAu active-page="0"></TabbarAu>
    </div>
  </div>

</template>

<script>
// @ is an alias to /src
import * as wechatApi from '../api/wechatApi'

import { Swipe, SwipeItem, Lazyload, Toast } from 'vant'

import Vue from 'vue'
import { getToken } from '@/utils/auth'

Vue.use(Swipe)
Vue.use(SwipeItem)
Vue.use(Lazyload)

export default {
  name: 'Home',
  data () {
    return {
      activePage: 0,
      code: '',
      userInfo: ''
    }
  },
  components: {
  },
  methods: {
    // test () {
    //   // this.url = axios.get('http://10.161.7.87:8080/admin/test')
    //   window.location.href = 'https://mp.weixin.qq.com/s/ZClnmFzROEIgPUlkPqTAXw'
    // },
    route (a) {
      switch (a) {
        case 'AuActivities':
          this.$router.push({ path: '/' + a })
          break
        case 'AuDepartment':
          this.$router.push({ path: '/' + a })
          break
        case 'AssociationDescription':
          this.$router.push({ path: '/' + a })
          break
      }
    },
    getUrlCode () { // 截取url中的code方法
      var url = location.search
      // eslint-disable-next-line no-new-object
      var theRequest = new Object()
      // eslint-disable-next-line eqeqeq
      if (url.indexOf('?') != -1) {
        var str = url.substr(1)
        var strs = str.split('&')
        for (var i = 0; i < strs.length; i++) {
          theRequest[strs[i].split('=')[0]] = (strs[i].split('=')[1])
        }
      }
      return theRequest
    },
    getCode () {
      this.code = ''
      var local = 'http://wxtest.fran6k.live/'
      this.code = this.getUrlCode().code
      // eslint-disable-next-line eqeqeq
      if (this.code == null || this.code == '') {
        // var getTimestamp = new Date().getTime()
        window.location.href = 'https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx117b9376ed2498e4&redirect_uri=' +
          encodeURIComponent(local) +
          // encodeURIComponent(local + '?timestamp=' + getTimestamp) +
          '&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect'
      }
    },
    goTwitter (url) {
      window.location.href = url
    }

  },
  created () {
    let that = this
    var token = getToken()
    var avator = this.$store.getters.avatar
    var check = JSON.parse(sessionStorage.getItem('state'))
    console.log('检测本地存储变量')

    var curUrl = window.location.href
    console.log(curUrl)
    if (!this.code) {
      this.getCode()
    }
    console.log(check)
    // if (this.code && !check.user.avatar) {
    if (this.code && !this.$store.getters.code) {
      Toast.loading({
        message: '加载中',
        forbidClick: true,
        duration: 9999999
      })
      wechatApi.wxlogin(this.code).then(res => {
        // that.userInfo = res.data
        console.log(res)
        Toast.clear()
        if (res.code === 20000) {
          this.loading = true
          res.data.code = this.code
          this.$store.dispatch('user/login', res.data).then(() => {
            sessionStorage.setItem('state', JSON.stringify(that.$store.state))
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else if (res.code === 20001) {
          this.$router.push({ path: '/bindid', query: { userWxInfo: res.data } })
        } else if (res.code === 20003) {
          window.location.href = 'http://wxtest.fran6k.live'
        }
      })
    }
  }

}
</script>
<style>
.home{
  /*width: 100%;*/
  /*background: #F1F1F1;*/
}
.top-box-image{
  /*z-index: 1;*/
  /*width: 100vw;*/
  width: 100%;
  text-align: center;
}
.mid-box{
  display: flex;
  position: relative;
  justify-content: center;
  align-items: center;
  width: 22rem;
  height: 10rem;
  margin: 0rem auto 0rem auto;
  border-radius: 10px;
  border:1px solid #BBBBBB;
  background-color: #ffffff;
}
.mid-box .mid-box-content{
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}
.mid-box .mid-box-line{
  margin: 0 4rem 0 4rem;
  height: 65%;
  border: solid #BBBBBB 1px;
}

.bottom-box{
  margin: 0 auto;
  position: relative;
  height: 100%;
}
.bottom-box .bottom-box-content {
  width: 22rem;
  height: 10rem;
  position: relative;
  margin: 0.8rem auto;
  border-radius: 10px;
  background-size:100% 100%;
}

.bottom-box .bottom-box-content .bottom-box-content-cover {
  width: 100%;
  height: 100%;
  position: relative;
  border-radius: 10px;
  border:1px solid #BBBBBB;
  z-index: 3;
  background-color: #4C4C4C;
}
.video_fill{
  /*width: 100%;*/
  width: 22rem;
  margin: 0 auto;
  display: inline-block;
}
</style>
