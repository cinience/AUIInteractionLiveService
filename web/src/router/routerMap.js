import Vue from 'vue'
import Router from 'vue-router'

import Frame from '@/components/frames/frame'
// const Layout = () => import('@/components/frames/layout')
import Login from '@/views/common/login'
import Share from '@/views/common/share'
import ErrorPage from '@/views/common/404'

import LiveList from '@/views/live/liveList'
import LiveDetail from '@/views/live/liveDetail'
import DirectorDetail from '@/views/live/directorDetail'
import DirectorList from '@/views/live/directorList'
import Setting from '@/views/live/setting'
import MobileDownload from '@/views/live/mobileDownload'

import EnterpriseLiveList from '@/views/enterprise/liveList'
import EnterpriseLiveDetail from '@/views/enterprise/liveDetail'

import SelectScene from '@/views/selectScene'

Vue.use(Router)

export const routerMap = [{
  path: '/live',
  name: 'frame',
  component: Frame,
  children: [
    {
      path: '/live/liveList',
      name: 'liveList',
      component: LiveList,
      meta: {
        auth: true,
        title: '直播列表',
        icon: 'el-icon-tickets',
        scene: '直播-通用'
      }
    }, {
      path: '/live/liveDetail/:liveId',
      name: 'liveDetail',
      component: LiveDetail,
      meta: {
        auth: true,
        title: '直播详情',
        notMenu: true,
        scene: '直播-通用'
      }
    },
    {
      path: '/live/directorList',
      name: 'directorList',
      component: DirectorList,
      meta: {
        auth: true,
        icon: 'el-icon-tickets',
        title: '导播台列表',
        scene: '直播-通用'
      }
    },
    {
      path: '/live/directorDetail/:CasterId',
      name: 'directorDetail',
      component: DirectorDetail,
      meta: {
        auth: true,
        title: '导播台',
        notMenu: true,
        scene: '直播-通用'
      }
    },
    {
      path: '/live/MobileDownload',
      name: 'mobileDownload',
      component: MobileDownload,
      meta: {
        auth: true,
        title: '下载移动端',
        notMenu: true,
        scene: '直播-通用'
      }
    }, {
      path: '/live/setting',
      name: 'setting',
      component: Setting,
      meta: {
        auth: true,
        title: '配置',
        icon: 'el-icon-setting',
        scene: '直播-通用'
      }
    }
  ],
  redirect: '/live/liveList'
}, {
  path: '/enterprise',
  name: 'frame',
  component: Frame,
  children: [
    {
      path: '/enterprise/liveList',
      name: 'EliveList',
      component: EnterpriseLiveList,
      meta: {
        auth: true,
        title: '直播列表',
        icon: 'el-icon-setting',
        scene: '企业直播'
      }
    }, {
      path: '/enterprise/liveDetail/:liveId',
      name: 'EliveDetail',
      component: EnterpriseLiveDetail,
      meta: {
        auth: true,
        title: '直播详情',
        notMenu: true,
        scene: '企业直播'
      }
    }
  ],
  redirect: '/enterprise/liveList'
}, {
  path: '/',
  name: 'frame',
  component: Frame,
  children: [
    {
      path: '/selectScene',
      name: 'selectScene',
      component: SelectScene,
      meta: {
        auth: true,
        title: '场景选择',
        notMenu: true,
        hideMenu: true
      }
    },
  ],
  redirect: '/selectScene'
}]

const router = new Router({
  mode: 'hash',
  linkActiveClass: 'is-active',
  routes: [{
    path: '/login',
    name: 'login',
    component: Login
  }, {
    path: '/share',
    name: 'share',
    component: Share,
    meta: {
      title: localStorage.getItem('appName')
    }
  },
  {
    path: '*',
    name: 'error',
    component: ErrorPage
  }].concat(routerMap)
})

export default router
