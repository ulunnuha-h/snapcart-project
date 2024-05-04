import './assets/main.css'
import { createRouter, createWebHistory } from 'vue-router'
import MainPage from './pages/Main/Main-Index.vue'
import LoginPage from './pages/Auth/Login-Page.vue'
import RegisterPage from './pages/Auth/Register-Page.vue'
import { createApp } from 'vue'
import App from './App.vue/'

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
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes : routes
})

createApp(App).use(router).mount('#app')
