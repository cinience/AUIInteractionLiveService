<template>
  <div class="page">
    <div class="page-nav">
      <div class="bread">配置</div>
    </div>
    <div class="page-content">
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="修改密码" name="modifyPassword">
            <div class="modifyPassword">
          <el-form
            :model="form"
            status-icon
            :rules="rules"
            ref="ruleForm"
            label-width="100px"
            class="demo-ruleForm"
          >
            <el-form-item label="新密码" prop="pass">
              <el-input
                type="password"
                v-model="form.pass"
                autocomplete="off"
                placeholder="请输入6到18位字符，至少包含一个数字，一个字母和一个特殊字符"
              ></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
              <el-input
                type="password"
                v-model="form.checkPass"
                autocomplete="off"
                placeholder="请再次输入新的密码"
              ></el-input>
            </el-form-item>
          </el-form>
           <div class="btn" >
              <el-button type="primary" @click="submitForm('ruleForm')"
                >确认提交</el-button
              >
            </div>
            </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script>
import { liveApi } from '@/api';
export default {
  data() {
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (/^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{6,18}$/.test(value)) {
          if (this.form.checkPass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
        } else {
          callback(new Error('请输入正确密码格式'));
        }

        callback();
      }
    };
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.form.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, trigger: 'blur' }
        ]
      },
      activeName: 'modifyPassword',
      form: {
        pass: '',
        checkPass: '',
      },

    }
  },
  methods: {
    handleClick() {},
    submitForm() {
      if (this.form.checkPass === '') return this.$message.error('请输入完整')
      liveApi.updatePassword({
        accountId: 'admin',
        password: this.form.checkPass
      }).then(res => {
        console.log(res, '======res');
        if (res.data.Code === 'Success') {
          this.$alert('密码修改成功，请重新登录', '提示', {
            confirmButtonText: '确定',
            callback: action => {
              localStorage.removeItem('appId')
              localStorage.removeItem('userId')
              localStorage.removeItem('appName')
              localStorage.removeItem('appKey')
              localStorage.removeItem('userNick')
              this.$router.replace({
                name: 'login'
              })
            }
          });

        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  },
};
</script>
<style scoped lang="less">
/deep/ .el-form-item__label{
    text-align: left;
}
/deep/ .el-form-item__content{
margin-left: 0!important;
}
.modifyPassword{
    width: 460px;
    margin-left: 65px;
   /deep/ .el-button--primary{
     width: 100%;
    }
}
.bread{
    height: 40px;
    line-height: 40px;
}
</style>
