<template>
  <div style="padding:12px">
    <el-button type="primary" size="small" @click="imgVisible = true">添加图片组件</el-button>
    <!-- 图片展示 -->
    <el-row  :gutter="24" style="margin-top: 20px">
      <el-col  :span="24" v-for="item in imgList" :key="item.ComponentId">
        <div style="margin-bottom:5px">{{item.ComponentName}}</div>
        <div class="caster-text-box" @click="selectedImg(item)">
          <div class="caster-text-play" @click="imgEdit(item)"></div>
          <div  class="caster-text-delete" @click="deleteImg(item)"></div>
              <img :style="{
              position: 'absolute',
              top: (item.ComponentLayer.PositionNormalizeds.Position[0] * 100)+ '%',
              left:(item.ComponentLayer.PositionNormalizeds.Position[1] * 100) + '%',
              width:(item.ComponentLayer.WidthNormalized * 100) + '%',
              height:(item.ComponentLayer.HeightNormalized * 100) + '%',
              }" :src="item.ImageLayerContent.MaterialId"/>
        </div>
      </el-col>
    </el-row>
      <!-- 设置图片 -->
    <el-dialog append-to-body :title="imgTitle" :visible.sync="imgVisible" width="700px">
      <el-row :gutter="24" style="width:100%">
        <el-col :span="12">
          <div class="text-show">
             <img v-if="imgInfo.ImageLayerContent.MaterialId" ref="text" :src="imgInfo.ImageLayerContent.MaterialId" :style="{
               position: 'absolute',
               top: (top * 100) + '%',
               left: (left * 100) + '%',
               width:(imgInfo.ComponentLayer.WidthNormalized * 100) + '%',
               height:(imgInfo.ComponentLayer.HeightNormalized * 100) + '%',
             }"/>
          </div>
          <el-upload
              class="upload-demo"
              action="https://jsonplaceholder.typicode.com/posts/"
              :limit="1"
              :show-file-list="false"
              :on-success ="successUpload"
              :on-error ="errorUpload"
              :before-upload ="beforeUpload"
              :file-list="fileList">
            <el-button size="small" >+点击上传</el-button>
            </el-upload>
        </el-col>
        <el-col :span="12">
             <el-form
        label-position="right"
        :model="imgInfo"
        ref="form"
      >
        <el-form-item label="名称" required label-width="70px">
          <el-input size="small" v-model="imgInfo.ComponentName"></el-input>
        </el-form-item>
        <el-form-item label="缩放"  label-width="70px">
         <el-slider @change="changeScale" @input="changeScale" v-model="scale" :show-tooltip="false" :show-input-controls="false" :show-input="true" input-size="mini"></el-slider>
        </el-form-item>
        <el-form-item label="显示位置" label-width="70px">
          <div>
          <div class="button-group">
            <el-button size="mini" class="el-icon-arrow-up" @click="changePosition('up')"> 向上</el-button>
          </div>
          <div class="button-group">
            <el-button size="mini" class="el-icon-arrow-left"  @click="changePosition('left')"> 向左</el-button>
            <el-button size="mini" class="el-icon-arrow-down"  @click="changePosition('down')"> 向下</el-button>
            <el-button size="mini" class="el-icon-arrow-right"  @click="changePosition('right')"> 向右</el-button>
          </div>
          <div class="button-group">
            <el-button size="mini"  @click="changePosition('leftTop')">移至左上</el-button>
            <el-button size="mini" @click="changePosition('rightTop')">移至右上</el-button>
          </div>
            <div class="button-group">
            <el-button size="mini" @click="changePosition('leftBottom')">移至左下</el-button>
            <el-button size="mini" @click="changePosition('rightBottom')">移至右下</el-button>
          </div>
          </div>
        </el-form-item>
      </el-form>
           </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="addImg">确 定</el-button>
        <el-button @click="imgVisible = false">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { service } from '@/api';
export default {
  props: {
    imgList: {
      type: Array,
      required: false,
      default: () => [],
    },
    imgTotal: {
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
      default: ''
    }
  },
  data() {
    return {
      fileList: [],
      LocationId: '',
      imgVisible: false,
      currentPage: 1,
      imgTitle: '添加图片组件',
      imgInfo: {
        ComponentId: '',
        ComponentLayer: {
          HeightNormalized: '',
          PositionNormalizeds: {
            Position: []
          },
          PositionRefer: 'topLeft',
          WidthNormalized: ''
        },
        ComponentName: '',
        ComponentType: 'image',
        Effect: 'none',
        LocationId: '',
        ImageLayerContent: {
          MaterialId: ''
        },
        size: ''
      },
      top: 0,
      left: 0,
      scale: 30,
      pageSize: this.defaultPageSize,
    };
  },
  watch: {
    imgVisible: function(val) {
      if (!val) {
        this.top = 0
        this.left = 0
        this.scale = 30
        this.imgInfo = {
          ComponentId: '',
          ComponentLayer: {
            HeightNormalized: '',
            PositionNormalizeds: {
              Position: []
            },
            PositionRefer: 'topLeft',
            WidthNormalized: ''
          },
          ComponentName: '',
          ComponentType: 'image',
          Effect: 'none',
          LocationId: '',
          ImageLayerContent: {
            MaterialId: ''
          },
          size: ''
        }
      }
    }
  },
  methods: {
    changeScale(val) {
      this.imgInfo.ComponentLayer.HeightNormalized = this.scale / 100
      this.imgInfo.ComponentLayer.WidthNormalized = this.scale / 100
    },
    beforeUpload(file) {
      console.log(file, '==fileing');

    },
    errorUpload(err, file, fileList) {
      console.log(err, 'err', file.type, fileList);
    },
    successUpload(file, fileList) {
      console.log(file.type, this.fileList, '===successUpload', fileList.raw.type);
      const that = this
      let read = new FileReader()
      read.readAsDataURL(fileList.raw)
      read.onload = function() {
        that.imgInfo.ImageLayerContent.MaterialId = read.result
        that.imgInfo.ComponentLayer.HeightNormalized = 0.3
        that.imgInfo.ComponentLayer.WidthNormalized = 0.3
      }
    },
    addImg() {
      this.imgInfo.ComponentLayer.PositionNormalizeds.Position[0] = Math.floor(this.top * 100) / 100
      this.imgInfo.ComponentLayer.PositionNormalizeds.Position[1] = Math.floor(this.left * 100) / 100
      console.log(this.imgInfo, '图片');
      const ComponentLayer = {
        PositionNormalized: [0, 0],
        HeightNormalized: '1',
        WidthNormalized: '2',
        PositionRefer: this.imgInfo.ComponentLayer.PositionRefer
      }
      service.CasterApiPopService({
        ActionName: 'AddCasterComponent',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          CasterId: this.CasterId,
          ComponentLayer: JSON.stringify(ComponentLayer),
          ComponentType: 'image',
          LocationId: 'RC01',
          ComponentName: this.imgInfo.ComponentName,
          Effect: this.imgInfo.Effect,
          ImageLayerContent: JSON.stringify(this.imgInfo.ImageLayerContent)
        }
      }).then(res => {
        console.log(JSON.parse(res.data.result.response), '场景');
        // this.textList = JSON.parse(res.data.result.response).Components.Component.filter(item => { return item.ComponentType === 'text' })
        // this.imgList = JSON.parse(res.data.result.response).Components.Component.filter(item => { return item.ComponentType === 'image' })
        // console.log(res, this.imgList, '文本==组件', this.textList);
      }).catch(err => {
        console.log(err);
      })

    },
    selectedImg(item) {
      console.log(item);
    },
    imgEdit(item) {
      console.log(item);
      this.imgInfo = item
      this.top = item.ComponentLayer.PositionNormalizeds.Position[0]
      this.left = item.ComponentLayer.PositionNormalizeds.Position[1]
      this.imgVisible = true
    },
    deleteImg(item) {
      console.log(item);
    },
    changePosition(type) {
      console.log(this.$refs.text.offsetHeight, 'top');
      switch (type) {
        case 'up':
          this.top = this.top - 0.1
          break;
        case 'left':
          this.left = this.left - 0.1
          break;
        case 'down':
          this.top = this.top + 0.1
          break;
        case 'right':
          this.left = this.left + 0.1
          break;
        case 'leftTop':
          this.top = 0
          this.left = 0
          break;
        case 'rightTop':
          this.top = 0
          this.left = 1 - Math.floor((this.$refs.text.offsetWidth / 300) * 100) / 100
          break;
        case 'rightBottom':
          this.left = 1 - Math.floor((this.$refs.text.offsetWidth / 300) * 100) / 100
          this.top = 1 - Math.floor((this.$refs.text.offsetHeight / 200) * 100) / 100
          break;
        case 'leftBottom':
          this.top = 1 - Math.floor((this.$refs.text.offsetHeight / 200) * 100) / 100
          this.left = 0
          break;
        default:
          break;
      }
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
/deep/ .el-card__body {
  padding: 0;
  height: 320px;
  position: relative;
}
.caster-text-box{
  margin-bottom: 20px; padding: 0;
  height: 150px;
  position: relative;
  border:1px solid #a8a8a8;
  &:hover{
    background-color: rgb(214, 214, 214);
    cursor: pointer;
  }
 .caster-text-play{
              position:absolute;
              top:0;
              right: 35px;
              width: 30px;
              height: 30px;
              background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
              background-size: 315px;
              background-position: 36px -7px;
              z-index: 9999;
}
  .caster-text-delete{
              position:absolute;
              top:0;
              right: 0;
              width: 30px;
              height: 30px;
              background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
              background-size: 315px;
              background-position: -6px -7px;
              z-index: 9999;
            }
}
.text-show{
  width: 300px;
  height: 200px;
  background-color: #e2dddd59;
  position: relative;
  margin-bottom: 10px;
}
/deep/ .el-input-group__prepend{
    text-align: center;
    background-color: #fff;
}
.button-group{
    width: 100%;
    text-align: center;
    height: 32px;
    .text{
        transform: scale(0.55, 0.55)
    }
}
/deep/ .el-slider__runway.show-input{
  width: 60%;
}
/deep/ .el-slider__input{
  width: 66px;
}
.el-slider{
  position: relative;
  &::after{
    content: '%';
    display: inline-block;
    position: absolute;
    top: -1px;
    right: 9px;
  }
}
</style>
