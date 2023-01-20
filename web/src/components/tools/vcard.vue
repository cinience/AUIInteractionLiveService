<template>
  <div class="v-card">
    <el-card v-for="item in data" :key="item.id">
      <div class="box-card">
        <div class="coverUrl">
          <img alt="" :src="item.coverUrl" />
        </div>
        <div class="liveRoomInfo">
          <div class="title">{{ item.title }}</div>
          <div class="data">
            <div>
              <span>在线人数：{{ item.metrics.onlineCount }}</span>
              <span>UV：{{ item.metrics.uv }}</span>
              <span>PV：{{ item.metrics.pv }}</span>
            </div>
            <div>
            <span>直播ID：{{ item.id }}</span>
            </div>
            <div>
              <span class="nick"
                ><i class="el-icon-s-custom"></i
                >主播ID：{{ `${item.anchorId}` }}</span
              >
              <span>主播昵称：{{ item.anchorNick }}</span>
            </div>
          </div>
          <div class="info">
            <div class="left">
              <div class="craeteTime">
                创建时间：{{ formatDate(item.createdAt) }}
              </div>
            </div>
            <div class="right">
              <el-button
                type="primary"
                @click="handleButtonClick('detail', item)"
                size="small"
                >管理直播</el-button
              >
            </div>
          </div>
        </div>
      </div>
    </el-card>
    <div class="pagination">
      <el-pagination
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        :current-page="currentPage"
        layout="total, prev, pager, next, sizes"
        :page-sizes="[10, 20, 50]"
        :page-size="pageSize"
        :total="totalCount"
      >
      </el-pagination>
    </div>
  </div>
</template>
<script>
import { formatDate } from '@/libs/utils';
export default {
  props: {
    data: {
      type: Array,
      required: true,
      default: () => [],
    },
    totalCount: {
      type: Number,
      required: false,
      default: 0,
    },
    defaultPageSize: {
      type: Number,
      required: false,
      default: 10,
    },
  },
  data() {
    return {
      currentPage: 1,
      pageSize: this.defaultPageSize,
    };
  },
  methods: {
    formatDate(date) {
      return date ? formatDate(date, 'MM-dd hh:mm') : 0;
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
.v-card {
  .el-card {
    margin-bottom: 20px;
  }
  .el-card__body {
    height: 240px;
  }
}
.box-card {
  width: 100%;
  display: flex;
}
.coverUrl {
  width: 16%;
  height: 150px;
  box-shadow: 2px 2px 10px 0px #ccc;
  padding: 0 !important;
  img {
    width: 100%;
    height: 100%;
  }
}
.liveRoomInfo {
  margin-left: 24px;
  display: flex;
  width: 84%;
  flex-direction: column;
  justify-content: space-between;
  .title {
    font-weight: 700;
    font-size: 24px;
  }
  .data {
    color: #999;
    div {
      span {
        display: inline-block;
        margin-right: 10px;
      }
    }
  }
  .info {
    display: flex;
    justify-content: space-between;
    .left {
      display: flex;
      .status {
        margin-left: 0;
        margin-right: 10px;
        &::before {
          content: "";
          display: inline-block;
          height: 8px;
          width: 8px;
          border-radius: 4px;
          background-color: #333;
          margin-right: 2px;
        }
      }
    }
    .right {
      position: relative;
      /deep/ .el-button--primary {
        position: absolute;
        bottom: 0;
        right: 0;
      }
    }
  }
}
.pagination {
  padding: 20px 0 0;
  text-align: right;
}
</style>
