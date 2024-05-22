import axios from 'axios';
import { ElMessage } from "element-plus";
import {userCipherFrontend, userCipherBackend} from "@/utils/encrypt";
import {backendBaseUrl} from "@/utils/base-url-setting";

const axios_config = axios.create({
    baseURL: backendBaseUrl,
    timeout: 10000
})

// 设置Axios的请求拦截器
axios_config.interceptors.request.use(
    config => {
        config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
        console.log('Authorization >>> ', config.headers.Authorization);

        // 如果访问路径中包含upload，则不进行Base64编码
        if (!config.url.includes('upload')) {
            console.log(config.url);
            if (config.data) {
                // 将请求数据 config.data 进行加密
                console.log("加密前数据:", config.data);
                const dd = userCipherFrontend.encrypt(JSON.stringify(config.data));
                // config.data中存储数据格式为：
                // {
                //     data: '加密后的数据',
                //     pos: '加密位置',
                //     mac: '',
                //     length: ''
                // }
                config.data = {
                    data: dd.cipherText,
                    position: dd.position,
                    mac: '',
                    length: 0
                };
            }
        }
        return config;
    },
    error => {
        console.log('request error', error);
        return Promise.reject(error);
    }
);

// 设置Axios的响应拦截器
// 如果token过期（返回401），则跳转到登录页面
axios_config.interceptors.response.use(
    response => {
        // return response;
        // 将相应数据 response.data.data 进行Base64解码（使用UTF-8）
        if (response.data.data) {
            const encryptedData = response.data.en_data.data;
            const position = response.data.en_data.position;
            let decryptedData = '';
            try{
                decryptedData = userCipherBackend.decrypt(encryptedData, position);
            } catch (e) {
                console.log('解密失败：', e);
                console.log('解密失败：', encryptedData, position);
                ElMessage.error('解密失败');
            }
            response.data.data = JSON.parse(decryptedData);
            console.log("解密解码结果：", response.data.data);
        }
        return response;
    },
    error => {
        console.log('error', error)
        console.log('error response', error.response)
        if (error.response.status === 401) {
            console.log(error.response);
            ElMessage.error('登录无效或过期，请重新登录');
            localStorage.clear();
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);


export default axios_config
