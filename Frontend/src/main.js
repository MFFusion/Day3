import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import {createWebHashHistory,createRouter} from "vue-router";

const routes = [
    { path: '/', component: () => import('./components/Home.vue') },
    {path:'/catalog',component:()=>import('./components/catalog/CatalogPage.vue')},
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})



createApp(App)
    .use(router)
    .mount('#app')
