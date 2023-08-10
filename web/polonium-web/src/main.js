import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'

import mitt from 'mitt';
const emitter = mitt();
app.config.globalProperties.emitter = emitter;

createApp(App).mount('#app')
