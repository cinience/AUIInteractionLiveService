<template>
 <div class="page">
    <div class="page-nav">
      <div class="bread">下载移动端</div>
    </div>
    <search-field title="使用指引">
      <Guide :guideList="guideList"/>
    </search-field>
    <div class="page-content">
    <div>
      <span>下载APP</span>
      <el-row  :gutter="24">
        <el-col :span="6">
       <div class="qr-code">
        <img :src="img"/>
      </div>
      <div class="info">
        <span>Android</span>
      </div>
        </el-col>
            <el-col :span="6">
       <div class="qr-code">
        <img :src="img"/>
      </div>
      <div class="info">
        <span>ios</span>
      </div>
        </el-col>
      </el-row>
      </div>
      <div style="margin-top:20px">
       <span>配置项</span>
      <el-row  :gutter="24" style="margin-top: 20px">
        <el-col :span="6">
       <div class="qr-code">
        <img :src="img"/>
      </div>
      <div class="info">
        <span>Android</span>
      </div>

        </el-col>
            <el-col :span="6">

       <div class="qr-code">
        <img :src="img"/>
      </div>
      <div class="info">
        <span>ios</span>
      </div>

        </el-col>
      </el-row>
      </div>
    </div>
 </div>
</template>
<script>
import { liveApi } from '@/api';
import searchField from '@/components/tools/searchField';
import Guide from '@/components/functional/guide';
export default {
  data() {
    return {
      appName: document.title,
      img: '',
      guideList: [{
        title: '第一步下载APP',
        intro: 'Android版请使用浏览器扫码下载（操作系统要求：Android 4.4及以上），iOS版请使用手机系统相机扫码下载（操作系统要求：iOS 10.0及以上）。'
      }, {
        title: '第二步配置',
        intro: '扫码自动填充配置项'
      }, {
        title: '完成',
        intro: ''
      }]
    };
  },
  created() {
    this.getAppConfigQRcode();
  },
  components: {
    Guide,
    searchField
  },
  methods: {
    getAppConfigQRcode() {
      liveApi.getAppConfigQRcode({
        deviceType: 'android',
      }).then((res) => {
        if (res.data.result) {
          this.img = `data:image/png;base64,${res.data.result.data}`;
        }
      });
    },
  },
};
</script>
<style scoped lang="less">
.bread{
    height: 40px;
    line-height: 40px;
}
/deep/ .el-col{
    display: flex;
    flex-direction: column;
    align-items: center;
}
.qr-code{
  width: 200px;
        // text-align: center;
    }
    .info{
        // margin-top: 20px;
        // text-align: center;
    }
</style>
