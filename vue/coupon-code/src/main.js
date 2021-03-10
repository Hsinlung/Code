import Vue from 'vue'
import App from './App.vue'
// import './plugins/vant.js' //按需加载后就不允许再配置全局引入组件
import router from './router'
import { Lazyload } from 'vant';

Vue.use(Lazyload);

// 注册时可以配置额外的选项
Vue.use(Lazyload, {
  lazyComponent: true
});

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
