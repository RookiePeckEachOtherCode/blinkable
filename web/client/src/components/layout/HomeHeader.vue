<template>
  <div id="header">
    <el-menu
        :default-active="activeIndex2"
        class="el-menu-demo"
        mode="horizontal"
        router
        background-color="#78a355"
        text-color="#fff"

    >
      <div class="logo-box">CRK菜鸟营</div>
      <el-menu-item index="/home/main-view">主页</el-menu-item>
      <el-sub-menu index="2">
        <template #title  >组件</template>
        <el-menu-item index="/home/admin-home" >管理者界面</el-menu-item>
        <el-menu-item index="/home/paper-list">文章列表</el-menu-item>
        <el-menu-item index="/home/paper-edit">文章发布</el-menu-item>
        <el-sub-menu index="2-4">
          <template #title>item four</template>
          <el-menu-item index="2-4-1">item one</el-menu-item>
          <el-menu-item index="2-4-2">item two</el-menu-item>
          <el-menu-item index="2-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
      <el-menu-item index="3" disabled style="font-size: 17px">关于站点</el-menu-item>
      <el-menu-item index="4" style="color:#ffd04b; font-size: 17px" @click="goToDocumentation">菜鸟营文档库</el-menu-item>
      <el-menu-item index="5" @click="goOj" style="color: #7bff91;font-size: 17px">菜鸟营Oj</el-menu-item>
      <div class="right-box">
        <el-avatar v-if="loginstate" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        <el-menu-item index="/login" v-if="!loginstate" class="custom-menu-item" >Login</el-menu-item>
      </div>
    </el-menu>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import {useUserInfoStore} from "@/stores/userinfo";
import router from "@/router";
const  defaultAcitve=ref<String>(router.currentRoute.value.path)

const loginstate=useUserInfoStore().authFromLocal();
const activeIndex = ref('1')
const activeIndex2 = ref('1')
const goToDocumentation = () => {
  const documentationLink = "https://rookiepeckeachothercode.github.io/RookiableDoc/#/md/%E7%BC%96%E5%86%99%E5%85%A5%E9%97%A8";

  window.location.href =documentationLink
};
const goOj=()=>{
  const ojlink="http://122.51.56.135:8888/"
  window.location.href =ojlink
}
</script>

<style scoped lang="scss">
#header {
  background-color: transparent;
  color: #fff;
  display: flex;
  justify-content: space-between;
  align-items: center; /* 垂直居中对齐 */
}

.el-menu-demo {
  flex: auto;
  width: 100%;
  background-color: transparent;
  backdrop-filter: blur(10px); /* 调整模糊程度 */
  display: flex; /* 让子元素水平排列 */
  .el-sub-menu{
    mode:horizontal;
    &:hover{
      background-color:rgb(255,255,255,50%);
    }
    .el-menu-item {
        &:hover {
          color: #ffd04b;
          background-color:transparent;
        }
    }
  }

  .logo-box {
    width: 200px;
    text-align: center;
    line-height: 60px;
    font-size: 22px;
    font-weight: bold;
    filter: brightness(110%);
    text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #31b2e3, 0 0 35px #31b2e3, 0 0 40px #31b2e3, 0 0 50px #31b2e3, 0 0 75px #31b2e3;
    animation: pink 1.5s ease-in-out infinite alternate;

  }

  .el-menu-item {
    &:not(.is-disabled) {
      &:hover {
        color: #ffd04b;
        background-color:rgb(255,255,255,50%);
      }
    }
    color: white;
  }

  .right-box {
    display: flex;
    align-items: center; /* 垂直居中对齐 */
    font-size: 18px;
    margin-left: auto; /* 将 right-box 推到最右边 */

    .el-avatar {
      margin-right: 10px;
    }
  }
  .right-box .custom-menu-item {
    font-size: 18px;
    font-weight: bold;
    color: #60aee6;
  }

}
</style>
