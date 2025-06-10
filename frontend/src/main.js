import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue'
import router from './router'
import pinia from './store'
import App from './App.vue'

// 导入ArcoDesign图标
import ArcoVueIcon from '@arco-design/web-vue/es/icon'

// 导入ArcoDesign样式
import '@arco-design/web-vue/dist/arco.css'
import './style.css'

const app = createApp(App)

// 使用插件
app.use(ArcoVue)
app.use(ArcoVueIcon)
app.use(router)
app.use(pinia)

app.mount('#app')