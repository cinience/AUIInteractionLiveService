import Vue from 'vue'
import router from './routerMap.js'
import store from '@/store'
// import { splitSearch } from '@/libs/utils'
// const query = splitSearch(window.location.search)
// const setTitle = (route) => {
//   document.title = route.meta.title ? route.meta.title : (query.appName ? decodeURI(query.appName.toString()) : '低代码音视频工厂')
// }

router.beforeEach((to, from, next) => {
  let auth = to.meta.auth
  let isLogin = localStorage.getItem('appId')
  if (auth && !isLogin) {
    console.log('************')
    console.log(auth)
    console.log(isLogin)

    Vue.prototype.$message.info('尚未登录')
    return next({ name: 'login' })
  }
  if (!to.meta.noBread) store.dispatch('BREAD', to.matched)
  console.log(to, '===yo');
  if (to.name === 'selectScene') {
    return next({ name: 'liveList' })
  }
  next()
})
router.afterEach(route => {
  // setTitle(route)
})

export default router
