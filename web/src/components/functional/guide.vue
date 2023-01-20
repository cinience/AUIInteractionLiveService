<template>
  <div class="guide-container">
    <div class="guide-inner">
      <div class="guide-item" v-for="item in guideList" :key="item.title">
        <i class="el-icon-check"></i>
        <div class="title">{{ item.title }}</div>
        <div class="intro">{{ item.intro }}</div>
      </div>
      <div class="line"></div>
    </div>
  </div>
</template>

<script>

export default {
  data: () => ({}),
  props: ['guideList']
}
</script>

<style scoped lang="less">
  @primary-color: rgba(79,172,255,1);
  @icon-size: 32px;
  @icon-gap: 10px;
  @arrow-gap: 20px;
  .guide-container {
    display: flex;
    justify-content: center;
    padding: 5px 0;
    position: relative;
    .guide-inner {
      max-width: 2000px;
      min-width: 1000px;
      width: 100%;
      display: flex;
      justify-content: space-between;
      padding: 0 80px;
      position: relative;
      .guide-item:first-of-type {
        i::after {
          display: none;
        }
      }
      .guide-item {
        position: relative;
        z-index: 2;
        background-color: #fff;
        .title {
          font-size: 15px;
          font-weight: bold;
          margin-bottom: 10px;
          position: relative;
          &::after {
            content: '';
            width: @icon-gap;
            height: @icon-size;
            position: absolute;
            background-color: #fff;
            left: -@icon-gap;
            top: 0;
          }
          &::before {
            content: '';
            width: @arrow-gap;
            height: @icon-size;
            position: absolute;
            background-color: #fff;
            left: -@icon-gap - @icon-size - @arrow-gap;
            top: 0;
          }
        }
        .intro {
          width: 200px;
          color: #909399;
          font-size: 13px;
        }
        i {
          font-size: 20px;
          color: @primary-color;
          border: 1px solid @primary-color;
          border-radius: 50%;
          width: @icon-size;
          height: @icon-size;
          line-height: @icon-size;
          position: absolute;
          left: -@icon-size - @icon-gap;
          top: -3px;
          text-align: center;
          background-color: #fff;
          &::after {
            content: '';
            position: absolute;
            border-width: 6px 10px;
            border-style: solid;
            border-color: transparent transparent transparent @primary-color;
            left: -@icon-gap - @arrow-gap + 2px;
            top: 9px;
            z-index: 2;
          }
        }
      }
    }
    .line {
      width: 80%;
      height: 2px;
      background-color: @primary-color;
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
      top: 12px;
    }
  }
</style>
