<template>
  <div id="bindIdBody">
    <van-image width="40vw" fit="cover" :src="require('../assets/logo.png')" style="margin: 0 auto;display: block;"></van-image>
    <van-form @submit="onSubmit" style="margin-top: 6vh;background-color: #ffffff">
      <van-field
        v-model="addUser.studentId"
        name="学号"
        label="学号"
        left-icon="idcard"
        placeholder="请输入学号"
        :rules="[{ required: true, message: '' }]"
      />
      <van-field
        v-model="addUser.password"
        type="password"
        name="密码"
        label="密码"
        placeholder="请输入WIFI登陆密码"
        left-icon="closed-eye"
        :rules="[{ required: true, message: '' }]"
      />
      <van-field
        v-model="addUser.realName"
        name="真实姓名"
        label="真实姓名"
        left-icon="user-o"
        placeholder="请输入真实姓名"
        :rules="[{ required: true, message: '' }]"
      />
      <van-field
        v-model="addUser.phoneNum"
        name="手机号"
        label="手机号"
        left-icon="phone-o"
        placeholder="请输入手机号"
        :rules="[{ required: true, message: '' }]"
      />
      <van-field
        v-model="addUser.weChatId"
        name="微信号"
        label="微信号"
        left-icon="records"
        placeholder="请输入微信号"
        :rules="[{ required: true, message: '' }]"
      />
      <div style="margin: 0 auto;width: 80vw;color: #A6A6A6">
        <p>注意:</p>
        <p>1.初次使用本系统需绑定信息</p>
        <p>2.学号和密码为教学楼网络登陆系统用户密码，密码错误将无法绑定并且无法使用本系统</p>
        <p>3.默认密码为sztu@+身份证后6位</p>
        <p>4.请确保姓名，手机号和微信号为真实有效信息，否则我们将无法联系到您</p>
        <p>5.所有数据保存在深圳技术大学内网服务器中</p>
      </div>
      <div style="margin: 0 auto;width: 80vw">
        <van-button color="#264B87" round block type="info" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import * as wechatApi from '../api/wechatApi'

import Vue from 'vue'
import { Field, Form, Toast } from 'vant'
Vue.use(Form)
Vue.use(Field)
Vue.use(Toast)

export default {
  name: 'BindId',
  data () {
    return {
      addUser: {
        studentId: '',
        password: '',
        realName: '',
        weChatId: '',
        phoneNum: '',
        avatar: '',
        nickName: '',
        openId: ''
      }
    }
  },
  methods: {
    onSubmit () {
      // this.$router.push({ path: '', query: {} })
      Toast.loading({
        message: '处理中',
        forbidClick: true,
        duration: 9999999
      })
      wechatApi.addWxUser(this.addUser).then(res => {
        if (res.code === 20000) {
          Toast.success(res.msg)
          window.location.href = process.env.SERVER_IP
        } else if (res.code === 20002) {
          Toast.fail(res.msg)
        }
      })
    }
  },
  created () {
    this.addUser.avatar = this.$route.query.userWxInfo.headimgurl
    this.addUser.nickName = this.$route.query.userWxInfo.nickname
    this.addUser.openId = this.$route.query.userWxInfo.openid

    // console.log(this.$route.query.userWxInfo)
  }
}
</script>

<style scoped>
#bindIdBody{
  background-color: #ffffff;
  /*height: 100%;*/
  height: 100vh;
}

.van-cell{
  width: 80vw;
  margin: 2vh auto;
  border-radius: 80vw;
  border: 1px solid #F1F1F1;
}
/*.van-cell .van-field .van-field__label{*/
/*  width: 5em;*/
/*}*/
.van-cell::after{
  border-bottom: none!important;
}
</style>
