<template>
  <div id="login">
    <div class="login-box">
      <div class="logo">
        <el-image src="https://tse1-mm.cn.bing.net/th/id/OIP-C.4dcJ_AHTJ81dikKbJ_xBtgHaGw?pid=ImgDet&rs=1" fit="fill" :lazy="true">
        </el-image>
      </div>
      <el-form ref="formRef" :model="form" :rules="rules">
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
    </div>
  </div>
</template>
<script setup lang="ts">
import {User ,Lock} from "@element-plus/icons-vue";
import {ref} from "vue";
import type {FormInstance} from "element-plus";
import {rules} from "@/rules/userinfo";
import {loginApi} from "@/apis/login";
import {useUserInfoStore} from "@/stores/userinfo";
import {ElMessage} from "element-plus";
import router from "@/router";

const useUserinfoStore=useUserInfoStore();

interface Form{
  username:string;
  password:string;
}
const form=ref<Form>({
  username:"admin",
  password:"admin"
});
//生成封装表单对象
const formRef=ref<FormInstance>()
const lin=async ()=>{
  const res=await loginApi(form.value);
  useUserinfoStore.setAuth(res.data.token)
  ElMessage.success("登录成功")
  router.push("/home")
};
const reset=()=>{
  formRef.value?.resetFields();//el自带的清空表单函数
}
</script>
<style lang="scss" scoped>
#login {
  height: 100vh;
  background: linear-gradient(to right, #BD4512FF, #1e4958);
  display: flex;
  justify-content: center;
   align-items: center;//水平垂直居中

  .login-box {
    width: 360px;
    height: 450px;
    background-color: #fff;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    .logo {
      width: 120px;
      height: 120px;
      border-radius: 50%;
      box-shadow: 0 0 4px 2px ;
      margin-bottom: 30px;//外边距
     overflow: hidden;//防止图片超出圆框
    } .el-form{
    width: 90%;
    margin-top: 30px;
    .input{
      height: 40px;
      :deep(.el-input__wrapper){
        border-radius: 20px;
       border: 1px solid #79ff61;
      }
    }
    .btns{
      display: flex;
     flex-direction: column;//平行
      //justify-content: center;//并列
      align-items: center;
      .btn{
        width: 100%;
        height: 40px;
        border-radius: 20px;
        margin-bottom: 20px;
        cursor: pointer;
      }
      .reset{
        text-align: center;
      }

    }
    }
  }
}
</style>