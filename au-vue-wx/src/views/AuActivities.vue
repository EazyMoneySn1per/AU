<template>
  <div>
    <van-tabs color="#0D4E9A" sticky v-model="active">
      <van-tab title="2021">
        <twitter-group v-for="item in twitterList_2021" v-bind:key="item.time" :title="item.title" :url="item.url" :pic-url="item.picUrl" :time="item.time"></twitter-group>
      </van-tab>
      <van-tab title="2020">
        <twitter-group v-for="item in twitterList_2020" v-bind:key="item.time" :title="item.title" :url="item.url" :pic-url="item.picUrl" :time="item.time"></twitter-group>
      </van-tab>
      <van-tab title="2019">
      </van-tab>
      <van-tab title="2018">
      </van-tab>
      <van-tab title="2017">
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from 'vue'
import { Tab, Tabs, Sidebar, SidebarItem } from 'vant'
import TwitterGroup from '@/components/TwitterGroup'

Vue.use(Tab)
Vue.use(Tabs)
Vue.use(Sidebar)
Vue.use(SidebarItem)
export default {
  name: 'AuActivities',
  components: { TwitterGroup },
  data () {
    return {
      active: 0,
      twitterList_2021: [
        {
          time: '2021-06-24',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TGuiaIPSub7ibEebezete1Ljz03DkhVkmNzOSq1iazXnDfsqHyetSVrVs15tiarjb15u5I5mDteQdjeQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/m_90_3fNTr6Ingz2nk2Y2g',
          title: '不说再见 | 不说再见是我们的秘密喔，毕业快乐！'
        },
        {
          time: '2021-06-22',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77QkG24t2icYIQLGzWgEaXT4dg4KGfDmK18dNjCVCDbhrdF37PCUnXqicSzPFKicNeDQuBFla56iaib482g/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/m_90_3fNTr6Ingz2nk2Y2g',
          title: '不说再见 | 深圳技术大学首届毕业晚会'
        },
        {
          time: '2021-06-19',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77SgvpaM50IICtdxprze6UhLeIvHia4enEMlYkolOQdaJrTJDZh1ezzjFvvUiaV5eicDxwUWCuIHUsgwQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/bYP6tVlm8rylVV9KC7mmwQ',
          title: '新的起点由我接力，带你去看未来风景'
        },
        {
          time: '2021-06-13',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TwWaDXicNZc8DUE4NEW0nqHzaNlSyqv6qibTnmyTzM1dxgJJBI1Rj1Bre5MMCf3XPlYFk5XVVib4YdA/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/KgSmBf5iVi1QQbNJJux3fw',
          title: '建党百年 | 社联邀你看电影！！'
        },
        {
          time: '2021-05-31',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77RTshKfvbDiaCc1skd0UmAkyJAczfPPCa7FW8hrRKPneIbXBiaYuUialSbltE0YMZY8TAYyX0tKqVDSw/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/QNs_Eo9HlaZLzSZ-oSjOQA',
          title: '社团之夜 | 快乐星球一日游已结束，即将开启下一段旅途'
        },
        {
          time: '2021-05-28',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TgCLJ3V7hH7orcVriaRiam4Za5oV3Ric6Lm3dzgCcANq8YRJDwuKiagxD8HiaQicbEVgSeia2xticL6o10AQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/ZClnmFzROEIgPUlkPqTAXw',
          title: '社团之夜 | 即将抵达快乐星球，“童”行须知待你查收'
        },
        {
          time: '2021-05-25',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TIGeeoqImoAMOJNnPibqQeEyNGqcdZ7uMWCsofaPRlcQdJrFjtYK9pZse0cYukmf8qXyYE8Xqz7uA/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/b1FwDPtNH_QrgBBUl2wAMQ',
          title: '社团之夜 | 抓紧时间上社团号大巴，一“童”前往快乐星球！'
        },
        {
          time: '2021-04-20',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77SvkDxawA2DliblCJ0OAghaVOFtNMb5ib6apqyLUu8p9Th0UykQicIbYYjMTkE4zQcVN2PeuNz42pRJQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/CJmRkluC1r_i8nrrqHDY5Q',
          title: '青篮杯 | 你订阅的青篮杯周报已到！请查阅！'
        },
        {
          time: '2021-04-17',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Q7MhSmF00V2BTSOArfRqwHg5pE0eCUibib1mcVxhC1QYAIYNVf226VjZ2V5mxIpTFFojRiabvVXJkTA/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/OPGpN4TASm5Vnnx-woQwsw',
          title: '青篮杯 | 第三届青篮杯火热开赛！周末和篮球赛更配哦~'
        },
        {
          time: '2021-04-16',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77SK8doZibOSfSf7TjujC6B6AotfZJpCT4XE9SDD5IfQwx5NeYXL2uHzqUOjFk1YsfJ3TNAo6gVvu7g/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/p-xCCW4KHfjvfjyTkyaGkw',
          title: '青篮杯 | 详细赛程新鲜出炉，了解一下？'
        },
        {
          time: '2021-04-07',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77ShDSdibQaibs3K0GKcr5ricAufNLOxz6fOBrh8iadCS5oCLnp1ej2GI28dgNS91hdqLIRq5ns69AL0aw/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/1Z8-bG5kZBQ381U5MI9ILw',
          title: '青篮杯 | 比赛时间定了！还可以这样？篮球人们欲报从速速速速速'
        },
        {
          time: '2021-04-05',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Rib64FFgeZPOoOCXwgkl2TopCHHAN5P7DffdP52fXJhSMcR2fAceicthOIp41kOajOapIkiaqdvXvxw/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/5AGxHmu_1hvn3U4dbBQ9Hw',
          title: '青篮杯 | 交友通道已开启，请注意查收'
        },
        {
          time: '2021-04-01',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Tkrz4gTVszOgNhI2dRg1gq5eK79c92MMLt7HKibz3FGsxnzJmMA6PIDZUcmibMB2zmagWUoUicEatkQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/gE6SoOAWuoQy4l-m1WUezw',
          title: '青篮杯 | 第三届青篮杯报名开启！'
        },
        {
          time: '2021-04-13',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77Rss1ZYDuT7NI2lE4nY12zgqpXtDOXdPZ5H9PUtpVSI2zrvhA9Y40umKNWP6RHicWjqwssM6uiaiadTA/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/oyB4GKWAgnrJ64a_Stgq9A',
          title: 'Begin Again | 我们唱完了，你听到了吗？'
        },
        {
          time: '2021-04-08',
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77TibFHUhic9BBJ4lEw6ibKtialn88NvscUf2Skrn8xLjgXooibLmPQlHaNxTXnUm5EcnudcdaxQYHvXGBQ/0?wx_fmt=jpeg',
          url: 'https://mp.weixin.qq.com/s/uCagxcKICF4tX9s6uc48qg',
          title: 'Begin Again | 您收到一封来自十大歌手的邀请函'
        }
      ],
      twitterList_2020: [
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77R6c2GNR6ib4lQLbKAnCTiclTjSFbuZRzeGiciajRCrib7DprreDBHicPNOZ1QE5QAFhvvuxPqspnfpwkIw/0?wx_fmt=jpeg',
          time: '2020-12-31',
          title: '年度总结 | 期待有你的下一个四季',
          url: 'https://mp.weixin.qq.com/s/L-Jnux_8WKij-EuCxuRn1A'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/S2QIjGmZ77RP5L8wooRWeMCbf11YfCg4mkgtocw8icUsibBd4I5BiaKj444P7Uly1U8SjwFl46xic0SJS9qUv2u2lA/0?wx_fmt=jpeg',
          time: '2020-12-17',
          title: '十大歌手 | 邀你共赴一场乐宴',
          url: 'https://mp.weixin.qq.com/s/wwYqmXKWQbeM24g9PKJvCA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak1HYR4C29saMpmhhQROichnicGxNxzwz2Pibym7Vqww2G9u9fKTPf0PP4fMY3WiaI1zicusJM0aKc92Zpg/0?wx_fmt=jpeg',
          time: '2020-12-06',
          title: '十大歌手 | 明日歌王由你来定！',
          url: 'https://mp.weixin.qq.com/s/7QmLzLLSeQarp-LPehBdDw'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak2U3QJDFlpVru7592icibxs1Of3pb2eXN7nJy2MkzHkLg6SjhG7n3h1EFsrBRicX19M7ia7ic38j8TRicSw/0?wx_fmt=jpeg',
          time: '2020-11-29',
          title: '十大歌手 | One pick晋级选手人气投票进行中！',
          url: 'https://mp.weixin.qq.com/s/VG7eIM66aRjP611t9QS_IA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak1wTZmD3WvnGpHCpyOMNg4BibzibYiarOWZFb1VQeac1T6B5icJWT8jRvZ20Ze4ibiaprY9PXdGQ009x7KA/0?wx_fmt=jpeg',
          time: '2020-11-14',
          title: '十大歌手 | 你有一份观赛须知待查收',
          url: 'https://mp.weixin.qq.com/s/6DVINrcvJBaftwj8ZHJ2KA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak2ibNH9SclfBHbymE048ECBEa3LkEVOZsazeF4YInB9DlHDgB1z3QGNTVCWS44L43DKXIwfzsng0oA/0?wx_fmt=jpeg',
          time: '2020-11-10',
          title: '十大歌手 | 炽诚初心 放肆歌唱',
          url: 'https://mp.weixin.qq.com/s/PSNaMARiAlEeQuj8EK8BBA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak3N6xsf3yKXkNyExU4NHiaLwOGQeicibbXxDr1YJbqPxAxKgb57fKtHayeicpicG0icaOicteA2niaJrRG1bQ/0?wx_fmt=jpeg',
          time: '2020-10-25',
          title: '十大歌手 | 歌染秋韵 声而不凡',
          url: 'https://mp.weixin.qq.com/s/Xp6iI2VfmlAKU5CYQ0BJoQ'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak3gXYBf0godn5p5SeAtDeicOC5I07gMiax7iczu9unn3DIFmXVNh7m64VV91PJZhAfn5MOdTMrw1PnxQ/0?wx_fmt=jpeg',
          time: '2020-11-06',
          title: '硕果金秋，与你同行',
          url: 'https://mp.weixin.qq.com/s/ehw7pcPvzGMkYoFPgCvzEQ'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak3kQGa8iaS1c5s7zHru0sQ8w3VREdf57I5sXYXYVtzexPdwZyCnMkIfYmqBpXP5D6I5wZY1e6HAc3w/0?wx_fmt=jpeg',
          time: '2020-10-14',
          title: '深圳技术大学2020年招新嘉年华圆满落幕',
          url: 'https://mp.weixin.qq.com/s/9DBWs90ffHqgweZeluAe_Q'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak2rRjnrXxHiaP5ltjibvrP0tpZGle8NyFkLDx4oaAIBTf04c7B3xfSDJUnFfyhcGZpFPyfBOp5N2jmQ/0?wx_fmt=jpeg',
          time: '2020-10-02',
          title: '招新嘉年华 | 社彩秋韵 联迎青衿',
          url: 'https://mp.weixin.qq.com/s/APCXl72g_5H9CuYHbciq0A'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak0kO8KGAv7jrR8MM5w4O7TxC3tyErU8hNOqiaNBOpflLME0MH9cdFToWOicjmzA2gQ70ugd64tEfjNA/0?wx_fmt=jpeg',
          time: '2020-09-18',
          title: '社联活动回顾',
          url: 'https://mp.weixin.qq.com/s/pueNaMtaIe0o5L9XcnUd3A'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak3p38s0ecjnaNNPS7hWvIicXZORf8NF89hGQ9To7kc9f3SlPxe58PMWtZq69ITPdfwoiccVbIMLGBEA/0?wx_fmt=jpeg',
          time: '2020-06-29',
          title: '2019-2020社团联合会换届大会',
          url: 'https://mp.weixin.qq.com/s/NGh1M_FRizvHDlV3S0miqA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak0E31wE6jIZicklkJia8OfIBdYLlSoH8sRplibX092La9WlBSF71Dd9ia25pSOe7ZQ7eR9QxBNbxV9b4Q/0?wx_fmt=jpeg',
          time: '2020-06-22',
          title: '深圳技术大学第二届青篮杯顺利落幕',
          url: 'https://mp.weixin.qq.com/s/5gCSkKuay-LKlf1Q8VjwQQ'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak14TM9iaLjoYOVskCsoNkpHs2A2NXBqxArB8pmBrJsPRYPeKmxvLsEaVIBApAibQMgeDFSEbdliaGguQ/0?wx_fmt=jpeg',
          time: '2020-06-12',
          title: '青篮杯来啦！！！',
          url: 'https://mp.weixin.qq.com/s/ExyUiRw-hPj4zhjdVtatpg'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak0jzl6RFotLW5AGBeNVHS0GrUZrsIZZaIicoricoCY3YsiaZ55yXPVDwmlEClXDPdVBDGl0q9sicTuLUw/0?wx_fmt=jpeg',
          time: '2020-05-27',
          title: '青篮杯|青春战“疫”，篮网出击',
          url: 'https://mp.weixin.qq.com/s/kIzh2HrFJSvlBpt6-5E0tQ'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak26rYPToJCr8gCrib7QGN2USm6xoFMB8HqzXmdhrAWDRC8UsO6tBbG08gy8OmCpxG4wLJSedNibTZeg/0?wx_fmt=jpeg',
          time: '2020-06-06',
          title: '童年回忆录 | 闭幕',
          url: 'https://mp.weixin.qq.com/s/jAKJ5XVdsf_BRVunKYfR-Q'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak080iaLicUBcA242XBHolarwY2DLuc0eUyOloXvzzFnhQ0II5v0XGvicia2OazpibCKYYIohex36XHnRUg/0?wx_fmt=jpeg',
          time: '2020-06-02',
          title: '童年回忆录 | 配音篇',
          url: 'https://mp.weixin.qq.com/s/E_6jkjJ6UH8ruNhygIVS0g'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak3Lx7dCHZOpiarfprWSBbCErGEBq1J9TQNyn5JFus9MhozQgELBW1aGGibyRib3shrlMJYNxAEe9Rs2w/0?wx_fmt=jpeg',
          time: '2020-06-01',
          title: '童年回忆录 | 趣事篇',
          url: 'https://mp.weixin.qq.com/s/Z1CC6Ta5577O8oV0gYdXlA'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak2pnUMBInOuGSX5GZHGsuhkzvn8x4Tpfboiadhcw8DHXUyt2X2z9rP66DtkvBq2AorKLHDLAGAuiceQ/0?wx_fmt=jpeg',
          time: '2020-05-23',
          title: '童年回忆录',
          url: 'https://mp.weixin.qq.com/s/JDlYTD1lS2Y08rlULxXhWw'
        },
        {
          picUrl: 'http://mmbiz.qpic.cn/mmbiz_jpg/VBL2VDyyak0KibJebeib34jtbPPiaQ6P9aodIFBgNoFWeOB4P70X6Cho4cVzuGEraAEoWiaFL73qCibicNI64MgB0Yrg/0?wx_fmt=jpeg',
          time: '2020-05-06',
          title: '童年回忆录',
          url: 'https://mp.weixin.qq.com/s/HBSwHTTMY8eZfdZmJ6ReDQ'
        }
      ]
    }
  },
  methods: {
    lineToTwitter (e) {
      window.location.href = e
    }
  }
}
</script>

<style>
</style>
