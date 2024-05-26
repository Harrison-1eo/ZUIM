import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import installElementPlus from './plugins/element'
import 'element-plus/dist/index.css'
import store from './store/store'
import * as ElIconModules from '@element-plus/icons-vue'

// Vue.config.productionTip = false
const app = createApp(App)
installElementPlus(app)

app.use(router)
app.use(ElementPlus)
app.use(store)
app.mount('#app')

// 统一注册el-icon图标
for (let iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
}
// new Vue({
//     el: '#app',
//     router,
//     store, //使用store
//     components: { App },
//     template: '<App/>'
// })

// 为了解决resizeObserver报错问题, 添加防抖函数
// const debounce = (fn, delay) => {
//     let timer
//     return (...args) => {
//         if (timer) {
//             clearTimeout(timer)
//         }
//         timer = setTimeout(() => {
//             fn(...args)
//         }, delay)
//     }
// }
//
// const _ResizeObserver = window.ResizeObserver;
// window.ResizeObserver = class ResizeObserver extends _ResizeObserver{
//     constructor(callback) {
//         callback = debounce(callback, 200);
//         super(callback);
//     }
// }
