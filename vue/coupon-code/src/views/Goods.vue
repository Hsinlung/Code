<template>
<div class="goods">
  <van-nav-bar
      title="商品列表"
      left-text="返回"
      right-text="按钮"
      left-arrow
      @click-left="onClickLeft"
      @click-right="onClickRight"
  />
  <div ref="container">
    <van-search v-model="value" placeholder="请输入搜索关键词" />

    <van-sticky :container="container">
      <p v-bind:key="index" v-for="(n,index) in 100">dfkjsdlkfjsdk</p>
    </van-sticky>
  </div>


</div>
</template>
<script>

import {
  Sticky,
  TreeSelect,
  Divider,
  Card,
  Button,
  Col,
  Row,
  Swipe,
  SwipeItem,
  Search,
  Tag,
  NavBar,
  Toast,
  List,
  Cell,
  Image,
  Grid,
  GridItem,
  Sidebar,
  SidebarItem
} from 'vant';
export default {
  name: "Goods",
  components: {
    [Sticky.name]: Sticky,
    [TreeSelect.name]: TreeSelect,
    [Divider.name]: Divider,
    [Col.name]: Col,
    [Row.name]: Row,
    [Card.name]: Card,
    [Button.name]: Button,
    [Swipe.name]: Swipe,
    [SwipeItem.name]: SwipeItem,
    [Search.name]: Search,
    [Tag.name]: Tag,
    [NavBar.name]: NavBar,
    [Toast.name]: Toast,
    [List.name]: List,
    [Cell.name]: Cell,
    [Image.name]: Image,
    [Grid.name]: Grid,
    [GridItem.name]: GridItem,
    [Sidebar.name]: Sidebar,
    [SidebarItem.name]: SidebarItem,
  },
  mounted() {
    this.container = this.$refs.container;
  },
  data() {
    return {
      container: null,
      value: '',
      visible:true,
      activeKey: 0,
      active: 0,
      items: [{ text: '分组 1' }, { text: '分组 2' },{ text: '分组 1' }, { text: '分组 2' }],
    };
  },
  methods: {
    good_detail: function (){
      this.$router.push({path: '/gooddetails'})
    },
    onClickLeft() {
      Toast('返回');
    },
    onClickRight() {
      Toast('按钮');
    },
  }
}
</script>

<style scoped>
.sidebar-container{
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  width: 220px;
  overflow-y: auto;
}
.goods_list{
  text-align: left;
}
.goods_m{

}

.van-tag{
  margin-right: 5px;
}
</style>