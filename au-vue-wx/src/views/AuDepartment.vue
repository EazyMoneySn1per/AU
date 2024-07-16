<template>
  <div>
    <twitter-group v-for="item in twitterList" v-bind:key="item.time" :title="item.title" :url="item.url" :pic-url="item.picUrl" :time="item.time"></twitter-group>
  </div>
</template>

<script>
import TwitterGroup from '@/components/TwitterGroup'

export default {
  name: 'AuDepartment',
  components: { TwitterGroup },
  data () {
    return {
      twitterList: [
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77RDN0HyWicu2r32Hnrar6SoqibvNLd4J1yadz8iafxSGsFhJ1o6bXlLiaOiaJxPqCUCphqjrkBvlqqX5RQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/6mAIAKAUmmOfnKCfEuqoew',
          title: '社联招新 | 开往未来的“社联”列车即将进站~'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Sqoic5wsd0V2yyticXicW7mHJk3j7Z6VUNISMTX9etBCxS3gsR81s0c1s9G7CwaImccozMpFrxur4Cw/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/G22DR4kYm2q5sVUXSL8yoA',
          title: '社联招新 | 策划部 & 外联部'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TibX6YP0TnhVurkSlKhPNZ8qh3TkmxonBQU2Nt7NO9Nwe6Gv7pUE4NgBxdMeN9POde4h1lZjZguXQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/tfVZUWbe59Gxd2mMyILvs',
          title: '社联招新 | 秘书部 & 财务部'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TibX6YP0TnhVurkSlKhPNZ8YicibN3yd5vmCqFAldEydSHZNEdfEJmQ6muxhH6mznDNq1ZusejmajyQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/JgKXE9FFq4_tEFpcDcEeYw',
          title: '社联招新 | 宣传部 & 技术部'
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>
