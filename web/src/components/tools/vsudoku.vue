<template>
  <div>
    <el-row :gutter="24" style="margin-top: 20px">
      <el-col :span="6" v-for="item in data" :key="item.liveId">
        <el-card style="margin-bottom: 20px">
          <div class="coverUrl">
            <img
              :src="item.coverUrl"
              class="image"
            />
          </div>
          <div style="padding: 7px 16px; height: 50%; position: relative">
            <div class="title">{{ item.title }}</div>
            <div class="info">
              <span>在线人数：{{ item.metrics.onlineCount }}</span>
              <span>UV：{{ item.uv }}</span>
              <span>PV：{{ item.pv }}</span>
            </div>
            <div class="nick">
              <span
                ><i class="el-icon-s-custom"></i
                >{{ ` ${item.anchorNick}` }}</span
              >
            </div>
            <div class="bottom">
              <time class="time"
                >创建时间：{{ formatDate(item.createTime) }}</time
              >
              <el-button
                type="primary"
                @click="handleButtonClick('detail', item)"
                size="small"
                class="button"
                >管理直播间</el-button
              >
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <div class="pagination">
      <el-pagination
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        :current-page="currentPage"
        layout="total, prev, pager, next, sizes"
        :page-sizes="[12, 24, 36]"
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
.pagination {
  padding: 20px 0 0;
  text-align: right;
}
.el-card {
  border-radius: 5%;
  &:hover {
    transform: translate(0px, -9px);
    box-shadow: 10px 12px 20px #ccc;
    transition: all 0.2s ease-in-out;
  }
}
/deep/ .el-card__body {
  padding: 0;
  height: 250px;
}
.time {
  font-size: 13px;
  position: absolute;
  bottom: 12px;
}
.title {
  font-weight: 700;
  font-size: 24px;
}
.coverUrl {
  height: 50%;
  img {
    width: 100%;
    height: 100%;
  }
}
.nick {
  margin-top: 10px;
  color: #999;
}
.info {
  color: #999;
  span {
    margin-right: 5px;
  }
}
.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 5px;
  position: absolute;
  right: 16px;
  bottom: 8px;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}
</style>
