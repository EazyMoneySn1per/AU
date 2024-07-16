<template>
  <div>
    <div class="banner-box">
    </div>
    <div class="body-box">
      <!--头像-->
      <div class="body-box-logo">
        <van-image round width="6.25rem" height="6.25rem" :src="assLogoServer + assInfo.logo" style="background-color: #F1F1F1">11</van-image>
      </div>
      <!--脱标方块填充-->
      <div style="height: 4rem"></div>
      <!--名字-->
      <div class="body-box-assName">
        <div class="body-box-assName-one">{{assInfo.assName}}</div>
        <div class="body-box-assName-two">现任社长：{{assInfo.presidentName}}</div>
      </div>
      <!--分割线-->
      <div class="divide-line"></div>
      <!--社团简介-->
      <div class="body-box-subContent">
        <div class="body-box-subContent-title">
          <van-icon size="24px" :name="require('../assets/assActivity.png')"></van-icon>
          <div style="margin-left: 5px;font-size: 18px">社团简介</div>
        </div>
        <div class="body-box-subContent-description">{{assInfo.assDescription}}</div>
      </div>
      <!--社团活动-->
      <div class="body-box-subContent">
        <div class="body-box-subContent-title">
          <van-icon size="24px" :name="require('../assets/assActivity.png')"></van-icon>
          <div style="margin-left: 5px;font-size: 18px">社团活动</div>
        </div>
        <van-image style="margin: 0 auto;display: block" width="90%" fit="cover" src="http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TGuiaIPSub7ibEebezete1Ljz03DkhVkmNzOSq1iazXnDfsqHyetSVrVs15tiarjb15u5I5mDteQdjeQ/0?wx_fmt=jpeg"></van-image>
      </div>
    </div>
  </div>
</template>

<script>
import * as wechatApi from '../api/wechatApi'

export default {
  name: 'AssociationDetail',
  data () {
    return {
      assId: '',
      assInfo: '',
      assLogoServer: process.env.VUE_APP_ASSLOGO
    }
  },
  created () {
    this.assId = this.$route.query.assId
    wechatApi.getAssociationsByAssid(this.assId).then(res => {
      this.assInfo = res.data
      // console.log(this.assInfo)
    })
  }
}
</script>

<style scoped>
.banner-box{
  background-color: #0D4E9A;
  width: 100vw;
  height: 15vh;
}
.body-box{
  position: relative;
  background-color: #F1F1F1;
  width: 100vw;
  height: 85vh;
  border-top-left-radius: 7px;
  border-top-right-radius: 7px;
  margin-top: -1vh;
}
.body-box-logo{
  position: absolute;
  text-align: center;
  width: 100vw;
  top: -2.5rem;
}
.body-box-assName{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin: 0;
  width: 100vw;
}
.body-box-assName-one{
  margin: 2px;
  font-size: 18px;
}
.body-box-assName-two{
  margin: 2px;
  color: #5F5F5F;
}
.body-box-subContent{
  width: 100%;
  align-items: center;
}
.body-box-subContent-title{
  display: flex;
  align-items: center;
  justify-content: left;
  color: #101010;
  font-size: 16px;
  margin: 4px auto;
  width: 90%;
}
.body-box-subContent-description{
  margin: 0 auto;
  width: 90%;
  word-wrap: break-word;
  text-indent: 20px;
  font-size: 17px;
  color: #888888;
}
.divide-line{
  width: 90%;
  height: 1px;
  margin: 2vh auto;
  background-color: #bbbbbb;
  opacity: 40%;
}
</style>
