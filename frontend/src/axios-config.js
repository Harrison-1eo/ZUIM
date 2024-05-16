import axios from 'axios';
import {ElMessage} from "element-plus";

// 设置Axios的基础URL
axios.defaults.baseURL = 'http://localhost:8000';

var textDecoder = new TextDecoder('utf-8');
var textEncoder = new TextEncoder('utf-8');

// 设置Axios的请求拦截器

axios.interceptors.request.use(
    config => {
        // 除了login和register外，请求发送之前，将token放入请求头中
        if (config.url !== '/login' && config.url !== '/register') {
            config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
        }
        // 将请求数据 config.data 进行Base64编码
        if (config.data) {
            config.data = btoa(String.fromCharCode.apply(null, textEncoder.encode(JSON.stringify(config.data))));
            console.log("Base64编码结果：", config.data);
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// 设置Axios的响应拦截器
// 如果token过期（返回401），则跳转到登录页面
axios.interceptors.response.use(
    response => {
        // return response;
        // 将相应数据 response.data.data 进行Base64解码（使用UTF-8）
        if (response.data.data) {
            response.data.data = JSON.parse(textDecoder.decode(Uint8Array.from(atob(response.data.data), c => c.charCodeAt(0))));
            console.log("Base64解码结果：", response.data.data);
        }
        return response;
    },
    error => {
        if (error.response.status === 401) {
            ElMessage.error('登录无效或过期，请重新登录');
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            localStorage.removeItem('userId');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);


export default axios;
