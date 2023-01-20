<template>
  <div style="padding:12px">
    <el-button type="primary" size="small" @click="addlayout">添加布局样式</el-button>
    <!-- 布局展示 -->
     <el-row  :gutter="24" style="margin-top: 20px">
      <el-col  :span="24" v-for="item in layoutList" :key="item.LayoutId">
        <div class="caster-layout-box" >
          <div class="caster-layout-play" @click="layoutEdit(item)"></div>
          <div  class="caster-layout-delete" @click="deleteLayout(item)"></div>
          <div @click="checkedLayer(item)"  v-for="(layer,index) in item.VideoLayers.VideoLayer" :key="index">
              <div :style="{
              position: 'absolute',
              top: (layer.PositionNormalizeds.Position[0] * 100)+ '%',
              left:(layer.PositionNormalizeds.Position[1] * 100) + '%',
              width:(layer.WidthNormalized * 100) + '%',
              height:(layer.HeightNormalized * 100) + '%',
              border: '1px solid #a8a8a8',
              textAlign: 'center',
              }">{{String.fromCharCode(64 + parseInt(index + 1))}}</div>
          </div>
        </div>
      </el-col>
    </el-row>
    <!--设置布局 -->
  <el-dialog append-to-body :title="layoutTitle" :visible.sync="layoutVisible" width="700px">
     <div class="layer-show">
     <div class="layer-box" v-for="(item,i) in layoutList" :key="item.LayoutId"  @click="selectedLayer(item,i)">
              <div :class="{selected:i===isActive}" v-for="(layer,index) in item.VideoLayers.VideoLayer" :key="index" :style="{
              position: 'absolute',
              top: (layer.PositionNormalizeds.Position[0] * 100)+ '%',
              left:(layer.PositionNormalizeds.Position[1] * 100) + '%',
              width:(layer.WidthNormalized * 100) + '%',
              height:(layer.HeightNormalized * 100) + '%',
              border: '1px solid #a8a8a8'
              }"></div>
      </div>
      <div class="layer-custom" ref="layerCustom"  @click="customlayer">
              自定义布局
      </div>
      </div>
      <el-row :gutter="24">
    <el-col :span="12" >
      <div v-show="isCustom" class="canvas-box">
        <canvas  @click="canvasClick" @mousedown="mouseDownCanvas" id="layoutConfig" width="300" height="300"></canvas>
      </div>
       <div v-show="layoutItem.VideoLayers.VideoLayer.length > 0 && !isCustom" class="123" style="width:300px;height:300px;position:relative">
     <div class="index" v-for="(layer,index) in layoutItem.VideoLayers.VideoLayer" :key="index" :style="{
              position: 'absolute',
              top: (layer.PositionNormalizeds.Position[0] * 100)+ '%',
              left:(layer.PositionNormalizeds.Position[1] * 100) + '%',
              width:(layer.WidthNormalized * 100) + '%',
              height:(layer.HeightNormalized * 100) + '%',
              border: '1px solid #a8a8a8',
              textAlign: 'center',
              }">{{layoutItem.BlendList.LocationId[index] || String.fromCharCode(64 + parseInt(index + 1))}}</div>
     </div>
    </el-col>
    <el-col :span="12" class="relevanceVideo">
      <el-form >
        <el-form-item v-if="isCustom">
          <el-button type="text" icon="el-icon-plus" @click="addLayoutItem">添加</el-button>
        </el-form-item>
        <div v-if="layoutItem.VideoLayers.VideoLayer.length > 0">
          <el-form-item  v-for="(layer,index) in layoutItem.VideoLayers.VideoLayer" :key="index" :label="String.fromCharCode(64 + parseInt(index + 1))">
          <el-select v-model="layoutItem.BlendList.LocationId[index]" @change="changeLocationId">
            <el-option v-for="videoItem in mediaList" :key="videoItem.LocationId" :label="videoItem.ResourceName" :value="videoItem.LocationId"></el-option>
            <!-- <option value="123">123</option> -->
          </el-select>
        </el-form-item>
        </div>
      </el-form>
    </el-col>
      </el-row>
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="addlayoutVisible">确 定</el-button>
      <el-button @click="layoutVisible = false">取 消</el-button>
    </span>
  </el-dialog>
  <!-- 编辑布局 -->
   <el-dialog append-to-body title="编辑布局信息" :visible.sync="updateLayoutVisible" width="700px">
      <el-row :gutter="24">
    <el-col :span="12" >
      <div  v-if="layoutItem.VideoLayers" style="width:300px;height:300px;position:relative">
     <div v-for="(layer,index) in layoutItem.VideoLayers.VideoLayer" :key="index" :style="{
              position: 'absolute',
              top: (layer.PositionNormalizeds.Position[0] * 100)+ '%',
              left:(layer.PositionNormalizeds.Position[1] * 100) + '%',
              width:(layer.WidthNormalized * 100) + '%',
              height:(layer.HeightNormalized * 100) + '%',
              border: '1px solid #a8a8a8',
              textAlign: 'center',
              }">{{layoutItem.BlendList.LocationId[index] || String.fromCharCode(64 + parseInt(index + 1))}}</div>
     </div>
    </el-col>
    <el-col :span="12" class="relevanceVideo">
      <el-form >
        <div v-if="layoutItem.VideoLayers">
          <el-form-item  v-for="(layer,index) in layoutItem.VideoLayers.VideoLayer" :key="index" :label="String.fromCharCode(64 + parseInt(index + 1))">
          <el-select v-model="layoutItem.BlendList.LocationId[index]" @change="changeLocationId">
            <el-option v-for="videoItem in mediaList" :key="videoItem.LocationId" :label="videoItem.ResourceName" :value="videoItem.LocationId"></el-option>
          </el-select>
        </el-form-item>
        </div>
      </el-form>
    </el-col>
      </el-row>
    <span slot="footer" class="dialog-footer">
      <el-button type="primary" @click="updatelayoutVisible">确 定</el-button>
      <el-button @click="updateLayoutVisible = false">取 消</el-button>
    </span>
  </el-dialog>
  </div>
</template>
<script>
import { service } from '@/api';
import { fabric } from 'fabric'
export default {
  props: {
    mediaList: {
      type: Array,
      required: false,
      default: () => [],
    },
    layoutList: {
      type: Array,
      required: false,
      default: () => [],
    },
    layoutTotal: {
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
      updateLayoutVisible: false,
      canvas: '',
      customLayout: [],
      LocationId: '',
      layoutItem: {
        AudioLayers: {
          AudioLayer: []
        },
        BlendList: {
          LocationId: []
        },
        MixList: {
          LocationId: []
        },
        VideoLayers: {
          VideoLayer: []
        }
      },
      isCustom: false,
      isActive: 0,
      layoutVisible: false,
      currentPage: 1,
      layoutTitle: '设置布局',
      pageSize: this.defaultPageSize,
    };
  },
  components: {
    // Text
  },
  watch: {
    // layoutItem: function(val) {
    //   if (!val) {
    //     this.top = 0
    //     this.left = 0
    //     this.textInfo = {
    //       ComponentId: '',
    //       ComponentLayer: {
    //         HeightNormalized: '',
    //         PositionNormalizeds: {
    //           Position: []
    //         },
    //         PositionRefer: 'topLeft',
    //         WidthNormalized: ''
    //       },
    //       ComponentName: '',
    //       ComponentType: 'text',
    //       Effect: 'none',
    //       LocationId: '',
    //       TextLayerContent: {
    //         Color: '0xff0000',
    //         FontName: 'KaiTi',
    //         SizeNormalized: '',
    //         Text: ''
    //       },
    //       size: ''
    //     }
    //   } else {
    //     this.textInfo.TextLayerContent.Color = this.textInfo.TextLayerContent.Color.replace('0x', '#')
    //     console.log(this.textInfo.TextLayerContent.Color, '=this.textInfo.TextLayerContent.Color');
    //   }
    // }
  },
  methods: {
    checkedLayer(item) {
      console.log(item, 'webpack');
    },
    deleteLayout(item) {
      console.log('item删除', item);
    },
    layoutEdit(item) {
      console.log(item, '-编辑');
      // this.updateLayout = item
      this.updateLayoutVisible = true
      this.layoutItem = item
    },
    updatelayoutVisible() {
      console.log(this.layoutItem, '编辑');
      service
        .CasterApiPopService({
          ActionName: 'AddCasterVideoResource',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            ...this.layoutItem,
            CasterId: this.CasterId,
          }
        })
        .then((res) => {
          this.$message.success('添加成功')
          // this.$emit('describeCasterVideoResources')
          // this.videoResourceVisible = false
        }).catch(err => {
          this.$message.err('添加失败')
          console.log(err);
        })

    },
    addLayoutItem() {
      this.layoutItem.AudioLayers.AudioLayer.push({ ValidChannel: 'all', VolumeRate: 1 })
      this.layoutItem.VideoLayers.VideoLayer.push({
        FillMode: 'fit',
        HeightNormalized: 0.25,
        PositionNormalizeds: {
          Position: [0, 0]
        },
        PositionRefer: 'topLeft',
        WidthNormalized: 0.5
      })
      // let canvas = new fabric.Canvas('layoutConfig')
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1] = {
        rect: '',
        text: '',
        group: '',
        width: '',
        height: '',
      }
      const that = this
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1]['rect'] = new fabric.Rect({ width: 150, height: 75, fill: '#fff', hasRotatingPoint: false });
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1]['text'] = new fabric.Text(String.fromCharCode(64 + parseInt(this.layoutItem.VideoLayers.VideoLayer.length)), {
        fontSize: 24,
      })
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1]['group'] = new fabric.Group([this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].rect, this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].text], {
        lockRotation: true,
      })
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].group.controls = {
        ...fabric.Text.prototype.controls,
        mtr: new fabric.Control({ visible: false })
      }
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].group.on('selected', function(e) { // 选中监听事件
        console.log('selected a rectangle', e.target);
      });
      this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].group.on('modified', function(e) { // 选中监听事件
        // if (e.action === 'scale') {
        //   that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].width = e.target.width * e.target.scaleX
        //   that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].height = e.target.height * e.target.scaleY
        //   console.log(this, 'width:', e.target.width * e.target.scaleX, 'height:', e.target.height * e.target.scaleY);
        // } else if (e.action === 'scaleX') {
        //   that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].width = e.target.width * e.target.scaleX
        // } else if (e.action === 'scaleY') {
        //   that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].height = e.target.height * e.target.scaleY
        // }
        that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].width = e.target.width * e.target.scaleX
        that.customLayout[that.layoutItem.VideoLayers.VideoLayer.length - 1].height = e.target.height * e.target.scaleY
        console.log('selected a modified', e);
      });
      this.canvas.add(this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].group);
      console.log(this.customLayout, this.canvas, this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].group, '====', this.customLayout[this.layoutItem.VideoLayers.VideoLayer.length - 1].text);
    },
    changeLocationId() {
      this.layoutItem.MixList = this.layoutItem.BlendList
      // console.log(this.layoutItem, '===av')
    },
    customlayer() {
      this.$refs.layerCustom.style = 'border: 1px solid #409EFF; color: #409EFF'
      console.log('自定义布局');
      this.layoutItem = {
        AudioLayers: {
          AudioLayer: []
        },
        BlendList: {
          LocationId: []
        },
        MixList: {
          LocationId: []
        },
        VideoLayers: {
          VideoLayer: []
        }
      }
      this.customLayout = []
      this.isActive = ''
      this.isCustom = true
      this.$nextTick(() => {
        this.canvas = new fabric.Canvas('layoutConfig')
      })
      // this.$nextTick(() => {
      //   console.log('canvsa');
      //   let canvas = new fabric.Canvas('layoutConfig')
      //   let rect = new fabric.Rect({ width: 100, height: 50, fill: 'green' });
      //   rect.on('selected', function(e) { // 选中监听事件
      //     console.log('selected a rectangle', e.target);
      //   });
      //   rect.on('modified', function(e) { // 选中监听事件
      //     console.log('selected a modified', 'width:', e.target.width * e.target.scaleX, 'height:', e.target.height * e.target.scaleY);
      //   });
      //   canvas.add(rect);
      //   console.log(rect, '====');
      // })
    },
    selectedLayer(item, index) {
      this.isActive = index
      this.$refs.layerCustom.style = 'border: 1px solid #a8a8a8'
      this.isCustom = false
      this.layoutItem = item
      // this.$nextTick(() => {
      //   this.canvas = ''
      // })
      console.log(item, 'item', index, this.isActive, this.isCustom, this.layoutItem = item);
    },
    mouseDownCanvas(e) {
      console.log(e, 'x:', e.offsetX, 'y:', e.offsetY);
    },
    canvasClick(e) {
      console.log(e, '===点击', e.pageX, this.$refs.layoutConfig.getContext('2d').isPointInPath(e.offsetX, e.offsetY));
    },
    addlayout() {
      this.layoutItem = this.layoutList[0] || this.layoutItem
      console.log(this.layoutList[0], '===456', this.layoutItem, '=----', this.mediaList);
      this.layoutVisible = true

    },
    addlayoutVisible() {
      if (this.isCustom) {
        this.layoutItem.VideoLayers.VideoLayer.forEach((item, index) => {
          item.WidthNormalized = Math.floor((this.customLayout[index].width / 300) * 100) / 100
          item.HeightNormalized = Math.floor((this.customLayout[index].height / 300) * 100) / 100
          item.PositionNormalizeds.Position[0]
          = Math.floor(((this.customLayout[index].group.top) / 300) * 100) / 100 > 0
              ? Math.floor(((this.customLayout[index].group.top) / 300) * 100) / 100 : 0
          item.PositionNormalizeds.Position[1]
           = Math.floor(((this.customLayout[index].group.left) / 300) * 100) / 100 > 0
              ? Math.floor(((this.customLayout[index].group.left) / 300) * 100) / 100 : 0
          console.log('top:', Math.floor(((this.customLayout[index].group.top) / 300) * 100) / 100, this.customLayout[index].group.top);
        })
        console.log(this.layoutItem, this.customLayout, '=iscustom', this.casterId);
        const VideoLayer = []
        for (let index = 0; index < this.layoutItem.VideoLayers.VideoLayer.length; index++) {
          const parms = {
            FillMode: this.layoutItem.VideoLayers.VideoLayer[index].FillMode,
            HeightNormalized: this.layoutItem.VideoLayers.VideoLayer[index].HeightNormalized,
            WidthNormalized: this.layoutItem.VideoLayers.VideoLayer[index].WidthNormalized,
            PositionRefer: this.layoutItem.VideoLayers.VideoLayer[index].PositionRefer,
            PositionNormalized: this.layoutItem.VideoLayers.VideoLayer[index].PositionNormalizeds.Position,
            FixedDelayDuration: 5000
          }
          VideoLayer.push(parms)
        }
        service.CasterApiPopService({
          ActionName: 'AddCasterLayout',
          AppId: localStorage.getItem('appId'),
          Parameter: {
            BlendList: this.layoutItem.BlendList.LocationId,
            MixList: this.layoutItem.MixList.LocationId,
            VideoLayer: VideoLayer,
            AudioLayer: this.layoutItem.AudioLayers.AudioLayer,
            CasterId: this.CasterId
          }
        }).then(res => {
          console.log(res);
        }).catch(err => {
          console.log(err);
        })
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
.el-card {

}
/deep/ .prism-big-play-btn pause{
  display: none !important;
}

.canvas-box{
  border: 1px solid #a8a8a8;
  background-color: #4d4d4d;
  width: 300px;
  height: 300px;
}
.layer-show{
  margin-bottom: 20px;
  display: flex;
  .layer-box{
    margin-right: 5px;
    width:88px;
    height:45px;
    position: relative;
    cursor: pointer;
    .selected{
       border: 1px solid #409EFF!important;
    }
  }
  .layer-custom{
    width:88px;
    height:45px;
    cursor: pointer;
    line-height: 45px;
    text-align: center;
    border: 1px solid #a8a8a8;
    font-size: 12px;
    color: #a8a8a8
  }
}
.caster-layout-box{
  margin-bottom: 20px; padding: 0;
  height: 150px;
  position: relative;
  &:hover{
    background-color: rgb(214, 214, 214);
    cursor: pointer;
  }
 .caster-layout-play{
              position:absolute;
              top:0;
              left: 0;
              width: 30px;
              height: 30px;
              background-image: url(https://img.alicdn.com/tfs/TB1uMu_GL1TBuNjy0FjXXajyXXa-384-86.png);
              background-size: 315px;
              background-position: 36px -7px;
              z-index: 9999;
}
  .caster-layout-delete{
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
</style>
