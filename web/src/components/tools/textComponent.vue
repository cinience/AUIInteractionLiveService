<template>
  <div style="padding:12px">
    <el-button  type="primary" size="small" @click="textVisible = true">添加文字组件</el-button>
    <!-- 文本展示 -->
    <el-row :gutter="24" style="margin-top: 20px">
      <el-col  :span="24" v-for="item in textList" :key="item.ComponentId">
        <div style="margin-bottom:5px">{{item.ComponentName}}</div>
        <div class="caster-text-box" @click="selectedText(item)">
          <div class="caster-text-play" @click="textEdit(item)"></div>
          <div  class="caster-text-delete" @click="deleteText(item)"></div>
              <div :style="{
              position: 'absolute',
              top: (item.ComponentLayer.PositionNormalizeds ?  (item.ComponentLayer.PositionNormalizeds.Position[0] * 100) : 0)+ '%',
              left:(item.ComponentLayer.PositionNormalizeds ? (item.ComponentLayer.PositionNormalizeds.Position[1] * 100): 0) + '%',
              color: item.TextLayerContent.Color.replace('0x','#'),
              fontFamily: item.TextLayerContent.FontName,
              }">{{item.TextLayerContent.Text}}</div>
        </div>
      </el-col>
    </el-row>
  <!-- 设置文本 -->
 <el-dialog append-to-body :title="textTitle" :visible.sync="textVisible" width="700px">
      <el-row :gutter="24" style="width:100%">
        <el-col :span="12">
          <div class="text-show">
            <span ref="text" class="text" :style="{
              color: textInfo.TextLayerContent.Color,
              fontSize: textInfo.size + 'px',
              fontFamily: textInfo.TextLayerContent.FontName,
              position: 'absolute',
              top: (top * 100) + '%',
              left: (left * 100) + '%',
            }">{{textInfo.TextLayerContent.Text}}</span>
          </div>
          <el-input type="textarea"  v-model="textInfo.TextLayerContent.Text" maxlength="300" show-word-limit></el-input>
        </el-col>
        <el-col :span="12">
             <el-form
        label-position="right"
        :model="textInfo"
        ref="form"
      >
        <el-form-item label="名称" required label-width="70px">
          <el-input size="small" v-model="textInfo.ComponentName"></el-input>
        </el-form-item>
        <el-form-item label="字体" label-width="70px">
           <el-select style="width: 100%"  size="small" v-model="textInfo.TextLayerContent.FontName" placeholder="请选择">
             <el-option
               v-for="item in textType"
               :key="item.value"
               :label="item.label"
               :value="item.value">
             </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="尺寸" label-width="70px">
           <el-select style="width:30%" size="small" v-model="textInfo.size" >
             <el-option
               v-for="item in textSize"
               :key="item.value"
               :label="item.label"
               :value="item.value">
             </el-option>
            </el-select>
            <el-input style="width:68%;margin-top: 4px;margin-left: 2%" size="small" type="color" v-model="textInfo.TextLayerContent.Color">
              <template slot="prepend">{{textInfo.TextLayerContent.Color}}</template>
            </el-input>
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
        <el-button type="primary" @click="addText">确 定</el-button>
        <el-button @click="textVisible = false">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { service } from '@/api';
export default {
  props: {
    textList: {
      type: Array,
      required: false,
      default: () => [],
    },
    textTotal: {
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
    },
  },
  data() {
    return {
      textType: [
        {
          label: '楷体',
          value: 'KaiTi'
        },
        {
          label: '阿里巴巴普惠体-常规',
          value: 'AlibabaPuHuiTi-Regular'
        },
        {
          label: '阿里巴巴普惠体-粗体',
          value: 'AlibabaPuHuiTi-Bold'
        },
        {
          label: '阿里巴巴普惠体-细体',
          value: 'AlibabaPuHuiTi-Light'
        },
        {
          label: '思源黑体-常规',
          value: 'NotoSansHans-Regular'
        },
        {
          label: '思源黑体-粗体',
          value: 'NotoSansHans-Bold'
        },
        {
          label: '思源黑体-细体',
          value: 'NotoSansHans-Light'
        }
      ],
      textSize: [
        {
          label: '12',
          value: 12
        },
        {
          label: '14',
          value: 14
        },
        {
          label: '16',
          value: 16
        },
        {
          label: '18',
          value: 18
        },
        {
          label: '24',
          value: 24
        },
        {
          label: '32',
          value: 32
        },
        {
          label: '36',
          value: 36
        },
        {
          label: '42',
          value: 42
        },
        {
          label: '48',
          value: 48
        }
      ],
      LocationId: '',
      textVisible: false,
      currentPage: 1,
      textTitle: '添加文字组件',
      textInfo: {
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
        ComponentType: 'text',
        Effect: 'none',
        LocationId: '',
        TextLayerContent: {
          Color: '0xff0000',
          FontName: 'KaiTi',
          SizeNormalized: '',
          Text: ''
        },
        size: ''
      },
      top: 0,
      left: 0,
      pageSize: this.defaultPageSize,
    };
  },
  watch: {
    textVisible: function(val) {
      if (!val) {
        this.top = 0
        this.left = 0
        this.textInfo = {
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
          ComponentType: 'text',
          Effect: 'none',
          LocationId: '',
          TextLayerContent: {
            Color: '0xff0000',
            FontName: 'KaiTi',
            SizeNormalized: '',
            Text: ''
          },
          size: ''
        }
      } else {
        this.textInfo.TextLayerContent.Color = this.textInfo.TextLayerContent.Color.replace('0x', '#')
        console.log(this.textInfo.TextLayerContent.Color, '=this.textInfo.TextLayerContent.Color');
      }
    }
  },
  methods: {
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
    addText() {
      this.textInfo.ComponentLayer.PositionNormalizeds.Position[0] = Math.floor(this.top * 100) / 100
      this.textInfo.ComponentLayer.PositionNormalizeds.Position[1] = Math.floor(this.left * 100) / 100
      this.textVisible = false
      this.textInfo.TextLayerContent.Color = this.textInfo.TextLayerContent.Color.replace('#', '0x')
      console.log(this.textInfo, '====textInfo', 'AddCasterComponent', this.top, this.left);
      service.CasterApiPopService({
        ActionName: 'AddCasterComponent',
        AppId: localStorage.getItem('appId'),
        Parameter: {
          CasterId: this.CasterId,
          ComponentLayer: JSON.stringify(this.textInfo.ComponentLayer),
          ComponentType: 'text',
          LocationId: 'RC02',
          ComponentName: this.textInfo.ComponentName,
          Effect: this.textInfo.Effect,
          TextLayerContent: JSON.stringify(this.textInfo.TextLayerContent)
        }
      }).then(res => {
        console.log(JSON.parse(res.data.result.response), '场景');
      }).catch(err => {
        console.log(err);
      })

    },
    selectedText(item) {

    },
    textEdit(item) {
      this.textInfo = item
      this.top = item.ComponentLayer.PositionNormalizeds.Position[0]
      this.left = item.ComponentLayer.PositionNormalizeds.Position[1]
      this.textVisible = true
    },
    deleteText(item) {

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
/deep/ .prism-big-play-btn pause{
  display: none !important;
}
/deep/ .el-card__body {
  padding: 0;
  height: 320px;
  position: relative;
}
.time {
  font-size: 13px;
  position: absolute;
  bottom: 12px;
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
/deep/ .el-textarea{
  width: 300px;
  height: 120px;
  margin-top: 5px;
}
/deep/ .el-textarea__inner{
    height: 120px;
    width: 300px;
}
/deep/ .el-input-group__prepend{
    // width: 50%;
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
</style>
