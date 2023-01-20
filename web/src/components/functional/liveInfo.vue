<template>
  <info-container title="直播信息">
    <div class="live-info-container">
      <div class="info-column">
        <div class="info-item">
          <div class="key">主播ID</div>
          <div class="value">{{liveDetail.anchorId}}</div>
        </div>
        <div class="info-item">
          <div class="key">主播昵称</div>
          <div class="value">{{liveDetail.anchorNick}}</div>
        </div>
        <div class="info-item">
          <div class="key">直播间ID</div>
          <div class="value">{{liveDetail.id}}</div>
        </div>
      </div>
      <div class="info-column">
        <div class="info-item">
          <div class="key">创建时间</div>
          <div class="value">{{formatDate(liveDetail.createdAt)}}</div>
        </div>
        <div class="info-item">
          <div class="key">结束时间</div>
          <div class="value">{{liveDetail.status != 2 ? formatDate(liveDetail.stoppedAt) : '未结束'}}</div>
        </div>
        <div class="info-item">
          <div class="key">推流地址</div>
          <div class="value">{{liveDetail.pushUrlInfo.rtmpUrl}}</div>
        </div>
        <div class="info-item">
          <div class="key">回放地址</div>
          <div class="value">{{liveDetail.playbackUrl}}</div>
        </div>
      </div>
      <div class="info-column">
        <div class="info-item">
          <div class="key">hls播流地址</div>
          <div class="value">{{liveDetail.pullUrlInfo.hlsUrl}}</div>
        </div>
        <div class="info-item">
          <div class="key">flv播流地址</div>
          <div class="value">{{liveDetail.pullUrlInfo.flvUrl}}</div>
        </div>
        <div class="info-item">
          <div class="key">低延迟H5Url</div>
          <div class="value">{{liveDetail.artcInfo ? liveDetail.artcInfo.artcH5Url : '未开启低延迟'}}</div>
        </div>
        <div class="info-item">
          <div class="key">低延迟Url</div>
          <div class="value">{{liveDetail.pullUrlInfo.rtsUrl ? liveDetail.pullUrlInfo.rtsUrl : '未开启低延迟'}}</div>
        </div>
        <div class="info-item">
          <div class="key">其他参数</div>
          <div class="value">{{JSON.stringify(liveDetail.extends)}}</div>
        </div>
      </div>
    </div>
  </info-container>
</template>

<script>
import InfoContainer from './infoContainer'
import { formatDate } from '@/libs/utils'

export default {
  props: {
    liveDetail: {
      type: Object,
      require: true,
      default: () => ({})
    }
  },

  methods: {
    formatDate(date) {
      return date
    },
  },

  components: {
    InfoContainer
  }
}
</script>

<style scoped lang="less">
.live-info-container {
  display: flex;
  justify-content: space-between;
  font-size: 13px;
  .info-column {
    flex: 1;
    padding: 20px;
    .info-item {
      display: flex;
      justify-content: space-between;
      margin-bottom: 8px;
      .key {
        width: 80px;
        flex-shrink: 0;
        color: #242f42;
        font-weight: 500;
      }
      .value {
        flex-grow: 1;
      }
    }
  }
}
</style>
