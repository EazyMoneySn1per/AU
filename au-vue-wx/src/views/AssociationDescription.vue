<template>
  <div>
    <twitter-group v-for="item in twitterList" v-bind:key="item.time" :title="item.title" :url="item.url" :pic-url="item.picUrl" :time="item.time"></twitter-group>
  </div>
</template>

<script>
import TwitterGroup from '@/components/TwitterGroup'
import * as wechatApi from '../api/wechatApi'

export default {
  name: 'AssociationDescription',
  components: { TwitterGroup },
  data () {
    return {
      twitterList: [
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Suv0F6hGLVU92EryAGyE2gmzd9buFIgg7fO5sXhec3XRygicliaCWrWfp3Z2uRwJ1ictZbEOefjdmZQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/25dmdKJiRmRVlwvqW3ulgQ',
          title: '社团介绍 | 准备好迎接新生活了吗？'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TkxyF6JYFSG0SIfVpHT1ytm9K89aia4EEKPqlcaNQLz4knedIApxjLSol9icCG1Ua1ABWpq5ABpsibg/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/27lJQlJNmIh33GMvdemhvg',
          title: '社团介绍 | 快乐的秘诀是——加入社团！'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TtjZZicBxZicYQUo2s5uDFxsPNxGSu9YXiccYxuzJfmHjB0cuG3jlCZNGRhy0Hb6ACM4skMKZo5USJg/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/llZVedHIMeqI6o8QcHH4yQ',
          title: '社团介绍 | 准备就绪 就等你来!'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TtjZZicBxZicYQUo2s5uDFxsVkphPZPCdzzjKO7OlaeXNsUDrwSdfLAiahcRz6zyUxlwNK0DHjBhXrg/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/25-Ngy5ba2Vox8xKufd2Rg',
          title: '社团介绍 | 各式各样的社团里 总有一个适合你！'
        }
      ]
    }
  },
  created () {
    this.setTwitter()
  },
  methods: {
    setTwitter () {
      this.twitterList.forEach(function (value, index) {
        if (value.picUrl === '' && value.title === '') {
          wechatApi.getTwitter(value.url).then(res => {
            value.picUrl = res.data.picUrl
            value.title = res.data.title
          })
          console.log(value)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
