import './assets/main.css'
import { createRouter, createWebHistory } from 'vue-router'
import MainPage from './pages/Main/Main-Index.vue'
import LoginPage from './pages/Auth/Login-Page.vue'
import RegisterPage from './pages/Auth/Register-Page.vue'
import { createApp } from 'vue'
import App from './App.vue/'
import CreatePage from './pages/Product/Create-Page.vue'

const routes = [
    {
        path : '/',
        component: MainPage
    },
    {
        path : '/login',
        component: LoginPage
    },
    {
        path : '/register',
        component: RegisterPage
    },
    {
        path : '/create',
        component: CreatePage
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes : routes
})

createApp(App).use(router).mount('#app')
