import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import installElementPlus from './plugins/element'
import 'element-plus/dist/index.css'


const app = createApp(App)
installElementPlus(app)

app.use(router)
app.use(ElementPlus)
app.mount('#app')
