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

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/register'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('user');
  
    // trying to access a restricted page + not logged in
    // redirect to login page
    if (authRequired && !loggedIn) {
      next('/login');
    } else if(!authRequired && loggedIn) {
      next('/');
    }else {
      next();
    }
  });
export default router;
