<template>
  <div class="page">
    <div class="page-nav">
      <div class="bread">导播列表</div>
    </div>
    <search-field title="使用指引">
      <Guide :guideList="guideList"/>
    </search-field>
    <div class="page-content">
      <div class="button-field">
        <el-button type="primary" size="small" @click="createDirector"
          >创建导播台</el-button
        >
      </div>
      <LiveTable ref="liveTable" from="live" :switchView="switchView" />
    </div>
    <el-dialog title="创建导播" :visible.sync="dialogVisible" width="700px">
      <el-form
        label-position="right"
        label-width="105px"
        :model="createDirectorModel"
        ref="form"
        :rules="rules"
      >
        <el-form-item label="导播台名称" required prop="CasterName">
          <el-input
            v-model="createDirectorModel.CasterName"
            placeholder="最大长度32位"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="onSubmit">确 定</el-button>
        <el-button @click="dialogVisible = false">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import searchField from '@/components/tools/searchField';
import Guide from '@/components/functional/guide';
import LiveTable from '@/components/functional/liveTable';
import { service } from '@/api';
export default {
  data: () => ({
    switchView: 0,
    guideList: [{
      title: '创建导播台',
      intro: '在导播台列表界面，点击创建导播台按钮，填写相关信息，点击确定创建导播台。'
    }, {
      title: '管理导播台',
      intro: '在导播台列表点击进入导播台'
    }, {
      title: '进入直播观看页',
      intro: '在导播台详情页中，点击进入直播地址按钮复制链接，打开观看，输入观众ID，进入直播观看端。'
    }, {
      title: '管理导播',
      intro: '回到导播台中，对当前的直播进行进行管理。'
    }],
    createDirectorModel: {
      CasterName: '',
    },
    rules: {
      CasterName: [
        { required: true, message: '请输入导播台名称', trigger: 'blur' },
        { max: 32, message: '最大长度32位', trigger: 'blur' },
      ],
    },
    dialogVisible: false,
  }),

  methods: {
    createDirector() {
      this.dialogVisible = true;
    },

    onSubmit() {
      service.CasterApiPopService({
        ActionName: 'CreateCaster',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          ChargeType: 'PostPaid',
          ClientToken: localStorage.getItem('userId'),
          NormType: 1,
          CasterName: this.createDirectorModel.CasterName
        }
      }).then(res => {
        this.$message.success('创建成功');
        this.$refs.liveTable.listLiveRooms()
        this.createDirectorModel = {
          CasterName: '',
        }
        this.dialogVisible = false;
      }).catch((err) => {
        this.$message.error('创建失败');
        console.error(err);
      })
        .finally(() => {
          this.$refs.form.resetFields();
        });
    },
  },

  components: {
    LiveTable,
    searchField,
    Guide,
  },
};
</script>

<style scoped lang="less">
.tips {
  display: block;
  font-size: 13px;
  color: #bbb;
  margin-bottom: 10px;
  text-align: center;
}
.ext-item {
  display: flex;
}
.button-field {
  position: relative;
}
/deep/ .el-radio-group {
  position: absolute;
  bottom: 18px;
  right: 0;
}
.dialog-title {
  .title {
    font-size: 14px;
    font-weight: 500;
    display: inline-block;
    margin-right: 30px;
  }
}
.bread{
    height: 40px;
    line-height: 40px;
}
</style>
