<template>
  <div class="list">
    <div
      :class="`notice-bar-container ${hoverdNotice ? 'notice-hoverd' : ''}`"
      v-if="!!notice"
      @mouseleave="hoverdNotice = false"
    >
      <div class="notice-inner">
        【公告】{{ notice }}
        <i class="el-icon-arrow-down" @mouseover="hoverdNotice = true"></i>
      </div>
    </div>
    <div class="comment-list">
      <div
        class="comment-item"
        v-for="(item, index) in commentList"
        :key="index"
      >
        <div class="name">
          {{ item.name }} <span v-show="item.userId === userId">(我)</span>
        </div>
        <div class="content">{{ item.content }}</div>
        <el-dropdown
          class="more"
          trigger="click"
          size="small"
          @command="handleCommand"
          v-show="item.userId !== userId"
        >
          <i class="el-icon-more"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item
              :command="`cancelBanComment//${item.userId}`"
              v-if="banList.includes(item.userId)"
              >解除禁言</el-dropdown-item
            >
            <el-dropdown-item :command="`banComment//${item.userId}`" v-else
              >禁止聊天</el-dropdown-item
            >
            <el-dropdown-item :command="`kickUserShort//${item.userId}`"
              >踢出房间</el-dropdown-item
            >
            <el-dropdown-item :command="`kickUserPermanently//${item.userId}`"
              >永久踢出</el-dropdown-item
            >
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
    <div class="send">
      <el-input
        type="textarea"
        placeholder="请输入内容"
        v-model="comment"
        maxlength="50"
        class="textarea"
        resize="none"
        @keydown.enter.native="sendComment"
      >
      </el-input>
      <el-button type="primary" size="small" @click="sendComment" class="btn"
        >发送</el-button
      >
      <el-checkbox v-model="isMuteAll" class="mute-checkbox" @change="muteAll"
        >全体禁言</el-checkbox
      >
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
export default {
  data: () => ({
    hoverdNotice: false,
    isMuteAll: false,
    comment: '',
  }),

  props: {
    commentList: {
      type: Array,
      require: true,
      default: () => [],
    },

    notice: {
      type: String,
      require: true,
      default: '',
    },
    parentIsMuteAll: {
      type: Boolean,
      require: true,
      default: false,
    },

    chatService: {
      type: Object,
      require: true,
      default: () => ({}),
    },

    banList: {
      type: Array,
      require: true,
      default: () => [],
    },
  },

  methods: {
    async sendComment() {
      let e = window.event || arguments[0];
      if (e.key === 'Enter' || e.code === 'Enter' || e.keyCode === 13) {
        e.returnValue = false;
      }
      if (!this.comment || !this.chatService) return;
      try {
        await this.chatService.sendComment(this.comment);
        this.comment = '';
      } catch (err) {
        this.$message.error('发送失败：' + err.body.reason);
        console.error(err);
      }
    },

    muteAll(e) {
      if (e) {
        this.chatService
          .banAllComment()
          .then(() => {
            this.$message.success('已全体禁言');
          })
          .catch((err) => {
            console.error(err);
            this.isMuteAll = false;
            this.$message.error('禁言失败');
          });
      } else {
        this.chatService
          .cancelBanAllComment()
          .then(() => {
            this.$message.success('已取消全体禁言');
          })
          .catch((err) => {
            console.error(err);
            this.isMuteAll = true;
            this.$message.error('取消全体禁言失败');
          });
      }
    },

    handleCommand(command) {
      const [commandName, userId] = command.split('//');
      if (commandName === 'banComment') {
        this.chatService.banComment(userId, 86400).then(() => {
          this.$message.success('禁言成功');
          this.$emit('banComment');
        });
      }
      if (commandName === 'cancelBanComment') {
        this.chatService.cancelBanComment(userId).then(() => {
          this.$message.success('已取消禁言');
          this.$emit('banComment');
        });
      }
      if (commandName === 'kickUserPermanently') {
        this.$emit('kickUserPermanently', userId);
      }
      if (commandName === 'kickUserShort') {
        this.$emit('kickUserShort', userId);
      }
    },
  },

  computed: {
    ...mapGetters(['userId', 'userNick']),
  },
  created() {
    this.isMuteAll = this.parentIsMuteAll;
  },
  watch: {
    parentIsMuteAll(newVal) {
      this.isMuteAll = newVal;
    },
  },
};
</script>

<style scoped lang="less">
.list {
  height: 100%;
  overflow: auto;
}
.notice-bar-container {
  height: 40px;
  padding: 5px 25px 0 10px;
  font-size: 12px;
  overflow: hidden;
  background-color: #f2f8fe;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  cursor: default;
  user-select: none;
  color: #0085ff;
  letter-spacing: 0.2px;
  opacity: 0.9;
  box-shadow: 2px 0 1px 1px #dfdfdf;
  z-index: 101;
  .notice-inner {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    line-height: 30px;
    word-break: break-all;
    background-color: #f2f8fe;
    z-index: 100;
    padding: 6px 25px 5px 10px;
    box-shadow: 2px 0 2px 1px #ccc;
    i {
      position: absolute;
      right: 7px;
      top: 14px;
      font-size: 14px;
      color: #0085ff;
      cursor: pointer;
      transform: rotate(0);
      transition: all linear 0.2s;
    }
  }
  &.notice-hoverd {
    overflow: visible;
    .notice-inner {
      i {
        transform: rotate(-180deg);
      }
    }
  }
}
.send {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100px;
  .textarea {
    height: 100%;
  }
  .btn {
    position: absolute;
    bottom: 5px;
    right: 5px;
  }
  .mute-checkbox {
    position: absolute;
    bottom: 10px;
    left: 10px;
  }
}
.comment-list {
  .comment-item {
    &:hover {
      background-color: #e0f7fa;
    }
    margin-bottom: 10px;
    position: relative;
    padding: 3px 5px;
    .name {
      font-size: 13px;
      color: #aaa;
      margin-bottom: 5px;
    }
    .more {
      position: absolute;
      right: 10px;
      top: 10px;
      cursor: pointer;
    }
  }
}
</style>
