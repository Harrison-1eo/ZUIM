import axios from 'axios';
import {ElMessage} from "element-plus";
import backendBaseUrl from "@/utils/base-url-setting";

const axios_origin = axios.create({
    baseURL: backendBaseUrl,
    timeout: 3000
});

axios_origin.interceptors.response.use(
    response => {
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

export default axios_origin
