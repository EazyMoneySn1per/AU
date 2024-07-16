<template>
  <div class="background">
    <div class="title">
      加入心仪的社团
    </div>
    <div style="margin: 0 auto;text-align: center">
      <van-image width="80%" :src="require('../assets/InterviewSubmitAsset/banner.png')"></van-image>
    </div>
    <van-form @submit="onSubmit" style="margin-top: 6vh;background-color: #ffffff">
      <!--性别选择-->
      <div class="cell-common">
        <van-icon size="24px" :name="require('../assets/InterviewSubmitAsset/input_sex.png')"></van-icon>
        <van-field name="radio" class="cell-common-style">
          <template #input>
            <van-radio-group v-model="info.sex" direction="horizontal">
              <van-radio name="男">男</van-radio>
              <van-radio name="女">女</van-radio>
            </van-radio-group>
          </template>
        </van-field>
      </div>
      <div class="cell-description">
        <!--        <van-icon size="24px" :name="require('../assets/InterviewSubmitAsset/input_name.png')"></van-icon>-->
        <van-icon size="24px" color="#435FA6" name="phone-o" />
        <van-field
          v-model="info.phoneNum"
          placeholder="请输入手机"
          input-align="left"
          type="text"
          class="cell-common-style"
          autosize
        />
      </div>
      <div class="cell-description">
        <van-icon size="24px" color="#435FA6" name="smile-comment-o"  />
        <van-field
          v-model="info.wxNum"
          placeholder="请输入微信号"
          input-align="left"
          type="text"
          class="cell-common-style"
          autosize
        />
      </div>

      <!--个人简介-->
      <div class="cell-description">
<!--        <van-icon size="24px" :name="require('../assets/InterviewSubmitAsset/input_name.png')"></van-icon>-->
        <van-icon size="24px" color="#435FA6" name="records" />
        <van-field
          v-model="info.description"
          placeholder="请输入个人简介"
          input-align="left"
          type="textarea"
          maxlength="100"
          class="cell-description-style"
          show-word-limit
          autosize
          row="3"
        />
      </div>
      <!--社团-->
      <div class="cell-common">
        <van-icon size="24px" :name="require('../assets/InterviewSubmitAsset/input_chooseAss.png')"></van-icon>
        <van-field
          readonly
          clickable
          name="picker"
          :value="curAss"
          placeholder="选择社团或社联"
          @click="showPicker = true"
          class="cell-common-style"
        />
        <van-popup v-model="showPicker" position="bottom">
          <van-picker
            show-toolbar
            :columns="pickList"
            @confirm="onConfirm"
            @cancel="showPicker = false"
          />
        </van-popup>
      </div>
      <!--提交按钮-->
      <div style="margin: 0 auto;width: 80vw">
        <van-button color="#264B87" round block type="info" native-type="submit">提交</van-button>
      </div>
      <div style="height: 2vh"></div>
    </van-form>
  </div>
</template>

<script>
import Vue from 'vue'
import { Field, Form, Toast, Radio, RadioGroup, Popup, Picker } from 'vant'
import * as wechatApi from '../api/wechatApi'

Vue.use(Form)
Vue.use(Field)
Vue.use(Toast)
Vue.use(Radio)
Vue.use(Popup)
Vue.use(Picker)
Vue.use(RadioGroup)

export default {
  name: 'InterviewSubmit',
  data () {
    return {
      info: {
        sex: '',
        description: '',
        name: '',
        studentId: '',
        assId: '',
        phoneNum: '',
        wxNum: '',
        department: ''
      },
      assList: [],
      pickList: [
        {
          text: '社团',
          children: []
        },
        {
          text: '社联',
          children: [{ text: '策划部' }, { text: '宣传部' }, { text: '外联部' }, { text: '财务部' }, { text: '秘书部' }]
        }
      ],
      curAss: '',
      curSubmitType: '',
      showPicker: false
    }
  },
  created () {
    this.info.studentId = this.$store.getters.studentId
    this.info.name = this.$store.getters.realName
    wechatApi.getAssociationsNameMapAssid().then(res => {
      this.assList = res.data
      for (const i in this.assList) {
        var object = {
          text: this.assList[i]
        }
        this.pickList[0].children.push(object)
      }
    })
  },
  methods: {
    onSubmit: function () {
      Toast.loading({
        message: '提交中',
        forbidClick: true,
        duration: 9999999
      })
      switch (this.curSubmitType) {
        case '0':
          wechatApi.addInterviewUser(this.info).then(res => {
            if (res.data.status === 0) {
              Toast.fail(res.msg)
            } else {
              Toast.setDefaultOptions('success', { forbidClick: true })
              Toast.success(res.msg)
              setTimeout(() => {
                this.$router.push({ path: '/my' })
              }, 2000)
            }
          })
          break
        case '1':
          wechatApi.AddInterviewUserAu(this.info).then(res => {
            if (res.data.status === 0) {
              Toast.fail(res.msg)
            } else {
              Toast.setDefaultOptions('success', { forbidClick: true })
              Toast.success(res.msg)
              setTimeout(() => {
                this.$router.push({ path: '/my' })
              }, 2000)
            }
          })
          break
        default:
          Toast.clear()
          break
      }
    },
    onConfirm (value) {
      switch (value[0]) {
        case '社联':
          this.info.department = value[1]
          this.curSubmitType = '1'
          this.curAss = value[0] + ' ' + value[1]
          break
        case '社团':
          this.curSubmitType = '0'
          this.curAss = value[1]
          for (const i in this.assList) {
            if (this.assList[i] === value[1]) {
              this.info.assId = i
            }
          }
          break
      }
      this.showPicker = false
    }
  }
}

</script>

<style scoped>
.background{
  background-color: white!important;
  width: 100%;
  height: 100%;
}
.title{
  text-align: center;
  font-size: 20px;
  padding: 20px;
  font-weight: 700;
}
/*.van-cell{*/
/*  width: 80vw;*/
/*  margin-left: 2vw;*/
/*  padding: 0;*/
/*  border-bottom: 1px solid #F1F1F1;*/
/*}*/
.cell-common-style{
  width: 80vw;
  margin-left: 5vw;
  padding: 0;
}
.cell-description-style{
  width: 80vw;
  margin-left: 5vw;
  padding: 0;
  border: 1px solid #F1F1F1;
}
.cell-description{
  display: flex;
  align-items: flex-start;
  justify-self: center;
  margin: 20px auto;
  width: 80vw;
}
.cell-common{
  display: flex;
  align-items: center;
  justify-self: center;
  margin: 20px auto;
  width: 80vw;
}
.van-cell::after{
  right: 0;
  left: 0;
}
</style>
