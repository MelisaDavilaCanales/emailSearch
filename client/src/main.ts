import "vue-toastification/dist/index.css";
import './assets/style.css'
import "./assets/customToast.scss";

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Toast from "vue-toastification";
import type { PluginOptions } from "vue-toastification";

import App from './App.vue'
import router from './router'

const app = createApp(App)

const options: PluginOptions = {
  hideProgressBar: true,
  timeout: 3000,

};

app.use(Toast, options)

app.use(createPinia())
app.use(router)

app.mount('#app')
