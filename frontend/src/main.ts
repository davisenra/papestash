import './assets/main.css';

import { createApp } from 'vue';
import App from '@/App.vue';
import { VueQueryPlugin } from '@tanstack/vue-query';

const app = createApp(App);
// TODO: add Vue Router
// TODO: add Pinia
app.use(VueQueryPlugin);
app.mount('#app');
