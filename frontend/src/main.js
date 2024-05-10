import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import installElementPlus from './plugins/element'
import 'element-plus/dist/index.css'

import * as ElIconModules from '@element-plus/icons-vue'


const app = createApp(App)
installElementPlus(app)

app.use(router)
app.use(ElementPlus)
app.mount('#app')

// 统一注册el-icon图标
for(let iconName in ElIconModules){
    app.component(iconName,ElIconModules[iconName])
}
