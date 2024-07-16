import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue'),
    meta: {
      title: '社联管理平台'
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/my',
    name: '我的',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/My.vue')
  },
  {
    path: '/association',
    name: '社团',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Association.vue')
  },
  {
    path: '/bindid',
    name: '绑定信息',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/BindId.vue')
  },
  {
    path: '/AssociationDetail',
    name: '无',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AssociationDetail.vue')
  },
  {
    path: '/AssociationDescription',
    name: '无',
    component: () => import(/* webpackChunkName: "about" */ '../views/AssociationDescription.vue')
  },
  {
    path: '/InterviewSubmit',
    name: '无',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/InterviewSubmit.vue')
  },
  {
    path: '/MyInterview',
    name: '无',
    component: () => import(/* webpackChunkName: "about" */ '../views/MyInterview.vue')
  },
  {
    path: '/AuActivities',
    name: '无',
    component: () => import(/* webpackChunkName: "about" */ '../views/AuActivities.vue')
  },
  {
    path: '/AuDepartment',
    name: '无',
    component: () => import(/* webpackChunkName: "about" */ '../views/AuDepartment.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
