import './assets/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import "./assets/main.scss"
import 'element-plus/dist/index.css'
import App from './App.vue'
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';

import Prism from 'prismjs';

VueMarkdownEditor.use(vuepressTheme, {
    Prism,
});

const app = createApp(App)

async function  asyncRegister(){
    const createPinia=(await import("pinia")).createPinia;
    app.use(createPinia());
    const router=(await import("@/router")).default;
    app.use(router);
    app.use(VueMarkdownEditor);
    app.mount('#app')

}
asyncRegister();
