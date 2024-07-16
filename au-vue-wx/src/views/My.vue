<template>
  <div style="height:100%;width:100%;overflow-x:hidden;">
    <router-view/>
<!--    table-layout: fixed;-->
    <!--个人信息-->
    <div class="top-box">
      <div class="top-box-square"></div>
      <div class="top-box-radius"></div>
      <div class="top-box-info">
        <van-image round width="6.25rem" height="6.25rem" :src="avatar">11</van-image>
        <div style="font-weight: 700;font-size: 1.38rem;margin-top: 0.875rem">{{realName}}</div>
        <div style="font-size: 0.88rem;color: #797979;">学号: {{studentId}}</div>
      </div>
      <!--加入的社团-->
    </div>
    <div class="body-box">
      <div class="mid-box">
        <!--左侧社团-->
        <div class="mid-box-content">
          <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 40vw">
            <van-image round style="margin-bottom: 0.5rem" height="3.06rem" width="3.06rem"
                       :src="assName && assName[0].assLogo !== '-1'? assLogoServer + assName[0].assLogo : assLogoServer + 'kongbai.png' "></van-image>
            <div style="font-size: 0.88rem;color: #797979;">{{assName && assName[0].assName !== '-1'? assName[0].assName:"无社团"}}</div>
          </div>
        </div>
        <div class="mid-box-line"></div>
        <!--右侧社团-->
        <div class="mid-box-content">
          <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 40vw">
            <van-image round style="margin-bottom: 0.5rem" height="3.06rem" width="3.06rem"
                       :src="assName && assName[1].assLogo !== '-1'? assLogoServer + assName[1].assLogo : assLogoServer + 'kongbai.png' "></van-image>
            <div style="font-size: 0.88rem;color: #797979;">{{assName && assName[1].assName !== '-1'? assName[1].assName:"无社团"}}</div>
          </div>
        </div>
      </div>
      <!--功能卡片-->
      <div class="feature-box">
        <div class="feature-box-card" @click="route('InterviewSubmit')">
          <van-icon size="1.94rem" name="edit" />
          <div style="color: #797979;font-size: 1.1rem;margin-top: -1rem;">报名入口</div>
        </div>
        <div class="feature-box-card" @click="route('MyInterview')" >
          <van-icon size="1.94rem" name="completed" />
          <div style="color: #797979;font-size: 1.1rem;margin-top: -1rem;">我的报名</div>
        </div>
        <div class="feature-box-card">
          <van-icon size="1.94rem" name="bulb-o" />
          <div style="color: #797979;font-size: 1.1rem;margin-top: -1rem;">意见反馈</div>
        </div>
      </div>
      <!--后台入口-->
      <div class="bottom-box" @click="route('auPlatform')">
        <div class="bottom-box-content">
          后台登陆入口
        </div>
      </div>
    </div>
    <!--底部bar-->
    <div style="z-index: 99;position: relative;height: 3.2rem;">
      <TabbarAu active-page="2"></TabbarAu>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { Image as VanImage } from 'vant'

Vue.use(VanImage)

export default {
  name: 'My',

  data () {
    return {
      activePage: 2,
      avatar: '',
      studentId: '',
      realName: '',
      assName: null,
      assLogoServer: process.env.VUE_APP_ASSLOGO
    }
  },
  created () {
    setTimeout(this.initFunc, 20)
    // this.initFunc()
  },
  methods: {
    route (a) {
      switch (a) {
        case 'InterviewSubmit':
          this.$router.push({ path: '/InterviewSubmit' })
          break
        case 'MyInterview':
          this.$router.push({ path: '/MyInterview' })
          break
        case 'auPlatform':
          window.location.href = 'http://10.1.20.136'
          break
      }
    },
    initFunc () {
      this.avatar = this.$store.getters.avatar
      this.studentId = this.$store.getters.studentId
      this.realName = this.$store.getters.realName
      this.assName = this.$store.getters.assName
      // console.log(this.avatar)
      console.log(this.assName)
    }
  },
  components: {}
}
</script>

<style scoped>
body {
  /*height: 100vh;*/
}

.top-box {
  position: relative;
}

.top-box-radius {
  position: relative;
}

.top-box-square {
  z-index: 0;
  width: 100%;
  /*width: 24.38rem;*/
  height: 16rem;
  background-color: #1E70C1;
}

.top-box-radius {
  z-index: 1;
  width: 140%;
  height: 30rem;
  position: absolute;
  left: -20%;
  top: 8rem;
  /*z-index: -1;*/
  border-radius: 50% 50% 0 0;
  background: #F1F1F1;
}

.top-box-info {
  width: 100%;
  position: absolute;
  top: 5rem;
  text-align: center;
  z-index: 2;
}

.mid-box {
  display: flex;
  /*position: relative;*/
  justify-content: center;
  align-items: center;
  z-index: 2;
  top: -1.5rem;
  width: 23rem;
  height: 10rem;
  border-radius: 10px;
  border: 1px solid #F1F1F1;
  background-color: #ffffff;
}

.mid-box .mid-box-content {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.mid-box .mid-box-line {
  margin: 0 1rem 0 1rem;
  height: 65%;
  border: solid #BBBBBB 1px;
  opacity: 0.3;
}

.feature-box {
  z-index: 2;
  top: 14rem;
  /*position: relative;*/
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
}

.feature-box-card {
  width: 6.5rem;
  height: 8rem;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  border-radius: 10px;
  /*border:1px solid #BBBBBB;*/
  background-color: #ffffff;

}

.bottom-box-content {
  width: 22rem;
  height: 4.25rem;
  background-color: white;
  border-radius: 10px;
  text-align: center;
  line-height: 4rem;
  margin: 1rem auto;
  color: #797979;
}

.bottom-box {
  z-index: 2;
  /*position: relative;*/
  /*top: 14rem;*/
  width: 100%
}

.body-box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 3rem;
  width: 100%;
  /*height: 100%;*/
}
</style>
