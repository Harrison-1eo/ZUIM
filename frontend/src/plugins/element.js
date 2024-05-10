import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
import locale from 'element-plus/es/locale/lang/zh-cn'

export default (app) => {
  app.use(ElementPlus, { locale })
}
