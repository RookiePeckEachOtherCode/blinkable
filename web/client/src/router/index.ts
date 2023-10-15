import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect:"/home"
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    }
  ]
})
//设置路由守卫，验证用户身份
//to 即将进入的目标路由对象，包含路径，参数，查询参数等
//from 当前导航正要离开的路由对象，包含了当前路由信息
//next 函数，用于控制导航页行为，接收一个参数用于指定目标路由
//router.beforeEach((to,from,next)=>{
  /*if(userlogin){
    next()
  }else{
    next()
  }*/
//})
export default router
