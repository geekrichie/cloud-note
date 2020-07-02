import Vue from 'vue'
import Login from '../components/Login'
import Register from '../components/Register'
import Main from '../components/Main'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/', component: Main }
]

const router = new VueRouter({
    routes // (缩写) 相当于 routes: routes
})

export default router;
