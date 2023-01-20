/* eslint-disable */ //Disable for process
import axios from 'axios'
import Vue from 'vue'

function setInterceptors (axiosInstance) {
  // request interceptor
  axiosInstance.interceptors.request.use(req => {
    req.headers['x-app-id'] = localStorage.getItem('appId')
    // 让服务端默认风格为小驼峰
    req.headers['x-key-style'] = 'LowerCamel'

    if (process.env.NODE_ENV !== 'production') {
      console.time(req.method.toUpperCase() + ' ' + '/api' + req.url)
    }
    return req
  }, error => {
    console.error(error)
    return Promise.reject(error)
  })

  // response interceptor
  axiosInstance.interceptors.response.use(res => {
    console.log(res.code, '=code');
    if(res.config.url.search('userLogin') > 0 && !localStorage.getItem('appId')){
      localStorage.setItem('appId',res.headers['x-app-id'])
    }
    if (process.env.NODE_ENV !== 'production') {
      console.timeEnd(res.config.method.toUpperCase() + ' ' + res.config.url)
    }
    if(res.code === '401'){
      this.$router.replace({
        name: 'login'
      })
    }
    return res
  }, error => {
    console.error(error.response.data.code)
    if (error.code === 'ECONNABORTED') {
      console.error('网络错误，请检查网络连接刷新后重试')
    }if (error.response.data.code === 'ServiceUnavailable') {
      Vue.prototype.$message.error('密码错误')
    }else {
      console.dir(error)
      Vue.prototype.$message.error(error.response.data.message)
    }
    return Promise.reject(error)
  })
}

// create axios instance
function getAxios (settings, timeout = 10000) {
  const axiosInstance = axios.create(settings)
  axios.defaults.timeout = timeout
  setInterceptors(axiosInstance)
  return axiosInstance
}

export const api = getAxios({
  baseURL:'/',
  withCredentials: true,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})
