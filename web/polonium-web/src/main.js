import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import mitt from 'mitt';

const app = createApp(App)
const eventBus = mitt();
app.provide('$eventBus', eventBus);

app.mount('#app')
