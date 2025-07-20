import { createApp } from 'vue'
import App from './App.vue'
import vuetify from "./plugins/vuetify.ts";
import router from "./router/router.ts";
import {pinia} from "./plugins/pinia.ts";
import '@mdi/font/css/materialdesignicons.css'

const app = createApp(App)

app.use(vuetify)
app.use(router)
app.use(pinia)
app.mount('#app')
