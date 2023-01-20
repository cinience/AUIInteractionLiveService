<template>
  <div style="padding: 12px">
    <el-button type="primary" size="small" @click="videoResourceVisible = true"
      >添加媒体资源</el-button
    >
    <!--媒体展示  -->
    <el-row :gutter="24" style="margin-top: 20px">
      <el-col
        :span="24"
        v-for="(item, index) in mediaList"
        :key="item.ResourceId"
      >
        <el-card style="margin-bottom: 20px">
          <div
            class="caster-video-item"
            @mousemove="item.isShow = true"
            @mouseleave="item.isShow = false"
          >
            <div class="caster-video-body">
              <div :id="item.LocationId" class="caster-video-player"></div>
              <div
                v-if="item.isShow && !item.isShowStop"
                class="caster-video-play"
                @click="playVideo(item, index)"
              ></div>
              <div
                v-if="item.isShow && item.isShowStop"
                class="caster-video-stopPlay"
                @click="stopPlayVideo(item)"
              ></div>
              <div
                v-if="item.isShow"
                class="caster-video-delete"
                @click="deleteVide(item)"
              ></div>
            </div>
            <div class="caster-video-item-info">
              <div class="item-info">
                <span>{{ index + 1 }}</span>
                <span>{{ item.ResourceName }}</span>
              </div>
              <div class="item-config">
                <el-tooltip
                  content="编辑资源"
                  placement="bottom"
                  effect="light"
                >
                  <span
                    class="item-config-edit"
                    @click="updatevideoResource(item)"
                  >
                    <img
                      src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIiIGhlaWdodD0iMTIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPjxkZWZzPjxwYXRoIGQ9Ik03LjIzNSAwYy4xNTQgMCAuMjg0LjEwOC4zMDIuMjUybC4yMzUgMS41OWMuMzc2LjE1LjcyMS4zNDggMS4wNDIuNTg4bDEuNTM2LS42YS4yOTUuMjk1IDAgMDEuMTA1LS4wMTguMzEuMzEgMCAwMS4yNzEuMTVsMS4yMzQgMi4wNzZhLjI5Ny4yOTcgMCAwMS0uMDc0LjM4NGwtMS4zMDIuOTljLjAyNS4xOTIuMDQ0LjM4NC4wNDQuNTg4IDAgLjIwNC0uMDE5LjM5Ni0uMDQ0LjU4OGwxLjMwMi45OWMuMTE3LjA5LjE0OC4yNTIuMDc0LjM4NGwtMS4yMzQgMi4wNzZhLjMwNy4zMDcgMCAwMS0uMjY1LjE1LjM1OC4zNTggMCAwMS0uMTExLS4wMThsLTEuNTM2LS42Yy0uMzIuMjM0LS42NjYuNDM4LTEuMDQyLjU4OGwtLjIzNSAxLjU5YS4yOTguMjk4IDAgMDEtLjMwMi4yNTJINC43NjhhLjI5OC4yOTggMCAwMS0uMzAzLS4yNTJsLS4yMzQtMS41OUE0LjUzOSA0LjUzOSAwIDAxMy4xOSA5LjU3bC0xLjUzNi42YS4yOTUuMjk1IDAgMDEtLjEwNS4wMTguMzEuMzEgMCAwMS0uMjcyLS4xNUwuMDQzIDcuOTYyYS4yOTcuMjk3IDAgMDEuMDc0LS4zODRsMS4zMDEtLjk5QTQuNjMgNC42MyAwIDAxMS4zNzUgNmMwLS4xOTguMDE5LS4zOTYuMDQzLS41ODhsLTEuMzAxLS45OWEuMjkuMjkgMCAwMS0uMDc0LS4zODRsMS4yMzMtMi4wNzZhLjMwNy4zMDcgMCAwMS4yNjYtLjE1Yy4wMzcgMCAuMDc0LjAwNi4xMS4wMThsMS41MzcuNmMuMzItLjIzNC42NjYtLjQzOCAxLjA0Mi0uNTg4bC4yMzQtMS41OUEuMjk4LjI5OCAwIDAxNC43NjggMHpNNiAzLjc1YTIuMjUgMi4yNSAwIDEwLS4wMDEgNC40OTlBMi4yNSAyLjI1IDAgMDA2IDMuNzV6IiBpZD0iYSIvPjwvZGVmcz48dXNlIGZpbGw9IiM4ODgiIHhsaW5rOmhyZWY9IiNhIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiLz48L3N2Zz4"
                    />
                  </span>
                </el-tooltip>
                <el-tooltip
                  content="点击试听音频"
                  placement="bottom"
                  effect="light"
                >
                  <span class="item-config-volumn">
                    <img
                      src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTIiIGhlaWdodD0iMTIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZD0iTTEyIDYuMDA3QTUuOTk3IDUuOTk3IDAgMDA2LjAxMyAwIDUuOTk0IDUuOTk0IDAgMDAuMzYxIDguMDM5IDIuMjI3IDIuMjI3IDAgMDAwIDkuMjUzIDIuNzQ5IDIuNzQ5IDAgMDAyLjc1IDEyYS43Ni43NiAwIDAwLjc1LS43NVY4LjAwNWExIDEgMCAwMC0xLS45OThoLS4yNWEyLjE4IDIuMTggMCAwMC0uNjEuMDk1IDQuNDk3IDQuNDk3IDAgMDEzLjI0Ni01LjQ3IDQuNSA0LjUgMCAwMTUuNDc0IDMuMjQzIDQuNDkgNC40OSAwIDAxMCAyLjIyNiAyLjI5MSAyLjI5MSAwIDAwLS42MS0uMDk2SDkuNWExIDEgMCAwMC0xIC45OTl2My4yNDZhLjc2Ljc2IDAgMDAuNzUuNzVDMTAuNzcgMTIgMTIgMTAuNzcgMTIgOS4yNTNhMi4yMiAyLjIyIDAgMDAtLjM2LTEuMjE0Yy4yMzctLjY1Mi4zNTktMS4zNC4zNi0yLjAzMnoiIGZpbGw9IiM4ODgiLz48L3N2Zz4="
                    />
                  </span>
                </el-tooltip>
                <el-popover
                  placement="top-start"
                  title=""
                  width="200"
                  trigger="click"
                >
                  <div class="caster-volumn">
                    <input type="range" max="2" min="0" step="0.1" value="0" />
                  </div>
                  <span class="item-change-volumn" slot="reference"> </span>
                </el-popover>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- 媒体 -->
    <el-dialog
      append-to-body
      :title="videoResourcTitle"
      :visible.sync="videoResourceVisible"
      width="700px"
    >
      <el-form
        label-position="right"
        label-width="105px"
        :model="videoResource"
        ref="form"
      >
        <el-form-item label="添加方式" required>
          <el-radio-group v-model="videoType">
            <!-- <el-radio :label="1">从直播台选取</el-radio>
              <el-radio :label="2">从媒资库选取</el-radio> -->
            <el-radio :label="3">输入直播URL</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" required>
          <el-input v-model="videoResource.ResourceName"></el-input>
        </el-form-item>
        <el-form-item label="直播URL">
          <el-input v-model="videoResource.LiveStreamUrl"></el-input>
        </el-form-item>
        <el-form-item v-if="videoType === 1" label="域名">
          <el-input v-model="videoResource.BeginOffset"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="addCasterVideoResourceVisible"
          >确 定</el-button
        >
        <el-button @click="videoResourceVisible = false">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
// import Text from '@/tools/text'
import { service } from '@/api';
export default {
  props: {
    mediaList: {
      type: Array,
      required: false,
      default: () => [],
    },
    mediaTotal: {
      type: Number,
      required: false,
      default: 0,
    },
    defaultPageSize: {
      type: Number,
      required: false,
      default: 10,
    },
    CasterId: {
      type: String,
      required: true,
      default: '',
    },
  },
  data() {
    return {
      LocationId: '',
      videoType: 3,
      isShow: false,
      isShowStop: false,
      currentPage: 1,
      videoResourcTitle: '添加媒体资源',
      videoResource: {
        LiveStreamUrl: '',
        MaterialId: '',
        VodUrl: '',
        BeginOffset: '',
        EndOffset: '',
        RepeatNum: '',
        ResourceName: '',
        FixedDelayDuration: '',
        PtsCallbackInterval: '',
      },
      videoResourceVisible: false,
      pageSize: this.defaultPageSize,
    };
  },
  components: {
    // Text
  },
  watch: {
    videoResource: function (val) {
      if (!val) {
        this.videoResource = {
          LiveStreamUrl: '',
          MaterialId: '',
          VodUrl: '',
          BeginOffset: '',
          EndOffset: '',
          RepeatNum: '',
          ResourceName: '',
          FixedDelayDuration: '',
          PtsCallbackInterval: '',
        };
      }
    },
  },
  methods: {
    changeLocationId() {
      this.layoutItem.MixList = this.layoutItem.BlendList;
    },
    mouseDownCanvas(e) {
      console.log(e, 'x:', e.offsetX, 'y:', e.offsetY);
    },
    deleteVide(item) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          service
            .CasterApiPopService({
              ActionName: 'DeleteCasterVideoResource',
              AppId: localStorage.getItem('appId'),
              Parameter: {
                CasterId: this.CasterId,
                ResourceId: item.ResourceId,
              },
            })
            .then((res) => {
              this.$message.success('删除成功');
              this.$emit('describeCasterVideoResources');
            })
            .catch((err) => {
              this.$message.err('删除失败');
              console.log(err);
            });
          this.$emit('describeCasterVideoResources');
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除',
          });
        });
    },
    playVideo(item, index) {
      console.log(item);
      item.isShowStop = true;
      // eslint-disable-next-line no-undef
      item.player = new Aliplayer(
        {
          id: item.LocationId,
          source: item.LiveStreamUrl,
        },
        function () {
          console.log('The player is created.');
        }
      );
    },
    stopPlayVideo(item) {
      console.log(item, 'zanz');
      item.isShowStop = false;
      item.player.pause();
    },
    updatevideoResource(item) {
      this.videoResourcTitle = '编辑';
      this.videoResource.ResourceName = item.ResourceName;
      this.videoResourceVisible = true;
    },
    addCasterVideoResourceVisible() {
      this.videoResource.LocationId
        = this.mediaList.length >= 9
          ? 'RV' + (this.mediaList.length + 1)
          : 'RV0' + (this.mediaList.length + 1);
      this.videoResource.CasterId = this.CasterId;
      service
        .CasterApiPopService({
          ActionName: 'AddCasterVideoResource',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            CasterId: this.videoResource.CasterId,
            ResourceName: this.videoResource.ResourceName,
            LiveStreamUrl: this.videoResource.LiveStreamUrl,
            LocationId: this.videoResource.LocationId,
          },
        })
        .then((res) => {
          this.$message.success('添加成功');
          this.$emit('describeCasterVideoResources');
          this.videoResourceVisible = false;
        })
        .catch((err) => {
          this.$message.err('添加失败');
          console.log(err);
        });
    },
    handleButtonClick(button, row) {
      this.$emit('buttonClick', {
        button: button,
        data: Object.assign({}, row),
      });
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;
      this.$emit('handleCurrentChange', currentPage);
    },
    handleSizeChange(pageSize) {
      this.pageSize = pageSize;
      this.$emit('handleSizeChange', pageSize);
    },
  },
};
</script>
<style lang="less" scoped>
.pagination {
  padding: 20px 0 0;
  text-align: right;
}
.el-card {
}
/deep/ .prism-big-play-btn pause {
  display: none !important;
}
.caster-video-item {
  box-sizing: content-box !important;
  width: 100%;
  height: 100%;
  display: inline-block;
  padding-right: 24px;
  position: relative;
  .caster-video-body {
    width: 100%;
    background-color: #4d4d4d;
    text-align: center;
    cursor: pointer;
    height: 300px;
    position: relative;
    .caster-video-player {
      width: 100%;
      height: 300px;
    }
    .caster-video-play {
      position: absolute;
      top: 0;
      left: 0;
      width: 30px;
      height: 30px;
      background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
      background-size: 315px;
      background-position: -37px -7px;
      z-index: 999;
    }
    .caster-video-stopPlay {
      position: absolute;
      top: 0;
      left: 0;
      width: 30px;
      height: 30px;
      background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
      background-size: 315px;
      background-position: 248px -7px;
      z-index: 999;
    }
    .caster-video-delete {
      z-index: 999;
      position: absolute;
      top: 0;
      right: 0;
      width: 30px;
      height: 30px;
      background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
      background-size: 315px;
      background-position: -6px -7px;
    }
  }
  .caster-video-item-info {
    height: 20px;
    line-height: 20px;
    color: #c3c3c6;
    padding: 0 3px 0 10px;

    .item-info {
      float: left;
      span {
        display: inline-block;
        margin-right: 15px;
      }
    }
    .item-config {
      float: right;
      .item-config-edit {
        display: inline-block;
        margin-right: 5px;
        cursor: pointer;
      }
      .item-config-volumn {
        display: inline-block;
        cursor: pointer;
        margin-right: 5px;
      }
      .item-change-volumn {
        display: inline-block;
        background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
        background-size: 161px;
        width: 18px;
        height: 18px;
        line-height: 20px;
        vertical-align: bottom;
        cursor: pointer;
        background-position: -85px -17px;
      }
    }
  }
}
.caster-volumn {
  // position: absolute;
  right: 0;
  bottom: 4px;
  height: 20px;
  width: 20px;
  line-height: 20px;
  input {
    margin-left: 9px;
    cursor: inherit;
    width: 170px;
    display: inline-block;
  }
}
/deep/ .el-card__body {
  padding: 0;
  height: 320px;
  position: relative;
}
.relevanceVideo {
  float: right;
}
</style>
