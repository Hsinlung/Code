import axios from 'axios'            //引入axios
import router from '../router'        //引入路由
import { Toast } from 'vant' //引入提示层

export function post(obj) {              //调用时参数为一个对象
    return new Promise((resolve, reject) => { //穿件Promise实例
        axios({                          //单例js 不允要this
            method: "POST",              //请求方式
            url: obj.url,                //请求地址
            data: obj.data               //请求参数
        }).then(res => {                 //请求成功操作
            if (obj.IS !== false) {      //是否需要弹框提示
                successState(res)        //调用弹框封装
            }

            resolve(res)                 //新Promise对象实例（成功）
        }).catch(err => {                //请求失败操作
            let status = err.response.status;  //错误状态码
            switch (status) {
                case 400:                 //状态码400

                    Toast.fail('数据格式错误，请重新填写！');

                    break;
                case 401:
                    Toast.fail('登录过期，已为您跳转到登录页！');

                    localStorage.removeItem("token")
                    router.push({ path: "/" });
                    break;
                case 500:
                    Toast.fail('服务器异常，请稍后再试！');

                    break;
                default:
                    Toast.fail('未知错误');//不在上述范围内的错误处理

            }
            reject(err)                    //新Promise对象实例（失败）
        })
    })
}
//拦截器
axios.interceptors.request.use(function (config) { // 每次请求时会从localStorage中获取token
    let token = localStorage.getItem("Authorization")

    if (token) {
        config.headers.common['Authorization'] = token

    }
    return config
}, function (error) {
    console.log(error)
    return Promise.reject(error);
})

//定义弹出层
function successState(res) {

    if (res.status == 200) {
        if (res.data.code == "200") {  //请求成功状态码

            Toast.success(res.data.msg);
            return
        } else if (res.data.code == "-1") { //请求失败状态码
            Toast.fail(res.data.data.msg,);

            return
        }
    }

}
