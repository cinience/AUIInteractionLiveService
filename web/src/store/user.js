import Vue from 'vue'

export const USER_NICK = 'USER_NICK'
export const USER_ID = 'USER_ID'
export const APP_ID = 'APP_ID'
export const APP_KEY = 'APP_KEY'
export default {
  state: {
    userId: '',
    userNick: '',
    menus: '',
    appId: '',
    appKey: ''
  },
  getters: {
    userNick: state => {
      return state.userNick || localStorage.getItem('userNick')
    },
    userId: state => {
      return state.userId || localStorage.getItem('userId')
    },
    menus: state => {
      return state.menus || {}
    },
    appId: state => {
      return state.appId || localStorage.getItem('appId')
    },
    appKey: state => {
      return state.appKey || localStorage.getItem('appKey')
    }
  },
  mutations: {
    [USER_NICK](state, userNick) {
      state.userNick = userNick
      localStorage.setItem('userNick', userNick)
    },
    [APP_ID](state, appId) {
      state.appId = appId
    },
    [APP_KEY](state, appkey) {
      state.appKey = appkey
      localStorage.setItem('appKey', appkey)
    },
    [USER_ID](state, userId) {
      state.userId = userId
      localStorage.setItem('userId', userId)
    }
  }
}
