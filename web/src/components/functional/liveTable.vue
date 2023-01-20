<template>
  <div>
    <el-tabs v-model="activeName" @tab-click="tabHandleClick">
      <el-tab-pane
        v-for="tab in tabInfo"
        :key="tab.value"
        :label="tab.text"
        :name="tab.name"
        v-loading="loading"
      >
        <v-card
          :ref="'vCard' + tab.name"
          v-if="switchView === 1"
          :data="tableData"
          @buttonClick="handleTableButtonClick"
          :defaultPageSize="10"
          isBackPage
          :totalCount="totalCount"
          @handleCurrentChange="handleCurrentChange"
          @handleSizeChange="handleSizeChange"
        ></v-card>
        <v-sudoku
          v-else-if="switchView === 2"
          :ref="'vSudoku' + tab.name"
          :data="tableData"
          @buttonClick="handleTableButtonClick"
          :defaultPageSize="12"
          isBackPage
          :totalCount="totalCount"
          @handleCurrentChange="handleCurrentChange"
          @handleSizeChange="handleSizeChange"
        ></v-sudoku>
        <v-table
          v-else
          :data="tableData"
          :columns="columns"
          :loading="loading"
          isBackPage
          @handleCurrentChange="handleCurrentChange"
          @handleSizeChange="handleSizeChange"
          @buttonClick="handleTableButtonClick"
          :defaultPageSize="10"
          :totalCount="totalCount"
          ref="table"
        ></v-table>
      </el-tab-pane>
    </el-tabs>
    <el-dialog title="导播台配置" :visible.sync="dialogVisible" width="700px">
      <el-form
        label-position="right"
        label-width="105px"
        :model="directorConfig"
        ref="form"
        :rules="rules"
      >
        <el-form-item label="导播台名称" required prop="CasterName">
          <el-input
            v-model="directorConfig.CasterName"
            placeholder="最大长度32位"
          ></el-input>
        </el-form-item>
        <el-form-item label="主播流域名" prop="DomainName">
          <el-input v-model="directorConfig.DomainName"></el-input>
        </el-form-item>
        <el-form-item label="转码格式">
          <el-select v-model="directorConfig.TranscodeConfig.CasterTemplate">
            <el-option
              v-for="item in transcodeConfigOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-divider></el-divider>
        <el-form-item label="录制配置">
          <el-switch
            v-model="isRecord"
            active-color="#13ce66"
            inactive-color="#ff4949"
          >
          </el-switch>
        </el-form-item>
        <el-form-item v-if="isRecord" label="ossBucket">
          <el-input v-model="directorConfig.RecordConfig.OssBucket"></el-input>
        </el-form-item>
        <el-form-item v-if="isRecord" label="endpoint">
          <el-input v-model="directorConfig.RecordConfig.endpoint"></el-input>
        </el-form-item>
        <el-form-item label="时间间隔">
          <el-input-number
            size="samll"
            v-model="directorConfig.RecordConfig.interval"
            controls-position="right"
          ></el-input-number
          >毫秒
        </el-form-item>
        <el-form-item v-if="isRecord" label="存储格式">
          <el-checkbox-group v-model="directorConfig.RecordConfig.videoFormat">
            <el-checkbox label="mp4">MP4</el-checkbox>
            <el-checkbox label="flv">FLV</el-checkbox>
            <el-checkbox label="m3u8">M3U8</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="延时时长">
          <el-input-number
            size="samll"
            v-model="directorConfig.Delay"
            controls-position="right"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="备播视频">
          <el-input v-model="directorConfig.UrgentMaterialId"></el-input>
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
import vTable from '@/components/tools/vtable';
import vCard from '@/components/tools/vcard.vue';
import vSudoku from '@/components/tools/vsudoku.vue';
import { mapGetters } from 'vuex';
import { liveApi, service } from '@/api';

const chargeTypeFilters = [
  {
    text: '后付费',
    value: 'PostPaid',
  },
  {
    text: '预付费',
    value: 'PrePaid',
  },
];
export default {
  props: {
    from: {
      type: String,
      require: true,
      default: 'live',
    },
    switchView: {
      type: Number,
      require: true,
      default: 1,
    },
  },

  data: () => ({
    transcodeConfigOptions: [
      {
        value: 'lp_ld',
        label: '流畅',
      },
      {
        value: 'lp_sd',
        label: '标清',
      },
      {
        value: 'lp_hd',
        label: '高清',
      },
      {
        value: 'lp_ud',
        label: '超清',
      },
    ],
    directorConfig: {
      CasterId: '',
      CasterName: '',
      DomainName: '',
      TranscodeConfig: {
        CasterTemplate: '',
      },
      RecordConfig: {
        OssBucket: '',
        endpoint: '',
        interval: 0,
        videoFormat: [],
      },
      Delay: 0,
      UrgentMaterialId: '',
    },
    rules: {
      CasterName: [
        { required: true, message: '请输入直播标题', trigger: 'blur' },
        { max: 32, message: '最大长度32位', trigger: 'blur' },
      ],
    },
    dialogVisible: false,
    tabInfo: [
      {
        text: '未开始',
        name: 'startLive',
        value: 0,
      },
      {
        text: '直播中',
        name: 'living',
        value: 1,
      },
      {
        text: '已结束',
        name: 'stopLive',
        value: 2,
      },
      {
        text: '暂停中',
        name: 'stopLive',
        value: 3,
      },
    ],
    columns: [
      {
        label: '导播台名称',
        key: 'CasterName',
      },
      {
        label: '付费方式',
        key: 'ChargeType',
      },
      {
        label: '操作',
        type: 'action',
        width: '250',
        selectButton: true,
        buttonInfos: [
          {
            name: 'detailDirector',
            label: '进入导播台',
            color: 'primary',
          },
          {
            name: 'view',
            label: '导播台配置',
            color: 'success',
          },
        ],
      },
    ],
    isRecord: false,
    activeName: 'living',
    tableData: [],
    loading: false,
    pageNumber: 1,
    pageSize: 10,
    status: 1,
    totalCount: 0,
  }),

  methods: {
    onSubmit() {
      console.log(this.directorConfig, '=form');
      service
        .CasterApiPopService({
          ActionName: 'SetCasterConfig',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            ...this.directorConfig,
          },
        })
        .then((res) => {
          this.$message.success('配置成功');
          this.dialogVisible = false;
        })
        .catch((err) => {
          this.$message.error('配置失败');
          console.error(err);
        })
        .finally(() => {
          this.$refs.form.resetFields();
        });
    },
    tabHandleClick() {
      switch (this.activeName) {
        case 'startLive':
          this.status = 0;
          break;
        case 'living':
          this.status = 1;
          break;
        case 'stopLive':
          this.status = 2;
          break;
        default:
          break;
      }
      this.pageSize = this.switchView === 1 ? 10 : 12;
      this.pageNumber
        = this.switchView === 1
          ? this.$refs[`vCard${this.activeName}`][0].currentPage
          : this.$refs[`vSudoku${this.activeName}`][0].currentPage;
      this.listLiveRooms();
    },
    async listLiveRooms() {
      this.loading = true;
      try {
        if (this.switchView) {
          const res = await liveApi.listLiveRooms({
            status: this.status,
            pageNum: this.pageNumber,
            pageSize: this.pageSize,
          });
          this.totalCount = res.data.totalCount;
          this.tableData = res.data.liveList;
        } else {
          const res = await service.CasterApiPopService({
            ActionName: 'DescribeCasters',
            AppId: localStorage.getItem('appId'),
            Parameter: {
              Status: 0,
              PageNum: 1,
              PageSize: 20,
            },
          });
          console.log(res, JSON.parse(res.data.result.response));
          this.totalCount = JSON.parse(res.data.result.response).Total;
          this.tableData = JSON.parse(
            res.data.result.response
          ).CasterList.Caster;
        }
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    },

    jumpToStandard(liveId) {
      liveApi
        .getStandardRoomJumpUrl({
          bizId: liveId,
          userId: this.userId,
          userNick: this.userNick,
        })
        .then((res) => {
          window.open(res.data.result.standardRoomJumpUrl);
        });
    },

    handleCurrentChange(e) {
      this.pageNumber = e;
      this.listLiveRooms();
    },

    handleSizeChange(e) {
      this.pageSize = e;
      this.listLiveRooms();
    },
    handleTableButtonClick(e) {
      const { button, data } = e;
      if (button === 'detail') {
        this.$router.push({
          path: `/${this.from}/liveDetail/${data.id}`,
        });
      }
      if (button === 'detailDirector') {
        this.$router.push({
          path: `/${this.from}/directorDetail/${data.CasterId}`,
        });
      }
      if (button === 'view') {
        service
          .CasterApiPopService({
            ActionName: 'DescribeCasterConfig',
            AppId: localStorage.getItem('appId'),
            Parameter: {
              CasterId: data.CasterId,
            },
          })
          .then((res) => {
            console.log('配置', JSON.parse(res.data.result.response));
            // this.directorConfig = JSON.parse(res.data.result.response)
            this.directorConfig = {
              CasterId: JSON.parse(res.data.result.response).CasterId,
              CasterName: JSON.parse(res.data.result.response).CasterName,
              DomainName: JSON.parse(res.data.result.response).DomainName,
              TranscodeConfig: {
                CasterTemplate: JSON.parse(res.data.result.response)
                  .TranscodeConfig.CasterTemplate,
              },
              RecordConfig: {
                OssBucket: JSON.parse(res.data.result.response).RecordConfig
                  .OssBucket,
                endpoint: JSON.parse(res.data.result.response).RecordConfig
                  .OssEndpoint,
                interval: 0,
                videoFormat: JSON.parse(res.data.result.response).RecordConfig
                  .RecordFormat.Format
                  ? [
                    JSON.parse(res.data.result.response).RecordConfig
                      .RecordFormat.Format,
                  ]
                  : [],
              },
              Delay: 0,
              UrgentMaterialId: '',
            };
            this.dialogVisible = true;
          });
      }
    },
  },

  created() {
    this.listLiveRooms();
  },

  computed: {
    ...mapGetters(['userId', 'userNick']),
  },
  watch: {
    switchView(newVal) {
      newVal === 1 ? (this.pageSize = 10) : (this.pageSize = 12);
      this.listLiveRooms();
    },
  },
  components: {
    vTable,
    vCard,
    vSudoku,
  },
};
</script>

<style scoped lang="less"></style>
