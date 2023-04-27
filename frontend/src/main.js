import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { router } from './router'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from "vuex";
import { elNotification, elMessageBox } from './utils/ele'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'virtual:windi.css'
import store from './store'
import App from './App.vue'


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
