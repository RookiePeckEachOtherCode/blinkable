<template>
  <div id="login">
    <div class="login-box" >
      <div class="pre-box" @click="mySwitch">
        <h1>Welcome</h1>
        <p>join us</p>
        <div class="img-box">
        <img src="http://localhost:8080/image">
        </div>
      </div>
      <el-form ref="formRef" :model="form" :rules="rules" >
        <div class="logo">
          <el-image src="https://tse1-mm.cn.bing.net/th/id/OIP-C.4dcJ_AHTJ81dikKbJ_xBtgHaGw?pid=ImgDet&rs=1" fit="fill" :lazy="true">
          </el-image>
        </div>
        <el-form-item prop="username">
          <el-input class="input" v-model="form.username" placeholder="用户名" :prefix-icon="User"></el-input>
        </el-form-item>
        <el-form-item prop="password">
        <el-input class="input" v-model="form.password" placeholder="密码" :prefix-icon="Lock"></el-input>
         </el-form-item>
        <div class="btns">
          <el-button type="primary" class="btn" @click="lin" >登录</el-button>
           <div class="btn reset" @click="reset">重置</div>
        </div>
      </el-form>
      <el-form ref="formRef" :model="form" :rules="rules">
        <div class="logo">
          <el-image src="https://tse1-mm.cn.bing.net/th/id/OIP-C.4dcJ_AHTJ81dikKbJ_xBtgHaGw?pid=ImgDet&rs=1" fit="fill" :lazy="true">
          </el-image>
        </div>
        <el-form-item prop="username">
          <el-input class="input" v-model="form.username" placeholder="用户名" :prefix-icon="User"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input class="input" v-model="form.password" placeholder="密码" :prefix-icon="Lock"></el-input>
        </el-form-item>
        <div class="btns">
          <el-button type="primary" class="btn" @click="reg" >注册</el-button>
          <div class="btn reset" @click="reset">重置</div>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script setup lang="ts">
import {User ,Lock} from "@element-plus/icons-vue";
import {ref} from "vue";
import type {FormInstance} from "element-plus";
import {rules} from "@/rules/userinfo";
import {loginApi} from "@/apis/login";
import  {registerApi} from "@/apis/registe";
import {useUserInfoStore} from "@/stores/userinfo";
import {ElMessage} from "element-plus";
import router from "@/router";
import $ from 'jquery';
const useUserinfoStore=useUserInfoStore();

interface Form{
  username:string;
  password:string;
}
const form=ref<Form>({
  username:"",
  password:""
});
//生成封装表单对象
const formRef=ref<FormInstance>()
const lin=async ()=>{
const res=await loginApi(form.value);
 const user_id= res.user_id.toString()
   useUserinfoStore.setAuth(res.token,user_id,res.icon_url);
  ElMessage.success("登录成功")
  router.push("/home/main-view")
};
const reg=async ()=>{
  const res=await registerApi(form.value);
  const user_id=res.user_id.toString()
  useUserinfoStore.setAuth(res.token,user_id)
  ElMessage.success("注册成功")
  router.push("/home/main-view")
}
const reset=()=>{
  formRef.value?.resetFields();//el自带的清空表单函数
}
let flag=true;
const mySwitch=()=>{
  if(flag){
    $(".pre-box").css("transform","translateX(100%)")
    $(".pre-box").css("background-color","#c9e0ed")
    $(".img").css("src","https://ts1.cn.mm.bing.net/th/id/R-C.62a4420f03d0594f14c4cedb4ca92f55?rik=eoZdvlU%2bTEZCnQ&riu=http%3a%2f%2fimg95.699pic.com%2fphoto%2f50100%2f1450.jpg_wh860.jpg&ehk=UL6jUo5ODeseZEucJ0Z9yz%2fTk7HuvshjF94HW0i%2fEvM%3d&risl=&pid=ImgRaw&r=0")
  }else {
    $(".pre-box").css("transform","translateX(0%)")
    $(".pre-box").css("background-color","#edc9c9")
    $(".img").css("src","http://localhost:8080/image")
  }
  flag=!flag;
}
</script>
<style lang="scss" scoped>
#login {
  height: 100vh;
  background: linear-gradient(to right, rgb(246, 191, 206), rgb(153, 219, 245));
  display: flex;
  justify-content: center;
   align-items: center;//水平垂直居中.
  position: relative;


  .pre-box {
    border-radius: 10px;
    width: 360px;
    height: 450px;
    background-color: rgb(246, 191, 206);
    z-index: 1;
    top: 0;
    left: 0; // 设置左侧对齐
    border: 1px solid darkgrey;
    position: absolute;
    box-shadow: 4px 4px 3px rgba(0, 0, 4, 0.3);
    transition: 0.5s ease-in-out;
  }
  .pre-box h1{
    margin-top: 75px;
    text-align: center;
    letter-spacing: 5px;
    color: white;
    text-shadow:4px 4px 3px rgba(0, 0, 4, 0.1) ;
  }
  .pre-box p{
    color: white;
    height: 30px;
    line-height: 30px;
    text-align: center;
    margin: 20px 0;
    font-weight: bold;
    text-shadow:4px 4px 3px rgba(0, 0, 4, 0.1) ;
  }
  .img-box{
    width:130px;
    height: 130px;
    border-radius: 50%;
    box-shadow: 0 0 4px 2px ;
    margin-bottom: 30px;//外边距
    margin-top: 0px;
    overflow: hidden;//防止图片超出圆框
    margin: 20px auto;
  }
  .img-box img{
    width: 100%;
  }
  .login-box {
    //background-color: white;
    width: 720px;
    height: 450px;
    border: 1px solid white;
    box-shadow: 4px 4px 3px rgba(0, 0, 4, 0.2);
    border-radius: 10px;
    display: flex;
    flex-direction: row;
    justify-content: center; // 水平居中
    align-items: center; // 垂直居中
    position: relative;
    z-index: 2; // 提高 .login-box 的层叠顺序
    .el-form {
            display: flex;
            flex-direction: column; /* 垂直排列 */
            flex: 1;
            width: 50%;
            margin-top: 80px;
            align-items: center;
            .input {
              display: flex;
              flex-direction: column;
              width: 280px;
              height: 40px;

              :deep(.el-input__wrapper) {
                border-radius: 20px;
                border: 1px solid #79ff61;
              }
            }
      .logo {
        background-color: white;
        width: 80px;
        height: 80px;
        border-radius: 50%;
        box-shadow: 0 0 4px 2px ;
        margin-bottom: 30px;//外边距
        margin-top: 0px;
        overflow: hidden;//防止图片超出圆框
      }
            .btns {
              display: flex;
              flex-direction: column; /* 垂直排列 */
              align-items: center;
              margin-top: 15px; /* 控制按钮和输入框之间的垂直间隔 */

              .btn {
                width: 100px;
                height: 40px;
                border-radius: 20px;
                margin-bottom: 20px;
                cursor: pointer;
              }

              .reset {
                text-align: center;
              }
            }
          }

  }
}
</style>