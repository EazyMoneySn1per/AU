<template>
  <div class="background">
<!--    <div class="title">我的报名列表</div>-->

    <div @click="goTwitter('https://mp.weixin.qq.com/s/rNT_g5wnlcoV6SSjbJUW3A')">
      <van-image style="display:block;" fit="cover" lazy-load src="http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Ry1I1bhVHherSopWqmQ7foAhY19cNL0X7AqP66I5NLmOICIIykticN8rau4YNTllpzwQqpph3Tv7Q/0?wx_fmt=jpeg"></van-image>
    </div>

    <van-tabs type="card" animated :swipeable="true">
      <!--学生社团-->
      <van-tab title="学生社团">
        <!--卡片-->
        <div class="ass-card" v-for="item in InterviewList" v-bind:key="item.id">
          <!--左侧logo和社团名称-->
          <div class="ass-card-left">
            <van-image round width="6.25rem" height="6.25rem" :src="assLogoServer+item.assLogo"></van-image>
            <div class="ass-card-left-assName">{{item.assName}}</div>
          </div>
          <!--右侧信息-->
          <div class="ass-card-right">
            <div class="ass-card-right-assInfo">状态：<span class="ass-card-right-assInfo-status">{{item.interViewStatusMessage}}</span></div>
            <div class="ass-card-right-assInfo">社长：{{ item.presidentName }}</div>
            <div class="ass-card-right-assInfo">联系方式：{{ item.presidentWechat }}</div>
            <van-collapse v-model="activeNames" accordion style="margin-top: 8px">
              <van-collapse-item title="反馈信息" :name="item.id">{{ item.backMessage }}</van-collapse-item>
            </van-collapse>
            <div class="ass-card-right-button" v-if="item.buttonControl !== 2">
              <van-button disabled color="#0D4E9A">确认加入</van-button>
              <van-button disabled>放弃加入</van-button>
            </div>
            <div class="ass-card-right-button" v-else>
              <van-button color="#0D4E9A" @click="studentConfirm(item.id)">确认加入</van-button>
              <van-button @click="studentRefuse(item.id)">放弃加入</van-button>
            </div>
          </div>
        </div>
        <!--空图片占位符-->
        <van-empty v-if="isEmpty === true" description="没有报名信息哦" />
      </van-tab>
      <!--社联-->
      <van-tab title="社联">
        <van-notice-bar color="#888888" background="#f4f3f4" left-icon="volume-o" :scrollable="false">
          <van-swipe
            vertical
            class="notice-swipe"
            :autoplay="2500"
            :show-indicators="false"
          >
            <van-swipe-item>报名2个部门进入社联的概率更大哦!</van-swipe-item>
            <van-swipe-item>面试相关通知会在这里进行更新！请及时关注</van-swipe-item>
            <van-swipe-item>报名成功后记得加入社联招新咨询群！</van-swipe-item>
          </van-swipe>
        </van-notice-bar>
        <!--卡片-->
        <div class="ass-card" v-for="item in AuInterviewList" v-bind:key="item.uuid">
          <!--左侧logo和部门名称-->
          <div class="ass-card-left">
            <van-image round width="6.25rem" height="6.25rem" :src="require('../assets/AuDepartmentLogo/'+item.departmentName + '.png')"></van-image>
            <div class="ass-card-left-assName">{{ item.departmentName }}</div>
          </div>
          <!--右侧信息-->
          <div class="ass-card-right">
            <van-steps :active="auSetInterviewStep(item.status)"
                       :active-icon="item.status === '8' || item.status === '9' ? 'cross':'like-o'"
                       :active-color="item.status === '8' || item.status === '9' ? '#e54242':'#0D4E9A'">
              <van-step>{{ item.status === '9' ? "一面失败" : "一面" }}</van-step>
              <van-step>{{ item.status === '8' ? "二面失败" : "二面" }}</van-step>
              <van-step>成功</van-step>
            </van-steps>
            <div class="ass-card-right-assInfo">面试时间：<span class="ass-card-right-assInfo-status">{{item.time}}</span></div>
            <div class="ass-card-right-assInfo">面试地点：<span class="ass-card-right-assInfo-status">{{item.location}}</span></div>
            <van-collapse v-model="AuActiveNames" accordion style="margin-top: 8px">
              <van-collapse-item title="面试通知" :name="item.uuid">{{ item.backMessage }}</van-collapse-item>
            </van-collapse>
            <div style="height: 2vh"></div>
          </div>
        </div>
        <!--空图片占位符-->
        <van-empty v-if="AuIsEmpty === true" description="没有报名信息哦" />
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from 'vue'
import { Tab, Tabs, Collapse, CollapseItem, Dialog, Toast, Empty, Loading, Step, Steps, NoticeBar, Swipe, SwipeItem } from 'vant'
import * as wechatApi from '../api/wechatApi'
import { removeToken } from '@/utils/auth'

Vue.use(Step)
Vue.use(Steps)
Vue.use(Tab)
Vue.use(Tabs)
Vue.use(Collapse)
Vue.use(CollapseItem)
Vue.use(Dialog)
Vue.use(Empty)
Vue.use(Loading)
Vue.use(NoticeBar)
Vue.use(Swipe)
Vue.use(SwipeItem)
export default {
  name: 'MyInterview',
  data () {
    return {
      activeNames: [],
      AuActiveNames: [],
      InterviewList: [],
      studentId: '',
      isEmpty: false,
      AuIsEmpty: false,
      AuInterviewList: [],
      assLogoServer: process.env.VUE_APP_ASSLOGO
    }
  },
  created () {
    var json = JSON.parse(sessionStorage.getItem('state'))
    this.studentId = json.user.studentId
    // this.studentId = '201904020217'
    this.getUserInterviewList(this.studentId)
    this.getAuInterviewList(this.studentId)
  },
  methods: {
    // 获取面试的社团列表
    getUserInterviewList (studentId) {
      Toast.loading({
        message: '加载中',
        forbidClick: true,
        duration: 9999999
      })
      wechatApi.getUserInterviews(studentId).then(res => {
        this.InterviewList = res.data
        this.InterviewList.length === 0 ? this.isEmpty = true : null
        Toast.clear()
      })
    },
    getAuInterviewList (studentId) {
      wechatApi.FindAuStudentInterviewLists(studentId).then(res => {
        this.AuInterviewList = res.data
        !this.AuInterviewList || this.AuInterviewList.length === 0 ? this.AuIsEmpty = true : null
      })
    },
    // 确认加入
    studentConfirm (id) {
      Dialog.confirm({
        title: '确认加入该社团吗',
        message: '一个人最多只能加入两个社团，并且无法撤销操作哦'
      })
        .then(() => {
          Toast.loading({
            message: '请稍等',
            forbidClick: true,
            duration: 9999999
          })
          wechatApi.studentConfirm(id).then(res => {
            if (res.data.status === 1) {
              Toast.setDefaultOptions('success', { forbidClick: true })
              Toast.success('加入成功，即将重新进入系统')
              removeToken()
              setTimeout(() => {
                window.location.href = process.env.SERVER_IP
              }, 2000)
            } else {
              Toast.fail(res.msg)
            }
          })
        })
        .catch(() => {
        })
    },
    // 放弃加入
    studentRefuse (id) {
      const that = this
      Dialog.confirm({
        title: '确定放弃加入该社团吗',
        message: ''
      })
        .then(() => {
          Toast.loading({
            message: '请稍等',
            forbidClick: true,
            duration: 9999999
          })
          wechatApi.studentRefuse(id).then(res => {
            that.getUserInterviewList(that.studentId)
            Toast.clear()
          })
        })
        .catch(() => {
        })
    },
    // 设置面试step
    auSetInterviewStep (step) {
      switch (step) {
        case '1':
          return 0
        case '2':
          return 1
        case '3':
          return 2
        case '8':
          return 1
        case '9':
          return 0
      }
    },
    // 部门logo
    chooseLogo (name) {
      switch (name) {
        case '策划部':
          return '../assets/AuDepartmentLogo/策划部.jpeg'
      }
    },
    goTwitter (url) {
      window.location.href = url
    }
  }
}
</script>

<style>
.notice-swipe {
  height: 40px;
  line-height: 40px;
}
.ass-card{
  display: flex;
  align-items: center;
  justify-self: center;
  width: 100%;
  height: 100%;
  border-bottom: 2vh solid #f4f3f4;
  background-color: white;
}
.ass-card-left-assName{
  padding-top: 10px;
  font-size: 14px;
}
.ass-card-right-assInfo{
  color: #888888;
  font-size: 12px;
}
.ass-card-right-assInfo-status{
  font-weight: 700;
  color: black;
}
.ass-card-right-button{
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1.5vh 0;
  width: 100%;
 }
.ass-card-right{
  /*padding: 10px;*/
  padding-top: 1vh;
  width: 50%;
  height: 100%;
}
.ass-card-left{
  text-align: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  /*height: 100%;*/
  height: 20vh;
  width: 40%;
}

.background{
  background-color: white;
  height: 100vh;
}
.title{
  text-align: center;
  font-size: 24px;
  padding: 10px;
}

.van-tabs__nav--card .van-tab.van-tab--active{
  background-color: #0D4E9A;
}
.van-tabs__nav--card{
  margin: 0;
  box-shadow: 0px 10px 10px #888888;
  border: none;
}
.van-tabs__nav--card .van-tab{
  color: black;
  border-right: #0D4E9A;
}
.van-cell{
  padding: 0;
  color: #888888;
  font-size: 12px;
}
/*.van-hairline--top-bottom::after{*/
/*  !*border: 0px solid #ebedf0;*!*/
/*}*/
</style>
