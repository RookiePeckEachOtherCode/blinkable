import { createRouter, createWebHistory } from 'vue-router'
import {useUserInfoStore} from "@/stores/userinfo";
import pinia from "@/stores/pinia";
const  userInfoStore=useUserInfoStore(pinia);
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect:"/home/main-view"
    },
    {
      path: '/home',
      name: 'home',
      component: ()=>import("@/views/HomeView.vue"),
      meta:{
        noAuth:true,
      },
      children:[{
        path:"admin-home",
        name:"admin-home",
        component:()=>import("@/components/adminhome/AdminHome.vue")
      },
        {
          path:"main-view",
          name:"main-view",
          component:()=>import("@/components/Mainview/Main.vue")
        },
        {
          path:"paper-list",
          name:"paper-list",
          component:()=>import("@/components/paper/paperlist.vue")
        },
        {
          path:"paper-edit",
          name:"paper-edit",
          component:()=>import ("@/components/paper/editpaper.vue")
        },
      ]},
    {
      path: '/login',
      name: 'login',
      meta:{
        noAuth:true,
      },
      component: () => import('../views/Login.vue')
    }
  ]
})
//设置路由守卫，验证用户身份
//to 即将进入的目标路由对象，包含路径，参数，查询参数等
//from 当前导航正要离开的路由对象，包含了当前路由信息
//next 函数，用于控制导航页行为，接收一个参数用于指定目标路由
router.beforeEach((to,from,next)=>{
  if(to.meta.noAuth||userInfoStore.authFromLocal()){
    next()
  }else{
    router.push("/login")
  }
})
export default router
