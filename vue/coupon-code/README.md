<van-row>
    <van-col span="6">
      <van-sidebar v-model="activeKey">
        <van-sidebar-item title="标签名称" />
        <van-sidebar-item title="标签名称" />
        <van-sidebar-item title="标签名称" />
      </van-sidebar>
    </van-col>
    <van-col span="16">
      <!-- 商品卡片 -->
      <div v-bind:key="index" v-for="(n,index) in 10">
        <van-card
            title="“26名科学家”，公开联名攻击中国和世卫组织"
            thumb="https://img01.yzcdn.cn/vant/ipad.jpeg"
            style="padding: 5px;"
        >
          <template #footer>
            <van-button size="small" type="warning" round>添加</van-button>
          </template>
        </van-card>
        <van-divider />
      </div>
      
    </van-col>
  </van-row>
  
  
  
  
  
  
  
  
<van-tree-select ref="container" height="100%" :items="items" :main-active-index.sync="active">

      <template #content>
        <van-sticky :container="container">
        <!-- 商品卡片 -->
        <div v-bind:key="index" v-for="(n,index) in 10">
          <van-card
              title="“26名科学家”，公开联名攻击中国和世卫组织"
              thumb="https://img01.yzcdn.cn/vant/ipad.jpeg"
          >
            <template #footer>
              <van-button size="small"  type="warning" round>添加</van-button>
            </template>
          </van-card>
          <van-divider />
        </div>
        </van-sticky>
      </template>

</van-tree-select>