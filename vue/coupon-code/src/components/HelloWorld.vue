<template>
  <div class="hello">
    <p></p>
    <van-row type="flex" justify="center" >
      <van-col span="18">
        <van-form @submit="onSubmit">
          <van-field
              clearable
              v-model="code"
              name="优惠码"
              label="优惠码"
              placeholder="请填写优惠码"
              :rules="[{ required: true}]"
          />
          <div style="margin: 16px;">
            <van-button round block type="info"  size="small" native-type="submit">兑换</van-button>
          </div>
        </van-form>
      </van-col>
    </van-row>
  </div>
</template>

<script>
import { Col, Row, Grid, GridItem, Field, Form, Button  } from 'vant';

export default {
  name: 'HelloWorld',
  data() {
    return {
      value: "",
      code: "",
    }
  },
  components: {
    //你用的时候还得按照这样的格式再次注册,不然报错
    [Col.name]: Col,
    [Row.name]: Row,
    [Grid.name]: Grid,
    [GridItem.name]: GridItem,
    [Field.name]: Field,
    [Form.name]: Form,
    [Button.name]: Button,
  },
  methods: {
    onSubmit(values) {
      console.log('submit', values);
      if (values){
        this.$router.push({path: '/goods'})
      }
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
