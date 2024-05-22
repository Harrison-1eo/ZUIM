import axios from 'axios';
import {ElMessage} from "element-plus";
import {backendBaseUrl} from "@/utils/base-url-setting";

console.log('backendBaseUrl >>> ', backendBaseUrl);

const axios_origin = axios.create({
    baseURL: backendBaseUrl,
    timeout: 10000
});

axios_origin.interceptors.request.use(
    config => {
        console.log('URL >>> ', config.url);
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

export default axios_origin
