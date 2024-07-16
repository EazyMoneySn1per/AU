<template>
  <div style="">
    <router-view/>
    <div style="background-color: #F1F1F1">
      <!--顶部搜索栏-->
      <div class="top-search">
        <van-icon name="apps-o " @click="showPop" size="2rem" style="margin-right: 0.1rem"></van-icon>
        <van-search shape="round" v-model="value" placeholder="请输入搜索关键词" background="#fff"
                    style="display: inline-block;width: 20rem;" @input="onSearch()" @focus="() => {this.showPopup = false}"/>
      </div>
      <!--优秀社团字样-->
      <div style="background: #fff;height: 10vh;">
        <h3 style="margin: 0 5vw">{{ currentSelectType }}</h3>
      </div>
      <div style="display: flex;align-items: flex-start;">
        <!--左侧社团种类选择栏-->
        <div style="margin-top: -12vh;height: 100%;background-color: #F1F1F1;">
          <transition name="fade">
            <van-sidebar @change="filterAssociations" v-model="activeType" v-if="showPopup"
                         style="white-space: nowrap;height: 100%;">
              <van-sidebar-item v-for="(item,key) in associationType" v-bind:key="key" :title="item"/>
            </van-sidebar>
          </transition>
        </div>
        <!--社团列表-->
        <!--        <div class="mid-list">-->
        <!--          <div v-for="(item,key) in currentAssocciations" :key="key" class="mid-list-card">-->
        <!--            <div class="mid-list-card-left">-->
        <!--              <div class="mid-list-card-left-img">-->
        <!--                <van-image style="margin-left: 1.6rem" round height="3.5rem" width="3.5rem" :src="assLogoServer+item.logo"></van-image>-->
        <!--              </div>-->
        <!--              <div class="mid-list-card-left-info">-->
        <!--                <div><h3 style="margin: 0;">{{item.assName}}</h3></div>-->
        <!--                <div style="color: #A6A6A6;width: 38vw;word-wrap: break-word">{{item.description}}</div>-->
        <!--              </div>-->
        <!--            </div>-->
        <!--            <div class="mid-list-card-right" @click="routeToAss(item.assId)">-->
        <!--              <van-icon size="2.2rem" :name="require('../assets/associations-finger.png')"></van-icon>-->
        <!--              <div style="color: #A6A6A6;margin-top:0.2rem;">点击查看</div>-->
        <!--            </div>-->
        <!--          </div>-->
        <!--        </div>-->
        <div class="mid-list">
          <div v-for="(item,key) in currentAssocciations" :key="key" class="mid-list-card">
            <div class="mid-list-card-left-img">
              <van-image style="margin-left: 2vw" round height="3.5rem" width="3.5rem"
                         :src="assLogoServer+item.logo"></van-image>
            </div>
            <div class="mid-list-card-left-info">
              <div><h3 style="margin: 0">{{ item.assName }}</h3></div>
              <div style="color: #A6A6A6;word-wrap: break-word;width: 100%">{{ item.description }}</div>
            </div>
            <div class="mid-list-card-right" @click="routeToAss(item.assId)">
              <van-icon size="2.2rem" :name="require('../assets/associations-finger.png')"></van-icon>
              <div style="color: #A6A6A6;margin-top:0.2rem;">点击查看</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div style="z-index: 99;position: relative;height: 5vh">
      <TabbarAu active-page="1"></TabbarAu>

    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { Search, Popup, Sidebar, SidebarItem, Toast } from 'vant'
import * as wechatApi from '../api/wechatApi'

Vue.use(Search)
Vue.use(Popup)
Vue.use(Sidebar)
Vue.use(SidebarItem)

export default {
  name: 'Association',
  created () {
    this.getTotalAssociation()
  },
  data () {
    return {
      activePage: 1,
      showPopup: false,
      value: '',
      totalAssociations: [],
      ExecllentAssociations: [],
      currentAssocciations: [],
      Type: [],
      currentSelectType: '优秀社团',
      temSelectType: '',
      activeType: '',
      associationType: ['优秀社团', '学术科技类', '文化体育类', '创新创业类', '学生互助类', '思想政治类', '志愿公益类'],
      assLogoServer: process.env.VUE_APP_ASSLOGO
    }
  },
  methods: {
    showPop () {
      this.showPopup = !this.showPopup
    },
    // 获取社团种类
    getAssociationType () {
    },
    // 获取所有的社团
    getTotalAssociation () {
      Toast.loading({
        message: '加载中',
        forbidClick: true,
        duration: 9999999
      })
      wechatApi.getTotalAssociation().then(res => {
        var data = res.data
        data.forEach(function (value, index) {
          if (value.assName === '社联') {
            data.splice(index, 1)
          }
        })
        this.totalAssociations = data
      })
      wechatApi.getExecllentAssociation().then(res => {
        this.currentAssocciations = this.ExecllentAssociations = res.data
        Toast.clear()
      })
    },
    // 根据社团种类筛选
    filterAssociations (index) {
      if (index !== 0) {
        this.currentSelectType = this.associationType[index]
        this.currentAssocciations = []
        for (const i in this.totalAssociations) {
          if (this.totalAssociations[i].assType === this.associationType[index]) {
            this.currentAssocciations.push(this.totalAssociations[i])
          }
        }
      } else {
        this.currentSelectType = '优秀社团'
        this.currentAssocciations = this.ExecllentAssociations
        console.log(this.currentAssocciations)
      }
      // console.log(this.currentAssocciations)
      this.showPopup = false
    },
    // 筛选优秀社团
    OutstandingAssociations () {
      // ......
      // this.currentAssocciations = ...
    },
    routeToAss (e) {
      this.$router.push({
        path: '/AssociationDetail',
        query: { assId: e }
      })
    },
    onSearch () {
      this.currentAssocciations = []
      // console.log(this.value)
      if (this.value === '') {
        console.log(this.temSelectType)
        if (this.temSelectType !== '') {
          this.currentSelectType = this.temSelectType
          this.filterAssociations(this.associationType.indexOf(this.currentSelectType))
          this.temSelectType = ''
        }
        return
      }
      // console.log(this.totalAssociations)
      this.temSelectType === '' ? this.temSelectType = this.currentSelectType : null
      this.currentSelectType = '搜索结果'
      for (var i = 0; i < this.totalAssociations.length; i++) {
        if (this.totalAssociations[i].assName.indexOf(this.value) >= 0) {
          this.currentAssocciations.push(this.totalAssociations[i])
        }
      }
    }
  }
}
</script>

<style scoped>
.top-search {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: white;
  height: 10vh;
}

.mid-list {
  margin: -6vh auto 0 auto;
  padding: 1vh 0;
  width: 100%;
  /*height: 42rem;*/
  /*background-color: #F1F1F1;*/
  display: flex;
  height: 75vh;
  overflow-y: scroll;
  /*flex-shrink: 0;*/
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}

.mid-list-card {
  padding: 1vh 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1vh 0;
  width: 90vw;
  border-radius: 15px;
  box-shadow: 0px 0px 10px -7px #242424;
  background-color: #ffffff;
}

.mid-list-card-left-info {
  width: 60%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  flex-wrap: wrap;
}

.mid-list-card-left-img {
  width: 20%;
  text-align: center;
}

.mid-list-card-right {
  /*width: 20vw;*/
  width: 20%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

/*.mid-list-card-left-info{*/
/*  margin: 0 2vw;*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  align-items: flex-start;*/
/*  justify-content: center;*/
/*}*/
/*.mid-list-card-right{*/
/*  display: flex;*/
/*  flex-direction: column;*/
/*  align-items: center;*/
/*  justify-content: center;*/
/*  right: 0;*/
/*  width: 30%;*/
/*  height: 100%;*/
/*  border-radius: 0px 15px 15px 0px;*/
/*  background-color: #ffffff;*/
/*}*/

.van-search__content {
  height: 2.5rem;
}

.van-search .van-cell {
  line-height: 2rem;
}

.van-sidebar {
  width: 25vw;
  /*height: 80vh;*/
  background-color: #F1F1F1;
  position: absolute;
  z-index: 4;

}

.van-sidebar-item {
  font-size: 14px;
  color: #797979;
  background-color: transparent;
}

.van-sidebar-item--select {
  background-color: #F1F1F1;
}

.van-sidebar-item--select::before {
  background-color: #3873AE;
  height: 1.3rem;
  width: 0.13rem;
  margin-left: 0.3rem;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
  transition: width .5s;

}

.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */
{
  opacity: 0;
  width: 0;
}
</style>
