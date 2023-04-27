import { createApp } from 'vue'

//#region css
import 'virtual:windi.css'
import 'element-plus/dist/index.css'
//#endregion

//#region package
import ElementPlus from 'element-plus'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from "vuex";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
//#endregion

//#region 自定义
import { elNotification, elMessageBox } from './utils/ele'
import store from './store'
import { router } from './router'
import App from './App.vue'
//#endregion


const app = createApp(App)

//#region 全局注册
app.config.globalProperties.$elNotification = elNotification
app.config.globalProperties.$elMessageBox = elMessageBox
app.config.globalProperties.$router = useRouter()
app.config.globalProperties.$route = useRoute()
app.config.globalProperties.$store = useStore()
//#endregion

app.use(store)
app.use(ElementPlus)
app.use(router)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.mount('#app')
