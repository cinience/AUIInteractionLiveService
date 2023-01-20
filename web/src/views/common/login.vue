<template>
  <div class="login">
    <div class="login-form" @keyup.enter="login">
      <div class="login-title">
        <h1>{{appName}}</h1>
      </div>
      <label for="userNick">账号</label>
      <input disabled type="userNick" name="userNick" v-model="form.userNick" id="userNick" placeholder="请输入中文字符、字母或数字，32位以内">
      <label for="passWord">密码</label>
       <div class="passWord">
      <input type="passWord" name="passWord" v-model="form.passWord" id="passWord" placeholder="请输入密码">
        <!-- <el-popover
        placement="top-start"
        title=""
        width="230"
        trigger="hover">
         <a href="https://pre-imp.console.aliyun.com" target="_blank">查看密码(低代码集成服务密钥)</a>
       <i slot="reference">?</i>
      </el-popover> -->
    </div>
      <button class="login-btn" @click="login">登录</button>
      <div class="link">
        <el-button type="text" @click="openDoc">如何对接你的账号，进行登录？</el-button>
      </div>
    </div>
  </div>
</template>
<script>
import MD5 from 'crypto-js/md5'
import { liveApi } from '@/api'

const userReg = /^[\u4e00-\u9fa5a-zA-Z0-9]+$/
export default {
  data: () => ({
    form: {
      userNick: 'admin',
      userId: '',
      passWord: '',
    },
    appName: document.title
  }),
  methods: {
    login() {
      if (!this.form.userNick) return this.$message.error('请输入账号昵称')
      if (!userReg.test(this.form.userNick) || this.form.userNick.length > 32) return this.$message.error('请输入中英文与数字，32位以内')
      this.form.userId = MD5(this.form.userNick).toString().substr(0, 20)
      this.$store.commit('USER_ID', this.form.userId)
      this.$store.commit('USER_NICK', this.form.userNick)
      console.log(this.$store.getters.userNick);
      liveApi.login({
        accountId: 'admin',
        password: this.form.passWord
      }).then(res => {
        console.log(res, '登录成功', res.data.scene)
        this.$router.push({ path: `/${res.data.scene}/liveList` })
        liveApi.getUserInfo({ userId: 'admin' }).then(res => {
          console.log(res, '=userInfo');
          this.$store.commit('APP_ID', res.data.appId)
          this.$store.commit('APP_KEY', res.data.appKey)
        })
      })
    },
    openDoc() {
      this.$message.info('敬请期待')
    }
  }
}
</script>

<style lang="less" scoped>
.publicInput{
  display: block;
  width: 400px;
  height: 45px;
  margin: 0 0 15px;
  border-radius: 5px;
}
.login{
  width: 100%;
  height: 100%;
  background: #324057;
  .login-title{
    margin-bottom: 60px;
    h1{
      color: #fff;
      font-size: 40px;
      text-align: center;
    }
  }
  .login-form{
    position: absolute;
    top: 40%;
    left: 50%;
    transform: translate(-50%,-50%);
    -webkit-transform: translate(-50%,-50%);
    -ms-transform: translate(-50%,-50%);
    label {
      color: #fff;
      font-size: 15px;
      line-height: 2.2;
    }
    input{
      .publicInput;
      border: 1px solid #fff;
      background: transparent;
      color: #fff;
      padding: 0 18px;
    }
    .passWord{
      position: relative;
      /deep/ .el-popover__reference-wrapper{
        position: absolute;
          top: 25%;
          right: 10px;
          cursor: pointer;
      }
      i {
        color: #fff;
        font-size: 15px
      }
    }
    .login-btn{
      .publicInput;
      background: #fff;
      color: #324057;
      border: none;
      text-align: center;
    }
    .link {
      text-align: center;
    }
  }
}
</style>
