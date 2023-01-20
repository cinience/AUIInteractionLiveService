<template>
  <div class="page">
    <div class="page-nav">
      <div class="bread">导播台列表 / 导播台</div>
    </div>
    <div class="page-content">
      <div class="title">
        <i
          class="el-icon-back"
          @click="() => $router.replace('/live/directorList')"
          style="cursor: pointer"
        ></i>
        {{ directorConfig.CasterName }}
      </div>
      <div class="btn-container">
        <el-button type="primary" size="small" @click="directorConfigInfo"
          >直播设置</el-button
        >
        <el-button type="success" size="small" @click="startLive"
          >开启直播</el-button
        >
        <el-button
          type="danger"
          size="small"
          @click="stopLive"
          v-show="liveDetail.status === 1"
          >停止直播</el-button
        >
        <el-button type="success" size="small" @click="syncPgm"
          >同步到PGM</el-button
        >
        <el-button type="info" size="small" @click="switchSources"
          >切换备播视频</el-button
        >
        <el-popover
          id="popover"
          placement="bottom"
          width="380"
          v-model="addressVisiable"
          style="margin-left: 10px"
        >
          <div class="address-list">
            <div class="item">
              <div class="name">Web端</div>
              <div class="content">
                <el-input
                  :readonly="true"
                  size="small"
                  id="web-share"
                ></el-input>
                <el-button
                  size="small"
                  type="primary"
                  class="copy"
                  >复制</el-button
                >
              </div>
            </div>
          </div>
          <el-button slot="reference" size="small"
            >播放地址 <i class="el-icon-caret-bottom"></i
          ></el-button>
        </el-popover>
      </div>
      <div class="main-container">
        <div class="play-info">
          <live-data :liveDetail="liveDetail" :countTime="countTime" />
        </div>
        <div class="player-box">
          <div class="pvm-box">
            <i class="pvm-icon"></i>
            <span>预监(PVM)</span>
            <div class="play-pvm">
              <div id="pvm"></div>
              <span>点击开启直播后选择布局可进行实时预监</span>
            </div>
          </div>
          <div class="pgm-box">
            <i class="pvm-icon"></i>
            <span>主监(PGM)</span>
            <div class="play-pgm">
              <div id="pgm"></div>
              <span>点击同步到PGM开启直播推流</span>
            </div>
          </div>
        </div>
        <div class="operation">
          <div class="item" @click="mediaConfig">
            <i class="el-icon-video-camera"></i>
            <span>媒体</span>
          </div>
          <div class="item" @click="layoutConfig">
            <i class="el-icon-menu"></i>
            <span>布局</span>
          </div>
          <div class="item" @click="imgConfig">
            <i class="el-icon-picture-outline"></i>
            <span>图片</span>
          </div>
          <div class="item" @click="textConfig">
            <i class="el-icon-tickets"></i>
            <span>文本</span>
          </div>
        </div>
      </div>
    </div>
    <el-drawer :title="drawerTitle" :visible.sync="drawer" >
      <text-component
      v-if="buttonType==='textConfig'"
        :textList="textList"
        :textTotal="textTotal"
        :CasterId="casterId"
      ></text-component>
      <img-component
      v-else-if="buttonType==='imgConfig'"
      :imgList="imgList"
        :imgTotal="imgTotal"
        :CasterId="casterId"
      ></img-component>
      <media-component
      v-else-if="buttonType==='mediaConfig'"
       :mediaList="mediaList"
        :mediaTotal="mediaTotal"
         :CasterId="casterId"
         @describeCasterVideoResources="describeCasterVideoResources"
      ></media-component>
      <layout-component
      v-else-if="buttonType==='layoutConfig'"
      :layoutList="layoutList"
      :mediaList="mediaList"
      :layoutTotal="layoutTotal"
      :CasterId="casterId"
      ></layout-component>
    </el-drawer>
    <el-dialog title="直播配置" :visible.sync="dialogVisible" width="700px">
      <el-form
        label-position="right"
        label-width="105px"
        :model="directorConfig"
        ref="form"
        :rules="rules"
      >
        <el-form-item label="导播标题" required prop="CasterName">
          <el-input
            v-model="directorConfig.CasterName"
            placeholder="最大长度32位"
          ></el-input>
        </el-form-item>
        <el-form-item label="域名" required prop="DomainName">
          <el-input
            v-model="directorConfig.DomainName"
          ></el-input>
        </el-form-item>
             <el-form-item label="输出配置">
              <el-select v-model="directorConfig.TranscodeConfig.CasterTemplate">
                    <el-option
                       v-for="item in transcodeConfigOptions"
                       :key="item.value"
                       :label="item.label"
                       :value="item.value">
                    </el-option>
              </el-select>
        </el-form-item>
            <!-- <el-form-item label="转码格式">
            <el-checkbox-group v-model="directorConfig.TranscodeConfig.LiveTemplateIds.LocationId">
                <el-checkbox label="lsd">标清</el-checkbox>
                <el-checkbox label="lld">流畅</el-checkbox>
                <el-checkbox label="lud">lud</el-checkbox>
                <el-checkbox label="lhd">高清自适应转码模板</el-checkbox>
          </el-checkbox-group>
        </el-form-item> -->
             <el-divider></el-divider>
                <el-form-item label="录制配置">
             <el-switch
                v-model="isRecord"
                active-color="#13ce66"
                inactive-color="#ff4949">
              </el-switch>
        </el-form-item>
            <el-form-item v-if="isRecord" label="存储位置">
                <el-input
            v-model="directorConfig.RecordConfig.OssBucket"
          ></el-input>
        </el-form-item>
         <el-form-item v-if="isRecord" label="存储格式">
               <el-checkbox-group v-model="RecordFormatValue">
                <el-checkbox label="mp4">MP4</el-checkbox>
                <el-checkbox label="flv">FLV</el-checkbox>
                <el-checkbox label="m3u8">M3U8</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-divider></el-divider>
         <el-form-item label="延时播放">
             <el-switch
                v-model="isDelay"
                active-color="#13ce66"
                inactive-color="#ff4949">
              </el-switch>
        </el-form-item>
        <el-form-item v-if="isDelay" label="延时时长">
                <el-input-number size="samll" v-model="directorConfig.Delay" controls-position="right"></el-input-number>
        </el-form-item>
        <el-divider></el-divider>
          <el-form-item label="备播视频" >
          <el-input
            v-model="directorConfig.UrgentMaterialId"
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
import { service } from '@/api';
import { mapGetters } from 'vuex';
import { downloadPCMixin } from '@/mixins/downloadPC';
import LiveInfo from '@/components/functional/liveInfo';
import LiveData from '@/components/functional/liveData';
import Chat from '@/components/functional/chat';
import BanList from '@/components/functional/banList';
import TextComponent from '@/components/tools/textComponent'
import ImgComponent from '@/components/tools/imgComponent'
import MediaComponent from '@/components/tools/mediaComponent'
import LayoutComponent from '@/components/tools/layoutComponent'

export default {
  mixins: [downloadPCMixin],

  data: () => ({
    drawerTitle: '',
    buttonType: '',
    transcodeConfigOptions: [{
      value: 'lp_ld',
      label: '流畅'
    },
    {
      value: 'lp_sd',
      label: '标清'
    },
    {
      value: 'lp_hd',
      label: '高清'
    },
    {
      value: 'lp_ud',
      label: '超清'
    }],
    drawer: false,
    mediaList: [],
    mediaTotal: 0,
    layoutList: [],
    layoutTotal: 0,
    imgList: [],
    imgTotal: 0,
    textList: [],
    textTotal: 0,
    casterId: '',
    directorConfig: {
      CasterName: '',
      ChannelEnable: '',
      RequestId: '',
      DomainName: '',
      ProgramEffect: 0,
      UrgentMaterialId: '',
      TranscodeConfig: {
        CasterTemplate: '',
        LiveTemplateIds: {
          LocationId: []
        }
      },
      CasterId: '',
      ProgramName: '',
      Delay: 0,
      RecordConfig: {
        OssEndpoint: '',
        RecordFormat: {
          RecordFormat: []
        },
        OssBucket: ''
      },
    },
    isRecord: false,
    isDelay: false,
    RecordFormatValue: [],
    dialogVisible: false,
    pushLink: '',
    playback: false,
    liveId: '',
    liveDetail: {},
    kicked: '',
    updateLiveModel: {
      title: '',
      notice: '',
      coverUrl: '',
      anchorId: '',
      anchorNick: '',
      roomTitle: '', // document.title
      playerBackground: '', // pc直播中止时的背景
      avatar: '', // 直播间头像
      favicon: '',
      shareUrl: '', // 分享链接
    },
    rules: {
      title: [
        { required: true, message: '请输入直播标题', trigger: 'blur' },
        { max: 32, message: '最大长度32位', trigger: 'blur' },
      ],
    },
    refreshInterval: null,
    countInterval: null,
    countTime: 0,
    status: ['未开始', '直播中', '已结束'],
    addressVisiable: false,
  }),

  methods: {
    getSence() {
      service.CasterApiPopService({
        ActionName: 'DescribeCasterScenes',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          CasterId: this.casterId,
        }
      }).then(res => {
        console.log(JSON.parse(res.data.result.response), '场景');
      }).catch(err => {
        console.log(err);
      })

    },
    switchSources() {
      console.log('切换备播视频');
    },
    syncPgm() {
      console.log('同步');
    },
    directorConfigInfo() {
      this.describeCasterConfig()
      this.dialogVisible = true
    },
    describeCasterConfig() {
      service
        .CasterApiPopService({
          ActionName: 'DescribeCasterConfig',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            CasterId: this.casterId
          }
        })
        .then((res) => {
          this.directorConfig = {
            CasterId: JSON.parse(res.data.result.response).CasterId,
            CasterName: JSON.parse(res.data.result.response).CasterName,
            DomainName: JSON.parse(res.data.result.response).DomainName,
            TranscodeConfig: {
              CasterTemplate: JSON.parse(res.data.result.response).TranscodeConfig.CasterTemplate,
            },
            RecordConfig: {
              OssBucket: JSON.parse(res.data.result.response).RecordConfig.OssBucket,
              endpoint: JSON.parse(res.data.result.response).RecordConfig.OssEndpoint,
              interval: 0,
              videoFormat: JSON.parse(res.data.result.response).RecordConfig.RecordFormat.Format ? [JSON.parse(res.data.result.response).RecordConfig.RecordFormat.Format] : []
            },
            Delay: 0,
            UrgentMaterialId: ''
          }
          this.directorConfig.RecordConfig.OssBucket ? this.isRecord = true : this.isRecord = false
          this.directorConfig.Delay || this.directorConfig.Delay === 0 ? this.isDelay = true : this.isDelay = false
          if (this.isRecord) {
            this.RecordFormatValue = this.directorConfig.RecordConfig.RecordFormat.RecordFormat.map(item => {
              return item.Format
            })
          }
          console.log(res, '配置', this.directorConfig);
        }).catch(err => {
          console.log(err);
        })
    },
    getLayouts() {
      service.CasterApiPopService({
        ActionName: 'DescribeCasterLayouts',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          CasterId: this.casterId,
        }
      }).then(res => {
        this.layoutList = JSON.parse(res.data.result.response).Layouts.Layout.map(item => {
          item.isShow = false
          return item
        })
        this.layoutTotal = JSON.parse(res.data.result.response).Total
        // console.log(this.drawerData, 'this.drawerData');

      }).catch(err => {
        console.log(err);
      })
    },
    describeCasterComponents() {
      service.CasterApiPopService({
        ActionName: 'DescribeCasterComponents',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          CasterId: this.casterId,
        }
      }).then(res => {
        this.textList = JSON.parse(res.data.result.response).Components.Component.filter(item => { return item.ComponentType === 'text' })
        this.imgList = JSON.parse(res.data.result.response).Components.Component.filter(item => { return item.ComponentType === 'image' })
        console.log(res, this.imgList, '文本==组件', this.textList);
      }).catch(err => {
        console.log(err);
      })
    },
    describeCasterVideoResources() {
      service
        .CasterApiPopService({
          ActionName: 'DescribeCasterVideoResources',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            CasterId: this.casterId,
          }
        })
        .then((res) => {
          console.log(JSON.parse(res.data.result.response), '=媒体');

          this.mediaList = JSON.parse(res.data.result.response).VideoResources.VideoResource
          this.mediaTotal = JSON.parse(res.data.result.response).Total
          this.mediaList = this.mediaList.map(item => {
            return { ...item, isShow: false, isShowStop: false }
          })
          // this.drawerData = res.data.data.VideoResources.VideoResource
          // this.drawerTotal = res.data.data.Total
          // console.log(this.mediaList, '==this.drawerData');
        });
    },
    mediaConfig() {
      console.log('mmmmmmm');
      this.describeCasterVideoResources()
      this.drawerTitle = '媒体资源';
      this.buttonType = 'mediaConfig'
      this.drawer = true;
    },
    layoutConfig() {
      this.getLayouts()
      this.drawerTitle = '布局样式';
      this.buttonType = 'layoutConfig'
      this.drawer = true;
    },
    textConfig() {
      this.describeCasterComponents()
      this.drawerTitle = '文本组件';
      this.drawer = true;
      this.buttonType = 'textConfig'
    },
    imgConfig() {
      this.describeCasterComponents()
      this.drawerTitle = '图片组件';
      this.drawer = true;
      this.buttonType = 'imgConfig'
    },
    downloadMobile() {
      window.open(window.location.origin + '#/mobileDownload');
    },
    onSubmit() {
      console.log(this.directorConfig, '==directorConfig');
      service.CasterApiPopService({
        ActionName: 'SetCasterConfig',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          ...this.directorConfig,
        }
      }).then(res => {
        this.$message.success('配置成功');
        this.dialogVisible = false;
        this.describeCasterConfig()
      }).catch((err) => {
        this.$message.error('配置失败');
        console.error(err);
      })
        .finally(() => {
          this.$refs.form.resetFields();
        });
    },

    stopLive() {
      this.$confirm('确定要停止该场直播吗？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'danger',
      })
        .then(() => {
          this.$message({
            type: 'success',
            message: '已停止直播',
          });
          // this.startPlayback();
        })
        .catch((err) => {
          console.error(err);
        });
    },

    startLive() {
      this.$confirm('确定要开启该场直播吗？', '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'info',
      })
        .then(() => {
          return service.CasterApiPopService({
            ActionName: 'StartCaster',
            AppId: localStorage.getItem('appId'),
            Parameter: {
              CasterId: this.casterId,
            }
          });
        })
        .then((res) => {
          console.log(res, '====kaishi');
          // this.$message({
          //   type: 'success',
          //   message: '已开启直播，可以开始推流了',
          // });
        })
        .catch((err) => {
          console.error(err);
        });

    },

  },

  computed: {
    ...mapGetters(['userId', 'userNick']),
  },
  components: {
    LiveInfo,
    LiveData,
    Chat,
    BanList,
    TextComponent,
    ImgComponent,
    MediaComponent,
    LayoutComponent

  },

  created() {
    this.casterId = this.$route.params.CasterId;
    this.describeCasterConfig();
    this.describeCasterVideoResources()
    this.describeCasterComponents()
    this.getLayouts()
    this.getSence()

  },

  beforeDestroy() {
    if (this.roomChannel) {
      this.roomChannel.leaveRoom();
      this.roomEngine.logout();
    }
  },
  watch: {
    dialogVisible: function(val) {
      if (!val) {
        // this.directorConfig = {}
      }
    }
  }
};
</script>

<style scoped lang="less">
.container-common {
  // background-color: #1e2326;
  // border-radius: @border-radius;
  & > div {
    width: 100%;
  }
  display: flex;
  flex-direction: column;
  height: 800px;
}
.push {
  cursor: pointer;
  width: 100%;
  line-height: 32px;
  text-align: center;
  background-color: #409eff;
  color: #fff;
  border-radius: 5px;
  margin-top: 5px;
}
.el-container {
  .el-tabs__content {
    background-color: #000;
  }
  .el-main {
    .page {
      min-width: 1000px;
      .page-content {
        padding: 20px 0;
        .title {
          font-size: 20px;
          font-weight: 500;
          margin-bottom: 13px;
          padding: 0 20px;
        }
        .btn-container {
          padding: 0 20px;
        }
        .main-container {
          position: relative;
          width: 100%;
          align-items: center;
          height: 100%;
          padding: 10px 0;
          .play-info {
            width: 100%;
          }
          .player-box {
            width: 100%;
            display: flex;
            .pvm-box {
              width: 50%;
              padding: 20px;
              .play-pvm {
                height: 500px;
                border: 1px solid #efefef;
                span {
                  display: block;
                  text-align: center;
                  line-height: 500px;
                }
              }
              .pvm-icon {
                display: inline-block;
                width: 10px;
                height: 10px;
                border-radius: 50%;
                margin-right: 2px;
                background-color: #35b24b;
              }
            }
            .pgm-box {
              width: 50%;
              padding: 20px;
              .play-pgm {
                height: 500px;
                border: 1px solid #efefef;
                span {
                  display: block;
                  text-align: center;
                  line-height: 500px;
                }
              }
              .pvm-icon {
                display: inline-block;
                width: 10px;
                height: 10px;
                border-radius: 50%;
                margin-right: 2px;
                background-color: #f15533;
              }
            }
          }
        }
        .operation {
          height: 66px;
          border-top: 1px solid #eee;
          padding: 0 20px;
          display: flex;
          .item {
            height: 66px;
            width: 66px;
            text-align: center;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            cursor: pointer;
            &:hover {
              background-color: #f2f2f2;
            }
            i {
              font-size: 28px;
              display: block;
            }
            span {
              font-size: 13px;
            }
          }
        }
      }
    }
  }
}
.ext-item {
  display: flex;
}
.message {
  position: relative;
  height: 100%;
  padding-bottom: 110px;
  overflow: auto;
  &.notice-hoverd-message {
    padding-top: 45px;
  }
}
.dialog-title {
  .title {
    font-size: 14px;
    font-weight: 500;
    display: inline-block;
    margin-right: 30px;
  }
}
.good-list {
  max-height: 280px;
  overflow: auto;
  .list-item {
    display: flex;
    margin-bottom: 12px;
    align-items: center;
    justify-content: space-between;
    .info-container {
      display: flex;
      .img {
        width: 60px;
        height: 60px;
        background-size: cover;
        background-repeat: no-repeat;
      }
      .info {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        flex: 1;
        padding: 0 30px 0 5px;
        .name {
          font-size: 14px;
        }
        .price {
          font-size: 17px;
          font-weight: 500;
          color: #ff0000;
        }
      }
    }
    .btn {
      height: 34px;
    }
  }
}
.address-list {
  padding: 8px 4px;
  .item {
    .name {
      margin-bottom: 5px;
    }
    .content {
      display: flex;
      button {
        margin-left: 5px;
      }
    }
  }
}
</style>
