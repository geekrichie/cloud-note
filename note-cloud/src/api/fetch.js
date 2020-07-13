import axios from 'axios'
import { Message,Loading } from 'element-ui'
import Cookies from 'js-cookie'
import router from '../router'

let loadingInstance;

axios.interceptors.request.use(
    config => {
        loadingInstance=Loading.service({
            lock: true,
            text: '加载中,请稍后...',
            background: 'rgba(255,255,255,0.7)'
        });
        if(Cookies.get('jwt-token')){
            config.headers.Authorization = "Bearer " + Cookies.get("jwt-token")
        }
        return config
    },
    error => {
        Message({                  //使用element-ui的message进行信息提示
            showClose: true,
            message: error,
            type: "warning"
          });
        loadingInstance && loadingInstance.close();
        return Promise.reject(error)
    }
)

axios.interceptors.response.use(
    res => {
        const { data } = res.data
        if(data && data.token) {
           Cookies.set("jwt-token",data.token, {
                expires: 1/48
           })
        }
        loadingInstance && loadingInstance.close();
        return res;
    },
    error=> {
      loadingInstance && loadingInstance.close();
      switch(error.response.status)
      {
        case 400:
            localStorage.removeItem('user');
            Cookies.remove("jwt-token");
            router.push("/login")
            break;
        case 401:
            localStorage.removeItem('user');
            Cookies.remove("jwt-token");
            router.push("/login")
            break;
        default:
            console.log("未知错误")
            router.push("/login")
      }
      return Promise.reject(error)
    }
)
export default axios;