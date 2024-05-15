import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import installElementPlus from './plugins/element'
import 'element-plus/dist/index.css'
// import store from './store/store'
import * as ElIconModules from '@element-plus/icons-vue'

// Vue.config.productionTip = false
const app = createApp(App)
installElementPlus(app)

app.use(router)
app.use(ElementPlus)
// app.use(store)
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