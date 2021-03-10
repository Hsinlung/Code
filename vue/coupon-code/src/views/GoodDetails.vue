<template>
<div class="gooddetails">
  <van-nav-bar
      title="标题"
      left-text="返回"
      right-text="按钮"
      left-arrow
      @click-left="onClickLeft"
      @click-right="onClickRight"
  />
  <van-row>
    <van-col span="24">
      <van-image
          width="100%"
          height="220"
          lazy-load
          src="https://img01.yzcdn.cn/vant/cat.jpeg"
      />
    </van-col>
  </van-row>
  <van-goods-action>
    <van-goods-action-icon icon="cart-o" text="购物车" v-model="num" :badge="num" />
    <van-goods-action-button type="warning" v-on:click="add_cart" text="加入购物车" />
    <van-goods-action-button type="danger" text="立即购买" />
  </van-goods-action>
</div>
</template>

<script>
import {Button, Col, Row, Image, GoodsAction, GoodsActionIcon, GoodsActionButton, NavBar, Toast} from 'vant';
export default {
  name: "GoodDetails",
  components: {
    [Col.name]: Col,
    [Row.name]: Row,
    [Button.name]: Button,
    [Image.name]: Image,
    [GoodsAction.name]: GoodsAction,
    [GoodsActionIcon.name]: GoodsActionIcon,
    [GoodsActionButton.name]: GoodsActionButton,
    [NavBar.name]: NavBar,
    [Toast.name]: Toast,
  },
  data() {
    return {
      num: '',
    };
  },
  methods: {
    add_cart: function (){
        if (this.num == ''){
          this.num = 1;
        }else{
          this.num += 1
        }
    },
    onClickLeft() {
      Toast('返回');
      this.$router.back()
    },
    onClickRight() {
      Toast('按钮');
    },
  }
}
</script>

<style scoped>

</style>